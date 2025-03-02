package user

import (
	"fmt"
	"libraryManagement/utility"
	"strconv"

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
	readerId, valid := getParamItem(c, "id")
	if !valid {
		utility.SendResponse(c, 400, false, "Invalid Input", nil)
		return
	}
	libId := getContextItem(c, "libId")
	userId := getContextItem(c, "id")
	user, err := uc.userService.GetUser(userId, readerId, libId)
	if err != nil {
		utility.SendResponse(c, 500, false, err.Message, nil, err.Details)
		return
	}
	utility.SendResponse(c, 200, true, "Successfully fetched user", user)
}
func (uc *UserController) GetAllUser(c *gin.Context) {
	libId := getContextItem(c, "libId")
	userId := getContextItem(c, "id")

	users, err := uc.userService.GetAllUser(libId, userId)
	if err != nil {
		utility.SendResponse(c, 500, true, "failed to fetch all users", nil, err.Error())
		return

	}
	utility.SendResponse(c, 200, true, "Successfully fetched all users", users)
}
func (uc *UserController) MakeAdmin(c *gin.Context) {
	var requestBody RequestMakeAdmin
	ownerId := getContextItem(c, "id")
	ownerLibId := getContextItem(c, "libId")

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		fmt.Println(requestBody)
		utility.SendResponse(c, 400, false, "invalid body", nil)
		return
	}

	err := uc.userService.UpdateUser(requestBody.Id, ownerId, ownerLibId, requestBody.Role)
	if err != nil {
		utility.SendResponse(c, 500, false, "failed to update the user", nil, err.Error())
		return
	}
	utility.SendResponse(c, 201, true, "Successfully Change Role of the user", nil)
}

func getQueryItem(c *gin.Context, id string) uint {
	idQuery := c.Query(id)
	libId, _ := strconv.Atoi(idQuery)
	return uint(libId)
}

func getContextItem(c *gin.Context, id string) uint {
	userId, _ := c.Get(id)
	fmt.Println(userId, "hai")
	result := uint(userId.(float64))
	return result
}
func getParamItem(c *gin.Context, id string) (uint, bool) {
	idParams, exist := c.Params.Get(id)
	result, _ := strconv.Atoi(idParams)
	return uint(result), exist

}
