package actions

import "github.com/gofiber/fiber"

func HelloWorldAction(c *fiber.Ctx) {
	c.Send("Hello, World!")
}
