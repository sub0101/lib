package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func init() {
// 	setupTestDB()
// }

var tempBookId uint

func TestAddIssueRequest(test *testing.T) {
	setupTestDB()
	router := SetupRouter()
	libId, _ := SetupTestLibrary()
	SetupTestUsers()
	bookId := CreateTempBook(libId)
	jwtToken := GetJwtUserToken()
	GetJwtOwnerToken()
	CreateTempOwner()
	tempReaderToken, _ := CreateTempReader()

	tests := []struct {
		name          string
		request       map[string]interface{}
		token         string // JWT token (valid or empty)
		expectedCode  int
		expectedError string
	}{
		{
			name: "Add Valid Issue Request Details with valid Book ID",
			request: map[string]interface{}{
				"bookId":      bookId,
				"requestType": "issue",
			},
			token:        jwtToken,
			expectedCode: 201,
		},
		{
			name: "Add Issue Request With Invalid Book ID",
			request: map[string]interface{}{
				"bookId":      12,
				"requestType": "issue",
			},
			token:         jwtToken,
			expectedCode:  404,
			expectedError: "book not found",
		},
		{
			name: " Invalid Issue wihout Login",
			request: map[string]interface{}{
				"bookId":      12,
				"requestType": "issue",
			},
			token:         "",
			expectedCode:  401,
			expectedError: "Unauthorized",
		},
		{
			name: " Invalid Issue with Unauthorized User",
			request: map[string]interface{}{
				"bookId":      bookId,
				"requestType": "issue",
			},
			token:         tempReaderToken,
			expectedCode:  403,
			expectedError: "Unauthorized",
		},
	}

	for _, tt := range tests {
		test.Run(tt.name, func(t *testing.T) {
			reqBody, _ := json.Marshal(tt.request)
			req, _ := http.NewRequest("POST", "/api/v1/request/", bytes.NewBuffer(reqBody))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+tt.token)
			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)
			assert.Equal(t, tt.expectedCode, res.Code)
		})
	}
	// TearDownTestDB()

}

