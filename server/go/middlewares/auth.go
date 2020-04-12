package middlewares

import (
	"fmt"
	"github.com/gofiber/fiber"
)

func Auth(c *fiber.Ctx) {
	fmt.Println("Auth work!")

	if c.Get("Content-Type") != "application/json" {
		c.JSON(fiber.Map{
			"status":  100,
			"message": "Wrong content type response",
		})

		return
	}

	c.Next()
}
