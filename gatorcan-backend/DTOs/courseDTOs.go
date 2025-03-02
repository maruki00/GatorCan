package dtos

import "time"

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
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
}

type EnrollRequestDTO struct {
	CourseID int `json:"courseID"` // The courseID field that will be sent in the request body
}

type EnrollmentResponseDTO struct {
	Message string `json:"message"` // Success message
}
