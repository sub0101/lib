package user

import (
	"fmt"
	"libraryManagement/internal/models"
	"libraryManagement/utility/errorss"

	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func (userService *UserService) UpdateUser(userId uint, ownerId uint, ownerLibId uint, role string) error {
	db := userService.DB
	var reader = models.User{}
	if userId == ownerId {
		return fmt.Errorf("Owner can not change his own role")
	}
	if err := db.Where("id=? AND lib_id=?", userId, ownerLibId).First(&reader); err.Error != nil {
		return fmt.Errorf("Can not find user with given data")
	}
	db.Where("id=?", userId).Updates(&models.User{Role: role})
	return nil
}

func (userService *UserService) GetAllUser(libId, userId uint) ([]models.User, *errorss.AppError) {
	var users = []models.User{}
	db := userService.DB

	if result := db.Where("lib_id=? AND id != ?", libId, userId).Find(&users); result.Error != nil {
		return nil, errorss.InternalServerError("No User Fouund", result.Error.Error())
	}
	return users, nil
}

func (UserService *UserService) GetUser(userId, readerId, libId uint) (*models.User, *errorss.AppError) {
	var user = models.User{}
	db := UserService.DB
	result := db.Where("id=?", userId).First(&user)
	if result.Error != nil {
		return nil, errorss.InternalServerError("No User Found", result.Error.Error())
	}
	fmt.Println(readerId, libId)
	if user.Role == "admin" || user.Role == "owner" {
		var reader models.User
		if result := db.Model(&models.User{}).Where("id=? AND lib_id=?", readerId, libId).First(&reader); result.Error != nil {
			return nil, errorss.InternalServerError("No User Found", result.Error.Error())
		}
		return &reader, nil
	}
	if userId == readerId {
		return &user, nil
	} else {
		return nil, errorss.InternalServerError("You are not authorized to view this user", "You are not authorized to view this user")
	}

}
