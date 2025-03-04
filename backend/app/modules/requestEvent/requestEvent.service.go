package requestevent

import (
	"fmt"
	issueregistery "libraryManagement/app/modules/issueRegistery"
	"libraryManagement/internal/dto"
	"libraryManagement/internal/models"
	"libraryManagement/utility"
	"libraryManagement/utility/errorss"
	"time"

	"gorm.io/gorm"
)

type RequestEventService struct {
	db *gorm.DB
}

type RequestEventTemp struct {
	models.RequestEvent
	Book models.BookInventory
}

func (requestService *RequestEventService) AddRequest(libId uint, requestPayload dto.RequestEventDTO) *errorss.AppError {

	db := requestService.db

	if !utility.RequestTypes[requestPayload.RequestType] {
		return errorss.BadRequest("Invalid Request Type", "invalid request type")
	}

	var book models.BookInventory
	if err := db.Where("id=? AND lib_id=?", requestPayload.BookID, libId).First(&book).Error; err != nil {
		return errorss.NotFound("Book Not Found", "book not found")
	}
	fmt.Println(requestPayload.RequestType)
	if requestPayload.RequestType == "return" {

		result := db.Where("isbn = ? AND reader_id = ? AND issue_status = ? ", book.ISBN, requestPayload.ReaderID, "issued").First(&models.IssueRegistery{})
		if err := result.Error; err != nil {

			return errorss.NotFound("No Issue Found", "no issued book found")
		}

	} else {

		result := db.Where("book_id=?  AND reader_id=? AND request_status =? ", requestPayload.BookID, requestPayload.ReaderID, "pending").First(&models.RequestEvent{})
		if err := result.Error; err == nil {
			return errorss.BadRequest("Request Already Exists", "request already exists")
		}
		result = db.Where("isbn =? AND reader_id=? AND issue_status=?", book.ISBN, requestPayload.ReaderID, "issued").First(&models.IssueRegistery{})

		if err := result.Error; err == nil {
			return errorss.BadRequest("Book Already Issued", "book already issued")
		}
		if err := db.Where(" id=? AND available_copies >= ?", book.ID, 1).First(&models.BookInventory{}).Error; err != nil {

			return errorss.NotFound("Book Not Available", "book not available")
		}
	}

	response := db.Table("request_events").Create(&requestPayload)
	if response.Error != nil {
		return errorss.InternalServerError("Failed to create request", "failed to create request")
	}
	return nil
}

func (requestService *RequestEventService) GetAllRequest(userId, libId uint) ([]map[string]interface{}, *errorss.AppError) {
	DB := requestService.db
	var requests []map[string]interface{}
	result := DB.Table("request_events r").Select("r.*", "b	.title").
		Joins("JOIN book_inventories b ON b.id = r.book_id").
		Where("b.lib_id = ?", libId).
		Find(&requests)
	_ = result
	if result.Error != nil {
		return nil, errorss.InternalServerError("No Requests Found", "no requests found")
	}
	return requests, nil
}
func (requestService *RequestEventService) GetRequest(reader_id, reqId uint) (*models.RequestEvent, *errorss.AppError) {

	DB := requestService.db
	var request models.RequestEvent
	var user models.User
	DB.Where("id=?", reader_id).Find(&user)

	if user.Role != "reader" {
		if err := DB.Preload("Book").Where("req_id = ?", reqId).First(&request).Error; err != nil {
			return nil, errorss.NotFound("Request Not Found", "request not found")
		}
	} else if err := DB.Preload("Book", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, title, isbn")
	}).Omit("approver_id").
		Where(" reader_id=?  AND req_id = ?", reader_id, reqId).
		First(&request).Error; err != nil {
		return nil, errorss.Forbidden("Not Authorized", "you are not authorized to view this request")
	}

	// var book models.BookInventory
	if !isRequestBelong(DB, user, request) {
		return nil, errorss.NotFound("Request Not Found", "request not found")
	}

	return &request, nil
}

