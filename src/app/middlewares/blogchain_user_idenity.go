package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zikwall/blogchain/src/app/lib"
	"github.com/zikwall/blogchain/src/app/models/user"
	"github.com/zikwall/blogchain/src/platform/service"
)

func WithBlogchainUserIdentity(blogchain *service.Instance) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		instance := &user.User{}

		claims := ctx.Locals("claims")

		if token, ok := claims.(*lib.TokenClaims); ok {
			u := user.CreateUserConnection(blogchain.GetBlogchainDatabaseInstance())

			if i, err := u.FindById(token.UUID); err == nil {
				instance = &i
			}
		}

		ctx.Locals("user", instance)

		return ctx.Next()
	}
}
