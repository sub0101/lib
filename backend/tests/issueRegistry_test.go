package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

//	func init() {
//		setupTestDB()
//	}
func TestGetAllIssues(t *testing.T) {
	setupTestDB()
	router := SetupRouter()
	SetupTestLibrary()
	SetupTestUsers()

	test := []struct {
		name          string
		token         string
		expectedCode  int
		expectedError string
		params        uint
	}{
		{
			name:          "Get All Issues",
			token:         testOwnerToken,
			expectedCode:  202,
			expectedError: "",
		},
		{
			name:          "Get All Issues  By Unauthorized",
			token:         testReaderToken,
			expectedCode:  403,
			expectedError: "Forbidden",
		},
		{
			name:          "Get All Issues by invalid token",
			token:         "",
			expectedCode:  401,
			expectedError: "Unauthorized",
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/api/v1/issue_register/", nil)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+tt.token)
			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)
			assert.Equal(t, tt.expectedCode, res.Code)
		})
	}
	// TearDownTestDB()
}
