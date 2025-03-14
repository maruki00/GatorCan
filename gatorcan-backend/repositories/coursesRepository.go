package repositories

import (
	"context"
	"errors"
	"gatorcan-backend/models"
	"time"

	"gorm.io/gorm"
)

type CourseRepository interface {
	GetEnrolledCourses(ctx context.Context, userID int) ([]models.Enrollment, error)
	GetCourses(ctx context.Context, page, pageSize int) ([]models.Course, error)
	GetCourseByID(ctx context.Context, courseID int) (models.ActiveCourse, error)
	RequestEnrollment(ctx context.Context, userID, activeCourseID uint) error
	ApproveEnrollment(ctx context.Context, enrollmentID uint) error
	RejectEnrollment(ctx context.Context, enrollmentID uint) error
	GetPendingEnrollments(ctx context.Context) ([]models.Enrollment, error)
}

type courseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &courseRepository{db: db}
}

func (r *courseRepository) GetEnrolledCourses(ctx context.Context, userID int) ([]models.Enrollment, error) {
	var enrollments []models.Enrollment
	if err := r.db.WithContext(ctx).
		Preload("ActiveCourse.Course").
		Where("user_id = ?", userID).
		Find(&enrollments).Error; err != nil {
		return nil, errors.New("failed to fetch enrolled courses")
	}

	return enrollments, nil
}

func (r *courseRepository) GetCourseByID(ctx context.Context, courseID int) (models.ActiveCourse, error) {
	var course models.ActiveCourse
	// Fetch course from courses table without loading related ActiveCourse
	if err := r.db.WithContext(ctx).First(&course, courseID).Error; err != nil {
		return models.ActiveCourse{}, errors.New("course not found")
	}

	return course, nil
}

func (r *courseRepository) GetCourses(ctx context.Context, page, pageSize int) ([]models.Course, error) {
	var courses []models.Course
	offset := (page - 1) * pageSize
	if err := r.db.WithContext(ctx).
		Limit(pageSize).
		Offset(offset).
		Find(&courses).Error; err != nil {
		return nil, errors.New("failed to fetch courses")
	}
	return courses, nil
}

// RequestEnrollment - User requests enrollment (Pending Approval)
func (r *courseRepository) RequestEnrollment(ctx context.Context, userID, CourseID uint) error {
	var existingEnrollment models.Enrollment
	if err := r.db.WithContext(ctx).
		Where("user_id = ? AND active_course_id = ?", userID, CourseID).
		First(&existingEnrollment).Error; err == nil {
		return errors.New("enrollment request already exists")
	}

	enrollment := models.Enrollment{
		UserID:         userID,
		ActiveCourseID: CourseID,
		Status:         models.Pending, // Default status
		EnrollmentDate: time.Now(),
	}

	if err := r.db.WithContext(ctx).
		Create(&enrollment).Error; err != nil {
		return errors.New("failed to create enrollment request")
	}

	return nil
}

// ApproveEnrollment - Admin approves a pending enrollment
func (r *courseRepository) ApproveEnrollment(ctx context.Context, enrollmentID uint) error {
	var enrollment models.Enrollment
	if err := r.db.WithContext(ctx).
		First(&enrollment, enrollmentID).Error; err != nil {
		return errors.New("enrollment request not found")
	}

	if enrollment.Status != models.Pending {
		return errors.New("only pending enrollments can be approved")
	}

	enrollment.Status = models.Approved
	enrollment.ApprovalDate = time.Now()
	enrollment.EnrollmentDate = time.Now()

	if err := r.db.WithContext(ctx).
		Save(&enrollment).Error; err != nil {
		return errors.New("failed to approve enrollment")
	}

	return nil
}

// RejectEnrollment - Admin rejects a pending enrollment
func (r *courseRepository) RejectEnrollment(ctx context.Context, enrollmentID uint) error {
	if err := r.db.WithContext(ctx).
		Delete(&models.Enrollment{}, enrollmentID).Error; err != nil {
		return errors.New("failed to reject enrollment")
	}
	return nil
}

// GetPendingEnrollments - Fetch all pending enrollments for admin approval
func (r *courseRepository) GetPendingEnrollments(ctx context.Context) ([]models.Enrollment, error) {
	var enrollments []models.Enrollment
	if err := r.db.WithContext(ctx).
		Where("status = ?", models.Pending).
		Find(&enrollments).Error; err != nil {
		return nil, errors.New("failed to fetch pending enrollments")
	}
	return enrollments, nil
}
