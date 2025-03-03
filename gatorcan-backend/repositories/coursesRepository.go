package repositories

import (
	"errors"
	"gatorcan-backend/database"
	"gatorcan-backend/models"
	"time"
)

type CourseRepository interface {
	GetEnrolledCourses(userID int) ([]models.Enrollment, error)
	GetCourses(page, pageSize int) ([]models.Course, error)
	GetCourseByID(courseID int) (models.ActiveCourse, error)
	RequestEnrollment(userID, activeCourseID uint) error
	ApproveEnrollment(enrollmentID uint) error
	RejectEnrollment(enrollmentID uint) error
	GetPendingEnrollments() ([]models.Enrollment, error)
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

func (r *courseRepository) GetCourseByID(courseID int) (models.ActiveCourse, error) {
	var course models.ActiveCourse
	// Fetch course from courses table without loading related ActiveCourse
	if err := database.DB.First(&course, courseID).Error; err != nil {
		return models.ActiveCourse{}, errors.New("course not found")
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

// RequestEnrollment - User requests enrollment (Pending Approval)
func (r *courseRepository) RequestEnrollment(userID, CourseID uint) error {
	var existingEnrollment models.Enrollment
	if err := database.DB.Where("user_id = ? AND active_course_id = ?", userID, CourseID).First(&existingEnrollment).Error; err == nil {
		return errors.New("enrollment request already exists")
	}

	enrollment := models.Enrollment{
		UserID:         userID,
		ActiveCourseID: CourseID,
		Status:         models.Pending, // Default status
		EnrollmentDate: time.Now(),
	}

	if err := database.DB.Create(&enrollment).Error; err != nil {
		return errors.New("failed to create enrollment request")
	}

	return nil
}

// ApproveEnrollment - Admin approves a pending enrollment
func (r *courseRepository) ApproveEnrollment(enrollmentID uint) error {
	var enrollment models.Enrollment
	if err := database.DB.First(&enrollment, enrollmentID).Error; err != nil {
		return errors.New("enrollment request not found")
	}

	if enrollment.Status != models.Pending {
		return errors.New("only pending enrollments can be approved")
	}

	enrollment.Status = models.Approved
	enrollment.ApprovalDate = time.Now()
	enrollment.EnrollmentDate = time.Now()

	if err := database.DB.Save(&enrollment).Error; err != nil {
		return errors.New("failed to approve enrollment")
	}

	return nil
}

// RejectEnrollment - Admin rejects a pending enrollment
func (r *courseRepository) RejectEnrollment(enrollmentID uint) error {
	if err := database.DB.Delete(&models.Enrollment{}, enrollmentID).Error; err != nil {
		return errors.New("failed to reject enrollment")
	}
	return nil
}

// GetPendingEnrollments - Fetch all pending enrollments for admin approval
func (r *courseRepository) GetPendingEnrollments() ([]models.Enrollment, error) {
	var enrollments []models.Enrollment
	if err := database.DB.Where("status = ?", models.Pending).Find(&enrollments).Error; err != nil {
		return nil, errors.New("failed to fetch pending enrollments")
	}
	return enrollments, nil
}
