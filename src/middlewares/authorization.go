package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zikwall/blogchain/src/models/user"
)

func Authorization(c *fiber.Ctx) error {
	userInstance := c.Locals("user").(*user.User)

	if userInstance.IsGuest() {
		return c.Status(403).JSON(fiber.Map{
			"status":  100,
			"message": "Кажется у Вас нет доступа...",
		})
	}

	return c.Next()
}
