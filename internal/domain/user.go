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
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// type UserRole struct {
// 	ID       uint   `gorm:"primaryKey"`
// 	UserID   uint   `gorm:"column:user_id"`
// 	RoleName string `gorm:"column:role_name"`
// }

type UserRepository interface {
	List() ([]User, error)
	FindByEmail(email string) (*User, error)
	FindByUsername(username string) (*User, error)
	FindByEmailOrUsername(username string) (*User, error)
	FindByID(id uint) (*User, error)
	Save(user *User) (*User, error)
	Update(user *User) (*User, error)
	Delete(*User) error
}
