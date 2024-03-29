package models

import (
	"time"

	"gorm.io/gorm"
)

type OrderStatus string

const (
	Pending    OrderStatus = "pending"
	Processing OrderStatus = "processing"
	Shipped    OrderStatus = "shipped"
	Delivered  OrderStatus = "delivered"
	Cancelled  OrderStatus = "cancelled"
)

type Orders struct {
	ID              uint `json:"id"`
	CustomerID      uint `json:"customer_id"`
	OrderDate       time.Time
	OrderStatus     OrderStatus `json:"order_status"`
	TotalPrice      float64     `json:"total_price"`
	ShippingAddress string      `json:"shipping_addres"`
	BillingAddress  string      `json:"billing_address"`
	PaymentMethod   string      `json:"payment_method"`
	PaymentStatus   string      `json:"payment_status"`
	DeliveryMethod  string      `json:"delivery_method"`
	Discount        string      `json:"discount"`
	Notes           string      `json:"notes"`
	TrackingNumber  string      `json:"tracking_number"`
	PromoCode       string      `json:"promo_code"`

	CreatedAt time.Time
	UpdatedAt time.Time
	Items     []*OrderItems `json:"items"`
}

type OrderItems struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement;"`
	OrderID   uint      `json:"orders_id" gorm:"foreignKey;column:orders_id;" validate:"required"`
	ProductID uint      `json:"products_id" gorm:"foreignKey;column:products_id;" validate:"required"`
	Quantity  int       `json:"quantity" gorm:"column:quantity;" validate:"required"`
	Price     float64   `json:"price" validate:"required"`
	Products  *Products `json:"products"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
