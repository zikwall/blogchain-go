package actions

import (
	"github.com/gofiber/fiber/v2"
	user2 "github.com/zikwall/blogchain/models/user"
)

func Profile(c *fiber.Ctx) error {
	user, err := user2.FindByUsername(c.Params("username"))

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  100,
			"message": "Что-то пошло не так...",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status": 200,
		"user":   user.Properties(),
	})
}
