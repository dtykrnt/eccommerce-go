package models

import "time"

type Products struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Description string  `json:"description" binding:"required"`
	IsActive    bool    `json:"is_active" gorm:"default:true"`
	Stocks      int     `json:"stocks"`
	Image       string  `json:"image"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
