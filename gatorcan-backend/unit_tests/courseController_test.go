package unit

import (
	"bytes"
	"encoding/json"
	"errors"
	dtos "gatorcan-backend/DTOs"
	"gatorcan-backend/controllers"

	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockCourseService is a mock implementation of the course service
type MockCourseService struct {
	mock.Mock
}

func (m *MockCourseService) GetEnrolledCourses(logger *log.Logger, username string) ([]dtos.CourseResponseDTO, error) {
	args := m.Called(logger, username)
	return args.Get(0).([]dtos.CourseResponseDTO), args.Error(1)
}

func (m *MockCourseService) GetCourses(logger *log.Logger, username string, page, pageSize int) ([]dtos.CourseResponseDTO, error) {
	args := m.Called(logger, username, page, pageSize)
	return args.Get(0).([]dtos.CourseResponseDTO), args.Error(1)
}

func TestGetEnrolledCourses(t *testing.T) {
	gin.SetMode(gin.TestMode)
	logger := log.New(bytes.NewBuffer([]byte{}), "", log.LstdFlags)

	tests := []struct {
		name           string
		username       string
		mockCourses    []dtos.CourseResponseDTO
		mockError      error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:     "Success",
			username: "testuser",
			mockCourses: []dtos.CourseResponseDTO{
				{ID: 1, Name: "Course 1"},
				{ID: 2, Name: "Course 2"},
			},
			mockError:      nil,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "User Not Found",
			username:       "nonexistent",
			mockCourses:    nil,
			mockError:      errors.New("user not found"),
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"error":"User not found"}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup
			mockService := new(MockCourseService)
			mockService.On("GetEnrolledCourses", logger, tc.username).Return(tc.mockCourses, tc.mockError)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Set("username", tc.username)

			// Execute
			controllers.GetEnrolledCourses(c, logger)

			// Assert
			assert.Equal(t, tc.expectedStatus, w.Code)
			if tc.expectedBody != "" {
				assert.JSONEq(t, tc.expectedBody, w.Body.String())
			}
			mockService.AssertExpectations(t)
		})
	}
}

func TestGetCourses(t *testing.T) {
	gin.SetMode(gin.TestMode)
	logger := log.New(bytes.NewBuffer([]byte{}), "", log.LstdFlags)

	tests := []struct {
		name           string
		username       string
		page           string
		pageSize       string
		mockCourses    []dtos.CourseResponseDTO
		mockError      error
		expectedStatus int
	}{
		{
			name:     "Success",
			username: "testuser",
			page:     "1",
			pageSize: "10",
			mockCourses: []dtos.CourseResponseDTO{
				{ID: 1, Name: "Course 1"},
			},
			mockError:      nil,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Invalid Page",
			username:       "testuser",
			page:           "invalid",
			pageSize:       "10",
			mockCourses:    nil,
			mockError:      nil,
			expectedStatus: http.StatusOK, // Uses default values
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup
			mockService := new(MockCourseService)
			mockService.On("GetCourses", logger, tc.username, mock.Anything, mock.Anything).
				Return(tc.mockCourses, tc.mockError)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Set("username", tc.username)

			// Create a new request
			req := httptest.NewRequest("GET", "/courses", nil)
			// Add query parameters
			q := req.URL.Query()
			q.Add("page", tc.page)
			q.Add("pageSize", tc.pageSize)
			req.URL.RawQuery = q.Encode()
			c.Request = req

			// Execute
			controllers.GetCourses(c, logger)

			// Assert
			assert.Equal(t, tc.expectedStatus, w.Code)
			mockService.AssertExpectations(t)
		})
	}
}

func TestEnrollInCourse(t *testing.T) {
	gin.SetMode(gin.TestMode)
	logger := log.New(bytes.NewBuffer([]byte{}), "", log.LstdFlags)

	tests := []struct {
		name           string
		username       string
		requestBody    dtos.EnrollRequestDTO
		mockError      error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:     "Success",
			username: "testuser",
			requestBody: dtos.EnrollRequestDTO{
				CourseID: 1,
			},
			mockError:      nil,
			expectedStatus: http.StatusCreated,
			expectedBody:   `{"message":"Enrollment requested successfully"}`,
		},
		{
			name:     "Course Not Found",
			username: "testuser",
			requestBody: dtos.EnrollRequestDTO{
				CourseID: 999,
			},
			mockError:      errors.New("course not found"),
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"error":"Course not found"}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Set("username", tc.username)

			// Create request body
			jsonBody, _ := json.Marshal(tc.requestBody)
			c.Request = httptest.NewRequest("POST", "/enroll", bytes.NewBuffer(jsonBody))
			c.Request.Header.Set("Content-Type", "application/json")

			// Execute
			controllers.EnrollInCourse(c, logger)

			// Assert
			assert.Equal(t, tc.expectedStatus, w.Code)
			if tc.expectedBody != "" {
				assert.JSONEq(t, tc.expectedBody, w.Body.String())
			}
		})
	}
}
