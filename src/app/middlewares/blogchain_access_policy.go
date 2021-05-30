package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zikwall/blogchain/src/app/exceptions"
	"github.com/zikwall/blogchain/src/app/models/user"
)

func UseBlogchainAccessControlPolicy(c *fiber.Ctx) error {
	userInstance := c.Locals("user").(*user.User)

	if userInstance.IsGuest() {
		return exceptions.Wrap("access control", fiber.NewError(403, "Кажется у Вас нет доступа..."))
	}

	return c.Next()
}
