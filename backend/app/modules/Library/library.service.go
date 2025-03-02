package library

import (
	"fmt"
	"libraryManagement/internal/dto"
	"libraryManagement/internal/models"

	"gorm.io/gorm"
)

type LibraryService struct {
	db *gorm.DB
}

func (service *LibraryService) GetAllLibrary() ([]dto.ResponseGetLibrary, error) {
	db := service.db
	var libraries []dto.ResponseGetLibrary
	libObject := db.Model(&[]models.Library{}).Select("name", "id").Find(&libraries)
	if err := libObject.Error; err != nil {
		return nil, err
	}

	return libraries, nil

}

func (service *LibraryService) AddLibrary() error {

	return nil

}

func (service *LibraryService) UpdateLibrary(userId, libId uint, payload dto.ResponseGetLibrary) error {
	db := service.db
	var user models.User
	if err := db.Where("lib_id=? AND id=?", libId, userId).First(&user).Error; err != nil {
		return fmt.Errorf("You are not Belongs to this library")
	}
	fmt.Println(user)
	if err := db.Where("id = ?", libId).Updates(&models.Library{Name: payload.Name}).Error; err != nil {
		return err
	}
	return nil

}
func (service *LibraryService) DeleteLibrary() {
	db := service.db
	_ = db

}
func (service *LibraryService) GetLibrary(userId uint, id uint) (*dto.ResponseGetLibrary, error) {
	db := service.db

	var library dto.ResponseGetLibrary

	libObject := db.Where("ID = ? ", id).First(&library)
	libObject = db.Model(&models.Library{}).Where("ID = ? ", id).First(&library)
	if err := libObject.Error; err != nil {
		return nil, err
	}

	return &library, nil

}
