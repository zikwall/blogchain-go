package actions

import (
	"github.com/gofiber/fiber"
	user2 "github.com/zikwall/blogchain/models/user"
)

func Profile(c *fiber.Ctx) {
	user, err := user2.FindByUsername(c.Params("username"))

	if err != nil {
		c.Status(404).JSON(fiber.Map{
			"status":  100,
			"message": "Что-то пошло не так...",
		})

		return
	}

	c.Status(200).JSON(fiber.Map{
		"status": 200,
		"user":   user.Properties(),
	})
}
