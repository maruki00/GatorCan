package dtos

import (
	"gatorcan-backend/models"
	"time"
)

type EnrolledCoursesResponseDTO struct {
	Name            string
	ID              uint
	StartDate       time.Time
	EndDate         time.Time
	Description     string
	InstructorName  string
	InstructorEmail string
}

type CourseResponseDTO struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Convert model courses to DTOs
func ConvertToCourseResponseDTOs(courses []models.Course) []CourseResponseDTO {
	courseResponseDTOs := make([]CourseResponseDTO, len(courses))
	for i, course := range courses {
		courseResponseDTOs[i] = CourseResponseDTO{
			ID:          course.ID,
			Name:        course.Name,
			Description: course.Description,
			CreatedAt:   course.CreatedAt,
			UpdatedAt:   course.UpdatedAt,
		}
	}
	return courseResponseDTOs
}

type EnrollRequestDTO struct {
	CourseID int `json:"courseID"` // The courseID field that will be sent in the request body
}

type EnrollmentResponseDTO struct {
	Message string `json:"message"` // Success message
}
