package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zikwall/blogchain/src/app/exceptions"
	"github.com/zikwall/blogchain/src/app/repositories"
)

func UseBlogchainAccessControlPolicy(c *fiber.Ctx) error {
	userInstance, ok := c.Locals("user").(*repositories.User)

	if !ok {
		return exceptions.Wrap("access control", fiber.NewError(500, "Что-то пошло не так..."))
	}

	if userInstance.IsGuest() {
		return exceptions.Wrap("access control", fiber.NewError(403, "Кажется у Вас нет доступа..."))
	}

	return c.Next()
}
