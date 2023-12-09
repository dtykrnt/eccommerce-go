package requests

type CreateOrderRequest struct {
	CustomerID uint `json:"customer_id"`
	Products   []struct {
		ProductID uint    `json:"product_id"`
		Quantity  int     `json:"quantity"`
		Price     float64 `json:"price"`
	} `json:"products"`
}
