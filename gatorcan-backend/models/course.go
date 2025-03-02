package models

import (
	"time"

	"gorm.io/gorm"
)

type Course struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Description string    `json:"description"`
	Capacity    int       `gorm:"not null" json:"capacity"`
	Enrolled    int       `gorm:"default:0" json:"enrolled"`
	StartDate   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"start_date"`
	EndDate     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"end_date"`
}

func (c *Course) Create(db *gorm.DB) error {
	return db.Create(c).Error
}
