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

	// Insert a test role into the database
	userRole := models.Role{Name: "user"}
	database.DB.Create(&userRole) // Ensure the role exists in DB

	// Insert a test user with the role
	adminToken, _ := utils.GenerateToken("adminuser", []string{"admin"})
	testUser := models.User{
		Username: "testuser",
		Email:    "testuser@example.com",
		Password: "hashedpassword",
		Roles:    []*models.Role{&userRole}, // Assign role correctly
	}
	database.DB.Create(&testUser)

	// Request for user details with valid token
	req, _ := http.NewRequest("GET", "/admin/testuser", nil)
	req.Header.Set("Authorization", "Bearer "+adminToken)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	// Validate response
	assert.Equal(t, http.StatusOK, w.Code)

	// Decode response
	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	// Assertions
	assert.Equal(t, "testuser", response["username"])
	assert.Equal(t, "testuser@example.com", response["email"])

	// Ensure roles are correctly formatted
	roles, ok := response["roles"].([]interface{})
	assert.True(t, ok, "Roles should be a list")

	// Extract role names safely
	var roleNames []string
	for _, role := range roles {
		switch v := role.(type) {
		case string:
			roleNames = append(roleNames, v)
		case map[string]interface{}:
			if name, exists := v["name"].(string); exists {
				roleNames = append(roleNames, name)
			}
		}
	}

	// Validate that the role "user" exists in the list
	assert.Contains(t, roleNames, "user", "Expected role 'user' in response")
}

// TestGetUserDetailsFailUnauthorized tests unauthorized access when no token is provided
func TestGetUserDetailsFailUnauthorized(t *testing.T) {
	SetupTestDB()
	router := SetupTestRouter()

	// Request for user details without any token
	req, _ := http.NewRequest("GET", "/admin/testuser", nil)
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
	req, _ := http.NewRequest("GET", "/admin/nonexistentuser", nil)
	req.Header.Set("Authorization", "Bearer "+adminToken)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Validate response
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "User not found")
}
