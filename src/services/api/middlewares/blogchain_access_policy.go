package middlewares

import (
	"github.com/zikwall/blogchain/src/pkg/exceptions"
	"github.com/zikwall/blogchain/src/services/api/repositories"

	"github.com/gofiber/fiber/v2"
)

var (
	ErrNotAccess      = fiber.NewError(403, "Кажется у Вас нет доступа...")
	ErrSomethingWrong = fiber.NewError(500, "Что-то пошло не так...")
)

func UseBlogchainAccessControlPolicy(ctx *fiber.Ctx) error {
	userInstance, ok := ctx.Locals("user").(*repositories.User)
	if !ok {
		return exceptions.Wrap("access control", ErrSomethingWrong)
	}
	if userInstance.IsGuest() {
		return exceptions.Wrap("access control", ErrNotAccess)
	}
	return ctx.Next()
}
