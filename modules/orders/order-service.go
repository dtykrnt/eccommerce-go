package orders

import (
	"context"
	"golang-basic/requests"
)

type IOrderService interface {
	GetAllOrder(ctx context.Context) ([]Orders, error)
	GetAllOrderItems(ctx context.Context) ([]OrderItems, error)
	GetOrderById(ctx context.Context, order Orders) (*Orders, error)
	UpdateOrder(ctx context.Context, order Orders) (*Orders, error)
	CreateOrder(ctx context.Context, order requests.CreateOrderRequest) (*Orders, error)
	DeleteOrder(ctx context.Context, order Orders) (any, error)
}

type orderService struct {
	orderRepository IOrderRepository
}

// DeleteOrder implements IOrderService.
func (*orderService) DeleteOrder(ctx context.Context, order Orders) (any, error) {
	panic("unimplemented")
}

// GetOrderById implements IOrderService.
func (*orderService) GetOrderById(ctx context.Context, order Orders) (*Orders, error) {
	panic("unimplemented")
}

// UpdateOrder implements IOrderService.
func (*orderService) UpdateOrder(ctx context.Context, order Orders) (*Orders, error) {
	panic("unimplemented")
}

func NewOrderService(orderRepository IOrderRepository) *orderService {
	return &orderService{
		orderRepository: orderRepository,
	}
}

func (o *orderService) GetAllOrder(ctx context.Context) ([]Orders, error) {
	return o.orderRepository.GetAllOrder(ctx)
}

func (o *orderService) CreateOrder(ctx context.Context, order requests.CreateOrderRequest) (*Orders, error) {
	return o.orderRepository.CreateOrder(ctx, order)
}

func (o *orderService) GetAllOrderItems(ctx context.Context) ([]OrderItems, error) {
	return o.orderRepository.GetAllOrderItems(ctx)
}
