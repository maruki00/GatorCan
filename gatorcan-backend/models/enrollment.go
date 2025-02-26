package models

import "gorm.io/gorm"

type Enrollment struct {
	ID             uint   `gorm:"primary_key"`
	UserID         uint   `gorm:"not null"`
	CourseID       uint   `gorm:"not null"`
	Status         string `gorm:"default:'pending';not null"` // 'pending', 'approved', or 'rejected'
	EnrollmentDate string `gorm:"default:current_timestamp"`
	ApprovalDate   string
	User           User   `gorm:"foreignkey:UserID"`
	Course         Course `gorm:"foreignkey:CourseID"`
}

func (e *Enrollment) Create(db *gorm.DB) error {
	return db.Create(e).Error
}
