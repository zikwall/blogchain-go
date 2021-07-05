package middlewares

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/zikwall/blogchain/src/app/lib/jwt"
	"github.com/zikwall/blogchain/src/platform/container"
	"io/ioutil"
	"net/http/httptest"
	"testing"
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

		body, _ := ioutil.ReadAll(resp.Body)

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
