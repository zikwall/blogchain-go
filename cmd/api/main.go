package main

import (
	"context"
	"os"

	"github.com/zikwall/blogchain/src/pkg/clickhouse"
	"github.com/zikwall/blogchain/src/pkg/container"
	"github.com/zikwall/blogchain/src/pkg/database"
	"github.com/zikwall/blogchain/src/pkg/log"
	"github.com/zikwall/blogchain/src/pkg/maxmind"
	metaV1 "github.com/zikwall/blogchain/src/pkg/meta/v1"
	"github.com/zikwall/blogchain/src/pkg/signal"
	"github.com/zikwall/blogchain/src/services/api/actions"
	"github.com/zikwall/blogchain/src/services/api/middlewares"
	"github.com/zikwall/blogchain/src/services/api/service"

	"github.com/gofiber/fiber/v2"
	"github.com/urfave/cli/v2"
	"github.com/zikwall/fsclient"
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
			// listeners
			&cli.StringFlag{
				Name:     "bind-address",
				Required: true,
				Usage:    "IP and port for TCP listener, example: 0.0.0.0:3001",
				EnvVars:  []string{"BIND_ADDRESS", "PORT"},
			},
			&cli.StringFlag{
				Name:     "bind-socket",
				Usage:    "Path to unix socket file for UDS listener",
				Required: false,
				Value:    "/tmp/blogchain.sock",
				EnvVars:  []string{"BIND_SOCKET"},
			},
			&cli.IntFlag{
				Name:     "listener",
				Usage:    "UDS or TCP, default TCP",
				Required: false,
				Value:    signal.ListenerTCP,
				EnvVars:  []string{"LISTENER"},
			},

			// database
			&cli.StringFlag{
				Name:     "database-host",
				Required: true,
				Usage:    "Database host",
				EnvVars:  []string{"DATABASE_HOST"},
				FilePath: "/srv/bc_secret/database_host",
			},
			&cli.StringFlag{
				Name:     "database-user",
				Required: true,
				Usage:    "Database user",
				EnvVars:  []string{"DATABASE_USER"},
				FilePath: "/srv/bc_secret/database_user",
			},
			&cli.StringFlag{
				Name:     "database-password",
				Required: true,
				Usage:    "Database password",
				EnvVars:  []string{"DATABASE_PASSWORD"},
				FilePath: "/srv/bc_secret/database_password",
			},
			&cli.StringFlag{
				Name:     "database-name",
				Required: true,
				Usage:    "Database name",
				EnvVars:  []string{"DATABASE_NAME"},
				FilePath: "/srv/bc_secret/database_name",
			},
			&cli.StringFlag{
				Name:     "database-dialect",
				Required: true,
				Usage:    "Database dialect: mysql, postgres, sqlite3, sqlserver etc",
				EnvVars:  []string{"DATABASE_DIALECT"},
				FilePath: "/srv/bc_secret/database_dialect",
			},
			&cli.StringFlag{
				Name:     "rsa-public-key",
				Required: false,
				Usage:    "Container secret key for JWT, and etc.",
				EnvVars:  []string{"RSA_PUBLIC_KEY"},
				FilePath: "/srv/bc_secret/rsa_public_key",
			},
			&cli.StringFlag{
				Name:     "rsa-private-key",
				Required: false,
				Usage:    "Container secret key for JWT, and etc.",
				EnvVars:  []string{"RSA_PRIVATE_KEY"},
				FilePath: "/srv/bc_secret/rsa_private_key",
			},

			// clickhouse
			&cli.StringFlag{
				Name:     "clickhouse-address",
				Usage:    "Clickhouse server address",
				Required: true,
				EnvVars:  []string{"CLICKHOUSE_ADDRESS"},
				FilePath: "/srv/bc_secret/clickhouse_address",
			},
			&cli.StringFlag{
				Name:     "clickhouse-user",
				Usage:    "Clickhouse server user",
				EnvVars:  []string{"CLICKHOUSE_USER"},
				FilePath: "/srv/bc_secret/clickhouse_user",
			},
			&cli.StringFlag{
				Name:     "clickhouse-password",
				Usage:    "Clickhouse server user password",
				EnvVars:  []string{"CLICKHOUSE_PASSWORD"},
				FilePath: "/srv/bc_secret/clickhouse_password",
			},
			&cli.StringFlag{
				Name:     "clickhouse-database",
				Usage:    "Clickhouse server database name",
				EnvVars:  []string{"CLICKHOUSE_DATABASE"},
				FilePath: "/srv/bc_secret/clickhouse_database",
			},
			&cli.StringFlag{
				Name:     "clickhouse-alt-hosts",
				Usage:    "Comma separated list of single address host for load-balancing",
				EnvVars:  []string{"CLICKHOUSE_ALT_HOSTS"},
				FilePath: "/srv/bc_secret/clickhouse_alt_hosts",
			},

			// geo
			&cli.StringFlag{
				Name:    "maxmind-mmdb",
				Usage:   "Path to City.mmdb file for Maxmind",
				Value:   "./share/geo/GeoLite2-City.mmdb",
				EnvVars: []string{"MAXMIND_FILEPATH"},
			},

			// cdn
			&cli.StringFlag{
				Name:     "cdn-host",
				Usage:    "",
				Value:    "localhost:1337",
				EnvVars:  []string{"CDN_HOST"},
				FilePath: "/srv/bc_secret/cdn_host",
			},
			&cli.StringFlag{
				Name:     "cdn-user",
				Usage:    "",
				EnvVars:  []string{"CDN_USER"},
				FilePath: "/srv/bc_secret/cdn_user",
			},
			&cli.StringFlag{
				Name:     "cdn-password",
				Usage:    "",
				EnvVars:  []string{"CDN_PASSWORD"},
				FilePath: "/srv/bc_secret/cdn_password",
			},
			&cli.StringFlag{
				Name:     "statistic-address",
				Required: true,
				Value:    "0.0.0.0:7000",
				Usage:    "Storage gRPC host",
				EnvVars:  []string{"STATISTIC_ADDRESS"},
			},
			// dev
			&cli.BoolFlag{
				Name:    "debug",
				Usage:   "Debug mode - details the stages of operation of the service, also in this mode, all logs are sent to stdout",
				EnvVars: []string{"DEBUG"},
			},
		},
		Action: Main,
	}

	if err := application.Run(os.Args); err != nil {
		log.Error(err)
	}
}

