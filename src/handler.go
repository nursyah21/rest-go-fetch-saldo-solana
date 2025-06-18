package main

import (
	"github.com/gofiber/fiber/v2"
)

type BalanceRequest struct {
	Wallets []string `json:"wallets"`
}

func GetBalance(c *fiber.Ctx) error {
	var req BalanceRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if len(req.Wallets) == 0 {
		return c.Status(400).JSON(fiber.Map{
			"error": "Missing required fields: wallets: []",
		})
	}

	return c.JSON(req)
}
