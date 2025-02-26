package models

import "gorm.io/gorm"

type Course struct {
	ID           uint   `gorm:"primary_key"`
	Name         string `gorm:"not null"`
	Description  string
	InstructorID uint   `gorm:"not null"` // Just store the InstructorID here
	StartDate    string `gorm:"not null"`
	EndDate      string `gorm:"not null"`
	CreatedAt    string `gorm:"default:current_timestamp"`
	UpdatedAt    string `gorm:"default:current_timestamp"`
	Instructor   User   `gorm:"foreignkey:InstructorID"`
}

func (c *Course) Create(db *gorm.DB) error {
	return db.Create(c).Error
}
