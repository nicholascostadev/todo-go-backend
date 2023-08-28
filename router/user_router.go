package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nicholascostadev/todo-backend/controller"
)

var userController = controller.NewUserController{}

func UserRouter(router fiber.Router) {
	router.Get("/", userController.GetAllUsers)
}
