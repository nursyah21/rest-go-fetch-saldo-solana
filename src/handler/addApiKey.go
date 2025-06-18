package handler

import (
	"fetch-saldo/src/helper"
	"fetch-saldo/src/models"

	"github.com/gofiber/fiber/v2"
)

func AddApiKey(c *fiber.Ctx) error {
	xSecret := c.Get("X-Secret")
	if xSecret != helper.SECRET_KEY {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid or Missing X-Secret"})
	}

	type request struct {
		ApiKey string `json:"api_key"`
	}
	var req request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if req.ApiKey == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Missing required fields: api_key",
		})
	}

	if exists := helper.GetAPIKeyCache(req.ApiKey); exists {
		return c.Status(400).JSON(fiber.Map{
			"error": "Api key already exists",
		})
	}

	if exists := models.ApiExist(req.ApiKey); exists {
		helper.SetAPIKeyCache(req.ApiKey, exists)

		return c.Status(400).JSON(fiber.Map{
			"error": "Api key already exists",
		})
	}

	err := models.CreateAPI(req.ApiKey)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to save API key",
		})
	}

	return c.JSON(fiber.Map{
		"message": "API key created successfully",
	})
}
