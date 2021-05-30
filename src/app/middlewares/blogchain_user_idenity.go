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
			i, err := user.ContextConnection(ctx.Context(), blogchain.GetDatabaseInstance()).FindById(token.UUID)

			if err == nil {
				instance = &i
			}
		}

		ctx.Locals("user", instance)

		return ctx.Next()
	}
}
