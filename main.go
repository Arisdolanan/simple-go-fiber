package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"simple-go-fiber-crud/database"
	"simple-go-fiber-crud/routers"
	"time"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	database.InitDatabase()

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cache.New())

	// Or extend your config for customization
	app.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Query("refresh") == "true"
		},
		Expiration:   30 * time.Minute,
		CacheControl: true,
	}))

	routers.SetupRoutes(app)
	app.Listen(":8000")
}
