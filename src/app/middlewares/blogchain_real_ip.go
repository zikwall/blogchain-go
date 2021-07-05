package middlewares

import "github.com/gofiber/fiber/v2"

func UseBlogchainRealIP(c *fiber.Ctx) error {
	ip := c.IP()

	if c.Get("X-Real-IP") != "" {
		ip = c.Get("X-Real-IP")
	} else if c.Get("X-Forwarded-For") != "" {
		ips := c.IPs()
		ip = ips[len(ips)-1]
	}

	c.Locals("ip", ip)
	return c.Next()
}
