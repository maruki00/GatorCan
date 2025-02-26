package models

import "gorm.io/gorm"

type Course struct {
	ID          uint   `gorm:"primary_key"`
	Name        string `gorm:"not null"`
	Description string
	CreatedAt   string `gorm:"default:current_timestamp"`
	UpdatedAt   string `gorm:"default:current_timestamp"`
}

func (c *Course) Create(db *gorm.DB) error {
	return db.Create(c).Error
}
