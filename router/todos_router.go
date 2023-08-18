package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nicholascostadev/todo-backend/controller"
	"github.com/nicholascostadev/todo-backend/model"
)

var todoController = controller.NewTodoController{}

func TodosRouter(router fiber.Router, todos []model.Todo) {
	router.Get("/", func(c *fiber.Ctx) error {
		return todoController.GetTodos(c, todos)
	})

	router.Post("/", func(c *fiber.Ctx) error {
		return todoController.AddTodo(c, &todos)
	})

	router.Delete("/:id", func(c *fiber.Ctx) error {
		return todoController.DeleteTodoById(c, &todos)
	})

	router.Patch("/:id", func(c *fiber.Ctx) error {
		return todoController.UpdateTodoById(c, &todos)
	})
}
