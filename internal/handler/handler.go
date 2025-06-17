package handler

import (
	"fetch-saldo/internal/model"

	"github.com/gofiber/fiber/v2"
)

func GetBalance(c *fiber.Ctx) error {
	req := c.Locals("req").(model.BalanceRequest)

	return c.JSON(req)
}
