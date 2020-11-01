// +build integration

package ci

import (
	"testing"
)

func TestPingServer(t *testing.T) {
	t.Run("it should be always test passed", func(t *testing.T) {
		app := fiber.New()

		app.Get("/ping", func(c *fiber.Ctx) {
			c.Status(200).JSON(fiber.Map{
				"status": "OK",
			})
		})

		req, _ := http.NewRequest("GET", "/ping", nil)
		res, err := app.Test(req)

		if err != nil {
			t.Fatalf(`%s: %s`, t.Name(), err)
		}

		if res.StatusCode != http.StatusOK {
			t.Fatalf(`%s: Is not OK HTTP request`, t.Name())
		}
	})
}
