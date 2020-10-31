package middlewares

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/zikwall/blogchain/types"
)

func XHeader(c *fiber.Ctx) error {
	xHeader := types.NewXHeader(c)

	if xHeader.IsBlogchainApp() == false {
		// todo
	}

	requestPath := c.Path()

	fmt.Println(fmt.Sprintf("Request %s from platform: %s@%s",
		requestPath,
		xHeader.XPlatform,
		xHeader.XAppVersion,
	))

	return c.Next()
}
