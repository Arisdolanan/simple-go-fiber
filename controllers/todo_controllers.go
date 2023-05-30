package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"simple-go-fiber-crud/models"
	"simple-go-fiber-crud/repository"
	"strconv"
)

type TodoController struct {
	todoRepository *repository.TodoRepository
}

func NewTodoController() *TodoController {
	return &TodoController{
		todoRepository: repository.NewTodoRepository(),
	}
}

func (r *TodoController) GetTodos(c *fiber.Ctx) error {
	data, err := r.todoRepository.FindAll()
	if err != nil {
		return c.JSON(
			fiber.Map{
				"Status":  fiber.StatusNotFound,
				"Message": err.Error(),
			})
	}
	return c.JSON(
		fiber.Map{
			"Status":  fiber.StatusOK,
			"Message": "Get data successfully",
			"Data":    data,
		})
}

func (r *TodoController) GetTodoById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	data, err := r.todoRepository.FindById(int(id))
	if err != nil {
		return c.JSON(
			fiber.Map{
				"Status":  fiber.StatusNotFound,
				"Message": err.Error(),
			})
	}
	return c.JSON(
		fiber.Map{
			"Status":  fiber.StatusOK,
			"Message": "Get data successfully",
			"Data":    data,
		})
}

func (r *TodoController) CreateTodo(c *fiber.Ctx) error {
	todo := new(models.Todo)
	err := c.BodyParser(&todo)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "check your input"})
	}
	data := r.todoRepository.Create(*todo)
	return c.JSON(&data)
}

func (r TodoController) UpdateTodo(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	payload := new(models.Todo)
	err := c.BodyParser(&payload)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "review your input"})
	}
	data := r.todoRepository.Update(id, *payload)
	return c.JSON(&data)
}

func (r *TodoController) DeleteTodo(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	result, err := r.todoRepository.Delete(id)
	if err != nil {
		return c.JSON(fiber.Map{
			"Status":  fiber.StatusNotFound,
			"Message": err.Error(),
		})
	}
	return c.JSON(
		fiber.Map{
			"Status":  fiber.StatusOK,
			"Message": fmt.Sprintf("Delete data id %s", id),
			"Data":    result,
		})
}
