package issueregistery

import (
	"fmt"
	"libraryManagement/internal/models"
	"libraryManagement/utility/errorss"
	"time"

	"gorm.io/gorm"
)

type IssueRegistryService struct {
	DB *gorm.DB
}

var issueStatuses = []string{"accepted"}

func (issueService *IssueRegistryService) CreateIssueRequest(readerId uint, expectedReturn time.Time, issueStatus string, bookId uint, issuerId uint) error {
	var book = models.BookInventory{}
	DB := issueService.DB

	bookResult := DB.Where("ID= ?", bookId).Find(&book)

	checkPreviousIssue := DB.Where("isbn = ? AND reader_id = ?  AND return_date IS  NULL", book.ISBN, readerId).First(&models.IssueRegistery{})

	if err := checkPreviousIssue.Error; err == nil {

		return fmt.Errorf("this request is already Exist")
	}
	if bookResult.Error != nil {

		return bookResult.Error
	}
	currentTime := time.Now()
	var issueRequest = models.IssueRegistery{
		ReaderID:           readerId,
		IssueStatus:        issueStatus,
		ISBN:               book.ISBN,
		IssueApproverID:    issuerId,
		IssueDate:          &currentTime,
		ExpectedReturnDate: &expectedReturn,
	}

	if err := DB.Create(&issueRequest).Error; err != nil {

		return err
	}
	fmt.Println("issue")
	return nil
}

func (issueService *IssueRegistryService) UpdateIssueRequest(isbn string, readerId uint, aprooverId uint, requestDate time.Time) error {
	db := issueService.DB
	var issueModel = models.IssueRegistery{}

	result := db.Where("isbn =? AND reader_id=? AND issue_status=?", isbn, readerId, "issued").First(&issueModel)
	if result.Error != nil {
		return result.Error
	}
	result = db.Where("isbn =? AND reader_id=? AND issue_status=?", isbn, readerId, "issued").Updates(&models.IssueRegistery{ReturnDate: &requestDate, ReturnApproverID: aprooverId, IssueStatus: "returned"})

	return result.Error
}

func (issueService *IssueRegistryService) GetAllIssueRequest(libId uint) ([]map[string]interface{}, *errorss.AppError) {
	db := issueService.DB
	var result []map[string]interface{}
	if err := db.Model(&models.IssueRegistery{}).
		Joins("join book_inventories b on  b.isbn = issue_registeries.isbn").
		Where("b.lib_id =?", libId).Find(&result).Error; err != nil {

		return nil, errorss.BadRequest("No Issue Found", "no issues founds")
	}
	return result, nil
}
