package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"libraryManagement/app/middleware"
	book "libraryManagement/app/modules/Book"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// func init() {
// 	setupTestDB()
// }

func SetupBookRouter(router *gin.Engine, db *gorm.DB) {
	bookController := book.NewBookController(db)

	bookRouter := router.Group("/api/v1")
	{

		bookRouter.GET("/books", middleware.IsAuth(), bookController.GetAllBook)
		bookRouter.GET("/books/issued", middleware.IsAuth(), bookController.GetIssuedBooks)
		bookRouter.GET("/books/:id", middleware.IsAuth(), bookController.GetBook)
		bookRouter.POST("/books", middleware.IsAuth("owner", "admin"), bookController.AddBook)
		bookRouter.PATCH("/:id", middleware.IsAuth("owner", "admin"), bookController.UpdateBook)
		bookRouter.DELETE("/:id", middleware.IsAuth("owner", "admin"), bookController.DeleteBook)
		bookRouter.GET("/search", middleware.IsAuth(), bookController.SearchBooks)
	}
}

func TestAddBooks(t *testing.T) {
	setupTestDB()
	router := SetupRouter()
	SetupTestLibrary()
	SetupTestUsers()

	jwtToken := GetJwtOwnerToken()

	tests := []struct {
		name          string
		book          map[string]interface{}
		token         string // JWT token (valid or empty)
		expectedCode  int
		expectedError string
	}{
		{
			name: "Add Valid Book Details",
			book: map[string]interface{}{
				"isbn":            "9756523612",
				"title":           "bookqw",
				"authors":         "authoro",
				"publisher":       "publishero",
				"version":         "1.1.1",
				"totalCopies":     150,
				"availableCopies": 150,
			},
			token:        jwtToken,
			expectedCode: 201,
		},
		{
			name: "Add Already Existing Book",
			book: map[string]interface{}{
				"isbn":            "9756523612",
				"title":           "bookqw",
				"authors":         "authoro",
				"publisher":       "publishero",
				"version":         "1.1.1",
				"totalCopies":     100,
				"availableCopies": 100,
			},
			token:        jwtToken,
			expectedCode: 201,
		},
		{
			name: "Add Book with invalid different copies",
			book: map[string]interface{}{
				"isbn":            "9756521212",
				"title":           "bookqw",
				"authors":         "authoro",
				"publisher":       "publishero",
				"version":         "1.1.1",
				"totalCopies":     70,
				"availableCopies": 150,
			},
			token:        jwtToken,
			expectedCode: 400,
		},
		{
			name: "Unauthorized User Trying to Add Book",
			book: map[string]interface{}{
				"isbn":            "9756523611",
				"title":           "bookqw",
				"authors":         "authoro",
				"publisher":       "publishero",
				"version":         "1.1.1",
				"totalCopies":     150,
				"availableCopies": 150,
			},
			token:         "",
			expectedCode:  401,
			expectedError: "Invalid Token",
		},
		{
			name: "Invalid book Payload",
			book: map[string]interface{}{
				"isbn":            "9756523",
				"title":           "bookqw12",
				"authors":         "authoro1",
				"publisher":       "publishero",
				"version":         "1.1.1",
				"totalCopies":     150,
				"availableCopies": 150,
			},
			token:         jwtToken,
			expectedCode:  400,
			expectedError: "invalid payload",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			jsonData, err := json.Marshal(tc.book)
			if err != nil {
				t.Fatalf("Failed to marshal book: %v", err)
			}

			req, _ := http.NewRequest("POST", "/api/v1/books/", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")

			if tc.token != "" {
				req.Header.Set("Authorization", "Bearer "+tc.token)
			}

			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedCode, w.Code)

		})
	}

	// TearDownTestDB()
}

