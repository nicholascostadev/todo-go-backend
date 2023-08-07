package services

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/nicholascostadev/todo-backend/controllers"
	"github.com/nicholascostadev/todo-backend/models"
)

func GetTodos(c *fiber.Ctx, todos []models.Todo) error {
	return controllers.GetTodos(c, todos)
}

type AddTodoRequestBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func AddTodo(c *fiber.Ctx, todos *[]models.Todo) error {
	var requestBody AddTodoRequestBody

	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Todo Data"})
	}

	newTodo := models.Todo{
		Id:          len(*todos) + 1,
		Title:       requestBody.Title,
		Description: requestBody.Description,
		Completed:   false,
	}

	return c.JSON(controllers.AddTodo(todos, newTodo))
}

func DeleteTodoById(c *fiber.Ctx, todos *[]models.Todo) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		fmt.Println(c.Params("id"), id, err)

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Todo Id"})
	}

	return c.JSON(controllers.DeleteTodoById(todos, id))
}
