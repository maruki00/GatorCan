package repositories

import (
	"errors"
	"gatorcan-backend/database"
	"gatorcan-backend/models"
)

type CourseRepository interface {
	GetEnrolledCourses(userID int) ([]models.Enrollment, error)
	GetCourses(page, pageSize int) ([]models.Course, error)
	GetCourseByID(courseID int) (models.Course, error)
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

func (r *courseRepository) GetCourseByID(courseID int) (models.Course, error) {
	var course models.Course
	if err := database.DB.Preload("ActiveCourse").First(&course, courseID).Error; err != nil {
		return models.Course{}, errors.New("course not found")
	}

	return course, nil
}

func (r *courseRepository) GetCourses(page, pageSize int) ([]models.Course, error) {
	var courses []models.Course
	offset := (page - 1) * pageSize
	if err := database.DB.Limit(pageSize).Offset(offset).Find(&courses).Error; err != nil {
		return nil, errors.New("failed to fetch courses")
	}
	return courses, nil
}
