package middlewares

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/zikwall/blogchain/src/utils"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestWithBlogchainJWTAuthorization(t *testing.T) {
	const SECRET = "secret"

	t.Run("it should be a valid token signature with middleware", func(t *testing.T) {
		app := fiber.New()
		test := app.Group("/test")
		{
			test.Get("/jwt",
				WithBlogchainJWTAuthorization([]byte(SECRET)),
				func(c *fiber.Ctx) error {
					valid := true

					claims, ok := c.Locals("claims").(*utils.TokenClaims)

					if !ok {
						valid = false
					}

					return c.JSON(fiber.Map{
						"valid":  valid,
						"claims": claims,
					})
				},
			)
		}

		claims := utils.TokenClaims{
			UUID: 100,
		}

		createdToken, err := utils.CreateJwtToken(
			utils.TokenRequiredAttributes{
				Claims:   claims,
				Duration: 999999,
				Secret:   SECRET,
			},
		)

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
			Claims utils.TokenClaims
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
