package middlewares

import (
	"github.com/gofiber/fiber"
	"github.com/zikwall/blogchain/models/user"
)

func Authorization(c *fiber.Ctx) {
	userInstance := c.Locals("user").(*user.User)

	if userInstance.IsGuest() {
		c.JSON(fiber.Map{
			"status":  300,
			"message": "Кажется у Вас нет доступа...",
		})

		return
	}

	c.Next()
}
