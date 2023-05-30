package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"simple-go-fiber-crud/config"
	"simple-go-fiber-crud/models"
)

var DBConn *gorm.DB

func InitDatabase() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		config.POSTGRES_HOST,
		config.POSTGRES_USERNAME,
		config.POSTGRES_PASSWORD,
		config.POSTGRES_DBNAME,
		config.POSTGRES_PORT,
	)
	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database@")
	}
	fmt.Println("Database connected")
	DBConn.AutoMigrate(&models.Todo{})
	fmt.Println("Migrate DB")
}
