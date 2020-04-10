package actions

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"github.com/zikwall/blogchain/types"
	"time"
)

// HelloWorldAction godoc
// @Summary Show Hello World text
// @Router / [get]
// @Success 200 "Hello, World!"
func HelloWorldAction(c *fiber.Ctx) {
	c.Send("Hello, World!")
}

func GenerateJWTAction(c *fiber.Ctx) {
	mySigningKey := []byte("secret")

	claims := types.TokenClaims{
		UUID: "UUID-HERE",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1000 * time.Second).Unix(),
			Issuer: "blogchain-auth",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		panic(err)
	}

	c.JSON(fiber.Map{
		"status": 200,
		"jwt": tokenString,
	})
}