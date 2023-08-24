package router

import (
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {
	// Todo routes
	TodosRouter(app.Group("/todos"))
}