func TestUpdateRequest(test *testing.T) {
	setupTestDB()
	router := SetupRouter()
	libId, _ := SetupTestLibrary()
	readerId := SetupTestUsers()
	bookId := CreateTempBook(libId)
	tempBookId = bookId
	jwtToken := GetJwtUserToken()
	_ = jwtToken
	reqId := CreateTempRequest(readerId, bookId, "issue")
	newReqId := CreateTempRequest(readerId, bookId, "return")

	ownerToken := GetJwtOwnerToken()
	// CreateTempOwner()
	// tempReaderToken := CreateTempReader()

	tests := []struct {
		name          string
		request       map[string]interface{}
		params        uint
		token         string
		expectedCode  int
		expectedError string
	}{
		{
			name: "Update Request with Issue status Valid Request ID",
			request: map[string]interface{}{

				"requestType": "accepted",
			},
			token:        ownerToken,
			params:       reqId,
			expectedCode: 200,
		},
		{
			name: "Update Request with Reject status Valid Request ID",
			request: map[string]interface{}{

				"requestType": "rejected",
			},
			token:        ownerToken,
			params:       reqId,
			expectedCode: 200,
		},
		{
			name: "Update Request with InValid Request ID",
			request: map[string]interface{}{

				"requestType": "accepted",
			},
			token:         ownerToken,
			params:        212,
			expectedCode:  400,
			expectedError: "invalid params",
		},
		{
			name: "Update Request with Invalid Request Type",
			request: map[string]interface{}{
				"requestType": "invalid",
			},
			token:         ownerToken,
			params:        reqId,
			expectedCode:  400,
			expectedError: "invalid request type",
		},
	}
	for _, tt := range tests {
		test.Run(tt.name, func(t *testing.T) {
			reqBody, _ := json.Marshal(tt.request)
			req, _ := http.NewRequest("PATCH", fmt.Sprintf("/api/v1/request/%d", tt.params), bytes.NewBuffer(reqBody))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+tt.token)
			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)
			assert.Equal(t, tt.expectedCode, res.Code)
		})
	}

	tempReaderToken, _ := CreateTempReader()

	test2 := []struct {
		name          string
		request       map[string]interface{}
		token         string // JWT token (valid or empty)
		expectedCode  int
		expectedError string
	}{
		{
			name: "Add Return Request with Valid Book ID",
			request: map[string]interface{}{
				"bookId":      tempBookId,
				"requestType": "return",
			},
			token:        jwtToken,
			expectedCode: 201,
		},
		{
			name: "Add Return Request with Invalid Book ID",
			request: map[string]interface{}{
				"bookId":      12,
				"requestType": "return",
			},
			token:         jwtToken,
			expectedCode:  404,
			expectedError: "book not found",
		},
		{
			name: "Invalid Return Request without Login",
			request: map[string]interface{}{
				"bookId":      12,
				"requestType": "return",
			},
			token:         "",
			expectedCode:  401,
			expectedError: "Unauthorized",
		},
		{
			name: "Invalid Return Request with Unauthorized User",
			request: map[string]interface{}{
				"bookId":      tempBookId,
				"requestType": "return",
			},
			token:         tempReaderToken,
			expectedCode:  403,
			expectedError: "Unauthorized",
		},
	}

	for _, tt := range test2 {
		test.Run(tt.name, func(t *testing.T) {
			reqBody, _ := json.Marshal(tt.request)
			req, _ := http.NewRequest("POST", "/api/v1/request/", bytes.NewBuffer(reqBody))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+tt.token)
			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)
			assert.Equal(t, tt.expectedCode, res.Code)
			log.Printf(res.Body.String())
		})
	}
	test3 := []struct {
		name          string
		request       map[string]interface{}
		token         string
		expectedCode  int
		params        uint
		expectedError string
	}{
		{
			name: "Update Return Request with Issue status Valid Request ID",
			request: map[string]interface{}{
				"requestType": "accepted",
			},
			token:        ownerToken,
			params:       newReqId,
			expectedCode: 200,
		},
		{
			name: "Update Return Request with Reject status Valid Request ID",
			request: map[string]interface{}{
				"requestType": "rejected",
			},
			token:        ownerToken,
			params:       newReqId,
			expectedCode: 200,
		},
		{
			name: "Update Return Request with  Inavlid Request Type",
			request: map[string]interface{}{
				"requestType": "invalid request type",
			},
			token:        ownerToken,
			params:       newReqId,
			expectedCode: 400,
		},
	}
	for _, tt := range test3 {
		test.Run(tt.name, func(t *testing.T) {
			reqBody, _ := json.Marshal(tt.request)
			req, _ := http.NewRequest("PATCH", fmt.Sprintf("/api/v1/request/%d", newReqId), bytes.NewBuffer(reqBody))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+tt.token)
			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)
			assert.Equal(t, tt.expectedCode, res.Code)
		})
	}
	// TearDownTestDB()

	// TearDownTestDB()
}

func TestGetAllRequests(test *testing.T) {
	setupTestDB()
	router := SetupRouter()
	libId, _ := SetupTestLibrary()
	readerId := SetupTestUsers()
	bookId := CreateTempBook(libId)
	_ = bookId
	jwtToken := GetJwtUserToken()
	_ = jwtToken
	CreateTempRequest(readerId, bookId, "issue")

	ownerToken := GetJwtOwnerToken()
	// _, _ := CreateTempOwner()
	// tempReaderToken := CreateTempReader()

	tests := []struct {
		name          string
		token         string
		expectedCode  int
		expectedError string
	}{
		{
			name:         "Get All Requests",
			token:        ownerToken,
			expectedCode: 200,
		},
		{
			name:          "Get All Requests without Login",
			token:         "",
			expectedCode:  401,
			expectedError: "Unauthorized",
		},
		{
			name:          "Get All Requests with Unauthorized User",
			token:         GetJwtUserToken(),
			expectedCode:  403,
			expectedError: "Unauthorized",
		},
	}

	for _, tt := range tests {
		test.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/api/v1/request/all", nil)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+tt.token)
			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)
			assert.Equal(t, tt.expectedCode, res.Code)
		})
	}

	// TearDownTestDB()
}

