package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/api/get-balances", GetBalance)

	app.Listen(":5000")
}