func Main(ctx *cli.Context) error {
	appContext, cancel := context.WithCancel(ctx.Context)
	defer func() {
		cancel()
		log.Info("app context is canceled, service is down!")
	}()

	await, stop := signal.Notifier(func() {
		log.Info("received a system signal to shutdown TRAFFIC BALANCER server, start the shutdown process..")
	})

	blogchain, err := service.New(
		appContext,
		&service.Configuration{
			DatabaseConfiguration: &database.Configuration{
				Host:     ctx.String("database-host"),
				User:     ctx.String("database-user"),
				Password: ctx.String("database-password"),
				Name:     ctx.String("database-name"),
				Dialect:  ctx.String("database-dialect"),
				Debug:    ctx.Bool("debug"),
			},
			Container: container.Configuration{},
			ClickhouseConfiguration: &clickhouse.Configuration{
				Address:  ctx.String("clickhouse-address"),
				User:     ctx.String("clickhouse-user"),
				Password: ctx.String("clickhouse-password"),
				Database: ctx.String("clickhouse-database"),
				AltHosts: ctx.String("clickhouse-alt-hosts"),
				IsDebug:  ctx.Bool("debug"),
			},
			ReaderConfig: maxmind.ReaderConfig{
				Path: ctx.String("maxmind-mmdb"),
			},
			StatisticAddress: ctx.String("statistic-address"),
		},
	)
	if err != nil {
		return err
	}

	defer func() {
		blogchain.Shutdown(func(err error) {
			log.Info(err)
		})
		blogchain.Stacktrace()
	}()

	app := fiber.New(fiber.Config{
		ErrorHandler: middlewares.ErrorHandler,
	})

	app.Static("/docs", "./src/app/public/docs")
	app.Static("/uploads", "./src/app/public/uploads")
	app.Get("/metrics", actions.PrometheusWithFastHTTPAdapter())
	app.Get("/meta", metaV1.GetVersion)

	app.Use(
		middlewares.WithBlogchainCORSPolicy(&service.HTTPAccessControl{
			AllowOrigins:     "*",
			AllowMethods:     "*",
			AllowHeaders:     "*",
			AllowCredentials: false,
			ExposeHeaders:    "",
			MaxAge:           0,
		}),
		middlewares.WithBlogchainXHeaderPolicy(),
		middlewares.UseBlogchainRealIP,
	)

	rsa := container.NewRSAContainer(
		container.TestPublicKey, container.TestPrivateKey,
	)

	httpController, err := actions.CreateHTTPControllerWithCopy(&actions.HTTPController{
		RSA:              &rsa,
		DB:               blogchain.Database(),
		Clickhouse:       blogchain.Clickhouse,
		ClickhouseBuffer: blogchain.ChBuffer,
		GeoReader:        blogchain.Reader,
		FsClient: &fsclient.FsClient{
			Uri:        ctx.String("cdn-host"),
			SecureType: 1,
			User:       ctx.String("cdn-user"),
			Password:   ctx.String("cdn-password"),
		},
	})
	if err != nil {
		return err
	}

	api := app.Group("/api",
		middlewares.WithBlogchainJWTAuthorization(&rsa),
		middlewares.WithBlogchainUserIdentity(blogchain.Database()),
	)
	api.Get("/healthcheck", actions.HealthCheck)
	api.Get("/runtime", actions.RuntimeStatistic(
		blogchain.Container.GetStartedAt(),
	))

	v1 := api.Group("/v1")
	v1.Get("/profile/:username", httpController.Profile)
	v1.Get("/content/:id", httpController.Content)
	v1.Get("/contents/:page?", httpController.Contents)
	v1.Get("/tag/:tag/:page?", httpController.Contents)
	v1.Get("/tags", httpController.Tags)
	v1.Get("/contents/user/:id/:page?", httpController.ContentsUser)

	withAccessControlPolicy := api.Use(
		middlewares.UseBlogchainAccessControlPolicy,
	)

	editor := withAccessControlPolicy.Group("/editor")
	editor.Get("/content/:id", httpController.ContentInformation)
	editor.Post("/content/add", httpController.ContentCreate)
	editor.Post("/content/update/:id", httpController.ContentUpdate)

	// authorization & authentication endpoints
	auth := app.Group("/auth", middlewares.UseBlogchainSignPolicy)
	auth.Post("/register", httpController.Register)
	auth.Post("/login", httpController.Login)
	auth.Post("/logout", httpController.Logout)

	go func() {
		ln, err := signal.ResolveListener(
			blogchain.Context(),
			ctx.Int("listener"),
			ctx.String("bind-socket"),
			ctx.String("bind-address"),
		)

		if err != nil {
			stop(err)
			return
		}

		if err := app.Listener(ln); err != nil {
			stop(err)
		}
	}()

	return await()
}
