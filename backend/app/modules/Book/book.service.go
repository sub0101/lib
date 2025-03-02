package book

import (
	"fmt"
	auth "libraryManagement/app/modules/Auth"
	"libraryManagement/internal/dto"
	"libraryManagement/internal/models"
	"libraryManagement/utility/errorss"

	"gorm.io/gorm"
)

type BookService struct {
	db *gorm.DB
}

var book models.BookInventory
var response dto.ResponseBookInfo
var user *models.User

func (bookService *BookService) AddBook(userId, libId uint, book models.BookInventory) *errorss.AppError {

	db := bookService.db

	if err := db.Where("ID = ? AND lib_id = ?", userId, libId).First(&user).Error; err != nil {
		return errorss.Forbidden("Can not Addd Book", "You are not authorized to access this resource")
	}

	if book.AvailableCopies != book.TotalCopies {
		return errorss.BadRequest("Can not add Book", "Total copies are not equal to available copies")
	}

	var oldBook models.BookInventory
	isExist := db.Where("isbn=? AND lib_id=?", book.ISBN, book.LibID).First(&oldBook)
	if err := isExist.Error; err == nil {
		oldBook.TotalCopies += book.TotalCopies
		oldBook.AvailableCopies += book.AvailableCopies
		updatedResult := db.Where("ID = ?", oldBook.ID).Updates(&models.BookInventory{TotalCopies: oldBook.TotalCopies, AvailableCopies: oldBook.AvailableCopies})
		return errorss.BadRequest("Can not update the book", updatedResult.Error.Error())
	}
	book.LibID = libId
	if err := db.Create(&book).Error; err != nil {
		return errorss.BadRequest("failed to add book", err.Error())
	}

	return nil
}

func (bookService *BookService) GetAllBook(userId uint, libId uint) ([]map[string]interface{}, error) {
	// var books []models.BookInventory
	var books []map[string]interface{}
	db := bookService.db
	var user models.User

	if err := auth.IsLibraryAdmin(db, libId, userId); err != nil {
		return nil, err
	}
	if err := db.Where("id=? AND lib_id=?", userId, libId).First(&user).Error; err != nil {

		return nil, fmt.Errorf("can not find user")
	}

	if user.Role == "reader" {
		if err := db.Model(&models.BookInventory{}).Omit("total_copies", "available_copies", "lib_id", "CreatedAt", "DeletedAt", "UpdatedAt").
			Where("lib_id = ?", libId).Find(&books).Error; err != nil {
			return nil, err
		}
		return books, nil
	}

	bookResult := db.Model(&models.BookInventory{}).Where("lib_id = ?", libId).Find(&books)

	return books, bookResult.Error
}

func (bookService *BookService) GetBook(id uint) (map[string]interface{}, error) {

	db := bookService.db
	var response map[string]interface{}
	if err := db.Model(&models.BookInventory{}).Where("ID = ?", id).First(&response).Error; err != nil {
		return nil, fmt.Errorf("no book found")
	}

	return response, nil

}

func (bookService *BookService) UpdateBook(userId, bookId uint, updatedBook dto.RequestUpdateBook) error {
	DB := bookService.db

	var book models.BookInventory
	DB.Model(&models.BookInventory{}).Where("ID = ?", bookId).First(&book)

	if err := auth.IsLibraryAdmin(DB, book.LibID, userId); err != nil {
		fmt.Println(err)
		return err
	}

	if err := DB.Model(&models.BookInventory{}).Select("available_copies").Where("id = ?", bookId).Updates(&updatedBook); err.Error != nil {
		fmt.Println(err)
		return err.Error
	}

	return nil

}
func (bookService *BookService) DeleteBook(bookId, userId uint) error {
	DB := bookService.db
	if result := DB.Where("id =? ", bookId).First(&book); result.Error != nil {
		fmt.Println(result.Error)
		return fmt.Errorf("Book does not exist")
	}
	if err := auth.IsLibraryAdmin(DB, book.LibID, userId); err != nil {
		return err
	}
	if result := DB.Where("id=?", bookId).Select("TotalCopies", "AvailableCopies").Updates(&models.BookInventory{TotalCopies: 0, AvailableCopies: 0}); result.Error != nil {
		fmt.Println(result.Error)
		return result.Error
	}

	return nil

}

func (bookService *BookService) GetIssuedBook(readerId uint, libId uint) ([]map[string]interface{}, error) {
	DB := bookService.db
	var issuedBooks []map[string]interface{}
	// var allIssued []models.IssueRegistery
	DB.Model(&models.IssueRegistery{}).Select("b.*", "*").Joins("JOIN book_inventories b on b.isbn = issue_registeries.isbn").Where("reader_id = ? AND issue_registeries.issue_status=?", readerId, "issued").Find(&issuedBooks)
	if len(issuedBooks) < 1 {
		return nil, fmt.Errorf("no issue books found")
	}
	// return
	// var issuedBooks []models.BookInventory
	// DB.Where("lib_id =?", libId).Find(&issuedBooks)

	// var issueRequests []models.IssueRegistery

	// if err := DB.Where("reader_id=?", readerId).Find(&issueRequests).Error; err != nil {
	// 	return nil, err
	// }

	// var temp = make(map[string]string)

	// for _, item := range issuedBooks {
	// 	temp[item.ISBN] = item.Title
	// }
	// var response []dto.IssuedBooks
	// fmt.Println(issueRequests)
	// copier.Copy(&response, &issueRequests)
	// for _, item := range response {
	// 	item.Book.Name = temp[item.ISBN]
	// }
	// return response, nil
	return issuedBooks, nil
}

func (bookService *BookService) SearchBook(libId uint, payload *dto.SearchBookPayload) ([]models.BookInventory, error) {

	db := bookService.db
	var searchedBooks []models.BookInventory

	result := db.Where("lib_id=? AND (authors = ? OR title = ? OR publisher = ?)", libId, payload.Author, payload.Title, payload.Publisher).Find(&searchedBooks)
	if result.Error != nil {
		return nil, result.Error
	}
	return searchedBooks, nil
}
