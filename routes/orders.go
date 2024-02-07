package routes

import (
	"golang-basic/modules/orders"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func OrdersGroup(db *gorm.DB, r *gin.Engine) {
	ordersRepo := orders.NewOrderRepository(db)
	orderService := orders.NewOrderService(ordersRepo)
	orderHandler := orders.NewOrderHandler(orderService)

	orderRoutes := r.Group("api/v1/orders")

	{
		orderRoutes.GET("/", orderHandler.GetAllOrder)
		orderRoutes.POST("/", orderHandler.CreateOrder)
		orderRoutes.GET("/items/", orderHandler.GetAllOrderItems)
	}
}
