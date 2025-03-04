package book

import (
	"fmt"
	auth "libraryManagement/app/modules/auth"
	"libraryManagement/internal/dto"
	"libraryManagement/internal/models"
	"libraryManagement/utility/errorss"
	"log"

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

	// if err := db.Where("ID = ? AND lib_id = ?", userId, libId).First(&user).Error; err != nil {
	// 	return errorss.Forbidden("Can not Addd Book", "you are not authorized to access this resource")
	// }

	if book.AvailableCopies != book.TotalCopies {
		return errorss.BadRequest("Can not add Book", "Total copies are not equal to available copies")
	}

	var oldBook models.BookInventory
	isExist := db.Where("isbn=? AND lib_id=?", book.ISBN, libId).First(&oldBook)
	fmt.Print(oldBook)
	if err := isExist.Error; err == nil {
		oldBook.TotalCopies += book.TotalCopies
		oldBook.AvailableCopies += book.AvailableCopies
		updatedResult := db.Model(&models.BookInventory{}).Where("ID = ?", oldBook.ID).Updates(&models.BookInventory{TotalCopies: oldBook.TotalCopies, AvailableCopies: oldBook.AvailableCopies})
		// return errorss.BadRequest("Can not update the book", updatedResult.Error.Error())
		if updatedResult.Error != nil {
			fmt.Println(updatedResult.Error)
			return errorss.BadRequest("can not update the book", updatedResult.Error.Error())
		}
		return nil
	}
	book.LibID = libId
	if err := db.Create(&book).Error; err != nil {
		return errorss.BadRequest("failed to add book", err.Error())
	}

	return nil
}

func (bookService *BookService) GetAllBook(userId uint, libId uint) ([]map[string]interface{}, *errorss.AppError) {
	// var books []models.BookInventory
	var books []map[string]interface{}
	db := bookService.db
	var user models.User
	if err := db.Where("ID = ?", userId).First(&user).Error; err != nil {
		return nil, errorss.NotFound("User Not Found", err.Error())
	}
	if user.Role == "reader" {
		if err := db.Model(&models.BookInventory{}).
			Omit("total_copies", "available_copies", "lib_id", "CreatedAt", "DeletedAt", "UpdatedAt").
			Where("lib_id = ?", libId).Find(&books).Error; err != nil {
			return nil, errorss.NotFound("No Books Found", err.Error())
		}
		return books, nil
	}

	bookResult := db.Model(&models.BookInventory{}).Where("lib_id = ?", libId).Find(&books)
	if bookResult.Error != nil {
		fmt.Println(bookResult.Error)
		return nil, errorss.NotFound("No Books Found", bookResult.Error.Error())
	}
	return books, nil
}

func (bookService *BookService) GetBook(id uint) (map[string]interface{}, *errorss.AppError) {

	db := bookService.db
	var response map[string]interface{}
	if err := db.Model(&models.BookInventory{}).Where("ID = ? ", id).First(&response).Error; err != nil {
		return nil, errorss.NotFound("Book Not Found", "book not found")
	}

	return response, nil

}

func (bookService *BookService) UpdateBook(userId, bookId uint, updatedBook dto.RequestUpdateBook) *errorss.AppError {
	DB := bookService.db

	var book models.BookInventory
	if err := DB.Model(&models.BookInventory{}).Where("ID = ?", bookId).First(&book).Error; err != nil {
		return errorss.NotFound("book not Found", "book not found")
	}

	if err := auth.IsLibraryAdmin(DB, book.LibID, userId); err != nil {
		fmt.Println(err)
		return errorss.Forbidden("Can not Update Book", "you are not authorized to access this resource")
	}

	if err := DB.Model(&models.BookInventory{}).Select("available_copies").Where("id = ?", bookId).
		Updates(&updatedBook); err.Error != nil {

		return errorss.BadRequest("Can not update the book", "failed to update the book")
	}

	return nil

}
func (bookService *BookService) DeleteBook(bookId, userId uint) *errorss.AppError {
	DB := bookService.db
	if result := DB.Where("id =? ", bookId).First(&book); result.Error != nil {
		fmt.Println(result.Error)
		return errorss.NotFound("Book Not Found", "book not found")
	}
	if err := auth.IsLibraryAdmin(DB, book.LibID, userId); err != nil {
		return errorss.Forbidden("Can not Delete Book", "you are not authorized to access this resource")
	}

	currentAvailable := book.AvailableCopies
	currentTotal := book.TotalCopies
	if book.AvailableCopies > 0 {
		currentAvailable -= 1
		currentTotal -= 1
	} else if book.AvailableCopies == 0 && book.TotalCopies > 0 {
		currentTotal -= 1
	} else if book.TotalCopies == 0 {
		return errorss.BadRequest("Can not delete book", "no copies available")
	}

	if result := DB.Where("id=?", bookId).Select("TotalCopies", "AvailableCopies").Updates(&models.BookInventory{TotalCopies: currentTotal, AvailableCopies: currentAvailable}); result.Error != nil {
		fmt.Println(result.Error)
		return errorss.BadRequest("Can not delete book", "failed to delete book")
	}

	log.Printf("available copies %v", currentAvailable)
	log.Printf("total copies %v", currentTotal)

	return nil

}

func (bookService *BookService) GetIssuedBook(readerId uint, libId uint) ([]map[string]interface{}, *errorss.AppError) {
	DB := bookService.db
	var issuedBooks []map[string]interface{}

	log.Printf("reader id %v", readerId)
	log.Printf("lib id %v", libId)
	result := DB.Model(&models.IssueRegistery{}).Select("b.*", "*").
		Joins("JOIN book_inventories b on b.isbn = issue_registeries.isbn").
		Where("reader_id = ? AND issue_registeries.issue_status=?", readerId, "issued").
		Find(&issuedBooks)
	if result.Error != nil {
		return nil, errorss.InternalServerError("No Books Found", "no books found")
	}

	return issuedBooks, nil
}

func (bookService *BookService) SearchBook(libId uint, payload *dto.SearchBookPayload) ([]models.BookInventory, *errorss.AppError) {

	db := bookService.db
	var searchedBooks []models.BookInventory
	fmt.Println(libId)

	result := db.Where("lib_id=? AND (authors = ? OR title = ? OR publisher = ? OR isbn=?)", libId, payload.Author, payload.Title, payload.Publisher, payload.ISBN).Find(&searchedBooks)
	if result.Error != nil {

		return nil, errorss.InternalServerError("No Books Found", result.Error.Error())
	}
	if len(searchedBooks) == 0 {

		return nil, errorss.NotFound("No Books Found", "no books found")

	}
	return searchedBooks, nil
}
