package main

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/urfave/cli"
	service "github.com/zikwall/blogchain/di"
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
				Name:   "host",
				Value:  "localhost",
				Usage:  "Run service in host",
				EnvVar: "SERVER_HOST",
			},
			&cli.IntFlag{
				Name:   "port",
				Value:  3001,
				Usage:  "Run service in port",
				EnvVar: "SERVER_PORT",
			},
			// database
			&cli.StringFlag{
				Name:   "dbhost",
				Value:  "@",
				Usage:  "Database host",
				EnvVar: "DATABASE_HOST",
			},
			&cli.StringFlag{
				Name:   "dbuser",
				Value:  "root2",
				Usage:  "Database user",
				EnvVar: "DATABASE_USER",
			},
			&cli.StringFlag{
				Name:   "dbpass",
				Value:  "prizrak211",
				Usage:  "Database password",
				EnvVar: "DATABASE_PASSWORD",
			},
			&cli.StringFlag{
				Name:   "dbname",
				Value:  "blogchain",
				Usage:  "Database name",
				EnvVar: "DATABASE_NAME",
			},
			&cli.StringFlag{
				Name:   "dbdriv",
				Value:  "mysql",
				Usage:  "Database driver",
				EnvVar: "DATABASE_DRIVER",
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		host := c.String("host")
		port := c.Int("port")

		service.DI().Bootstrap()
		service.DI().Database.Open(service.DBConfig{
			Host: c.String("dbhost"),
			User: c.String("dbuser"),
			Pass: c.String("dbpass"),
			Port: "",
			Name: c.String("dbname"),
			Driv: c.String("dbdriv"),
		})

		defer service.DI().Database.Close()

		app := fiber.New()

		InitRoutes(app)

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
