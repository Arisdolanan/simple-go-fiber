package routers

import (
	"github.com/gofiber/fiber/v2"
	"simple-go-fiber-crud/controllers"
)

func SetupRoutes(app *fiber.App) {
	todoRoutes := controllers.NewTodoController()
	app.Get("/todos", todoRoutes.GetTodos)
	app.Get("/todos/:id", todoRoutes.GetTodoById)
	app.Post("/todos", todoRoutes.CreateTodo)
	app.Put("/todos/:id", todoRoutes.UpdateTodo)
	app.Delete("/todos/:id", todoRoutes.DeleteTodo)
}
