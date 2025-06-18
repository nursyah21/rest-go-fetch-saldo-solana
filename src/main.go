package main

import (
	"fetch-saldo/src/handler"
	"fetch-saldo/src/helper"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func init() {
	helper.LoadEnv()
	helper.InitCache()
	helper.ConnectDB()
}

func main() {
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

	app.Post("/api/add-api-key", handler.AddApiKey)
	app.Post("/api/get-balances", handler.GetBalance)

	app.Listen(":5000")
}
