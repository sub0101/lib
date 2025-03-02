package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddBooks(t *testing.T) {

	router := SetupRouter()
	// jwtToken := getJWTToken(router) // Get valid JWT token for authentication

	tests := []struct {
		name          string
		book          map[string]interface{}
		token         string // JWT token (valid or empty)
		expectedCode  int
		expectedError string
	}{
		{
			name: "Valid Book 1",
			book: map[string]interface{}{
				"title":  "Go Lang Basics",
				"author": "John Doe",
				"isbn":   "1234567890",
			},
			token:        jwtToken, // ✅ Valid token
			expectedCode: 201,
		},
		{
			name: "Unauthorized User Trying to Add Book",
			book: map[string]interface{}{
				"title":  "Unauthorized Book",
				"author": "Fake Author",
				"isbn":   "9999999999",
			},
			token:         "", // ❌ No JWT token
			expectedCode:  401,
			expectedError: "Unauthorized",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			jsonData, _ := json.Marshal(tc.book)

			req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")

			if tc.token != "" {
				req.Header.Set("Authorization", "Bearer "+tc.token) // 🔥 Use JWT if provided
			}

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			// ✅ Validate response
			assert.Equal(t, tc.expectedCode, w.Code)

			// ✅ Check error message for unauthorized access
			if tc.expectedCode == 401 {
				var response map[string]interface{}
				json.Unmarshal(w.Body.Bytes(), &response)
				assert.Contains(t, response["error"], tc.expectedError)
			}
		})
	}
}
