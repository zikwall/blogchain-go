package main

import (
	"encoding/json"
	"github.com/gofiber/fiber"
	"github.com/zikwall/blogchain/actions"
	service "github.com/zikwall/blogchain/di"
	"github.com/zikwall/blogchain/models/content"
	"net/http"
	"os"
	"testing"
)

func Init() {
	service.DI().Bootstrap()
	service.DI().Database.Open(service.DBConfig{
		Host: "@",
		User: os.Getenv("MYSQL_USER"),
		Pass: os.Getenv("MYSQL_PASSWORD"),
		Port: "3001",
		Name: os.Getenv("MYSQL_DATABASE"),
		Driv: "mysql",
	})
}

func TestApiContents(t *testing.T) {
	Init()
	defer service.DI().Database.Close()

	app := fiber.New()

	app.Get("/api/v1/contents/:page?", actions.GetContents)

	req, _ := http.NewRequest("GET", "/api/v1/contents", nil)
	res, err := app.Test(req)

	if err != nil {
		t.Fatalf(`%s: %s`, t.Name(), err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf(`%s: Is not OK HTTP request, given %d`, t.Name(), res.StatusCode)
	}

	var jsonApi struct {
		Contents []content.PublicContent `json:"contents"`
		Meta     struct {
			Pages int `json:"pages"`
		} `json:"meta"`
	}

	err = json.NewDecoder(res.Body).Decode(&jsonApi)

	if err != nil {
		t.Fatalf("Failed to load parse JSON in test: %s, given error: %v", t.Name(), err)
	}

	if len(jsonApi.Contents) == 0 {
		t.Fatalf("Opps... not content given in: %s", t.Name())
	}
}
