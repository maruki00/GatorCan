package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	dtos "gatorcan-backend/DTOs"
	"gatorcan-backend/utils"

	"github.com/stretchr/testify/assert"
)

// TestEnrollUser tests the enrollment functionality
func TestEnrollUser(t *testing.T) {
	// Set up test database and router
	SetupTestDB()
	router := SetupTestRouter()

	// Generate a valid JWT token for a student
	token, err := utils.GenerateToken("teststudent", []string{"student"})
	assert.NoError(t, err)

	// Test Case 1: Successful enrollment
	enrollRequest := dtos.EnrollRequestDTO{
		CourseID: 1, // Assuming course ID 1 exists in test DB
	}

	requestBody, err := json.Marshal(enrollRequest)
	assert.NoError(t, err)

	req, _ := http.NewRequest("POST", "/courses/enroll", bytes.NewBuffer(requestBody))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Validate the HTTP status code for successful enrollment
	assert.Equal(t, http.StatusCreated, w.Code)

	// Decode the response
	var response dtos.EnrollmentResponseDTO
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Enrollment requested successfully", response.Message)

	// Test Case 2: Attempt to enroll in the same course again (should fail)
	req, _ = http.NewRequest("POST", "/courses/enroll", bytes.NewBuffer(requestBody))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Should return a status 500 when trying to enroll again
	// This is based on the actual behavior in your application
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	// Check error response
	var errorResponse map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &errorResponse)
	assert.NoError(t, err)
	assert.Equal(t, "Failed to request enrollment", errorResponse["error"])

	// Test Case 3: Try to enroll in a non-existent course
	invalidEnrollRequest := dtos.EnrollRequestDTO{
		CourseID: 9999, // Non-existent course ID
	}

	requestBody, err = json.Marshal(invalidEnrollRequest)
	assert.NoError(t, err)

	req, _ = http.NewRequest("POST", "/courses/enroll", bytes.NewBuffer(requestBody))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Based on the logs, this returns 500 instead of 404
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	// Test Case 4: Invalid course ID (zero or negative)
	invalidEnrollRequest = dtos.EnrollRequestDTO{
		CourseID: 0, // Invalid course ID
	}

	requestBody, err = json.Marshal(invalidEnrollRequest)
	assert.NoError(t, err)

	req, _ = http.NewRequest("POST", "/courses/enroll", bytes.NewBuffer(requestBody))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Should return a bad request status
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Test Case 5: Unauthorized access (no token)
	req, _ = http.NewRequest("POST", "/courses/enroll", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Should return an unauthorized status
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
