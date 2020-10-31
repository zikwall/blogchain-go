package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/urfave/cli/v2"
	"github.com/zikwall/blogchain/src/actions"
	"github.com/zikwall/blogchain/src/middlewares"
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
	application := &cli.App{
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

	application.Action = func(c *cli.Context) error {
		blogchain, err := service.NewBlogchainServiceInstance(
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

		app := fiber.New()
		app.Use(cors.New(cors.Config{
			AllowOrigins:     "*",
			AllowMethods:     "*",
			AllowHeaders:     "*",
			AllowCredentials: false,
			ExposeHeaders:    "",
			MaxAge:           0,
		}))

		app.Static("/docs", "./src/public/docs")
		app.Static("/uploads", "./src/public/uploads")

		// only blogchain apps
		app.Use(middlewares.XHeader)

		// main endpoint group by `/api` prefix
		api := app.Group("/api", middlewares.JWT)

		v1 := api.Group("/v1")
		{
			v1.Get("/profile/:username", actions.Profile)
			v1.Get("/content/:id", actions.GetContent)
			v1.Get("/contents/:page?", actions.GetContents)
			v1.Get("/tags", actions.Tags)
			v1.Get("/contents/user/:id/:page?", actions.GetUserContents)
			v1.Get("/tag/:tag/:page?", actions.GetContents)
		}

		// not usage JWT middleware in Login & Register endpoints
		auth := app.Group("/auth", middlewares.Auth)
		{
			auth.Post("/register", actions.Register)
			auth.Post("/login", actions.Login)
			auth.Post("/logout", actions.Login)
		}

		// editor required authorization
		editor := api.Group("/editor", middlewares.Authorization)
		{
			editor.Get("/content/:id", actions.GetEditContent)
			editor.Post("/content/add", actions.AddContent)
			editor.Post("/content/update/:id", actions.UpdateContent)
		}

		api.Get("/", actions.HelloWorldAction)

		go func() {
			if err := app.Listen(c.String("bind-address")); err != nil {
				blogchain.GetInternalLogger().Error(err)
			}
		}()

		blogchain.WaitBlogchainSystemNotify()
		blogchain.ShutdownBlogchainServer()

		return nil
	}

	err := application.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
