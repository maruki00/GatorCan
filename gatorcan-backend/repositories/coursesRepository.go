package repositories

import (
	"errors"
	"gatorcan-backend/database"
	"gatorcan-backend/models"
)

type CourseRepository interface {
	// Interface Method Declarations
	GetEnrolledCourses(username string) ([]models.Course, error)
}

type courseRepository struct {
	//db *database.Database
}

func NewCourseRepository() CourseRepository {
	return &courseRepository{}
}

func (r *courseRepository) GetEnrolledCourses(username string) ([]models.Course, error) {
	// Fetch user from DB
	var user models.User
	if err := database.DB.Preload("Roles").Where("username = ?", username).First(&user).Error; err != nil {
		//c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return nil, errors.New("User not found")
	}

	// Fetch enrolled courses
	var enrollments []models.Enrollment
	if err := database.DB.Preload("ActiveCourse.Course").Where("user_id = ?", user.ID).Find(&enrollments).Error; err != nil {
		//c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch enrolled courses"})
		return nil, errors.New("Failed to fetch enrolled courses")
	}

	var courses []models.Course
	for _, enrollment := range enrollments {
		courses = append(courses, enrollment.ActiveCourse.Course)
	}

	return courses, nil
}
