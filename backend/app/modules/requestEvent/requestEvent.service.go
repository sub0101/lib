package requestevent

import (
	"fmt"
	issueregistery "libraryManagement/app/modules/issueRegistery"
	"libraryManagement/internal/dto"
	"libraryManagement/internal/models"
	"libraryManagement/utility"
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

func (requestService *RequestEventService) AddRequest(libId uint, requestPayload dto.RequestEventDTO) error {

	db := requestService.db

	if !utility.RequestTypes[requestPayload.RequestType] {
		return fmt.Errorf("Requet Type is Not valid")
	}

	var book models.BookInventory
	if err := db.Where("id=? AND lib_id=?", requestPayload.BookID, libId).First(&book).Error; err != nil {
		return fmt.Errorf("book Not Found")
	}
	fmt.Println(requestPayload.RequestType)
	if requestPayload.RequestType == "return" {
		result := db.Where("isbn = ? AND reader_id = ? AND issue_status = ? ", book.ISBN, requestPayload.ReaderID, "issued").First(&models.IssueRegistery{})
		if err := result.Error; err != nil {

			return fmt.Errorf("Not issued book found")
		}

	} else {

		result := db.Where("book_id=?  AND reader_id=? AND request_status =? ", requestPayload.BookID, requestPayload.ReaderID, "pending").First(&models.RequestEvent{})
		if err := result.Error; err == nil {
			return fmt.Errorf("Issue request already exist")
		}
		result = db.Where("isbn =? AND reader_id=? AND issue_status=?", book.ISBN, requestPayload.ReaderID, "issued").First(&models.IssueRegistery{})

		if err := result.Error; err == nil {
			return fmt.Errorf("Book is already Issued kindly return it first")
		}
		if err := db.Where(" id=? AND available_copies >= ?", book.ID, 1).First(&models.BookInventory{}).Error; err != nil {

			return fmt.Errorf("Not Enough Available Copies")
		}
	}

	response := db.Table("request_events").Create(&requestPayload)
	return response.Error
}

func (requestService *RequestEventService) GetAllRequest(userId, libId uint) ([]map[string]interface{}, error) {
	DB := requestService.db
	var requests []map[string]interface{}
	result := DB.Table("request_events r").Select("r.*").
		Joins("JOIN book_inventories b ON b.id = r.book_id").
		Where("b.lib_id = ?", libId).
		Find(&requests)
	_ = result
	fmt.Println(requests)
	return requests, nil
}
func (requestService *RequestEventService) GetRequest(reader_id, reqId uint) (*models.RequestEvent, error) {

	DB := requestService.db
	var request models.RequestEvent
	var user models.User
	DB.Where("id=?", reader_id).Find(&user)

	if user.Role != "reader" {
		if err := DB.Preload("Book").Where("req_id = ?", reqId).First(&request).Error; err != nil {
			return nil, fmt.Errorf("no request found")
		}
	} else if err := DB.Preload("Book", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, title, isbn")
	}).Omit("approver_id").
		Where(" reader_id=?  AND req_id = ?", reader_id, reqId).
		First(&request).Error; err != nil {
		return nil, fmt.Errorf("no request found")
	}

	// var book models.BookInventory
	if !isRequestBelong(DB, user, request) {
		return nil, fmt.Errorf("no record found")
	}

	return &request, nil
}

func (requestService *RequestEventService) UpdateRequest(libId, reqId, userId uint, statusType dto.RequestIssueStatus) error {

	DB := requestService.db
	currentTime := time.Now()

	var request = models.RequestEvent{}
	var book = models.BookInventory{}
	var user models.User
	DB.Where("id=?", userId).First(&user)
	if result := DB.Where("req_id = ?", reqId).First(&request); result.Error != nil {

		return result.Error
	}
	if !isRequestBelong(DB, user, request) {
		return fmt.Errorf("no request found")
	}

	res := DB.Transaction(func(tx *gorm.DB) error {

		result := tx.Model(&models.RequestEvent{}).Where("req_id =?", reqId).Updates(&models.RequestEvent{ApprovalDate: &currentTime, ApproverID: userId, RequestStatus: statusType.Type})
		if result.Error != nil {
			fmt.Println("not able to update the request")
			return result.Error
		}
		if statusType.Type == "accepted" {

			if err := tx.Where("id = ?", request.BookID).First(&book).Error; err != nil {
				return fmt.Errorf("Book Not Found")
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

	return res
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

func (requestService *RequestEventService) DeleteRequest(userId, reqId uint) error {

	DB := requestService.db
	var request models.RequestEvent
	var user models.User
	if err := DB.Where("req_id=?", reqId).First(&request).Error; err != nil {

		return fmt.Errorf("no request found")
	}
	DB.Where("id=?", userId).First(&user)
	if !isRequestBelong(DB, user, request) {
		return fmt.Errorf("No Record Found or Can not delete")
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
