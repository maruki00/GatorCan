package models

import "gorm.io/gorm"

// ActiveCourse represents a course that is currently being taught.
type ActiveCourse struct {
	ID           uint   `gorm:"primary_key"`
	InstructorID uint   `gorm:"not null"`
	CourseID     uint   `gorm:"not null"`
	StartDate    string `gorm:"not null"`
	EndDate      string `gorm:"not null"`
	CreatedAt    string `gorm:"default:current_timestamp"`
	UpdatedAt    string `gorm:"default:current_timestamp"`
	IsActive     bool   `gorm:"default:true"`
	Instructor   User   `gorm:"foreignkey:InstructorID"`
	Course       Course `gorm:"foreignkey:CourseID"`
}

func (c *ActiveCourse) Create(db *gorm.DB) error {
	return db.Create(c).Error
}
