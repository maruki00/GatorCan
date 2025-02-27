package services

import (
	"errors"
	"gatorcan-backend/DTOs"
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
