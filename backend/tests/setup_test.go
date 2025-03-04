package tests

import (
	"fmt"
	book "libraryManagement/app/modules/Book"
	user "libraryManagement/app/modules/User"
	auth "libraryManagement/app/modules/auth"
	issueregistery "libraryManagement/app/modules/issueRegistery"
	requestevent "libraryManagement/app/modules/requestEvent"
	"libraryManagement/internal/models"
	"libraryManagement/utility"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var testDB *gorm.DB
var dbb *gorm.DB
var testReaderToken string
var testOwnerToken string
var libId uint

func init() {
	setupTestDB()
}
func setupTestDB() {
	var err error
	godotenv.Load(".env")
	// dbURL := "postgres://postgres:Suraj0101@localhost:5432/LibraryTestProject"
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	// db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	err = db.AutoMigrate(&models.Library{}, &models.Auth{}, &models.User{}, &models.BookInventory{}, &models.IssueRegistery{}, &models.RequestEvent{})

	if err != nil {
		log.Fatal("Failed to connect to test database:", err)
	}

	db.AutoMigrate(&models.Auth{}, &models.BookInventory{}, &models.User{})
	dbb = db
	// SetupTestLibrary()
	// SetupTestUsers()

}

func startTransaction() *gorm.DB {
	testDB := dbb.Begin()
	return testDB
}

func SetupRouter() *gin.Engine {

	router := gin.Default()

	fmt.Println("book router intialized")
	mRouter := router.Group("/api/v1")
	testDB = startTransaction()

	auth.SetupAuthRouter(mRouter, testDB)
	book.SetupBookRouter(mRouter, testDB)
	requestevent.SetupRequestEventRouter(mRouter, testDB)
	user.SetupUserRouter(mRouter, testDB)
	issueregistery.SetupIssueRegistryRouter(mRouter, testDB)
	// routes.SetupRouter(mRouter, testDB)
	// SetupBookRouter(router, testDB)
	// SetupAuthRouter(router, testDB)

	return router
}

func stopTransaction() {
	testDB.Rollback()
}

func GetJwtOwnerToken() string {
	return testOwnerToken
}
func GetJwtUserToken() string {
	return testReaderToken
}

func SetupTestLibrary() (uint, uint) {
	password, err := utility.HashPassword("Tvsbk0101@")
	if err != nil {
		fmt.Println("Error creating at the time of hashing")
		return 0, 0

	}
	authUser := models.Auth{
		Email:    "owner1@example.com",
		Password: password,
	}
	library := models.Library{
		Name: "LibraryOne",
	}

	if result := testDB.Create(&library); result.Error != nil {
		fmt.Println("Error creating database")
		return 0, 0
	}

	testUser := models.User{
		Email:         "owner1@example.com",
		ContactNumber: "1234567890",
		Name:          "suraj",
		Role:          "owner",
	}
	testUser.LibId = library.ID
	libId = library.ID

	if result := testDB.Create(&authUser); result.Error != nil {
		fmt.Println("Error creating database")
		return 0, 0
	}
	if result := testDB.Create(&testUser); result.Error != nil {
		fmt.Println("Error creating database")
		return 0, 0
	}

	testOwnerToken, _ = utility.CreateJwtToken(utility.JwtPayload{Id: testUser.ID, Role: testUser.Role, LibId: testUser.ID})
	return library.ID, testUser.ID
}

func SetupTestUsers() uint {
	password, err := utility.HashPassword("Tvsbk0101@")
	if err != nil {
		fmt.Println("Error creating at the time of hashing")
		return 0

	}
	authUser := models.Auth{
		Email:    "user1@example.com",
		Password: password,
	}
	testUser := models.User{
		Email:         "user1@example.com",
		ContactNumber: "1234567891",
		LibId:         libId,
		Name:          "surajReader",
		Role:          "reader",
	}
	_ = authUser
	_ = testUser
	if result := testDB.Create(&authUser); result.Error != nil {
		fmt.Println("Error creating database")
	}
	if result := testDB.Create(&testUser); result.Error != nil {
		fmt.Println("Error creating database")
	}
	log.Printf("reader %v", testUser)
	testReaderToken, _ = utility.CreateJwtToken(utility.JwtPayload{Id: testUser.ID, Role: testUser.Role, LibId: testUser.LibId})
	return testUser.ID
}

// func TearDownTestDB() {
// 	ClearTestDB()
// }

func ClearTestDB() {
	err := testDB.Exec("TRUNCATE TABLE libraries, auths, users, book_inventories, issue_registeries, request_events RESTART IDENTITY CASCADE").Error
	if err != nil {
		log.Fatal("Failed to truncate tables:", err)
	}
	fmt.Println("Test database truncated successfully!")
}

func GetTestDB() *gorm.DB {
	return testDB
}
