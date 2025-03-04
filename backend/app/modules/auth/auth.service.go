package auth

import (
	"fmt"
	"libraryManagement/internal/dto"
	"libraryManagement/internal/models"
	"libraryManagement/utility"
	"libraryManagement/utility/errorss"

	"gorm.io/gorm"
)

type AuthService struct {
	DB *gorm.DB
}

func (service *AuthService) Login(body dto.RequestLoginBody) (uint, uint, string, *errorss.AppError) {
	var auth = models.Auth{}
	var user = models.User{}
	DB := service.DB.Session(&gorm.Session{NewDB: true})
	email, password := body.Email, body.Password
	verify := DB.Where("email = ?", email).First(&auth)
	fmt.Println(auth)
	if verify.Error != nil {
		return 0, 0, "", errorss.BadRequest("invalid email or password", "invalid email or password")

	}

	if utility.VerifyPassword(password, auth.Password) == false {

		return 0, 0, "", errorss.BadRequest("invalid email or password", "invalid email or password")
	}
	DB.Where("email = ?", email).First(&user)

	// DB.Model(&models.Auth{}).Where("email=?", email).Updates(&models.Auth{RefreshToken: refreshToken})
	return user.LibId, user.ID, user.Role, nil

}
func (service *AuthService) Signup(user models.User, auth models.Auth) *errorss.AppError {

	DB := service.DB
	// if !ValidatePassword(auth.Password) {
	// 	return fmt.Errorf("Password is Not valid must contain Atleast 1 special character, 1 Capital Letter ,  min length 8")
	// }
	if err := DB.Where("id= ?", user.LibId).First(&models.Library{}).Error; err != nil {
		return errorss.BadRequest("library does not exist", "library does not exist")

	}
	if err := DB.Where("email = ?", auth.Email).First(&models.Auth{}).Error; err == nil {
		return errorss.BadRequest("user already exist", "use already exist")
	}
	if errr := DB.Where("contact_number=?", user.ContactNumber).First(&models.User{}).Error; errr == nil {
		return errorss.BadRequest("user already exist", "user already exist")
	}

	err := DB.Transaction(func(tx *gorm.DB) error {

		if err := tx.Create(&user).Error; err != nil {
			return fmt.Errorf("not able to create account")
		}

		if err := tx.Create(&auth).Error; err != nil {
			return fmt.Errorf("not able to create account")
		}
		return nil
	})
	fmt.Println(err)
	if err != nil {
		return errorss.BadRequest("not able to create account", err.Error())

	}
	return nil

}

func (service *AuthService) SignupLibrary(user models.User, auth models.Auth, library models.Library) *errorss.AppError {

	DB := service.DB
	user.Role = "owner"
	DB.Where("email=?", auth.Email).Find(&models.Auth{})
	if err := DB.Where("email = ? OR contact_number=?", user.Email, user.ContactNumber).First(&models.User{}).Error; err == nil {
		return errorss.BadRequest("user already Exist", "user already exist")
	}

	DB.Where("name=?", library.Name).Find(&models.Library{})
	if err := DB.Where("name = ?", library.Name).First(&models.Library{}).Error; err == nil {
		return errorss.BadRequest("library already exist", "library already exist")
	}

	err := DB.Transaction(func(tx *gorm.DB) error {
		tx.Create(&library)

		user.LibId = library.ID
		if err := tx.Create(&user).Error; err != nil {
			return fmt.Errorf(err.Error())
		}

		if err := tx.Create(&auth).Error; err != nil {
			return fmt.Errorf(err.Error())
		}

		return nil
	})
	if err != nil {
		return errorss.BadRequest("not able to create account", err.Error())
	}
	return nil

}
