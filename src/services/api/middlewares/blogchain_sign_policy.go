package middlewares

import (
	"github.com/zikwall/blogchain/src/pkg/exceptions"

	"github.com/gofiber/fiber/v2"
)

var ErrWrongContentType = fiber.NewError(400, "wrong content type response")

func UseBlogchainSignPolicy(ctx *fiber.Ctx) error {
	if ctx.Get("Content-Type") != "application/json" {
		return exceptions.Wrap("content type", ErrWrongContentType)
	}
	return ctx.Next()
}