func TestGetAllBooks(t *testing.T) {
	setupTestDB()
	router := SetupRouter()
	SetupTestLibrary()
	SetupTestUsers()

	tests := []struct {
		name          string
		token         string
		expectedCode  int
		expectedError string
	}{
		{
			name:          "Get All Books By Owner",
			token:         GetJwtOwnerToken(),
			expectedCode:  200,
			expectedError: "Invalid Token",
		},
		{
			name:          "Get All Books By reader",
			token:         GetJwtUserToken(),
			expectedCode:  200,
			expectedError: "Invalid Token",
		},
		{
			name:          "Get All Books Without JWT",
			token:         "",
			expectedCode:  401,
			expectedError: "Invalid Token",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/api/v1/books/", nil)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+tc.token)
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedCode, w.Code)

			// if tc.expectedCode == 200 {
			// 	var books []models.BookInventory
			// 	json.Unmarshal(w.Body.Bytes(), &books)
			// 	assert.GreaterOrEqual(t, len(books), 0)
			// }
		})
	}

	// TearDownTestDB()
}

func TestGetBookDetail(t *testing.T) {
	setupTestDB()
	router := SetupRouter()
	libId, _ := SetupTestLibrary()
	SetupTestUsers()
	tempOwnerToken, _ := CreateTempOwner()

	bookId := CreateTempBook(libId)

	tests := []struct {
		name          string
		token         string
		expectedCode  int
		params        int
		expectedError string
	}{
		{
			name:          "Get Book By Valid Id",
			token:         GetJwtOwnerToken(),
			params:        int(bookId),
			expectedCode:  200,
			expectedError: "",
		},
		{
			name:          "Get Book By Without Login",
			token:         "",
			expectedCode:  401,
			expectedError: "invalid token",
			params:        int(bookId),
		},
		{
			name:          "Get Book By Unauthorized User",
			token:         tempOwnerToken,
			expectedCode:  403,
			expectedError: "not authorized to access this resource",
			params:        int(bookId),
		},
		{
			name:          "Get Book By Wihout Params Id",
			token:         GetJwtUserToken(),
			expectedCode:  400,
			expectedError: "invalid params",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", fmt.Sprintf("/api/v1/books/%d", tc.params), nil)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+tc.token)
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedCode, w.Code)

		})
	}
	// TearDownTestDB()
}

func TestUpdateBook(t *testing.T) {
	setupTestDB()
	router := SetupRouter()
	libId, _ := SetupTestLibrary()
	readerId := SetupTestUsers()
	GetJwtOwnerToken()
	_ = readerId
	tempOwnerToken, _ := CreateTempOwner()

	bookId := CreateTempBook(libId)

	tests := []struct {
		name          string
		token         string
		expectedCode  int
		book          map[string]interface{}
		expectedError string
		params        int
	}{
		{
			name:   "Update Book By Valid Id",
			token:  GetJwtOwnerToken(),
			params: int(bookId),
			book: map[string]interface{}{

				"availableCopies": 100,
			},
			expectedCode: 201,
		},
		{
			name:          "Update Book By Without Login",
			token:         "",
			expectedCode:  401,
			expectedError: "invalid token",
			params:        int(bookId),
		},
		{
			name:          "Update Book By Invalid Book Id",
			token:         testOwnerToken,
			expectedCode:  404,
			expectedError: "invalid token",
			params:        121,
		},
		{
			name:          "Update Book By Unauthorized User",
			token:         tempOwnerToken,
			expectedCode:  403,
			expectedError: "can not update book",
			params:        int(bookId),
		},
		{
			name:          "Update Book By Wihout Params Id",
			token:         GetJwtOwnerToken(),
			expectedCode:  400,
			expectedError: "invalid params",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			jsonData, err := json.Marshal(tc.book)
			if err != nil {
				t.Fatalf("Failed to marshal book: %v", err)
			}

			req, _ := http.NewRequest("PATCH", fmt.Sprintf("/api/v1/books/%d", tc.params), bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+tc.token)
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedCode, w.Code)

		})
	}

	// TearDownTestDB()
}

