package models

import "gorm.io/gorm"

type EnrollmentRequestStatus string

const (
	Pending  EnrollmentRequestStatus = "pending"
	Approved EnrollmentRequestStatus = "approved"
	Rejected EnrollmentRequestStatus = "rejected"
)

type Enrollment struct {
	ID             uint   `gorm:"primary_key"`
	UserID         uint   `gorm:"not null"`
	CourseID       uint   `gorm:"not null"`
	Status         string `gorm:"default:'pending';not null"` // 'pending', 'approved', or 'rejected'
	EnrollmentDate string `gorm:"default:current_timestamp"`
	ApprovalDate   string
}

func (e *Enrollment) Create(db *gorm.DB) error {
	return db.Create(e).Error
}
