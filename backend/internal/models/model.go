package models

import (
	"time"

	"gorm.io/gorm"
)

// Users model represents a user in the library system.

type Auth struct {
	gorm.Model
	Email        string `json:"email" binding:"required" gorm:"unique;not null"`
	Password     string `json:"password" binding:"require" gorm:"not null"`
	RefreshToken string `json:"refreshToken"`
}
type User struct {
	gorm.Model
	Name          string `json:"name"  binding:"required" gorm:"size:100;not null" `
	Email         string `json:"email"  binding:"required" gorm:"size:100;unique;not null"`
	ContactNumber string `json:"contact"  binding:"required" gorm:"unique;size:15;not null"`
	Role          string `json:"role" gorm:"size:15;not null"`
	LibId         uint   `json:"libraryId"  gorm:"not null"`
}

type Library struct {
	gorm.Model
	Name  string          `json:"name" binding:"required" gorm:"unique;size:100;not null"`
	Books []BookInventory `gorm:"foreignKey:LibID;references:ID"`
	Users []User          `gorm:"foreignKey:LibId;references:ID;OnDelete:CASCADE"`
}

type BookInventory struct {
	gorm.Model
	ISBN            string `json:"isbn" binding:"required" gorm:"unique;not null" validate:"isbn"`
	LibID           uint   `json:"libraryId"  gorm:"not null"`
	Title           string `json:"title" binding:"required" gorm:"size:255;not null" validate:"alpha_space"`
	Authors         string `json:"authors" binding:"required" gorm:"size:255" validate:"alpha_space"`
	Publisher       string `json:"publisher" binding:"required" gorm:"size:255" validate:"alpha_space"`
	Version         string `json:"version" binding:"required" gorm:"size:50"`
	TotalCopies     int    `json:"totalCopies" binding:"required" gorm:"default:0" validate:"required,min=1"`
	AvailableCopies int    `json:"availableCopies" binding:"required" gorm:"default:0" validate:"required,min=1"`

	Issues   []IssueRegistery `gorm:"foreignKey:ISBN;references:ISBN;OnDelete:CASCADE"`
	Requests []RequestEvent   `gorm:"foreignKey:BookID;references:ID;OnDelete:CASCADE"`
}
type RequestEvent struct {
	ReqID         uint          ` json:"requestId" gorm:"primaryKey;autoIncrement"`
	BookID        uint          `json:"bookId" binding:"required" gorm:"not null"`
	ReaderID      uint          `json:"readerId" gorm:"not null" `
	RequestDate   *time.Time    `json:"requestDate" gorm:"default:current_timestamp"`
	ApprovalDate  *time.Time    `json:"approvalDate"`
	ApproverID    uint          `json:"approverId"`
	RequestType   string        `json:"requestType" binding:"required"`
	RequestStatus string        `json:"requestStatus" gorm:"default:'pending'"`
	Book          BookInventory `gorm:"foreignKey:BookID"`
}

type IssueRegistery struct {
	IssueID            uint       `json:"issueId" gorm:"primaryKey;autoIncrement"`
	ISBN               string     `json:"isbn" gorm:"	not null"`
	ReaderID           uint       `json:"readerId" gorm:"size:50;not null"`
	IssueApproverID    uint       `json:"issueApproverId"`
	IssueStatus        string     `json:"issueStatus" gorm:"size:50;not null"`
	IssueDate          *time.Time `json:"issueDate"`
	ExpectedReturnDate *time.Time `json:"expectedReturnDate"`
	ReturnDate         *time.Time `json:"returnDate"`
	ReturnApproverID   uint       `json:"returnApproverId"`
	CreatedAt          *time.Time `gorm:"autoCreatedTime"`
	UpdatedAt          *time.Time `gorm:"autoUpdatedTime"`
}
