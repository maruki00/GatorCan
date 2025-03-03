package tests

import (
	"bytes"
	"gatorcan-backend/database"
	"gatorcan-backend/models"
	"gatorcan-backend/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdatePassword(t *testing.T) {
	SetupTestDB()
	router := SetupTestRouter()

	// Create test user
	hashedPassword, _ := utils.HashPassword("oldpassword")
	testUser := models.User{
		Username: "testuser",
		Email:    "testuser@example.com",
		Password: hashedPassword,
	}
	database.DB.Create(&testUser)

	// Generate JWT token
	token, _ := utils.GenerateToken("testuser", []string{"student"})

	// ✅ Test successful password update
	reqBody := `{"old_password": "oldpassword", "new_password": "newsecurepassword"}`
	req, _ := http.NewRequest("PUT", "/user/update", bytes.NewReader([]byte(reqBody)))
	req.Header.Set("Authorization", "Bearer "+token)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Password updated successfully")
}

func TestUpdateRoles(t *testing.T) {
	SetupTestDB()
	router := SetupTestRouter()

	// Create test users
	studentRole := models.Role{Name: "student"}
	adminRole := models.Role{Name: "admin"}
	database.DB.Create(&studentRole)
	database.DB.Create(&adminRole)

	// Create test users
	adminUser := models.User{Username: "admin", Email: "admin@example.com", Password: "hashedpassword"}
	regularUser := models.User{Username: "testuser", Email: "testuser@example.com", Password: "hashedpassword"}
	database.DB.Create(&adminUser)
	database.DB.Create(&regularUser)

	// Generate JWT tokens
	adminToken, _ := utils.GenerateToken("admin", []string{"admin"})
	userToken, _ := utils.GenerateToken("testuser", []string{"user"}) // userToken is used here

	// ✅ Test successful role update by admin
	reqBody := `{"username": "testuser", "roles": ["student"]}`
	req, _ := http.NewRequest("PUT", "/admin/update_role", bytes.NewReader([]byte(reqBody)))
	req.Header.Set("Authorization", "Bearer "+adminToken)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "User roles updated successfully")

	// ❌ Test non-admin trying to update roles
	req, _ = http.NewRequest("PUT", "/admin/update_role", bytes.NewReader([]byte(reqBody)))
	req.Header.Set("Authorization", "Bearer "+userToken) // userToken used here

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusForbidden, w.Code)
	assert.Contains(t, w.Body.String(), "Unauthorized access")

	// ❌ Test updating roles for a non-existent user
	reqBody = `{"username": "nonexistentuser", "roles": ["admin"]}`
	req, _ = http.NewRequest("PUT", "/admin/update_role", bytes.NewReader([]byte(reqBody)))
	req.Header.Set("Authorization", "Bearer "+adminToken)

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "User not found")

	// ❌ Test missing authentication
	req, _ = http.NewRequest("PUT", "/admin/update_role", bytes.NewReader([]byte(reqBody)))

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "Authorization token required")

	CloseTestDB()
}
