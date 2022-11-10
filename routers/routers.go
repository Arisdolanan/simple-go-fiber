package routers

import (
	"github.com/gofiber/fiber/v2"
	"simple-go-fiber-crud/controllers"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}

func SetupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)
	app.Get("/todos", controllers.GetTodos)
	app.Get("/todos/:id", controllers.GetTodoById)
	app.Post("/todos", controllers.CreateTodo)
	app.Put("/todos/:id", controllers.UpdateTodo)
	app.Delete("/todos/:id", controllers.DeleteTodo)
}
