package controller

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/nicholascostadev/todo-backend/model"
	"github.com/nicholascostadev/todo-backend/service"
)

type NewTodoController struct {
}

var todoService = service.NewTodoService{}

func (T *NewTodoController) GetTodos(c *fiber.Ctx, todos []model.Todo) error {
	return todoService.GetTodos(c, todos)
}

type AddTodoRequestBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (T *NewTodoController) AddTodo(c *fiber.Ctx, todos *[]model.Todo) error {
	var requestBody AddTodoRequestBody

	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Todo Data"})
	}

	return c.JSON(todoService.AddTodo(todos, service.AddTodoRequestBody{
		Title:       requestBody.Title,
		Description: requestBody.Description,
	}))
}

func (T *NewTodoController) DeleteTodoById(c *fiber.Ctx, todos *[]model.Todo) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		fmt.Println(c.Params("id"), id, err)

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Todo Id"})
	}

	return c.JSON(todoService.DeleteTodoById(todos, id))
}

type UpdateTodoRequestBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func (T *NewTodoController) UpdateTodoById(c *fiber.Ctx, todos *[]model.Todo) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		fmt.Println(c.Params("id"), id, err)

		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Todo Id"})
	}

	var requestBody UpdateTodoRequestBody

	if err := c.BodyParser(&requestBody); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Todo Data"})
	}

	return c.JSON(todoService.UpdateTodoById(todos, id, service.UpdateTodoRequestBody{
		Title:       requestBody.Title,
		Description: requestBody.Description,
		Completed:   requestBody.Completed,
	}))
}
