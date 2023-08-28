package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nicholascostadev/todo-backend/controller"
)

var todoController = controller.NewTodoController{}

func TodosRouter(router fiber.Router) {
	router.Get("/", todoController.GetTodos)
	router.Post("/", todoController.AddTodo)
	router.Delete("/clear-todos", todoController.ClearTodos)
	router.Delete("/:id", todoController.DeleteTodoById)
	router.Patch("/:id", todoController.UpdateTodoById)
}
