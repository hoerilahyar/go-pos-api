package domain

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint64         `gorm:"primaryKey;autoIncrement" json:"id"`
	Code        string         `gorm:"unique;not null;size:50" json:"code" binding:"required"`
	Name        string         `gorm:"not null;size:100" json:"name" binding:"required"`
	Description string         `gorm:"type:text" json:"description"`
	CategoryID  *uint64        `json:"category_id,omitempty"`
	Category    *Category      `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	Price       *float64       `gorm:"not null" json:"price" binding:"required"`
	CostPrice   float64        `json:"cost_price"`
	Stock       int            `gorm:"default:0" json:"stock"`
	IsActive    bool           `gorm:"default:true" json:"is_active"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
type ProductUpdate struct {
	ID          uint64    `json:"id"`
	Code        string    `json:"code"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CategoryID  *uint64   `json:"category_id,omitempty"`
	Category    *Category `json:"category,omitempty"`
	Price       *float64  `json:"price"`
	CostPrice   float64   `json:"cost_price"`
	Stock       int       `json:"stock"`
	IsActive    bool      `json:"is_active"`
}

type Category struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"not null;size:100;unique" json:"name" binding:"required"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
