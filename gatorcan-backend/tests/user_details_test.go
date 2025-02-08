package tests

import (
	"encoding/json"
	"gatorcan-backend/database"
	"gatorcan-backend/models"
	"gatorcan-backend/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetUserDetailsSuccess tests fetching user details successfully
func TestGetUserDetailsSuccess(t *testing.T) {
	SetupTestDB()
	router := SetupTestRouter()

	// Insert a test user into the database
	adminToken, _ := utils.GenerateToken("adminuser", []string{"admin"})
	var userRole models.Role
	testUser := models.User{
		Username: "testuser",
		Email:    "testuser@example.com",
		Password: "hashedpassword",
		Roles:    []*models.Role{&userRole}, // Correctly assign role as []*models.Role
	}
	database.DB.Create(&testUser)

	// Request for user details with valid token
	req, _ := http.NewRequest("GET", "/user/testuser", nil)
	req.Header.Set("Authorization", "Bearer "+adminToken)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Validate response
	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "testuser", response["username"])
	assert.Equal(t, "testuser@example.com", response["email"])
	assert.Contains(t, response["roles"], "user")
}

// TestGetUserDetailsFailUnauthorized tests unauthorized access when no token is provided
func TestGetUserDetailsFailUnauthorized(t *testing.T) {
	SetupTestDB()
	router := SetupTestRouter()

	// Request for user details without any token
	req, _ := http.NewRequest("GET", "/user/testuser", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Validate response
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "Authorization token required")
}

// TestGetUserDetailsFailUserNotFound tests when the requested user is not found
func TestGetUserDetailsFailUserNotFound(t *testing.T) {
	SetupTestDB()
	router := SetupTestRouter()

	// Generate a valid admin token
	adminToken, _ := utils.GenerateToken("adminuser", []string{"admin"})

	// Request for user details of a non-existing user
	req, _ := http.NewRequest("GET", "/user/nonexistentuser", nil)
	req.Header.Set("Authorization", "Bearer "+adminToken)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Validate response
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "User not found")
}
