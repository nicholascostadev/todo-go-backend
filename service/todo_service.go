package service

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/nicholascostadev/todo-backend/model"
)

type NewTodoService struct{}

func (T *NewTodoService) GetTodos(c *fiber.Ctx) error {
	todos, err := model.GetAllTodos()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(todos)
}

type AddTodoInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (T *NewTodoService) AddTodo(partialTodo AddTodoInput) model.Todo {
	newTodo := model.Todo{
		ID:          int(uuid.New().ID()),
		Title:       partialTodo.Title,
		Description: partialTodo.Description,
		Completed:   false,
	}

	createdTodo, err := model.CreateTodo(newTodo)
	if err != nil {
		return model.Todo{}
	}

	return createdTodo
}

func (T *NewTodoService) DeleteTodoById(id int) error {
	err := model.DeleteTodo(id)

	return err
}

type UpdateTodoInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func (T *NewTodoService) UpdateTodoById(id int, partialTodo UpdateTodoInput) model.Todo {
	newTodo := model.Todo{
		ID:          id,
		Title:       partialTodo.Title,
		Description: partialTodo.Description,
		Completed:   partialTodo.Completed,
	}

	todo, err := model.UpdateTodo(id, newTodo)
	if err != nil {
		fmt.Println(err)
		return model.Todo{}
	}

	return todo
}

func (T *NewTodoService) ClearTodos(completed bool) error {
	err := model.ClearTodos(completed)
	if err != nil {
		return err
	}

	return nil
}
