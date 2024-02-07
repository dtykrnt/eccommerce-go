package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CustomersGroup(r *gin.Engine) {
	customerRoutes := r.Group("v1/customers")

	{
		customerRoutes.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, &gin.H{"customer": "customer"})
		})
	}
}
