package middlewares

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/zikwall/blogchain/models/user"
	"strings"
)

func JWT(c *fiber.Ctx) error {
	mySigningKey := []byte("secret")
	tokenString := c.Get("Authorization")
	// default empty instance of user
	userInstance := &user.User{}

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

			if err == nil {
				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					if uuid, ok := claims["uuid"]; ok {
						// set new instance
						userInstance, _ = user.FindById(int64(uuid.(float64)))
					}
				}
			}
		}
	}

	// always set user instance
	c.Locals("user", userInstance)
	return c.Next()
}
