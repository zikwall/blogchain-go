package middlewares

import (
	"github.com/zikwall/blogchain/src/services/api/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func WithBlogchainCORSPolicy(http *service.HTTPAccessControl) fiber.Handler {
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
