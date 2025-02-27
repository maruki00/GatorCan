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
