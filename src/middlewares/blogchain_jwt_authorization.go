package middlewares

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/zikwall/blogchain/src/models/user"
	"github.com/zikwall/blogchain/src/service"
	"strings"
)

func WithBlogchainJWTAuthorization(blogchain *service.BlogchainServiceInstance) fiber.Handler {
	sigin := blogchain.Container.GetContainerSecret()

	return func(ctx *fiber.Ctx) error {
		tokenString := ctx.Get("Authorization")
		instance := &user.User{}

		if tokenString != "" {
			header := strings.Split(tokenString, "Bearer ")
			if len(header) == 2 {
				tokenString = header[1]

				token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
					}
					return sigin, nil
				})

				if err == nil {
					if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
						if uuid, ok := claims["uuid"]; ok {
							u := user.NewUserModel()
							instance, _ = u.FindById(int64(uuid.(float64)))
						}
					}
				}
			}
		}

		ctx.Locals("user", instance)
		return ctx.Next()
	}
}
