package middlewares

import (
	"github.com/zikwall/blogchain/src/pkg/database"
	"github.com/zikwall/blogchain/src/pkg/jwt"
	"github.com/zikwall/blogchain/src/services/api/repositories"

	"github.com/gofiber/fiber/v2"
)

func WithBlogchainUserIdentity(connection *database.Connection) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		instance := &repositories.User{}
		claims := ctx.Locals("claims")
		if token, ok := claims.(*jwt.TokenClaims); ok {
			user, err := repositories.UseUserRepository(ctx.Context(), connection).FindByID(token.UUID)
			if err == nil {
				instance = &user
			}
		}
		ctx.Locals("user", instance)
		return ctx.Next()
	}
}
