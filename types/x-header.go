package types

import "github.com/gofiber/fiber/v2"

type XHeader struct {
	XBlogchainApp string
	XPlatform     string
	XAppVersion   string
	XDeviceName   string
	XDeviceId     string
}

func NewXHeader(c *fiber.Ctx) XHeader {
	x := XHeader{
		XBlogchainApp: "",
		XPlatform:     "",
		XAppVersion:   "",
		XDeviceName:   "",
		XDeviceId:     "",
	}

	x.resolveHeaders(c)

	return x
}

func (x *XHeader) IsBlogchainApp() bool {
	return len(x.XBlogchainApp) != 0
}

func (x *XHeader) resolveHeaders(c *fiber.Ctx) {
	x.XBlogchainApp = c.Get("X-Blogchain-App")
	x.XPlatform = c.Get("X-Platform")
	x.XAppVersion = c.Get("X-App-Version")
	x.XDeviceName = c.Get("X-Device-Name")
	x.XDeviceId = c.Get("X-Device-Id")
}
