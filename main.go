package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
	"simple-go-fiber-crud/database"
	"simple-go-fiber-crud/routers"
	"strconv"
	"time"
)

func main() {
	database.InitDatabase()

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin,Authorization",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	app.Use(recover.New())
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "Hello World",
		})
	})
	app.Get("/health", func(c *fiber.Ctx) error {
		var startTime time.Time
		return c.JSON(map[string]string{"status": strconv.Itoa(fiber.StatusOK), "Route": c.Path(), "message": "Server is Up", "uptime": time.Since(startTime).String()})
	})

	routers.SetupRoutes(app)
	log.Fatal(app.Listen(fmt.Sprintf(":8000")))
}
