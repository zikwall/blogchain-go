package middlewares

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	user2 "github.com/zikwall/blogchain/models/user"
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

			token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return mySigningKey, nil
			})

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				if uuid, ok := claims["uuid"]; ok {
					user, _ := user2.FindById(int64(uuid.(float64)))
					if user.Exist() {
						c.Locals("user", user)
						fmt.Println("User auth by JWT, user is: %s", user.Username)
					}
				}
			}
		}
	}

	c.Next()
}
