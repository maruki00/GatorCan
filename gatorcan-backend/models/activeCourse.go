package models

import (
	"gorm.io/gorm"
	"time"
)

// ActiveCourse represents a course that is currently being taught.
type ActiveCourse struct {
	ID           uint      `gorm:"primary_key"`
	InstructorID uint      `gorm:"not null"`
	CourseID     uint      `gorm:"not null"`
	StartDate    time.Time `gorm:"not null"`
	EndDate      time.Time `gorm:"not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
	IsActive     bool      `gorm:"default:true"`
	Instructor   User      `gorm:"foreignKey:InstructorID"`
	Course       Course    `gorm:"foreignKey:CourseID"`
}

func (c *ActiveCourse) Create(db *gorm.DB) error {
	return db.Create(c).Error
}
