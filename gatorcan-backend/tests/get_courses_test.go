package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	dtos "gatorcan-backend/DTOs"
	"gatorcan-backend/utils"

	"github.com/stretchr/testify/assert"
)

// TestGetCoursesPagination tests the GetCourses endpoint with pagination.
func TestGetCoursesPagination(t *testing.T) {
	// Set up test database and router.
	SetupTestDB()
	router := SetupTestRouter()
	// Generate a valid JWT token for a student.
	token, err := utils.GenerateToken("teststudent", []string{"student"})
	assert.NoError(t, err)

	// ---- Test Page 1 ----
	req, _ := http.NewRequest("GET", "/courses/?page=1&pageSize=20", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Validate the HTTP status code.
	assert.Equal(t, http.StatusOK, w.Code)

	// Decode the response JSON into a slice of CourseResponseDTO.
	var coursesPage1 []dtos.CourseResponseDTO
	err = json.Unmarshal(w.Body.Bytes(), &coursesPage1)
	assert.NoError(t, err)
	// Expect 20 courses on page 1.
	assert.Equal(t, 20, len(coursesPage1), "Expected 20 courses on page 1")

	// ---- Test Page 2 ----
	req, _ = http.NewRequest("GET", "/courses/?page=2&pageSize=20", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var coursesPage2 []dtos.CourseResponseDTO
	err = json.Unmarshal(w.Body.Bytes(), &coursesPage2)
	assert.NoError(t, err)
	// Expect the remaining 10 courses on page 2.
	assert.Equal(t, 10, len(coursesPage2), "Expected 10 courses on page 2")
	CloseTestDB()
}
