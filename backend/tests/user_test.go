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

func TestMakeAdmin(t *testing.T) {
	setupTestDB()
	router := SetupRouter()
	SetupTestLibrary()
	readerId := SetupTestUsers()
	log.Println(readerId)
	tempOwnerToken, _ := CreateTempOwner()

	testReaderToken, tempReaderId := CreateTempReader()

	log.Printf("testReaderId: %v", tempReaderId)
	log.Printf("testReaderToken: %v", testReaderToken)
	test := []struct {
		name          string
		token         string
		expectedCode  int
		expectedError string
		params        uint
		body          map[string]interface{}
	}{
		{
			name:          "Make Admin",
			token:         testOwnerToken,
			expectedCode:  202,
			expectedError: "",
			body: map[string]interface{}{
				"id":   readerId,
				"role": "admin",
			},
		},
		{
			name:          "Make Admin By Unauthorized",
			token:         tempOwnerToken,
			expectedCode:  404,
			expectedError: "Forbidden",
			body: map[string]interface{}{
				"id":   readerId,
				"role": "admin",
			},
		},
		{
			name:          "Make Admin By invalid role",
			token:         testOwnerToken,
			expectedCode:  400,
			expectedError: "invalid role",
			body: map[string]interface{}{
				"id":   readerId,
				"role": "admisdfdsn",
			},
		},
		{
			name:          "Make Admin by invalid token",
			token:         "",
			expectedCode:  401,
			expectedError: "Unauthorized",
			body: map[string]interface{}{
				"id":   readerId,
				"role": "admin",
			},
		},
		{
			name:          "Make Admin to Unauthorized user",
			token:         testOwnerToken,
			expectedCode:  404,
			expectedError: "Forbidden",
			body: map[string]interface{}{
				"id":   tempReaderId,
				"role": "admin",
			},
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			jsonData, _ := json.Marshal(tt.body)
			req, _ := http.NewRequest("PATCH", "/api/v1/users/make_admin", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+tt.token)
			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)
			assert.Equal(t, tt.expectedCode, res.Code)
		})
	}
	// TearDownTestDB()
}

func TestGetUser(t *testing.T) {
	setupTestDB()
	router := SetupRouter()
	SetupTestLibrary()
	readerId := SetupTestUsers()
	tempOwnerToken, _ := CreateTempOwner()

	test := []struct {
		name          string
		token         string
		expectedCode  int
		expectedError string
		params        uint
	}{
		{
			name:          "Get User",
			token:         testOwnerToken,
			expectedCode:  202,
			expectedError: "",
			params:        readerId,
		},
		{
			name:          "Get User Details By reader",
			token:         testReaderToken,
			expectedCode:  202,
			expectedError: "",
			params:        readerId,
		},
		{
			name:          "Get User By Unauthorized",
			token:         tempOwnerToken,
			expectedCode:  403,
			expectedError: "Forbidden",
			params:        1,
		},
		{
			name:          "Get User By Invalid user id",
			token:         testOwnerToken,
			expectedCode:  404,
			expectedError: "not found",
			params:        1121,
		},
		{
			name:          "Get User by invalid token",
			token:         "",
			expectedCode:  401,
			expectedError: "unauthorized",
			params:        readerId,
		},
		{
			name:          "Get User by invalid params",
			token:         "",
			expectedCode:  401,
			expectedError: "unauthorized",
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", fmt.Sprintf("/api/v1/users/%d", tt.params), nil)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+tt.token)
			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)
			assert.Equal(t, tt.expectedCode, res.Code)
		})
	}
	// TearDownTestDB()
}

func TestGetAllUser(t *testing.T) {
	setupTestDB()
	router := SetupRouter()
	SetupTestLibrary()
	SetupTestUsers()
	// tempOwnerToken, _ := CreateTempOwner()

	test := []struct {
		name          string
		token         string
		expectedCode  int
		expectedError string
	}{
		{
			name:          "Get All User",
			token:         GetJwtOwnerToken(),
			expectedCode:  202,
			expectedError: "",
		},
		{
			name:          "Get All User By Unauthorized",
			token:         GetJwtUserToken(),
			expectedCode:  403,
			expectedError: "Forbidden",
		},
		{
			name:          "Get All User by invalid token",
			token:         "",
			expectedCode:  401,
			expectedError: "Unauthorized",
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/api/v1/users/", nil)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+tt.token)
			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)
			assert.Equal(t, tt.expectedCode, res.Code)
		})
	}
	// TearDownTestDB()
}

func TestDeleteUser(t *testing.T) {
	setupTestDB()
	router := SetupRouter()
	SetupTestLibrary()
	readerId := SetupTestUsers()
	tempOwnerToken, _ := CreateTempOwner()
	log.Printf("readerId: %v", readerId)
	test := []struct {
		name          string
		token         string
		expectedCode  int
		expectedError string
		query         uint
	}{
		{
			name:          "Delete User By Unauthorized",
			token:         tempOwnerToken,
			expectedCode:  403,
			expectedError: "Forbidden",
			query:         readerId,
		},
		{
			name:          "Delete User By Valid Owner",
			token:         testOwnerToken,
			expectedCode:  202,
			expectedError: "",
			query:         readerId,
		},

		{
			name:          "Delete User by invalid token",
			token:         "",
			expectedCode:  401,
			expectedError: "unauthorized",
			query:         readerId,
		},
		{
			name:          "Delete User by invalid params",
			token:         "",
			expectedCode:  401,
			expectedError: "unauthorized",
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("DELETE", "/api/v1/users/", nil)
			q := req.URL.Query()
			q.Add("id", fmt.Sprintf("%d", tt.query))
			req.URL.RawQuery = q.Encode() // Encode query back to URL

			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+tt.token)
			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)
			assert.Equal(t, tt.expectedCode, res.Code)
		})
	}
	// TearDownTestDB()
}
