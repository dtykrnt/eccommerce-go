package routes

import (
	"golang-basic/modules/products"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ProductGroup(db *gorm.DB, r *gin.Engine) {
	productRepo := products.NewProductRepository(db)
	productService := products.NewProductService(productRepo)
	productHandler := products.NewProductHandler(productService)

	productRoutes := r.Group("api/v1/products")
	{
		productRoutes.POST("/", productHandler.CreateProduct)
		productRoutes.GET("/", productHandler.GetAllProducts)
		productRoutes.GET("/:id", productHandler.GetProductById)
		productRoutes.PUT("/:id", productHandler.UpdateProductById)
		productRoutes.DELETE("/:id", productHandler.DeleteProductById)

	}
}
