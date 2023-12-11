package orders

import (
	"context"
	"fmt"
	"golang-basic/models"
	"golang-basic/requests"
	"time"

	"gorm.io/gorm"
)

type Orders models.Orders
type OrderItems models.OrderItems
type IOrderRepository interface {
	GetAllOrder(ctx context.Context) ([]Orders, error)
	GetAllOrderItems(ctx context.Context) ([]OrderItems, error)
	GetOrderById(ctx context.Context, order Orders) (*Orders, error)
	UpdateOrder(ctx context.Context, order Orders) (*Orders, error)
	CreateOrder(ctx context.Context, order requests.CreateOrderRequest) (*Orders, error)
	DeleteOrder(ctx context.Context, order Orders) (any, error)
}

type orderRepository struct {
	db *gorm.DB
}

// DeleteOrder implements IOrderRepository.
func (*orderRepository) DeleteOrder(ctx context.Context, order Orders) (any, error) {
	panic("unimplemented")
}

// GetOrderById implements IOrderRepository.
func (*orderRepository) GetOrderById(ctx context.Context, order Orders) (*Orders, error) {
	panic("unimplemented")
}

// UpdateOrder implements IOrderRepository.
func (*orderRepository) UpdateOrder(ctx context.Context, order Orders) (*Orders, error) {
	panic("unimplemented")
}

func NewOrderRepository(db *gorm.DB) *orderRepository {
	return &orderRepository{db}
}

func (p *orderRepository) GetAllOrder(ctx context.Context) ([]Orders, error) {
	var orders []Orders

	if err := p.db.WithContext(ctx).Preload("Items").Preload("Items.Products").Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func (p *orderRepository) GetAllOrderItems(ctx context.Context) ([]OrderItems, error) {
	var orders []OrderItems
	if err := p.db.WithContext(ctx).Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func (p *orderRepository) CreateOrder(ctx context.Context, order requests.CreateOrderRequest) (*Orders, error) {
	tx := p.db.Begin()

	if tx.Error != nil {
		return nil, tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	newOrder := Orders{
		CustomerID:  order.CustomerID,
		OrderDate:   time.Now(),
		OrderStatus: models.Pending,
	}
	if err := tx.Create(&newOrder).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	fmt.Println("SUCCESS CRETE ORDER")
	fmt.Println(&newOrder)

	for _, product := range order.Products {
		orderItem := OrderItems{
			OrderID:   newOrder.ID,
			ProductID: product.ProductID,
			Quantity:  product.Quantity,
		}
		fmt.Println("ORDER ITEM")

		if err := tx.Create(&orderItem).Error; err != nil {
			fmt.Println("FAILED CRETE ORDER ITEM")
			fmt.Println(&orderItem)
			tx.Rollback()
			return nil, err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	if err := p.db.WithContext(ctx).Preload("Items").First(&newOrder, newOrder.ID).Error; err != nil {
		return nil, err
	}

	return &newOrder, nil
}

func (p *orderRepository) createOrderItem(ctx context.Context, order OrderItems) (*OrderItems, error) {

	if err := p.db.WithContext(ctx).Create(&order).Error; err != nil {
		return nil, err
	}

	return &order, nil
}
