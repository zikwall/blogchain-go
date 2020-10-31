package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) error {
	if c.Get("Content-Type") != "application/json" {
		return c.JSON(fiber.Map{
			"status":  100,
			"message": "Wrong content type response",
		})
	}

	return c.Next()
}
