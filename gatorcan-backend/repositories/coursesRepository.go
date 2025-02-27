package repositories

import (
	"errors"
	"gatorcan-backend/database"
	"gatorcan-backend/models"
)

type CourseRepository interface {
	// Interface Method Declarations
	GetEnrolledCourses(userID int) ([]models.Enrollment, error)
}

type courseRepository struct {
	//db *database.Database
}

func NewCourseRepository() CourseRepository {
	return &courseRepository{}
}

func (r *courseRepository) GetEnrolledCourses(userID int) ([]models.Enrollment, error) {
	// Fetch enrolled courses
	var enrollments []models.Enrollment
	if err := database.DB.Preload("ActiveCourse.Course").Where("user_id = ?", userID).Find(&enrollments).Error; err != nil {
		return nil, errors.New("failed to fetch enrolled courses")
	}

	return enrollments, nil
}
