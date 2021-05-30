package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/zikwall/blogchain/src/platform/service"
)

func WithBlogchainCORSPolicy(http service.HttpAccessControl) fiber.Handler {
	return cors.New(
		cors.Config{
			AllowOrigins:     http.AllowOrigins,
			AllowMethods:     http.AllowMethods,
			AllowHeaders:     http.AllowHeaders,
			AllowCredentials: http.AllowCredentials,
			ExposeHeaders:    http.ExposeHeaders,
			MaxAge:           http.MaxAge,
		},
	)
}
