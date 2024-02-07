package models

import (
	"time"

	"gorm.io/gorm"
)

type Products struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Description string  `json:"description" binding:"required"`
	IsActive    bool    `json:"is_active" gorm:"default:true"`
	Stocks      int     `json:"stocks"`
	Image       *string `json:"image" gorm:"default:null"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