func TestDeleteBook(t *testing.T) {
	setupTestDB()
	router := SetupRouter()
	libId, _ := SetupTestLibrary()
	readerId := SetupTestUsers()
	_ = readerId
	tempOwnerToken, _ := CreateTempOwner()

	bookId := CreateTempBook(libId)

	tests := []struct {
		name          string
		token         string
		expectedCode  int
		expectedError string
		params        int
	}{
		{
			name:          "Delete Book By Valid Id",
			token:         GetJwtOwnerToken(),
			params:        int(bookId),
			expectedCode:  200,
			expectedError: "",
		},
		{
			name:          "Delete Book By Valid Id But No Copies Left",
			token:         GetJwtOwnerToken(),
			params:        int(bookId),
			expectedCode:  400,
			expectedError: "",
		},
		{
			name:          "Delete Book By Without Login",
			token:         "",
			expectedCode:  401,
			expectedError: "invalid token",
			params:        int(bookId),
		},
		{
			name:          "Delete Book By Unauthorized User",
			token:         tempOwnerToken,
			expectedCode:  403,
			expectedError: "you are not authorized to access this resource",
			params:        int(bookId),
		},
		{
			name:          "Delete Book By Wihout Params Id",
			token:         GetJwtOwnerToken(),
			expectedCode:  400,
			expectedError: "invalid params",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req, _ := http.NewRequest("DELETE", fmt.Sprintf("/api/v1/books/%d", tc.params), nil)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+tc.token)
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedCode, w.Code)

		})
	}

	// TearDownTestDB()
}

func TestGetIssuedBook(t *testing.T) {

	router := SetupRouter()
	SetupTestLibrary()
	SetupTestUsers()

	tests := []struct {
		name          string
		token         string
		expectedCode  int
		expectedError string
	}{
		{
			name:         "Get All Issued Books",
			token:        GetJwtUserToken(),
			expectedCode: 200,
		},
		{
			name:          "Get All Issued Books Without JWT",
			token:         "",
			expectedCode:  401,
			expectedError: "Invalid Token",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/api/v1/books/issued", nil)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+tc.token)
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedCode, w.Code)

			// if tc.expectedCode == 200 {
			// 	var books []models.BookInventory
			// 	json.Unmarshal(w.Body.Bytes(), &books)
			// 	assert.GreaterOrEqual(t, len(books), 0)
			// }
		})
	}
	stopTransaction()
	// TearDownTestDB()
}

func TestSeachBook(t *testing.T) {

	router := SetupRouter()
	libId, _ := SetupTestLibrary()
	SetupTestUsers()
	CreateTempBook(libId)
	tempToken, _ := CreateTempOwner()

	tests := []struct {
		name          string
		token         string
		searchBody    map[string]interface{}
		expectedCode  int
		expectedError string
	}{
		{
			name:  "Search Books",
			token: GetJwtUserToken(),
			searchBody: map[string]interface{}{
				"isbn":  "9756523612",
				"title": "Book1",
			},
			expectedCode: 200,
		},
		{
			name:  "Search Book BY Unauthorized User",
			token: tempToken,
			searchBody: map[string]interface{}{
				"isbn":  "9756523612",
				"title": "Book1",
			},
			expectedCode:  404,
			expectedError: "not found",
		},
		{
			name:          "Search Books Without JWT",
			token:         "",
			expectedCode:  401,
			expectedError: "Invalid Token",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			jsonData, _ := json.Marshal(tc.searchBody)
			req, _ := http.NewRequest("GET", "/api/v1/books/search", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+tc.token)
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedCode, w.Code)

			// if tc.expectedCode == 200 {
			// 	var books []models.BookInventory
			// 	json.Unmarshal(w.Body.Bytes(), &books)
			// 	assert.GreaterOrEqual(t, len(books), 0)
			// }
		})
	}
	stopTransaction()
	// TearDownTestDB()
}
