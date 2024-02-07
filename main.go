package main

import (
	"flag"
	"golang-basic/configs"
	"golang-basic/modules/orders"
	"golang-basic/routes"
	"golang-basic/seeds"
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

	//Setup Route
	routes.Route(db)

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
