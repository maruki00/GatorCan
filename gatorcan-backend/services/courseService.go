package services

import (
	"errors"
	dtos "gatorcan-backend/DTOs"
	"gatorcan-backend/models"
	"gatorcan-backend/repositories"
	"log"
)

func GetEnrolledCourses(logger *log.Logger, username string) ([]dtos.EnrolledCoursesResponseDTO, error) {

	user, err := repositories.NewUserRepository().GetUserByUsername(username)
	if err != nil {
		logger.Printf("user not found: %s %d", username, 404)
		return nil, errors.New("user not found")
	}

	enrollments, err := repositories.NewCourseRepository().GetEnrolledCourses(int(user.ID))
	if err != nil {
		logger.Printf("failed to fetch enrolled courses: %s %d", username, 500)
		return nil, errors.New("failed to fetch enrolled courses")
	}

	var enrolledCourses []dtos.EnrolledCoursesResponseDTO
	for _, enrollment := range enrollments {
		var course dtos.EnrolledCoursesResponseDTO
		course.ID = enrollment.ActiveCourse.Course.ID
		course.Name = enrollment.ActiveCourse.Course.Name
		course.Description = enrollment.ActiveCourse.Course.Description
		course.StartDate = enrollment.ActiveCourse.StartDate
		course.EndDate = enrollment.ActiveCourse.EndDate
		var instructor *models.User
		instructor, err = repositories.NewUserRepository().GetUserByID(enrollment.ActiveCourse.InstructorID)
		course.InstructorName = instructor.Username
		course.InstructorEmail = instructor.Email
		enrolledCourses = append(enrolledCourses, course)
	}

	return enrolledCourses, nil
}

func GetCourses(logger *log.Logger, username string, page int, pageSize int) ([]dtos.CourseResponseDTO, error) {

	_, err := repositories.NewUserRepository().GetUserByUsername(username)
	if err != nil {
		logger.Printf("user not found: %s %d", username, 404)
		return nil, errors.New("user not found")
	}
	// Fetch courses using pagination
	courses, err := repositories.NewCourseRepository().GetCourses(page, pageSize)
	if err != nil {
		// Log the error details along with the page parameters
		logger.Printf("Failed to fetch courses for page %d with pageSize %d: %v", page, pageSize, err)
		return nil, errors.New("failed to fetch courses")
	}

	// Convert models.Course to dtos.CourseResponseDTO
	var courseResponseDTOs []dtos.CourseResponseDTO
	for _, course := range courses {
		dto := dtos.CourseResponseDTO{
			ID:          course.ID,
			Name:        course.Name,
			Description: course.Description,
			StartDate:   course.StartDate,
			EndDate:     course.EndDate,
		}
		courseResponseDTOs = append(courseResponseDTOs, dto)
	}

	return courseResponseDTOs, nil
}
