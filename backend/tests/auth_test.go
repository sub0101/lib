package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"libraryManagement/internal/dto"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

//	func SetupAuthRouter(router *gin.Engine, db *gorm.DB) {
//		fmt.Println("auth router setup")
//		authController := auth.NewAuthController(db)
//		authRouter := router.Group("/api/v1/auth")
//		{
//			authRouter.POST("/login", authController.Login)
//			authRouter.POST("/signup", authController.Signup)
//			authRouter.POST("/library/signup", authController.SignupLibrary)
//			authRouter.POST("/forget", func(ctx *gin.Context) {})
//		}
//	}

func TestCreateLibrary(t *testing.T) {

	router := SetupRouter()
	// SetupTestLibrary()
	validTestCases := []struct {
		name string
		dto.RequestSignupLibraryBody

		expectedCode  int
		expectedError string
	}{
		{
			name: "Adding valid user",
			RequestSignupLibraryBody: dto.RequestSignupLibraryBody{
				Name:          "suraj singh",
				Email:         "owner1@gmail.com",
				LibraryName:   "LibraryOne",
				ContactNumber: "9801232331",
				Password:      "Tvsbk0101@",
			},
			expectedCode:  201,
			expectedError: "",
		},
		{
			name: "Request Signup Already Exist Test",
			RequestSignupLibraryBody: dto.RequestSignupLibraryBody{
				Name:          "suraj singh",
				Email:         "owner1@gmail.com",
				LibraryName:   "LibraryFive",
				ContactNumber: "1234567894",
				Password:      "Tvsbk0101@",
			},
			expectedCode:  400,
			expectedError: "invalid body",
		},
		{
			name: "Request Signup with Library Already Exist",
			RequestSignupLibraryBody: dto.RequestSignupLibraryBody{
				Name:          "suraj singh",
				Email:         "owner111@gmail.com",
				LibraryName:   "LibraryOne",
				ContactNumber: "1234567894",
				Password:      "Tvsbk0101@",
			},
			expectedCode:  400,
			expectedError: "User Already Exist",
		},
		{
			name: "Request Signup with invalid  body",
			RequestSignupLibraryBody: dto.RequestSignupLibraryBody{
				Name: "suraj singh",

				LibraryName:   "LibraryFive",
				ContactNumber: "1234567894",
				Password:      "Tvsbk0101@",
			},
			expectedCode:  400,
			expectedError: "User Already Exist",
		},
		{
			name: "Request Signup with Already Contact Exist",
			RequestSignupLibraryBody: dto.RequestSignupLibraryBody{
				Name:  "suraj singh",
				Email: "owner111@gmail.com",

				LibraryName:   "LibraryOne",
				ContactNumber: "1234567894",
				Password:      "Tvsbk0101@",
			},
			expectedCode:  400,
			expectedError: "user already exist",
		},
	}

	for _, tc := range validTestCases {
		t.Run(tc.name, func(t *testing.T) {
			jsonData, _ := json.Marshal(tc.RequestSignupLibraryBody)

			req, _ := http.NewRequest("POST", "/api/v1/auth/library/signup", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedCode, w.Code)
			fmt.Println(w.Body.String())

		})
	}
	stopTransaction()
	// TearDownTestDB()
}

