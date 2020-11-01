package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/zikwall/blogchain/src/service"
)

func WithBlogchainCORSPolicy(blogchain *service.BlogchainServiceInstance) fiber.Handler {
	return cors.New(
		cors.Config{
			AllowOrigins:     blogchain.AccessControls.AllowOrigins,
			AllowMethods:     blogchain.AccessControls.AllowMethods,
			AllowHeaders:     blogchain.AccessControls.AllowHeaders,
			AllowCredentials: blogchain.AccessControls.AllowCredentials,
			ExposeHeaders:    blogchain.AccessControls.ExposeHeaders,
			MaxAge:           blogchain.AccessControls.MaxAge,
		},
	)
}
