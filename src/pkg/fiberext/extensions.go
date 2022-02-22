package fiberext

import "github.com/gofiber/fiber/v2"

func RealIP(c *fiber.Ctx) string {
	// client can set the X-Forwarded-For or the X-Real-IP header to any arbitrary value it wants,
	// unless you have a trusted reverse proxy, you shouldn't use any of those values.
	if c.Get("X-Real-IP") != "" {
		return c.Get("X-Real-IP")
	} else if c.Get("X-Forwarded-For") != "" {
		ips := c.IPs()
		return ips[len(ips)-1]
	}
	// return remote addr
	return c.IP()
}
