package unit

import (
	"gatorcan-backend/models"
	"gatorcan-backend/services"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockCourseRepository is a mock implementation of the course repository
type MockCourseRepository struct {
	mock.Mock
}

func (m *MockCourseRepository) GetCourses(page int, pageSize int) ([]models.Course, error) {
	args := m.Called(page, pageSize)
	return args.Get(0).([]models.Course), args.Error(1)
}

func TestGetCoursesService(t *testing.T) {
	mockRepo := new(MockCourseRepository)
	courseService := services.NewCourseService(mockRepo)

	tests := []struct {
		name          string
		page          int
		pageSize      int
		mockCourses   []models.Course
		mockError     error
		expectedCount int
		expectError   bool
	}{
		{
			name:     "Success - Full Page",
			page:     1,
			pageSize: 20,
			mockCourses: []models.Course{
				{ID: 1, Name: "Course 1", Description: "Description 1"},
				{ID: 2, Name: "Course 2", Description: "Description 2"},
			},
			mockError:     nil,
			expectedCount: 2,
			expectError:   false,
		},
		{
			name:          "Success - Empty Page",
			page:          2,
			pageSize:      20,
			mockCourses:   []models.Course{},
			mockError:     nil,
			expectedCount: 0,
			expectError:   false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			logger := &log.Logger{}
			// Setup mock expectations
			mockRepo.On("GetCourses", logger, "", tc.page, tc.pageSize).Return(tc.mockCourses, tc.mockError)

			// Call the service
			courses, err := courseService.GetCourses(logger, "", tc.page, tc.pageSize)

			// Assertions
			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedCount, len(courses))

				// Verify the DTO conversion
				if len(courses) > 0 {
					assert.Equal(t, tc.mockCourses[0].Name, courses[0].Name)
					assert.Equal(t, tc.mockCourses[0].Description, courses[0].Description)
				}
			}

			// Verify that mock expectations were met
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestConvertToCourseDTO(t *testing.T) {
	course := models.Course{
		ID:          1,
		Name:        "Test Course",
		Description: "Test Description",
	}

	dto := models.Course{
		ID:          1,
		Name:        "Test Course",
		Description: "Test Description",
	}

	assert.Equal(t, course.ID, dto.ID)
	assert.Equal(t, course.Name, dto.Name)
	assert.Equal(t, course.Description, dto.Description)
}
