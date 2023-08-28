package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nicholascostadev/todo-backend/service"
)

type NewUserController struct{}

var userService = service.NewUserService{}

func (T *NewUserController) GetAllUsers(c *fiber.Ctx) error {
	users, err := userService.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": "There was an error when trying to retrieve all users, try again later"})
	}

	return c.JSON(users)
}
