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

func (userService *UserService) UpdateUser(userId uint, ownerId uint, ownerLibId uint, role string) *errorss.AppError {
	db := userService.DB
	var reader = models.User{}
	var roles = []string{"admin", "readaer"}
	include := false
	for _, r := range roles {
		if r == role {
			include = true
		}
	}
	if !include {
		return errorss.BadRequest("invalid role", "invalid role")
	}
	if userId == ownerId {
		return errorss.Forbidden("Owner can not update his own role", "can not update his own role")
	}
	if err := db.Where("id=? AND lib_id=?", userId, ownerLibId).First(&reader); err.Error != nil {
		return errorss.NotFound("user not found", err.Error.Error())
	}
	db.Where("id=?", userId).Updates(&models.User{Role: role})
	return nil
}

func (userService *UserService) GetAllUser(libId, userId uint) ([]models.User, *errorss.AppError) {
	var users = []models.User{}
	db := userService.DB

	if result := db.Where("lib_id=? AND id != ?", libId, userId).Find(&users); result.Error != nil {
		return nil, errorss.NotFound("no user fouund", result.Error.Error())
	}
	return users, nil
}

func (UserService *UserService) GetUser(userId, readerId, libId uint) (*models.User, *errorss.AppError) {
	var user = models.User{}
	db := UserService.DB
	result := db.Where("id=?", userId).First(&user)
	if result.Error != nil {
		return nil, errorss.NotFound("No User Found", result.Error.Error())
	}
	if err := db.Where("id=?", readerId).First(&models.User{}).Error; err != nil {
		return nil, errorss.NotFound("No User Found", err.Error())
	}
	fmt.Println(readerId, libId)
	if user.Role == "admin" || user.Role == "owner" {
		var reader models.User
		if result := db.Model(&models.User{}).Where("id=? AND lib_id=?", readerId, libId).First(&reader); result.Error != nil {
			return nil, errorss.Forbidden("can not access this resource", result.Error.Error())
		}
		return &reader, nil
	}
	if userId == readerId {
		return &user, nil
	} else {
		return nil, errorss.Forbidden("can not access this resource", "can not access this resource")
	}

}

func (userService *UserService) DeleteUser(userId, ownerId, ownerLibId uint) *errorss.AppError {
	db := userService.DB
	var user = models.User{}
	if userId == ownerId {
		return errorss.BadRequest("can not delete  own account", "can not delete  own account")
	}
	if err := db.Where("id=?", userId).First(&user); err.Error != nil {
		return errorss.NotFound("no user found", "no user found")
	}
	result := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id=? AND lib_id=?", userId, ownerLibId).First(&user); err.Error != nil {
			return fmt.Errorf("can not delete user")
		}

		tx.Delete(&user)
		tx.Where("email=?", user.Email).Delete(&models.Auth{})

		return nil
	})
	if result != nil {
		fmt.Println("")
		return errorss.Forbidden("can not delete user", result.Error())
	}
	return nil
}
