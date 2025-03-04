package database

import (
	"libraryManagement/internal/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {

}

func InitDB() (error, *gorm.DB) {

	dbURL := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	err = db.AutoMigrate(&models.Library{}, &models.Auth{}, &models.User{}, &models.BookInventory{}, &models.IssueRegistery{}, &models.RequestEvent{})

	if err != nil {
		log.Fatal(err)
		return err, nil

	}
	return nil, db
}
