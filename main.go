package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/urfave/cli"
	service "github.com/zikwall/blogchain/src/di"
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
				Name:   "bind-address",
				Value:  "localhost:3001",
				Usage:  "Run service in host",
				EnvVar: "SERVER_HOST",
			},
			// database
			&cli.StringFlag{
				Name:     "database-host",
				Required: true,
				Value:    "@",
				Usage:    "Database host",
				EnvVar:   "DATABASE_HOST",
			},
			&cli.StringFlag{
				Name:     "database-user",
				Required: true,
				Value:    "blogchain",
				Usage:    "Database user",
				EnvVar:   "DATABASE_USER",
			},
			&cli.StringFlag{
				Name:     "database-password",
				Required: true,
				Value:    "123456",
				Usage:    "Database password",
				EnvVar:   "DATABASE_PASSWORD",
			},
			&cli.StringFlag{
				Name:     "database-name",
				Required: true,
				Value:    "blogchain",
				Usage:    "Database name",
				EnvVar:   "DATABASE_NAME",
			},
			&cli.StringFlag{
				Name:     "database-driver",
				Required: true,
				Value:    "mysql",
				Usage:    "Database driver",
				EnvVar:   "DATABASE_DRIVER",
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		host := c.String("bind-address")

		service.DI().SetupCloseHandler()
		service.DI().Bootstrap()
		service.DI().Database.Open(service.DBConfig{
			Host: c.String("database-host"),
			User: c.String("database-user"),
			Pass: c.String("database-password"),
			Name: c.String("database-name"),
			Driv: c.String("database-driver"),
		})

		app := fiber.New()

		InitRoutes(app)

		err := app.Listen(host)

		if err != nil {
			return err
		}

		return nil
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
