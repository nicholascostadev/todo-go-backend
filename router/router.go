package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nicholascostadev/todo-backend/models"
)

func InitRoutes(app *fiber.App, todos []models.Todo) {
	// Todo routes
	TodosRouter(app.Group("/todos"), todos)
}
