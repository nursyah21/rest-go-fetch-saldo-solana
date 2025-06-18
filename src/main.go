package main

import (
	"fetch-saldo/src/config"
	"fetch-saldo/src/handler"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func main() {
	config.NewConfig()
	app := fiber.New()

	app.Use(limiter.New(limiter.Config{
		Max:      10,
		Duration: 1 * time.Minute,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Rate limit exceeded",
			})
		},
	}))

	app.Post("/api/get-balances", handler.GetBalance)

	app.Listen(":5000")
}
