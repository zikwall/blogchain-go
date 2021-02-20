package actions

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/zikwall/blogchain/src/app/models/user"
)

// HealthCheckAction godoc
// @Summary Show OK text
// @Router / [get]
// @Success 200 "OK"
func HealthCheck(c *fiber.Ctx) error {
	userInstance := c.Locals("user").(*user.User)

	if !userInstance.IsGuest() {
		return c.SendString(fmt.Sprintf("Hello, %s!", userInstance.Profile.Name))
	}

	return c.SendString("OK")
}
