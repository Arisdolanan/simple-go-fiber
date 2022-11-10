package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"simple-go-fiber-crud/models"
)

var (
	DBConn *gorm.DB
)

func InitDatabase() {
	var err error
	dsn := "host=localhost user=postgres password=postgres dbname=gotodo port=5432"
	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database@")
	}
	fmt.Println("Database connected")
	DBConn.AutoMigrate(&models.Todo{})
	fmt.Println("Migrate DB")
}