func (requestService *RequestEventService) UpdateRequest(libId, reqId, userId uint, statusType dto.RequestIssueStatus) *errorss.AppError {

	DB := requestService.db
	currentTime := time.Now()

	var request = models.RequestEvent{}
	var book = models.BookInventory{}
	var user models.User
	DB.Where("id=?", userId).First(&user)
	if result := DB.Where("req_id = ?", reqId).First(&request); result.Error != nil {

		return errorss.NotFound("Request Not Found", "request not found")
	}
	if request.RequestStatus != "pending" {
		return errorss.BadRequest("Request Already Processed", "request already processed")
	}

	// if err := DB.Where("req_id = ? AND request_status =?", reqId, "pending").
	// 	First(&models.RequestEvent{}).Error; err != nil {
	// 	return errorss.NotFound("no request found", "request not found")
	// }

	if !isRequestBelong(DB, user, request) {
		return errorss.Forbidden("not authorized", "you are not authorized to update this request")
	}

	res := DB.Transaction(func(tx *gorm.DB) error {

		result := tx.Model(&models.RequestEvent{}).Where("req_id =?", reqId).Updates(&models.RequestEvent{ApprovalDate: &currentTime, ApproverID: userId, RequestStatus: statusType.Type})
		if result.Error != nil {
			fmt.Println("not able to update the request")
			return errorss.BadRequest("failed to update request", "failed to update request")
		}
		if statusType.Type == "accepted" {

			if err := tx.Where("id = ?", request.BookID).First(&book).Error; err != nil {
				return errorss.NotFound("book bot found", "book not found")
			}

			if request.RequestType == "return" {
				if err := updateIssue(tx, book.ISBN, request.ReaderID, userId, currentTime); err != nil {
					return err
				}
				DB.Where("id =?", request.BookID).Updates(&models.BookInventory{AvailableCopies: book.AvailableCopies + 1})

			} else if request.RequestType == "issue" {
				if err := addIssue(tx, request.ReaderID, currentTime.Add(time.Hour*24*15), "issued", request.BookID, userId); err != nil {
					return err
				}
				fmt.Println(book.AvailableCopies)
				tx.Where("id =?", request.BookID).Updates(&models.BookInventory{AvailableCopies: book.AvailableCopies - 1})
			}
		}
		return nil
	})
	if res != nil {
		return errorss.InternalServerError("Failed to update request", "failed to update request")
	}
	return nil
}

func (requestService *RequestEventService) GetUserRequests(userId uint) ([]models.RequestEvent, error) {

	DB := requestService.db
	var requests []models.RequestEvent

	DB.Model(&models.RequestEvent{}).Preload("Book", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, title, isbn")
	}).
		Where("request_events.reader_id = ?", userId).
		Find(&requests)

	fmt.Println(requests)
	return requests, nil

}

func (requestService *RequestEventService) DeleteRequest(userId, reqId uint) *errorss.AppError {

	DB := requestService.db
	var request models.RequestEvent
	var user models.User
	if err := DB.Where("req_id=?", reqId).First(&request).Error; err != nil {

		return errorss.NotFound("Request Not Found", "request not found")
	}
	DB.Where("id=?", userId).First(&user)
	if !isRequestBelong(DB, user, request) {
		return errorss.Forbidden("Not Authorized", "you are not authorized to delete this request")
	}
	DB.Where("req_id = ?", reqId).Delete(&request)
	return nil
}

func isRequestBelong(DB *gorm.DB, user models.User, request models.RequestEvent) bool {

	if err := DB.Where("id=? AND lib_id=?", request.BookID, user.LibId).First(&models.BookInventory{}).Error; err != nil {
		return false
	}
	return true
}

func addIssue(db *gorm.DB, readerId uint, expectedReturn time.Time, issueStatus string, bookId uint, issuerId uint) error {
	issueService := issueregistery.IssueRegistryService{DB: db}
	err := issueService.CreateIssueRequest(readerId, expectedReturn, issueStatus, bookId, issuerId)
	return err
}

func updateIssue(db *gorm.DB, isbn string, readerId uint, aprooverId uint, returnDate time.Time) error {
	issueService := issueregistery.IssueRegistryService{DB: db}
	err := issueService.UpdateIssueRequest(isbn, readerId, aprooverId, returnDate)
	return err
}
