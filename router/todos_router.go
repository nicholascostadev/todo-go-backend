package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nicholascostadev/todo-backend/models"
	"github.com/nicholascostadev/todo-backend/services"
)

func TodosRouter(router fiber.Router, todos []models.Todo) {
	router.Get("/", func(c *fiber.Ctx) error {
		return services.GetTodos(c, todos)
	})

	router.Post("/", func(c *fiber.Ctx) error {
		return services.AddTodo(c, &todos)
	})

	router.Delete("/:id", func(c *fiber.Ctx) error {
		return services.DeleteTodoById(c, &todos)
	})

	router.Patch("/:id", func(c *fiber.Ctx) error {
		return services.UpdateTodoById(c, &todos)
	})
}
