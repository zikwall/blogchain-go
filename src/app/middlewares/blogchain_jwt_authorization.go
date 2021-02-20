package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zikwall/blogchain/src/app/lib"
	"github.com/zikwall/blogchain/src/platform/container"
)

func WithBlogchainJWTAuthorization(rsa container.RSA) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		token, ok := lib.ParseAuthHeader(ctx.Get(lib.AuthHeaderName))

		if ok {
			claims, err := lib.VerifyJwtToken(token, rsa)
			if err == nil {
				ctx.Locals(lib.ClaimsCtxKey, claims)
			}
		}

		return ctx.Next()
	}
}
