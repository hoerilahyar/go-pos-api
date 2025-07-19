package domain

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"unique"`
	Users     []User `gorm:"many2many:user_roles;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
