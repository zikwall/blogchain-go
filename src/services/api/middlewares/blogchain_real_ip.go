package middlewares

import (
	"github.com/zikwall/blogchain/src/pkg/fiberext"

	"github.com/gofiber/fiber/v2"
)

func UseBlogchainRealIP(ctx *fiber.Ctx) error {
	ctx.Locals("ip", fiberext.RealIP(ctx))
	return ctx.Next()
}
