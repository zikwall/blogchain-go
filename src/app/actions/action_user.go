package actions

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zikwall/blogchain/src/app/models/user"
)

func (a BlogchainActionProvider) Profile(c *fiber.Ctx) error {
	u := user.CreateUserConnection(a.db)

	result, err := u.FindByUsername(c.Params("username"))

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  100,
			"message": "Что-то пошло не так...",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status": 200,
		"user":   result.Properties(),
	})
}
