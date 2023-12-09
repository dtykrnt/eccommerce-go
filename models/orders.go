package models

import (
	"time"
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
	Items     []Products `json:"items" gorm:"many2many:order_items;foreignkey:id;association_foreignkey:products_id;"`
}

type OrderItems struct {
	ID        uint    `json:"id" gorm:"primaryKey;autoIncrement;"`
	OrderID   uint    `json:"orders_id" gorm:"foreignKey;column:orders_id;"`
	ProductID uint    `json:"products_id" gorm:"foreignKey;column:products_id;"`
	Quantity  int     `json:"quantity" gorm:"column:quantity;"`
	Price     float64 `json:"price"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
