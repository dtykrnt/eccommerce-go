package models

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	ID          uint   `json:"id"`
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	IsActive    bool   `json:"is_active" gorm:"default:true"`
	Password    []byte `json:"-"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
