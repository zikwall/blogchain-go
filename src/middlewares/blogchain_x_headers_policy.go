package middlewares

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/zikwall/blogchain/src/service"
)

const (
	XHeaderLogFormat           = "Blogchain Client request %s from platform %s@%s"
	XHeaderBlogchainApp        = "X-Blogchain-App"
	XHeaderBlogchainPlatform   = "X-Platform"
	XHeaderBlogchainAppVersion = "X-App-Version"
	XHeaderBlogchainDeviceName = "X-Device-Name"
	XHeaderBlogchainDeviceId   = "X-Device-Id"
)

type (
	BlogchainXHeaders struct {
		xBlogchainApp string
		xPlatform     string
		xAppVersion   string
		xDeviceName   string
		xDeviceId     string
	}
)

func (x BlogchainXHeaders) IsBlogchainOriginalApp() bool {
	return len(x.xBlogchainApp) != 0
}

func WithBlogchainXHeaderPolicy(blogchain *service.BlogchainServiceInstance) fiber.Handler {
	formatted := func(request, platform, version string) string {
		colorizer := blogchain.GetInternalLogger().GetColorizer()

		request = colorizer.Colored(request, service.Yellow)
		platform = colorizer.Colored(platform, service.Cyan)
		version = colorizer.Colored(version, service.Green)

		return fmt.Sprintf(XHeaderLogFormat, request, platform, version)
	}

	return func(ctx *fiber.Ctx) error {
		x := BlogchainXHeaders{
			xBlogchainApp: ctx.Get(XHeaderBlogchainApp),
			xPlatform:     ctx.Get(XHeaderBlogchainPlatform),
			xAppVersion:   ctx.Get(XHeaderBlogchainAppVersion),
			xDeviceName:   ctx.Get(XHeaderBlogchainDeviceName),
			xDeviceId:     ctx.Get(XHeaderBlogchainDeviceId),
		}

		if x.IsBlogchainOriginalApp() {
			blogchain.GetInternalLogger().Info(
				formatted(ctx.Path(), x.xPlatform, x.xAppVersion),
			)
		}

		return ctx.Next()
	}
}
