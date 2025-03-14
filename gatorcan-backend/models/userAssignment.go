package models

import (
	"gorm.io/gorm"
)

/*
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	path TEXT NOT NULL,
	user_id INTEGER NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
*/

type UserAssignment struct {
	gorm.Model
	UserId uint   `gorm:"not null"`
	Path   string `gorm:"unique;not null"`
}
