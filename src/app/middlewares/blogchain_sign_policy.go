package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zikwall/blogchain/src/app/exceptions"
)

func UseBlogchainSignPolicy(c *fiber.Ctx) error {
	if c.Get("Content-Type") != "application/json" {
		return exceptions.Wrap("content type", fiber.NewError(400, "wrong content type response"))
	}

	return c.Next()
}
