package repositories

import (
	"errors"
	"gatorcan-backend/database"
	"gatorcan-backend/models"
)

type CourseRepository interface {
	GetEnrolledCourses(userID int) ([]models.Enrollment, error)
}

type courseRepository struct {
}

func NewCourseRepository() CourseRepository {
	return &courseRepository{}
}

func (r *courseRepository) GetEnrolledCourses(userID int) ([]models.Enrollment, error) {
	var enrollments []models.Enrollment
	if err := database.DB.Preload("ActiveCourse.Course").Where("user_id = ?", userID).Find(&enrollments).Error; err != nil {
		return nil, errors.New("failed to fetch enrolled courses")
	}

	return enrollments, nil
}
