package user

import (
	"fmt"
	"libraryManagement/utility"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	userService UserService
}
type RequestMakeAdmin struct {
	Id   uint   `json:"id" binding:"required"`
	Role string `json:"role" binding:"required"`
}

func NewUserController(db *gorm.DB) *UserController {

	userService := UserService{DB: db}
	return &UserController{userService: userService}
}

func (uc *UserController) UpdateUser(c *gin.Context) {

}

func (uc *UserController) GetUser(c *gin.Context) {
	readerId, valid := utility.GetParamItem(c, "id")
	if !valid {
		utility.SendResponse(c, 400, false, "invalid params", nil)
		return
	}
	libId := utility.GetContextItem(c, "libId")
	userId := utility.GetContextItem(c, "id")
	user, err := uc.userService.GetUser(userId, readerId, libId)
	if err != nil {
		// utility.SendResponse(c, 500, false, err.Message, nil, err.Details)
		utility.SendResponse(c, err.Code, false, err.Message, nil, err.Details)
		return
	}
	utility.SendResponse(c, 202, true, "successfully fetched user", user)
}
func (uc *UserController) GetAllUser(c *gin.Context) {
	libId := utility.GetContextItem(c, "libId")
	userId := utility.GetContextItem(c, "id")

	users, err := uc.userService.GetAllUser(libId, userId)
	if err != nil {
		// utility.SendResponse(c, 500, true, "failed to fetch all users", nil, err.Error())
		utility.SendResponse(c, err.Code, false, err.Message, nil, err.Details)
		return

	}
	utility.SendResponse(c, 202, true, "successfully fetched all users", users)
}
func (uc *UserController) MakeAdmin(c *gin.Context) {
	var requestBody RequestMakeAdmin
	ownerId := utility.GetContextItem(c, "id")
	ownerLibId := utility.GetContextItem(c, "libId")
	log.Printf("owner id %v", ownerId)
	log.Printf("owner lib id %v", ownerLibId)
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		fmt.Println(requestBody)
		utility.SendResponse(c, 400, false, "invalid body", nil)
		return
	}

	err := uc.userService.UpdateUser(requestBody.Id, ownerId, ownerLibId, requestBody.Role)
	if err != nil {
		// utility.SendResponse(c, 500, false, "failed to update the user", nil, err.Error())
		utility.SendResponse(c, err.Code, false, err.Message, nil, err.Details)
		return
	}
	utility.SendResponse(c, 202, true, "Successfully Change Role of the user", nil)
}

func (uc *UserController) DeleteUser(c *gin.Context) {

	userId := utility.GetQueryItem(c, "id")
	log.Printf("userIdd %v", userId)
	ownerId := utility.GetContextItem(c, "id")
	ownerLibId := utility.GetContextItem(c, "libId")
	err := uc.userService.DeleteUser(userId, ownerId, ownerLibId)
	if err != nil {
		utility.SendResponse(c, err.Code, false, err.Message, nil, err.Details)
		return
	}
	utility.SendResponse(c, 202, true, "successfully deleted user", nil)
}
