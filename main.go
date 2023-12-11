package main

import (
	"flag"
	"golang-basic/configs"
	"golang-basic/modules/orders"
	"golang-basic/modules/products"
	"golang-basic/seeds"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}
}

func main() {

	handleInit()
}

func handleInit() {

	db, err := configs.InitDB()
	if err != nil {
		panic(err)
	}

	handleSeeder(db)

	handleRouting(db)

}

func handleSeeder(db *gorm.DB) {
	flag.Parse()
	args := flag.Args()

	if len(args) >= 1 {
		switch args[0] {
		case "seed":
			seeds.Execute(db, args[1:]...)
			os.Exit(0)
		}
	}
}

func handleRouting(db *gorm.DB) {
	r := setupGin()

	productGroup(db, r)
	customerGroup(r)
	orderGroup(db, r)

	r.Run()
}

func setupGin() *gin.Engine {
	return gin.Default()
}

func productGroup(db *gorm.DB, r *gin.Engine) {
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

func customerGroup(r *gin.Engine) {
	customerRoutes := r.Group("v1/customers")

	{
		customerRoutes.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, &gin.H{"customer": "customer"})
		})
	}
}

func orderGroup(db *gorm.DB, r *gin.Engine) {
	orderRepo := orders.NewOrderRepository(db)
	orderService := orders.NewOrderService(orderRepo)
	orderHandler := orders.NewOrderHandler(orderService)

	customerRoutes := r.Group("api/v1/orders")

	{
		customerRoutes.POST("/", orderHandler.CreateOrder)
		customerRoutes.GET("/", orderHandler.GetAllOrder)
		customerRoutes.GET("/items", orderHandler.GetAllOrderItems)
	}
}
