package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zikwall/blogchain/src/app/lib/jwt"
	"github.com/zikwall/blogchain/src/platform/container"
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
