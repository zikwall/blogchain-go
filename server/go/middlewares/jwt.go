package middlewares

import (
	"fmt"
	"github.com/gofiber/fiber"
)

func JWT(c *fiber.Ctx) {
	fmt.Println("JWT is alive!")

	c.Next()
}
