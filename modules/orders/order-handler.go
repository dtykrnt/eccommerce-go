package orders

type IOrderHandler interface {
	CreateOrder()
	GetAllOrder()
	UpdateOrder()
	DeleteOrder()
}
