package auth

import (
	"fmt"
	"libraryManagement/internal/dto"
	"libraryManagement/internal/models"
	"libraryManagement/utility"
	"libraryManagement/validators"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
	authService AuthService
}

func NewAuthController(db *gorm.DB) *AuthController {

	authService := AuthService{DB: db}
	return &AuthController{authService: authService}
}

func (ac *AuthController) Login(c *gin.Context) {
	var body dto.RequestLoginBody
	if err := c.ShouldBindBodyWithJSON(&body); err != nil {
		utility.SendResponse(c, 401, false, "Invalid Input ", nil, err.Error())
		return

	}

	libId, id, role, err := ac.authService.Login(body)
	if err != nil {
		utility.SendResponse(c, 401, false, "Invalid Input ", nil, err.Error())
		return
	}
	token, err := utility.CreateJwtToken(utility.JwtPayload{LibId: libId, Id: id, Role: role})
	c.SetCookie("jwt", token, 3600, "/", "localhost", false, true)
	utility.SendResponse(c, 200, true, "Successfully Logedin", token)
}

func (ac *AuthController) SignupLibrary(c *gin.Context) {

	var body dto.RequestSignupLibraryBody
	err := c.ShouldBindJSON(&body)

	if err != nil {
		fmt.Println("err", err.Error())
		utility.SendResponse(c, 401, false, "Invalid Input Body", nil, err.Error())
		return
	}
	hashPassword, err := utility.HashPassword(body.Password)
	if err != nil {
		utility.SendResponse(c, 501, false, "Server Side Error", nil, err.Error())

	}
	auth := models.Auth{Email: body.Email, Password: hashPassword}
	library := models.Library{Name: body.LibraryName}
	user := models.User{Name: body.Name, Email: body.Email, ContactNumber: body.ContactNumber}
	fmt.Println("creating library")
	err = ac.authService.SignupLibrary(user, auth, library)
	if err != nil {
		utility.SendResponse(c, 501, false, "Server Side Error", nil, err.Error())
		return
	}
	utility.SendResponse(c, 201, true, "Account created Successfully", nil)

}

func (ac *AuthController) Signup(c *gin.Context) {
	var body dto.RequestSignupUserBody
	err := c.ShouldBindBodyWithJSON(&body)
	if err != nil {

		utility.SendResponse(c, 400, false, "Invalid Input Body", nil, err.Error())
		return
	}
	if err := validators.Validate.Struct(body); err != nil {
		utility.SendResponse(c, 400, false, "Validations Error", nil, err.Error())

		// c.JSON(http.StatusBadRequest, gin.H{"validation_error": err.Error()})
		return
	}

	hashPassword, err := utility.HashPassword(body.Password)
	if err != nil {
		utility.SendResponse(c, 501, false, "Server Side Error", nil, err.Error())

	}

	auth := models.Auth{Email: body.Email, Password: hashPassword}

	user := models.User{Name: body.Name, Email: body.Email, ContactNumber: body.ContactNumber, LibId: body.LibId, Role: "reader"}

	err = ac.authService.Signup(user, auth)
	if err != nil {
		utility.SendResponse(c, 501, false, "Server Side Error", nil, err.Error())
		return
	}
	utility.SendResponse(c, 201, true, "Account Created Successfully", nil)
}
