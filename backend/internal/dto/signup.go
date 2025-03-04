package dto

import (
	"gorm.io/gorm"
)

type RequestSignupLibraryBody struct {
	Name          string `json:"name"  binding:"required" gorm:"size:100;not null" validate:"required,min=2,max=100,alpha_space"`
	Email         string `json:"email"  binding:"required" gorm:"size:100;unique;not null" validate:"email"`
	ContactNumber string `json:"contact"  binding:"required" gorm:"unique;size:15;not null" validate:"required,phone"`
	Role          string `json:"role" gorm:"size:15;not null"`
	LibId         uint   `json:"libraryId"  gorm:"not null"`
	LibraryName   string `json:"libraryName" binding:"required,min=2" gorm:"not null"`
	Password      string `json:"password" binding:"required,min=8" gorm:"not null" validate:"required,password"`
}

type RequestSignupUserBody struct {
	gorm.Model
	Name          string `json:"name" gorm:"size:100;not null" validate:"required,min=2,max=100,alpha_space"`
	Email         string `json:"email"   gorm:"size:100;unique;not null" validate:"required,email"`
	ContactNumber string `json:"contact"  binding:"required" gorm:"size:15;not null" validate:"phone"`
	Role          string `json:"role" gorm:"size:15;not null"`
	LibId         uint   `json:"libraryId"  binding:"required" gorm:"not null"` // Could be foreign key for the library

	Password string `json:"password" binding:"required,min=8" gorm:"not null" validate:"required,password"`
}

type RequestLoginBody struct {
	Email    string `json:"email"  binding:"required" gorm:"size:100;unique;not null" validate:"required,email"`
	Password string `json:"password" binding:"required,min=8" gorm:"not null" validate:"required,password"`
}

type JwtResponse struct {
	token string
}

type BaseResponse struct {
	Success bool        `json:"success" example:"true"`
	Message string      `json:"message" example:"Operation successful"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty" example:"Invalid credentials"`
}
