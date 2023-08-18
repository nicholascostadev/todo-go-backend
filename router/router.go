package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nicholascostadev/todo-backend/model"
)

func InitRoutes(app *fiber.App, todos []model.Todo) {
	// Todo routes
	TodosRouter(app.Group("/todos"), todos)
}
