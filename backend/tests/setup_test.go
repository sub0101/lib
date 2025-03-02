package tests

import (
	"libraryManagement/app/middleware"
	"libraryManagement/app/routes"
	"libraryManagement/internal/models"
	"libraryManagement/utility"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupTestDB() {
	var err error
	testDB, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{}) // 🔥 In-memory DB
	if err != nil {
		log.Fatal("Failed to connect to test database:", err)
	}

	// Auto-migrate models for testing
	testDB.AutoMigrate(&models.BookInventory{}, &models.User{})

	// ✅ Create a test user
	testUser := models.User{
		Email:    "test@example.com",
		Password: utility.HashPassword("123"),
	}
	testDB.Create(&testUser)

	// ✅ Generate a JWT token for the test user
	testToken, _ = utils.GenerateJWT(testUser.Email) // Your JWT function
}

// setupRouter initializes the Gin router for testing.
func SetupRouter() *gin.Engine {
	router := gin.Default() // Create a new Gin router

	// Apply JWT authentication middleware
	router.Use(middleware.IsAuth())
	mrouter := router.RouterGroup("/api/v1")
	// Load all routes
	// routes.SetupRoutes(router, testDB) // Your routes
	routes.SetupRouter(mrouter, test)

	return router
}
