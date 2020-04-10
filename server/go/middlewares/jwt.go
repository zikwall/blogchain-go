package middlewares

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"strings"
)

func JWT(c *fiber.Ctx) {
	mySigningKey := []byte("secret")
	tokenString := c.Get("Authorization")

	if tokenString != "" {
		header := strings.Split(tokenString, "Bearer ")

		// strict mode
		if len(header) == 2 {
			tokenString = header[1]

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return mySigningKey, nil
			})

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				fmt.Println(claims["uuid"], claims["exp"])
			} else {
				fmt.Println(err)
			}
		}
	}

	c.Next()
}
