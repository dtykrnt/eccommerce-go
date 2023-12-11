package orders

import (
	"golang-basic/requests"
	"golang-basic/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IOrderHandler interface {
	CreateOrder(c *gin.Context)
	GetAllOrder(c *gin.Context)
	GetAllOrderItems(c *gin.Context)
	UpdateOrder(c *gin.Context)
	DeleteOrder(c *gin.Context)
}

type orderHandler struct {
	orderService IOrderService
}

func NewOrderHandler(orderService IOrderService) *orderHandler {
	return &orderHandler{orderService}
}

func (h *orderHandler) CreateOrder(c *gin.Context) {
	var order requests.CreateOrderRequest
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, responses.NewErrorResponse(err.Error()))
		return
	}

	createdProduct, err := h.orderService.CreateOrder(c.Request.Context(), order)

	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, responses.NewSuccessResponse("Success Create Products", createdProduct))

}

func (h *orderHandler) GetAllOrder(c *gin.Context) {
	order, err := h.orderService.GetAllOrder(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.NewErrorResponse(err.Error()))
	}
	c.JSON(http.StatusAccepted, responses.NewSuccessResponse("Success get Orders", order))
}

func (h *orderHandler) GetAllOrderItems(c *gin.Context) {
	order, err := h.orderService.GetAllOrderItems(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.NewErrorResponse(err.Error()))
	}
	c.JSON(http.StatusAccepted, responses.NewSuccessResponse("Success get Orders", order))
}
