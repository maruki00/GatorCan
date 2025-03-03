package tests

import (
	"bytes"
	//"log"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"gatorcan-backend/database"
	"gatorcan-backend/models"
	//"gatorcan-backend/tests"
	"gatorcan-backend/utils"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetEnrolledCourses(t *testing.T) {

	// Setup the database and router
	SetupTestDB()
	//database.Connect()
	r := SetupTestRouter()

	// Create a test user
	password, _ := utils.HashPassword("admin")
	testUser := models.User{
		Username: "student",
		Email:    "student@example.com",
		Password: password,
		Roles:    []*models.Role{{Name: string(models.Student)}},
	}
	database.DB.Create(&testUser)

	// Define test cases
	tests := []struct {
		name         string
		payload      gin.H
		expectedCode int
		expectedMsg  string
		token        string
	}{
		{
			name: "Unauthorized",
			payload: gin.H{
				"username": "testuser",
				"password": "testpassword",
			},
			token:        generateToken("admin", []string{string(models.Admin)}),
			expectedCode: http.StatusForbidden,
			expectedMsg:  "Unauthorized access",
		},
		{
			name: "Success",
			payload: gin.H{
				"username": "student",
				"password": "admin",
			},
			expectedCode: http.StatusOK,
			token:        generateToken("student", []string{string(models.Student)}),
			expectedMsg:  "",
			// expectedBody:   `[]`, // Assuming the user has no enrolled courses for simplicity
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create request body
			body, _ := json.Marshal(tt.payload)
			req, _ := http.NewRequest("GET", "/courses/enrolled", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+tt.token)

			// Record the response
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			// Assert the response
			assert.Equal(t, tt.expectedCode, w.Code)
			var response map[string]string
			json.Unmarshal(w.Body.Bytes(), &response)
			assert.Contains(t, response["error"], tt.expectedMsg)
		})
	}
}
