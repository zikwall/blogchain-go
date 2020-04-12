package actions

import (
	"fmt"
	"github.com/gofiber/fiber"
	user2 "github.com/zikwall/blogchain/models/user"
)

// HelloWorldAction godoc
// @Summary Show Hello World text
// @Router / [get]
// @Success 200 "Hello, World!"
func HelloWorldAction(c *fiber.Ctx) {
	userInstance := c.Locals("user").(*user2.User)

	if !userInstance.IsGuest() {
		c.Send(fmt.Sprintf("Hello, %s!", userInstance.Username))
		return
	}

	c.Send("Hello, World!")
}
