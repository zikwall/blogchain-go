package main

import (
	"fmt"
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/urfave/cli"
	"github.com/zikwall/blogchain/actions"
	service "github.com/zikwall/blogchain/di"
	"github.com/zikwall/blogchain/middlewares"
	"log"
	"os"
)

// @title Blog Chain swagger documentation for Go service
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support Blog Chain
// @contact.url http://www.blogchain.io/support
// @contact.email support@blogchain.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host blogchain.io
func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "host",
				Value: "localhost",
				Usage: "Run service in host",
			},
			&cli.IntFlag{
				Name:  "port",
				Value: 3001,
				Usage: "Run service in port",
			},
			// database
			&cli.StringFlag{
				Name:  "dbhost",
				Value: "@",
				Usage: "Database host",
			},
			&cli.StringFlag{
				Name:  "dbuser",
				Value: "root2",
				Usage: "Database user",
			},
			&cli.StringFlag{
				Name:  "dbpass",
				Value: "prizrak211",
				Usage: "Database password",
			},
			&cli.StringFlag{
				Name:  "dbname",
				Value: "blogchain",
				Usage: "Database name",
			},
			&cli.StringFlag{
				Name:  "dbdriv",
				Value: "mysql",
				Usage: "Database driver",
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		host := c.String("host")
		port := c.Int("port")

		service.DI().Bootstrap()
		service.DI().Database.ConnectDatabase(service.DBConfig{
			Host: c.String("dbhost"),
			User: c.String("dbuser"),
			Pass: c.String("dbpass"),
			Port: "",
			Name: c.String("dbname"),
			Driv: c.String("dbdriv"),
		})

		app := fiber.New()
		app.Use(cors.New(cors.Config{
			Filter:           nil,
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"*"},
			AllowHeaders:     []string{"*"},
			AllowCredentials: false,
			ExposeHeaders:    nil,
			MaxAge:           0,
		}))

		app.Static("/docs", "./docs")

		// Main endpoint group by `/api` prefix
		api := app.Group("/api", middlewares.JWT)
		api.Get("/", actions.HelloWorldAction)

		v1 := api.Group("/v1")
		v1.Get("/profile/:username", actions.Profile)

		// content
		v1.Post("/content/add", actions.AddContent)
		v1.Get("/content/:id", actions.GetContent)

		// not usage JWT middleware in Login & Register endpoints
		auth := app.Group("/auth", middlewares.Auth)
		auth.Post("/register", actions.Register)
		auth.Post("/login", actions.Login)
		auth.Post("/logout", actions.Login)

		err := app.Listen(fmt.Sprintf("%s:%d", host, port))

		if err != nil {
			log.Fatal(err)
		}

		return nil
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
