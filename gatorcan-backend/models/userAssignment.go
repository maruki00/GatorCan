package models

import (
	"gorm.io/gorm"
)

/*
	create table user_assignments (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		path TEXT NOT NULL,
		user_id INTEGER NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		deleted_at TIMESTAMP DEFAULT NULL
	);
*/

type UserAssignment struct {
	gorm.Model
	UserId uint   `gorm:"not null"`
	Path   string `gorm:"unique;not null"`
}
