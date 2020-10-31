package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/urfave/cli/v2"
	"github.com/zikwall/blogchain/src/service"
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
				Name:    "bind-address",
				Value:   "localhost:3001",
				Usage:   "Run service in host",
				EnvVars: []string{"SERVER_HOST"},
			},
			// database
			&cli.StringFlag{
				Name:     "database-host",
				Required: true,
				Value:    "@",
				Usage:    "Database host",
				EnvVars:  []string{"DATABASE_HOST"},
			},
			&cli.StringFlag{
				Name:     "database-user",
				Required: true,
				Value:    "blogchain",
				Usage:    "Database user",
				EnvVars:  []string{"DATABASE_USER"},
			},
			&cli.StringFlag{
				Name:     "database-password",
				Required: true,
				Value:    "123456",
				Usage:    "Database password",
				EnvVars:  []string{"DATABASE_PASSWORD"},
			},
			&cli.StringFlag{
				Name:     "database-name",
				Required: true,
				Value:    "blogchain",
				Usage:    "Database name",
				EnvVars:  []string{"DATABASE_NAME"},
			},
			&cli.StringFlag{
				Name:     "database-driver",
				Required: true,
				Value:    "mysql",
				Usage:    "Database driver",
				EnvVars:  []string{"DATABASE_DRIVER"},
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		_, err := service.NewBlogchainServiceInstance(
			service.BlogchainServiceConfiguration{
				BloghainDatabaseConfiguration: service.BloghainDatabaseConfiguration{
					Host:     c.String("database-host"),
					User:     c.String("database-user"),
					Password: c.String("database-password"),
					Name:     c.String("database-name"),
					Driver:   c.String("database-driver"),
				},
			},
		)

		if err != nil {
			return err
		}

		host := c.String("bind-address")

		app := fiber.New()

		InitRoutes(app)

		err = app.Listen(host)

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
