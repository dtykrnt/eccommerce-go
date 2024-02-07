package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupGin() *gin.Engine {
	return gin.Default()
}

func Route(db *gorm.DB) {
	r := setupGin()
	r.Static("/assets", "./assets")

	ProductGroup(db, r)
	CustomersGroup(r)
	OrdersGroup(db, r)

	r.Run(":3000")
}