func TestGetUserRequests(test *testing.T) {
	setupTestDB()
	router := SetupRouter()
	libId, _ := SetupTestLibrary()
	readerId := SetupTestUsers()
	bookId := CreateTempBook(libId)
	_ = bookId
	jwtToken := GetJwtUserToken()
	_ = jwtToken
	CreateTempRequest(readerId, bookId, "issue")

	// ownerToken := GetJwtOwnerToken()
	// _, _ := CreateTempOwner()
	// tempReaderToken := CreateTempReader()

	tests := []struct {
		name          string
		token         string
		expectedCode  int
		expectedError string
	}{
		{
			name:         "Get User Requests",
			token:        jwtToken,
			expectedCode: 202,
		},
		{
			name:          "Get User Requests without Login",
			token:         "",
			expectedCode:  401,
			expectedError: "Unauthorized",
		},
		{
			name:          "Get User Requests with Unauthorized User",
			token:         GetJwtOwnerToken(),
			expectedCode:  403,
			expectedError: "Unauthorized",
		},
	}

	for _, tt := range tests {
		test.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/api/v1/request/", nil)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+tt.token)
			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)
			assert.Equal(t, tt.expectedCode, res.Code)
		})
	}

	// TearDownTestDB()
}

func TestDeleteRequest(test *testing.T) {

	setupTestDB()
	router := SetupRouter()
	libId, _ := SetupTestLibrary()
	readerId := SetupTestUsers()
	bookId := CreateTempBook(libId)
	_ = bookId
	jwtToken := GetJwtUserToken()
	_ = jwtToken
	reqId := CreateTempRequest(readerId, bookId, "issue")

	ownerToken := GetJwtOwnerToken()
	tempOwnerToken, _ := CreateTempOwner()
	// tempReaderToken := CreateTempReader()

	tests := []struct {
		name          string
		token         string
		params        uint
		expectedCode  int
		expectedError string
	}{
		{
			name:          "Delete Request with Unauthorized User",
			token:         tempOwnerToken,
			params:        reqId,
			expectedCode:  403,
			expectedError: "Unauthorized",
		},
		{
			name:         "Delete Request with Valid Request ID",
			token:        ownerToken,
			params:       reqId,
			expectedCode: 200,
		},
		{
			name:          "Delete Request with Invalid Request ID",
			token:         ownerToken,
			params:        212,
			expectedCode:  404,
			expectedError: "invalid params",
		},
		{
			name:          "Delete Request without Login",
			token:         "",
			params:        reqId,
			expectedCode:  401,
			expectedError: "Unauthorized",
		},
	}

	for _, tt := range tests {
		test.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("DELETE", fmt.Sprintf("/api/v1/request/%d", tt.params), nil)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+tt.token)
			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)
			assert.Equal(t, tt.expectedCode, res.Code)
		})
	}

	// TearDownTestDB()
}

func TestGetRequestById(test *testing.T) {
	setupTestDB()
	router := SetupRouter()
	libId, _ := SetupTestLibrary()
	readerId := SetupTestUsers()
	bookId := CreateTempBook(libId)
	_ = bookId
	jwtToken := GetJwtUserToken()
	_ = jwtToken
	reqId := CreateTempRequest(readerId, bookId, "issue")

	ownerToken := GetJwtOwnerToken()
	// tempOwnerToken, _ := CreateTempOwner()
	// tempReaderToken := CreateTempReader()

	tests := []struct {
		name          string
		token         string
		params        uint
		expectedCode  int
		expectedError string
	}{
		{
			name:         "Get Request By Reader Id",
			token:        ownerToken,
			params:       reqId,
			expectedCode: 200,
		},
		{
			name:         "Get Request By Owner Id",
			token:        ownerToken,
			params:       reqId,
			expectedCode: 200,
		},
		{
			name:          "Get Request By Id with Invalid Request ID",
			token:         ownerToken,
			params:        212,
			expectedCode:  404,
			expectedError: "request not found",
		},
		{
			name:          "Get Request By Id without Login",
			token:         "",
			params:        reqId,
			expectedCode:  401,
			expectedError: "Unauthorized",
		},
	}

	for _, tt := range tests {
		test.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", fmt.Sprintf("/api/v1/request/%d", tt.params), nil)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+tt.token)
			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)
			assert.Equal(t, tt.expectedCode, res.Code)
		})
	}

	// TearDownTestDB()
}
