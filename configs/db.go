package configs

import (
	"fmt"
	"golang-basic/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, dbUser, dbName, dbPass)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	fmt.Println("Connection Database Success ...")

	db.AutoMigrate(&models.Products{})
	db.AutoMigrate(&models.Orders{})
	db.AutoMigrate(&models.Customers{})
	db.AutoMigrate(&models.OrderItems{})
	db.AutoMigrate(&models.Users{})

	fmt.Println("Database Migrate Success ...")

	return db, nil
}
