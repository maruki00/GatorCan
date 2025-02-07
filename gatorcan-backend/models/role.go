package models

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name  userRole `gorm:"unique;not null"`
	Users []*User  `gorm:"many2many:user_roles;"`
}
