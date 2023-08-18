package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/nicholascostadev/todo-backend/model"
)

type NewTodoService struct {
}

func (T *NewTodoService) GetTodos(c *fiber.Ctx, todos []model.Todo) error {
	return c.JSON(todos)
}

type AddTodoRequestBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (T *NewTodoService) AddTodo(todos *[]model.Todo, partialTodo AddTodoRequestBody) []model.Todo {
	newTodo := model.Todo{
		Id:          int(uuid.New().ID()),
		Title:       partialTodo.Title,
		Description: partialTodo.Description,
		Completed:   false,
	}

	*todos = append(*todos, newTodo)

	return *todos
}

func (T *NewTodoService) DeleteTodoById(todos *[]model.Todo, id int) []model.Todo {
	for i, todo := range *todos {
		if todo.Id == id {
			*todos = append((*todos)[:i], (*todos)[i+1:]...)
			break
		}
	}

	return *todos
}

type UpdateTodoRequestBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func (T *NewTodoService) UpdateTodoById(todos *[]model.Todo, id int, partialTodo UpdateTodoRequestBody) model.Todo {
	newTodo := model.Todo{
		Id:          id,
		Title:       partialTodo.Title,
		Description: partialTodo.Description,
		Completed:   partialTodo.Completed,
	}

	for i, todo := range *todos {
		if todo.Id == id {
			(*todos)[i] = newTodo
			break
		}
	}

	return newTodo
}
