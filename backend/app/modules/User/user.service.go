package user

import (
	"fmt"
	"libraryManagement/internal/models"

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

func (userService *UserService) GetAllUser(libId, userId uint) ([]models.User, error) {
	var users = []models.User{}
	db := userService.DB
	if result := db.Preload("Library").Where("lib_id=? AND id != ?", libId, userId).Find(&users); result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
