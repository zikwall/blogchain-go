package middlewares

import (
	"github.com/zikwall/blogchain/src/pkg/container"
	"github.com/zikwall/blogchain/src/pkg/jwt"

	"github.com/gofiber/fiber/v2"
)

func WithBlogchainJWTAuthorization(rsa container.RSA) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		token, ok := jwt.ParseAuthHeader(ctx.Get(jwt.AuthHeaderName))
		if ok {
			claims, err := jwt.VerifyJwtToken(token, rsa)
			if err == nil {
				ctx.Locals(jwt.ClaimsCtxKey, claims)
			}
		}
		return ctx.Next()
	}
}
