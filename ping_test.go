package main

import (
	"github.com/gofiber/fiber"
	"net/http"
	"testing"
)

func TestPingServer(t *testing.T) {
	app := fiber.New()

	app.Get("/ping", func(c *fiber.Ctx) {
		c.Status(200).JSON(fiber.Map{
			"status": "ok",
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
}
