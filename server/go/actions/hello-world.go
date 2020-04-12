package actions

import (
	"github.com/gofiber/fiber"
)

// HelloWorldAction godoc
// @Summary Show Hello World text
// @Router / [get]
// @Success 200 "Hello, World!"
func HelloWorldAction(c *fiber.Ctx) {
	c.Send("Hello, World!")
}
