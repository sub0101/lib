package tests

import (
	"fmt"
	"libraryManagement/internal/models"
	"libraryManagement/utility"
	"log"
)

var tempLibId uint

func CreateTempReader() (string, uint) {

	db := GetTestDB()

	reader := models.User{
		Email:         "tempuser11@gmail.com",
		Name:          "tempowner1",
		Role:          "",
		ContactNumber: "1234567888",
		LibId:         tempLibId,
	}

	auth := models.Auth{
		Email:    "tempuser11@gmail.com",
		Password: "Tvsbk0101@",
	}

	db.Create(&reader)
	db.Create(&auth)
	tempReaderToken, _ := utility.CreateJwtToken(utility.JwtPayload{Id: reader.ID, Role: reader.Role, LibId: reader.LibId})
	log.Printf("user %v", reader)
	return tempReaderToken, reader.ID
}

func CreateTempOwner() (string, uint) {
	db := GetTestDB()
	library := models.Library{
		Name: "tempLibrary",
	}
	db.Create(&library)

	owner := models.User{
		Email:         "tempowner1@gmail.com",
		Name:          "tempowner1",
		Role:          "owner",
		ContactNumber: "1234567877",
	}
	owner.LibId = library.ID
	tempLibId = library.ID

	auth := models.Auth{
		Email:    "tempowner1@gmail.com",
		Password: "Tvsbk0101@",
	}

	db.Create(&owner)
	db.Create(&auth)
	tempOwnerToken, _ := utility.CreateJwtToken(utility.JwtPayload{Id: owner.ID, Role: owner.Role, LibId: owner.LibId})

	return tempOwnerToken, library.ID
}

func CreateTempBook(libdID uint) uint {
	db := GetTestDB()
	book := models.BookInventory{
		ISBN:            "1234567890",
		Title:           "Book1",
		Publisher:       "PublisherOne",
		Authors:         "AuthorOne",
		TotalCopies:     1,
		AvailableCopies: 1,
		Version:         "1.0",
		LibID:           libdID,
	}
	if err := db.Create(&book).Error; err != nil {
		return 0
	}
	log.Printf("book id %d", book.ID)
	log.Printf("library id %d", libdID)
	return book.ID
}

func CreateTempRequest(userId uint, bookId uint, reqType string) uint {

	db := GetTestDB()
	// request := dto.RequestEventDTO{
	// 	BookID:      bookId,
	// 	RequestType: "issue",
	// }
	var requestEvent = models.RequestEvent{
		BookID:        bookId,
		RequestType:   reqType,
		ReaderID:      userId,
		RequestStatus: "pending",
	}

	if result := db.Model(models.RequestEvent{}).Create(&requestEvent); result.Error != nil {
		fmt.Printf("Error creating request %v", result.Error)
		return 0
	}
	return requestEvent.ReqID
}
