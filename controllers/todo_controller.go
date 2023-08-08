package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nicholascostadev/todo-backend/models"
)

func GetTodos(c *fiber.Ctx, todos []models.Todo) error {
	return c.JSON(todos)
}

func AddTodo(todos *[]models.Todo, newTodo models.Todo) []models.Todo {
	*todos = append(*todos, newTodo)

	return *todos
}

func DeleteTodoById(todos *[]models.Todo, id int) []models.Todo {
	for i, todo := range *todos {
		if todo.Id == id {
			*todos = append((*todos)[:i], (*todos)[i+1:]...)
			break
		}
	}

	return *todos
}

func UpdateTodoById(todos *[]models.Todo, id int, newTodo models.Todo) models.Todo {
	for i, todo := range *todos {
		if todo.Id == id {
			(*todos)[i] = newTodo
			break
		}
	}

	return newTodo
}
