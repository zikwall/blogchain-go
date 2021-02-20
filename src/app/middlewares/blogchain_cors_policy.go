package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/zikwall/blogchain/src/platform/service"
)

func WithBlogchainCORSPolicy(blogchain *service.BlogchainServiceInstance) fiber.Handler {
	return cors.New(
		cors.Config{
			AllowOrigins:     blogchain.HttpAccessControls.AllowOrigins,
			AllowMethods:     blogchain.HttpAccessControls.AllowMethods,
			AllowHeaders:     blogchain.HttpAccessControls.AllowHeaders,
			AllowCredentials: blogchain.HttpAccessControls.AllowCredentials,
			ExposeHeaders:    blogchain.HttpAccessControls.ExposeHeaders,
			MaxAge:           blogchain.HttpAccessControls.MaxAge,
		},
	)
}
