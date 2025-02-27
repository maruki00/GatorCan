package models

import (
	"gorm.io/gorm"
	"time"
)

type EnrollmentRequestStatus string

const (
	Pending  EnrollmentRequestStatus = "pending"
	Approved EnrollmentRequestStatus = "approved"
	Rejected EnrollmentRequestStatus = "rejected"
)

type Enrollment struct {
	ID             uint                    `gorm:"primary_key"`
	UserID         uint                    `gorm:"not null"`
	ActiveCourseID uint                    `gorm:"not null"`
	Status         EnrollmentRequestStatus `gorm:"default:'pending';not null"`
	EnrollmentDate time.Time               `gorm:"autoCreateTime"`
	ApprovalDate   time.Time
	ActiveCourse   ActiveCourse `gorm:"foreignKey:ActiveCourseID"`
	User           User         `gorm:"foreignKey:UserID"`
}

func (e *Enrollment) Create(db *gorm.DB) error {
	return db.Create(e).Error
}
