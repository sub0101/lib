package auth

import (
	"fmt"
	"libraryManagement/internal/dto"
	"libraryManagement/internal/models"
	"libraryManagement/utility"

	"gorm.io/gorm"
)

type AuthService struct {
	DB *gorm.DB
}

func (service *AuthService) Login(body dto.RequestLoginBody) (uint, uint, string, error) {
	var auth = models.Auth{}
	var user = models.User{}
	DB := service.DB.Session(&gorm.Session{NewDB: true})
	email, password := body.Email, body.Password
	verify := DB.Where("email = ?", email).First(&auth)
	fmt.Println(auth)
	if verify.Error != nil {
		return 0, 0, "", fmt.Errorf("Invalid Email ID")
	}

	if utility.VerifyPassword(password, auth.Password) == false {

		return 0, 0, "", fmt.Errorf("Invalid email or  Password")
	}
	DB.Where("email = ?", email).First(&user)

	refreshToken, _ := utility.
		CreateJwtToken(utility.JwtPayload{LibId: user.LibId, Id: user.ID, Role: user.Role})

	DB.Model(&models.Auth{}).Where("email=?", email).Updates(&models.Auth{RefreshToken: refreshToken})
	return user.LibId, user.ID, user.Role, nil

}
func (service *AuthService) Signup(user models.User, auth models.Auth) error {

	DB := service.DB
	// if !ValidatePassword(auth.Password) {
	// 	return fmt.Errorf("Password is Not valid must contain Atleast 1 special character, 1 Capital Letter ,  min length 8")
	// }
	if err := DB.Where("id= ?", user.LibId).First(&models.Library{}).Error; err != nil {
		return fmt.Errorf("library does not exist")

	}
	if err := DB.Where("email = ?", auth.Email).First(&models.Auth{}).Error; err == nil {
		return fmt.Errorf("email is already exist")
	}

	err := DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&auth).Error; err != nil {
			return fmt.Errorf("not able to create account")
		}

		if err := tx.Create(&user).Error; err != nil {
			return fmt.Errorf("not able to create account")

		}
		return nil
	})
	fmt.Println(err)
	return err

}

func (service *AuthService) SignupLibrary(user models.User, auth models.Auth, library models.Library) error {

	DB := service.DB
	user.Role = "owner"

	err := DB.Transaction(func(tx *gorm.DB) error {
		tx.Create(&library)
		tx.Create(&auth)
		user.LibId = library.ID
		err := tx.Create(&user)
		fmt.Println(err.Error)

		return nil
	})
	fmt.Println(err)
	return err

}
