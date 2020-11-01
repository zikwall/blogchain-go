package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zikwall/blogchain/src/utils"
	"strings"
)

func WithBlogchainJWTAuthorization(sigin []byte) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		tokenString := ctx.Get("Authorization")

		if tokenString != "" {
			header := strings.Split(tokenString, "Bearer ")

			if len(header) == 2 {
				tokenString = header[1]
				token, err := utils.VerifyJwtToken(tokenString, sigin)

				if err == nil {
					ctx.Locals("claims", token)
				}
			}
		}

		return ctx.Next()
	}
}
