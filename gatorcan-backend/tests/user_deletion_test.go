package tests

import (
	"gatorcan-backend/database"
	"gatorcan-backend/models"
	"gatorcan-backend/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteUserSuccess(t *testing.T) {
	SetupTestDB()
	router := SetupTestRouter()

	// Insert a test user into the database
	adminToken, _ := utils.GenerateToken("admin", []string{"admin"})
	var userRole models.Role
	testUser := models.User{
		Username: "testuser",
		Email:    "testuser@example.com",
		Password: "hashedpassword",
		Roles:    []*models.Role{&userRole}, // Correctly assign role as []*models.Role
	}
	database.DB.Create(&testUser)

	// Send DELETE request with valid admin token
	req, _ := http.NewRequest("DELETE", "/admin/testuser", nil)
	req.Header.Set("Authorization", "Bearer "+adminToken)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Validate response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "User testuser has been deleted successfully")
}

func TestDeleteUserFailUnauthorized(t *testing.T) {
	SetupTestDB()
	router := SetupTestRouter()

	// Send DELETE request without any token
	req, _ := http.NewRequest("DELETE", "/admin/testuser", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Validate response
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "Authorization token required")

	CloseTestDB()
}

func TestDeleteUserFailUserNotFound(t *testing.T) {
	SetupTestDB()
	router := SetupTestRouter()

	// Generate admin token
	adminToken, _ := utils.GenerateToken("adminuser", []string{"admin"})

	// Send DELETE request for a non-existing user
	req, _ := http.NewRequest("DELETE", "/admin/nonexistentuser", nil)
	req.Header.Set("Authorization", "Bearer "+adminToken)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Validate response
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "not found")

	CloseTestDB()
}
