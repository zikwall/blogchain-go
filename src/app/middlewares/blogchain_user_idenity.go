package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zikwall/blogchain/src/app/lib/jwt"
	"github.com/zikwall/blogchain/src/app/repositories"
	"github.com/zikwall/blogchain/src/platform/service"
)

func WithBlogchainUserIdentity(blogchain *service.Instance) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		instance := &repositories.User{}

		claims := ctx.Locals("claims")

		if token, ok := claims.(*jwt.TokenClaims); ok {
			i, err := repositories.UseUserRepository(ctx.Context(), blogchain.GetDatabaseConnection()).FindByID(token.UUID)

			if err == nil {
				instance = &i
			}
		}

		ctx.Locals("user", instance)

		return ctx.Next()
	}
}
