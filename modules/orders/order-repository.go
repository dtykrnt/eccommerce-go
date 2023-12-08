package orders

import "gorm.io/gorm"

type IOrderRepository interface {
	GetAllOrder()
	GetOrderById()
	UpdateOrder()
	CreateOrder()
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *orderRepository {
	return &orderRepository{db}
}

