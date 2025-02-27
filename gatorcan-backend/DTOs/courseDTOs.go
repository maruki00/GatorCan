package dtos

type EnrolledCoursesResponseDTO struct {
	CourseName        string
	CourseID          uint
	CourseStartDate   string
	CourseEndDate     string
	CourseDescription string
	InstructorName    string
}
