package tests

import (
	"bytes"
	"encoding/json"
	"gatorcan-backend/database"
	"gatorcan-backend/models"
	"gatorcan-backend/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	// Setup the database and router
	SetupTestDB()
	//database.Connect()
	r := SetupTestRouter()

	// Create a test user
	password, _ := utils.HashPassword("testpassword")
	testUser := models.User{
		Username: "testuser",
		Email:    "testuser@example.com",
		Password: password,
		Roles:    []*models.Role{{Name: "user"}},
	}
	database.DB.Create(&testUser)

	// Define test cases
	tests := []struct {
		name         string
		payload      gin.H
		expectedCode int
		expectedMsg  string
	}{
		{
			name: "Valid login",
			payload: gin.H{
				"username": "testuser",
				"password": "testpassword",
			},
			expectedCode: http.StatusOK,
			expectedMsg:  "Login successful",
		},
		{
			name: "Invalid username",
			payload: gin.H{
				"username": "wronguser",
				"password": "testpassword",
			},
			expectedCode: http.StatusUnauthorized,
			expectedMsg:  "Invalid username or password",
		},
		{
			name: "Invalid password",
			payload: gin.H{
				"username": "testuser",
				"password": "wrongpassword",
			},
			expectedCode: http.StatusUnauthorized,
			expectedMsg:  "Invalid username or password",
		},
		{
			name: "Missing username",
			payload: gin.H{
				"password": "testpassword",
			},
			expectedCode: http.StatusBadRequest,
			expectedMsg:  "Key: 'Username' Error:Field validation for 'Username' failed on the 'required' tag",
		},
		{
			name: "Missing password",
			payload: gin.H{
				"username": "testuser",
			},
			expectedCode: http.StatusBadRequest,
			expectedMsg:  "Key: 'Password' Error:Field validation for 'Password' failed on the 'required' tag",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create request body
			body, _ := json.Marshal(tt.payload)
			req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			// Record the response
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			// Assert the response
			assert.Equal(t, tt.expectedCode, w.Code)
			var response map[string]string
			json.Unmarshal(w.Body.Bytes(), &response)
			assert.Contains(t, response["message"], tt.expectedMsg)
		})
	}

	// Clean up the test user
	database.DB.Delete(&testUser)
}
