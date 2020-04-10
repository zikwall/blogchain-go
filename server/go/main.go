package main

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/urfave/cli"
	"github.com/zikwall/blogchain/actions"
	"github.com/zikwall/blogchain/middlewares"
	"log"
	"os"
)

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
		},
	}

	app.Action = func(c *cli.Context) error {
		host := c.String("host")
		port := c.Int("port")

		app := fiber.New()

		app.Use(middlewares.JWT)
		app.Get("/", actions.HelloWorldAction)

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
