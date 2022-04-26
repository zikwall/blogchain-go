package middlewares

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/zikwall/blogchain/src/pkg/container"
	"github.com/zikwall/blogchain/src/pkg/jwt"

	"github.com/gofiber/fiber/v2"
)

func TestWithBlogchainJWTAuthorization(t *testing.T) {
	t.Run("it should be a valid token signature with middleware", func(t *testing.T) {
		app := fiber.New()
		rsa := &container.MockRSA{}
		app.Group("/test").Get("/jwt", WithBlogchainJWTAuthorization(rsa),
			func(c *fiber.Ctx) error {
				valid := true
				claims, ok := c.Locals("claims").(*jwt.TokenClaims)
				if !ok {
					valid = false
				}
				return c.JSON(fiber.Map{
					"valid":  valid,
					"claims": claims,
				})
			},
		)
		claims := &jwt.TokenClaims{
			UUID: 100,
		}

		createdToken, err := jwt.CreateJwtToken(claims, 999, rsa.GetPrivateKey())
		if err != nil {
			t.Fatal(err)
		}

		req := httptest.NewRequest("GET", "/test/jwt", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", createdToken))
		resp, _ := app.Test(req)

		if resp.StatusCode != 200 {
			t.Fatal("Failed check signature by response status code")
		}

		body, _ := io.ReadAll(resp.Body)

		response := struct {
			Valid  bool
			Claims jwt.TokenClaims
		}{}

		if err := json.Unmarshal(body, &response); err != nil {
			t.Fatal(err)
		}

		t.Log(response)

		if !response.Valid {
			t.Fatal("Failed check signature")
		}
		if response.Claims.UUID != 100 {
			t.Fatal("Failed check signature")
		}
	})
}
