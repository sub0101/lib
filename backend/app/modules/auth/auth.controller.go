package auth

import (
	"fmt"
	_ "libraryManagement/docs"
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
		utility.SendResponse(c, 400, false, "invalid body", nil, err.Error())
		return

	}
	if err := validators.Validate.Struct(body); err != nil {
		utility.SendResponse(c, 400, false, "invalid body", nil, err.Error())

		return
	}
	libId, id, role, err := ac.authService.Login(body)
	if err != nil {
		// utility.SendResponse(c, 401, false, "unauthorized", nil, err.Error())
		utility.SendResponse(c, err.Code, false, err.Message, nil, err.Details)
		return
	}
	token, _ := utility.CreateJwtToken(utility.JwtPayload{LibId: libId, Id: id, Role: role})

	c.SetCookie("jwt", token, 3600, "/", "localhost", false, true)
	utility.SendResponse(c, 200, true, "successfully logged in", token)
}

func (ac *AuthController) SignupLibrary(c *gin.Context) {

	var body dto.RequestSignupLibraryBody

	if err := c.ShouldBindJSON(&body); err != nil {
		fmt.Println("err", err.Error())

		utility.SendResponse(c, 400, false, "invalid input body", nil, err.Error())
		return
	}
	hashPassword, err := utility.HashPassword(body.Password)
	if err != nil {
		utility.SendResponse(c, 501, false, "server side error", nil, err.Error())

	}
	if err := validators.Validate.Struct(body); err != nil {
		utility.SendResponse(c, 400, false, "invalid  body", nil, err.Error())
		return
	}
	auth := models.Auth{Email: body.Email, Password: hashPassword}
	library := models.Library{Name: body.LibraryName}
	user := models.User{Name: body.Name, Email: body.Email, ContactNumber: body.ContactNumber}

	errr := ac.authService.SignupLibrary(user, auth, library)
	if errr != nil {

		utility.SendResponse(c, errr.Code, false, errr.Message, nil, errr.Details)

		return
	}
	utility.SendResponse(c, 201, true, "Account created Successfully", nil)

}

func (ac *AuthController) Signup(c *gin.Context) {
	var body dto.RequestSignupUserBody
	err := c.ShouldBindBodyWithJSON(&body)
	if err != nil {

		utility.SendResponse(c, 400, false, "invalid  body", nil, err.Error())
		return
	}
	if err := validators.Validate.Struct(body); err != nil {
		utility.SendResponse(c, 400, false, "invalid  body", nil, err.Error())

		// c.JSON(http.StatusBadRequest, gin.H{"validation_error": err.Error()})
		return
	}

	hashPassword, err := utility.HashPassword(body.Password)
	if err != nil {
		utility.SendResponse(c, 501, false, "Server Side Error", nil, err.Error())

	}
	user := models.User{Name: body.Name, Email: body.Email, ContactNumber: body.ContactNumber, LibId: body.LibId, Role: "reader"}

	auth := models.Auth{Email: body.Email, Password: hashPassword}

	errr := ac.authService.Signup(user, auth)
	if errr != nil {
		utility.SendResponse(c, errr.Code, false, errr.Message, nil, errr.Details)
		return
	}
	utility.SendResponse(c, 201, true, "Account Created Successfully", nil)
}
