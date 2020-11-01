package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zikwall/blogchain/src/models/user"
	"github.com/zikwall/blogchain/src/service"
	"github.com/zikwall/blogchain/src/utils"
	"strings"
)

func WithBlogchainJWTAuthorization(blogchain *service.BlogchainServiceInstance) fiber.Handler {
	sigin := blogchain.Container.GetContainerSecret()

	return func(ctx *fiber.Ctx) error {
		tokenString := ctx.Get("Authorization")
		instance := &user.User{}

		if tokenString != "" {
			header := strings.Split(tokenString, "Bearer ")

			if len(header) == 2 {
				tokenString = header[1]
				token, err := utils.VerifyJwtToken(tokenString, sigin)

				if err == nil {
					u := user.NewUserModel()
					instance, _ = u.FindById(token.UUID)
				}
			}
		}

		ctx.Locals("user", instance)
		return ctx.Next()
	}
}
