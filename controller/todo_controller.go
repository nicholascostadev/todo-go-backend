package controller

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/nicholascostadev/todo-backend/service"
)

type NewTodoController struct {
}

var todoService = service.NewTodoService{}

func (T *NewTodoController) GetTodos(c *fiber.Ctx) error {
	return todoService.GetTodos(c)
}

type AddTodoRequestBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (T *NewTodoController) AddTodo(c *fiber.Ctx) error {
	var requestBody AddTodoRequestBody

	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Todo Data"})
	}

	return c.JSON(todoService.AddTodo(service.AddTodoRequestBody{
		Title:       requestBody.Title,
		Description: requestBody.Description,
	}))
}

func (T *NewTodoController) DeleteTodoById(c *fiber.Ctx) error {
	var err error
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		fmt.Println(c.Params("id"), id, err)

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Todo Id"})
	}

	err = todoService.DeleteTodoById(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Todo Id"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Todo deleted successfully"})
}

type UpdateTodoRequestBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func (T *NewTodoController) UpdateTodoById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		fmt.Println(c.Params("id"), id, err)

		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Todo Id"})
	}

	var requestBody UpdateTodoRequestBody

	if err := c.BodyParser(&requestBody); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Todo Data"})
	}

	return c.JSON(todoService.UpdateTodoById(id, service.UpdateTodoRequestBody{
		Title:       requestBody.Title,
		Description: requestBody.Description,
		Completed:   requestBody.Completed,
	}))
}

func (T *NewTodoController) ClearTodos(c *fiber.Ctx) error {
	queryStatus := c.Query("status", "")
	fmt.Println(queryStatus)

	if queryStatus != "completed" && queryStatus != "incomplete" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid status"})
	}

	err := todoService.ClearTodos(queryStatus == "completed")

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Unknown error"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Todos cleared successfully"})
}
