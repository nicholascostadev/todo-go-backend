package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/nicholascostadev/todo-backend/model"
	"github.com/nicholascostadev/todo-backend/router"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())
	model.SetupDatabase()

	router.InitRoutes(app)

	app.Listen("localhost:8080")
}
