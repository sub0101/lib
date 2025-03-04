package auth

import (
	"fmt"
	"libraryManagement/internal/models"

	"gorm.io/gorm"
)

func IsLibraryAdmin(DB *gorm.DB, libraryId, userId uint) error {
	authService := AuthService{DB: DB}
	return authService.validate(libraryId, userId)

}

func (authService *AuthService) validate(libraryId, userId uint) error {
	DB := authService.DB
	if err := DB.Where("id=? AND lib_id=?", userId, libraryId).First(&models.User{}); err.Error != nil {
		return fmt.Errorf("user do not belong to this library")
	}
	return nil

}
