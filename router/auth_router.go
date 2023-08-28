package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nicholascostadev/todo-backend/controller"
)

var authController = controller.NewAuthController{}

func AuthRouter(router fiber.Router) {
	router.Get("/me", authController.GetSessionById)

	router.Post("/register", authController.RegisterUser)
	router.Post("/login", authController.LoginUser)
}
