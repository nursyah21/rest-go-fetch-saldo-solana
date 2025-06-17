package middleware

import "github.com/gofiber/fiber/v2"

func ParseBodyRequest[T any](c *fiber.Ctx) error {
	var req T
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	c.Locals("req", req)
	return c.Next()
}
