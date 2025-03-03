package tests

import (
	"encoding/json"
	"fmt"
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

	var studentRole models.Role
	if err := database.DB.Where("name = ?", "student").First(&studentRole).Error; err != nil {
		t.Fatalf("Failed to fetch student role: %v", err)
	}
	adminToken, _ := utils.GenerateToken("adminuser", []string{"admin"})
	password, _ := utils.HashPassword("testpassword")
	testUser := models.User{
		Username: "testuser",
		Email:    "testuser@example.com",
		Password: password,
		Roles:    []*models.Role{&studentRole}, // Assign role correctly
	}
	fmt.Println(testUser.Roles[0].Name)
	database.DB.Create(&testUser)

	// Request for user details with valid token
	req, _ := http.NewRequest("GET", "/user/testuser", nil)
	req.Header.Set("Authorization", "Bearer "+adminToken)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	// Validate response
	assert.Equal(t, http.StatusOK, w.Code)

	var response models.User
	json.Unmarshal([]byte(w.Body.String()), &response)

	// Assertions
	assert.Equal(t, "testuser", response.Username)
	assert.Equal(t, "testuser@example.com", response.Email)

	roles := response.Roles

	// Extract role names safely
	var roleNames []string
	for _, role := range roles {
		roleNames = append(roleNames, role.Name)
	}

	// Validate that the role "user" exists in the list
	assert.Contains(t, studentRole.Name, "student", "Expected role 'student' in response")
	CloseTestDB()
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
	CloseTestDB()
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
	assert.Contains(t, w.Body.String(), "not found")

	CloseTestDB()
}
