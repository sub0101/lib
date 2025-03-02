package book

import (
	"libraryManagement/internal/dto"
	"libraryManagement/internal/models"
	"libraryManagement/utility"
	"libraryManagement/validators"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookController struct {
	bookService BookService
}

func NewBookController(db *gorm.DB) *BookController {
	bookService := BookService{db: db}
	return &BookController{bookService: bookService}
}

func (bc *BookController) AddBook(c *gin.Context) {
	var book models.BookInventory

	id := utility.GetContextItem(c, "id")
	libId := utility.GetContextItem(c, "libId")

	if err := c.ShouldBindJSON(&book); err != nil {
		utility.SendResponse(c, http.StatusBadRequest, false, "Inavalid request", nil, err.Error())
		return
	}
	if err := validators.Validate.Struct(book); err != nil {
		utility.SendResponse(c, 400, false, "Validations Error", nil, err.Error())
		return
	}

	err := bc.bookService.AddBook(id, libId, book)
	if err != nil {
		utility.SendResponse(c, err.Code, false, err.Message, nil, err.Details)
		return
	}
	utility.SendResponse(c, 201, true, "Successfully Added Book", nil)

}
func (bookController *BookController) GetBook(c *gin.Context) {

	bookId, valid := getParamItem(c, "id")
	userLib := utility.GetContextItem(c, "libId")
	if !valid {
		utility.SendResponse(c, 400, false, "invalid params", nil)
		return

	}

	book, err := bookController.bookService.GetBook(bookId)
	if err != nil {
		utility.SendResponse(c, 404, false, "Book Not Found", nil, err.Error())
		return
	}

	libId := book["lib_id"].(uint)

	if libId != userLib {
		utility.SendResponse(c, 403, false, "Can not access book", nil, "You are not authorized to see this book")
		return
	}
	utility.SendResponse(c, 202, true, "Successfully get Book", book)

}

func (bc *BookController) GetAllBook(c *gin.Context) {

	libId := getContextItem(c, "libId")
	userId := getContextItem(c, "id")

	books, err := bc.bookService.GetAllBook(userId, libId)
	if err != nil {
		utility.SendResponse(c, 400, false, "Bad Request", nil, err.Error())
		return
	}
	utility.SendResponse(c, http.StatusOK, true, "All Books Fetched", books)

}

func (bookController *BookController) DeleteBook(c *gin.Context) {
	bookId, _ := getParamItem(c, "id")
	userId := getContextItem(c, "id")
	if err := bookController.bookService.DeleteBook(bookId, userId); err != nil {
		utility.SendResponse(c, 400, false, "Not able to Delete", nil, err.Error())
		return
	}
	utility.SendResponse(c, 200, true, "Successfully Deleted the  Book", nil)
}

func (bookController *BookController) UpdateBook(c *gin.Context) {
	var updatedBook dto.RequestUpdateBook

	libId := utility.GetContextItem(c, "libId")

	if err := c.ShouldBindBodyWithJSON(&updatedBook); err != nil {
		utility.SendResponse(c, 400, false, "invalid requrest body", nil)
		return
	}
	bookId, exist := getParamItem(c, "id")

	// if err := validators.Validate.Struct(updatedBook); err != nil {
	// 	utility.SendResponse(c, 400, false, "Validations Error", nil, err.Error())
	// 	return
	// }
	if !validators.IsValidateBook(updatedBook) {
		utility.SendResponse(c, 400, false, "Validations Error", nil)
		return

	}
	if !exist {
		utility.SendResponse(c, 400, false, "invalid requrest params", nil)
		return
	}
	userId := getContextItem(c, "id")
	book, err := bookController.bookService.GetBook(bookId)

	if err != nil {
		utility.SendResponse(c, 404, false, "Book Not Found", nil, err.Error())
		return
	}
	if book["lib_id"].(uint) != libId {
		utility.SendResponse(c, 404, false, "Can not Upadte the book", nil)
		return
	}

	if err := bookController.bookService.UpdateBook(userId, bookId, updatedBook); err != nil {
		utility.SendResponse(c, 400, false, "not able to update", nil, err.Error())
		return
	}
	utility.SendResponse(c, 201, true, "Updated Succesfully ", nil)
}

func (bookController *BookController) GetIssuedBooks(c *gin.Context) {

	id := getContextItem(c, "id")
	libId := getContextItem(c, "libId")
	res, err := bookController.bookService.GetIssuedBook(id, libId)
	if err != nil {
		utility.SendResponse(c, 400, false, "failed to fetched all books", nil, err.Error())

	}
	utility.SendResponse(c, 200, true, "fetched all books", res)
}

func (bc *BookController) SearchBooks(c *gin.Context) {
	var payload *dto.SearchBookPayload
	libId := utility.GetContextItem(c, "libId")
	if err := c.ShouldBindJSON(&payload); err != nil {
		utility.SendResponse(c, 400, false, "Error", nil, err.Error())
		return
	}
	res, err := bc.bookService.SearchBook(libId, payload)
	if err != nil {
		utility.SendResponse(c, 400, false, "failed to fetch books", nil, err.Error())
		return
	}
	utility.SendResponse(c, 200, true, "Successfully fetch all books", res)
}

func getQueryItem(c *gin.Context, id string) uint {
	idQuery := c.Query(id)
	libId, _ := strconv.Atoi(idQuery)
	return uint(libId)
}

func getContextItem(c *gin.Context, id string) uint {
	userId, _ := c.Get(id)
	result := uint(userId.(float64))
	return result
}
func getParamItem(c *gin.Context, id string) (uint, bool) {
	idParams, exist := c.Params.Get(id)
	result, _ := strconv.Atoi(idParams)
	return uint(result), exist

}
