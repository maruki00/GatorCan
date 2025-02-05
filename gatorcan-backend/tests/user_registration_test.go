package tests

import (
	"bytes"
	"encoding/json"
	"gatorcan-backend/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCreateUserSuccess ensures only an Admin can register new users
func TestCreateUserSuccess(t *testing.T) {
	SetupTestDB()
	router := SetupTestRouter()

	// Generate an admin JWT token
	adminToken, _ := utils.GenerateToken("adminuser", []string{"admin"})

	// Define request payload
	userData := map[string]interface{}{
		"username": "newuser",
		"email":    "newuser@example.com",
		"password": "securepass",
		"roles":    []string{"user"},
	}
	jsonData, _ := json.Marshal(userData)

	// Send request with valid Admin token
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+adminToken)

	// Execute request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Validate response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "created successfully")
}

// TestCreateUserFailNonAdmin ensures that non-admin users cannot register others
func TestCreateUserFailNonAdmin(t *testing.T) {
	SetupTestDB()
	router := SetupTestRouter()

	// Generate a regular user JWT token
	userToken, _ := utils.GenerateToken("normaluser", []string{"user"})

	userData := map[string]interface{}{
		"username": "testuser",
		"email":    "testuser@example.com",
		"password": "password123",
		"roles":    []string{"user"},
	}
	jsonData, _ := json.Marshal(userData)

	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+userToken)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Expect Forbidden (403)
	assert.Equal(t, http.StatusForbidden, w.Code)
	assert.Contains(t, w.Body.String(), "Access denied")
}

// TestCreateUserFailInvalidToken ensures registration fails with an invalid token
func TestCreateUserFailInvalidToken(t *testing.T) {
	SetupTestDB()
	router := SetupTestRouter()

	userData := map[string]interface{}{
		"username": "testuser",
		"email":    "testuser@example.com",
		"password": "password123",
		"roles":    []string{"user"},
	}
	jsonData, _ := json.Marshal(userData)

	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer invalid-token")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Expect Unauthorized (401)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid or expired token")
}

// TestCreateUserFailMissingFields ensures validation works when fields are missing
func TestCreateUserFailMissingFields(t *testing.T) {
	SetupTestDB()
	router := SetupTestRouter()

	adminToken, _ := utils.GenerateToken("adminuser", []string{"admin"})

	// Missing email field
	userData := map[string]interface{}{
		"username": "testuser",
		"password": "password123",
		"roles":    []string{"user"},
	}
	jsonData, _ := json.Marshal(userData)

	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+adminToken)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Expect Bad Request (400)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Missing username, email, password or role")
}

// TestCreateUserFailUserExists ensures that a user with the same username or email cannot be registered
func TestCreateUserFailUserExists(t *testing.T) {
	SetupTestDB()
	router := SetupTestRouter()

	// Insert an existing user into the database
	adminToken, _ := utils.GenerateToken("adminuser", []string{"admin"})
	existingUser := map[string]interface{}{
		"username": "existinguser",
		"email":    "existinguser@example.com",
		"password": "password123",
		"roles":    []string{"user"},
	}
	existingUserData, _ := json.Marshal(existingUser)

	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(existingUserData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+adminToken)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Now try registering the same user
	newUser := map[string]interface{}{
		"username": "existinguser",
		"email":    "existinguser@example.com",
		"password": "newpassword123",
		"roles":    []string{"user"},
	}
	newUserData, _ := json.Marshal(newUser)

	req, _ = http.NewRequest("POST", "/users", bytes.NewBuffer(newUserData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+adminToken)

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Expect Bad Request (400) with "User already exists" message
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "User already exists")
}

// TestCreateUserFailInvalidEmail ensures the email format is valid
func TestCreateUserFailInvalidEmail(t *testing.T) {
	SetupTestDB()
	router := SetupTestRouter()

	adminToken, _ := utils.GenerateToken("adminuser", []string{"admin"})

	// Invalid email format
	userData := map[string]interface{}{
		"username": "testuser3",
		"email":    "invalid-email",
		"password": "password123",
		"roles":    []string{"user"},
	}
	jsonData, _ := json.Marshal(userData)

	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+adminToken)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Expect Bad Request (400)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid email format")
}
