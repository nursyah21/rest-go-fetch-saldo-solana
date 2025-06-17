package main

import (
	"fetch-saldo/internal/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.SetupRoutes(app)

	app.Listen(":5000")
}
