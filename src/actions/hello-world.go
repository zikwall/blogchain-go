package actions

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	user2 "github.com/zikwall/blogchain/src/models/user"
)

// HelloWorldAction godoc
// @Summary Show Hello World text
// @Router / [get]
// @Success 200 "Hello, World!"
func HelloWorldAction(c *fiber.Ctx) error {
	userInstance := c.Locals("user").(*user2.User)

	if !userInstance.IsGuest() {
		return c.SendString(fmt.Sprintf("Hello, %s!", userInstance.Profile.Name))
	}

	return c.SendString("Hello, World!")
}
