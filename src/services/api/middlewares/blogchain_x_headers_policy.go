package middlewares

import (
	"fmt"

	"github.com/zikwall/blogchain/src/pkg/log"

	"github.com/gofiber/fiber/v2"
)

const (
	XHeaderLogFormat           = "Blogchain Client request %s from platform %s@%s"
	XHeaderBlogchainApp        = "X-Blogchain-App"
	XHeaderBlogchainPlatform   = "X-Platform"
	XHeaderBlogchainAppVersion = "X-App-Version"
	XHeaderBlogchainDeviceName = "X-Device-Name"
	XHeaderBlogchainDeviceID   = "X-Device-ID"
)

type BlogchainXHeaders struct {
	xBlogchainApp string
	xPlatform     string
	xAppVersion   string
	xDeviceName   string
	xDeviceID     string
}

func (x *BlogchainXHeaders) IsBlogchainOriginalApp() bool {
	return x.xBlogchainApp != ""
}

func WithBlogchainXHeaderPolicy() fiber.Handler {
	formatted := func(request, platform, version string) string {
		request = log.Colored(request, log.Yellow)
		platform = log.Colored(platform, log.Cyan)
		version = log.Colored(version, log.Green)

		return fmt.Sprintf(XHeaderLogFormat, request, platform, version)
	}
	return func(ctx *fiber.Ctx) error {
		x := BlogchainXHeaders{
			xBlogchainApp: ctx.Get(XHeaderBlogchainApp),
			xPlatform:     ctx.Get(XHeaderBlogchainPlatform),
			xAppVersion:   ctx.Get(XHeaderBlogchainAppVersion),
			xDeviceName:   ctx.Get(XHeaderBlogchainDeviceName),
			xDeviceID:     ctx.Get(XHeaderBlogchainDeviceID),
		}
		if x.IsBlogchainOriginalApp() {
			log.Info(
				formatted(ctx.Path(), x.xPlatform, x.xAppVersion),
			)
		}
		return ctx.Next()
	}
}