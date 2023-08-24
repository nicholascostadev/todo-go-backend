package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nicholascostadev/todo-backend/controller"
)

var todoController = controller.NewTodoController{}

func TodosRouter(router fiber.Router) {
	router.Get("/", func(c *fiber.Ctx) error {
		return todoController.GetTodos(c)
	})

	router.Post("/", func(c *fiber.Ctx) error {
		return todoController.AddTodo(c)
	})
  
	router.Delete("/clear-todos", func(c *fiber.Ctx) error {
		return todoController.ClearTodos(c)
	})

	router.Delete("/:id", func(c *fiber.Ctx) error {
		return todoController.DeleteTodoById(c)
	})

	router.Patch("/:id", func(c *fiber.Ctx) error {
		return todoController.UpdateTodoById(c)
	})

}
