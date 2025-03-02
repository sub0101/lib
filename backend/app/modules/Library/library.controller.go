package library

import (
	"fmt"
	"libraryManagement/internal/dto"
	"libraryManagement/utility"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LibraryController struct {
	libraryService LibraryService
}

func NewLibraryController(db *gorm.DB) *LibraryController {

	service := LibraryService{db: db}
	return &LibraryController{libraryService: service}
}

func (lc *LibraryController) GetAllLibrary(c *gin.Context) {

	libraries, err := lc.libraryService.GetAllLibrary()
	if err != nil {
		utility.SendResponse(c, 501, false, "Server Error", nil)
		return
	}

	utility.SendResponse(c, 200, true, "Successfully Fetched All libraries ", libraries)

}
func (lc *LibraryController) GetLibrary(c *gin.Context) {
	idParams, valid := c.Params.Get("id")
	userId := utility.GetContextItem(c, "id")
	if !valid {
		utility.SendResponse(c, 401, false, "Invalid Params", nil)

	}
	id, _ := strconv.Atoi(idParams)
	libraries, err := lc.libraryService.GetLibrary(userId, uint(id))
	if err != nil {
		utility.SendResponse(c, 501, false, "Server Error", nil)
		return
	}

	utility.SendResponse(c, 200, true, "Successfully Fetched  library ", libraries)
}

func (lc *LibraryController) AddLibrary(c *gin.Context) {
	err := lc.libraryService.AddLibrary()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Success": false, "message": err.Error()})
		return
	}
}

func (lc *LibraryController) DeleteLibrary(c *gin.Context) {

}

func (lc *LibraryController) UpdateLibrary(c *gin.Context) {

	var payload dto.ResponseGetLibrary
	userId := utility.GetContextItem(c, "id")
	libId, validate := utility.GetParamItem(c, "id")
	if err := c.ShouldBindJSON(&payload); err != nil || validate == false {
		utility.SendResponse(c, 400, false, "Invalid Payload", nil, err.Error())
		return
	}
	fmt.Println(userId, libId)
	err := lc.libraryService.UpdateLibrary(userId, libId, payload)
	if err != nil {
		utility.SendResponse(c, 400, false, "Unable to Update the Library", nil, err.Error())
		return
	}
	utility.SendResponse(c, 200, true, "Successfully Updated the Library", nil)

}