func TestLoginUser(t *testing.T) {

	router := SetupRouter()
	SetupTestLibrary()
	var validTestCases = []struct {
		name string
		dto.RequestLoginBody
		expectedCode  int
		expectedError string
	}{
		{
			name: "Login User",
			RequestLoginBody: dto.RequestLoginBody{
				Email:    "owner1@example.com",
				Password: "Tvsbk0101@",
			},
			expectedCode:  200,
			expectedError: "",
		},
		{
			name: "Login Invalid User",
			RequestLoginBody: dto.RequestLoginBody{
				Email:    "user2@gmail.com",
				Password: "Tvsbk0102@",
			},
			expectedCode:  400,
			expectedError: "invalid email or password",
		},
		{
			name: "Login User With Invalid Credentials",
			RequestLoginBody: dto.RequestLoginBody{
				Email:    "user1@gmail.com",
				Password: "Tvsbk0102@",
			},
			expectedCode:  400,
			expectedError: "invalid email or password",
		},
		{
			name: "Login wtih Invalid Body",
			RequestLoginBody: dto.RequestLoginBody{

				Password: "Tvsbk0102@",
			},
			expectedCode:  400,
			expectedError: "invalid body",
		},
		{
			name: "invalid body",
			RequestLoginBody: dto.RequestLoginBody{
				Email:    "dddgmail.com",
				Password: "Tvsbk0102@",
			},
			expectedCode:  400,
			expectedError: "invalid body",
		},
	}

	for _, tc := range validTestCases {
		t.Run(tc.name, func(t *testing.T) {
			jsonData, _ := json.Marshal(tc.RequestLoginBody)

			req, _ := http.NewRequest("POST", "/api/v1/auth/login", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedCode, w.Code)

			if tc.expectedCode == 400 || tc.expectedCode == 401 || tc.expectedCode == 403 {
				var response map[string]interface{}
				json.Unmarshal(w.Body.Bytes(), &response)
				assert.Contains(t, response["message"], tc.expectedError)
			}
		})
	}
	stopTransaction()
}

func TestCreateReader(t *testing.T) {

	router := SetupRouter()
	libId, _ := SetupTestLibrary()
	SetupTestUsers()
	var validTestCases = []struct {
		name string
		dto.RequestSignupUserBody
		expectedCode  int
		expectedError string
	}{
		{
			name: "Adding valid user",
			RequestSignupUserBody: dto.RequestSignupUserBody{
				Name:          "suraj singh",
				Email:         "suraj01@gmail.com",
				ContactNumber: "8734567895",
				Password:      "Tvsbk0101@",
				LibId:         libId,
			},
			expectedCode:  201,
			expectedError: "",
		},
		{
			name: "Request Signup Already Exist Test",
			RequestSignupUserBody: dto.RequestSignupUserBody{
				Name:          "suraj singh",
				Email:         "suraj01@gmail.com",
				ContactNumber: "9834567899",
				Password:      "Tvsbk0101@",
				LibId:         libId,
			},
			expectedCode:  400,
			expectedError: "user already exist",
		},
		{
			name: "request Signup With Invalid Library",
			RequestSignupUserBody: dto.RequestSignupUserBody{
				Name:          "suraj singh",
				Email:         "suraj02@gmail.com",
				ContactNumber: "7834567850",
				Password:      "Tvsbk0101@",
				LibId:         120,
			},
			expectedCode:  400,
			expectedError: "library does not exist",
		},
		{
			name: "request Signup With Invalid Password or email",
			RequestSignupUserBody: dto.RequestSignupUserBody{
				Name:  "suraj singh",
				Email: "surajgmail.com",

				ContactNumber: "12345678",
				Password:      "Tvsbk1",
				LibId:         libId,
			},
			expectedCode:  400,
			expectedError: "invalid  body",
		},
		{
			name: "request Signup With Invalid Password or email",
			RequestSignupUserBody: dto.RequestSignupUserBody{
				Name: "suraj singh",

				ContactNumber: "12345678",
				Password:      "Tvsbk1",
				LibId:         libId,
			},
			expectedCode:  400,
			expectedError: "invalid  body",
		},
	}
	for _, tc := range validTestCases {
		t.Run(tc.name, func(t *testing.T) {
			jsonData, _ := json.Marshal(tc.RequestSignupUserBody)
			req, _ := http.NewRequest("POST", "/api/v1/auth/signup", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, tc.expectedCode, w.Code)
			if tc.expectedCode == 400 {
				var response map[string]interface{}
				json.Unmarshal(w.Body.Bytes(), &response)
				assert.Contains(t, response["message"], tc.expectedError)
			}
		})
	}
	// TearDownTestDB()
	stopTransaction()
}
