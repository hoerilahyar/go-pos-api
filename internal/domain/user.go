package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Username  string         `gorm:"size:100;unique;not null" json:"username"`
	Name      string         `gorm:"size:100" json:"name"`
	Email     string         `gorm:"size:100;unique;not null" json:"email"`
	Password  string         `gorm:"size:255;not null" json:"password"`
	Roles     []Role         `gorm:"many2many:user_roles;" json:"roles,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type AuthUsecase interface {
	Register(user *User) error
	Login(username, password string) (*TokenResponse, error)
	GetProfile(userID uint) (*User, error)
}

type UserRepository interface {
	FindByEmail(email string) (*User, error)
	FindByEmailOrUsername(username string) (*User, error)
	FindByID(id uint) (*User, error)
	Save(user *User) error
	Update(user *User) error
	Delete(id uint) error
}
