package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/urfave/cli/v2"
	"github.com/zikwall/blogchain/src/app/actions"
	"github.com/zikwall/blogchain/src/app/lib/upload"
	"github.com/zikwall/blogchain/src/app/middlewares"
	"github.com/zikwall/blogchain/src/platform/clickhouse"
	"github.com/zikwall/blogchain/src/platform/container"
	"github.com/zikwall/blogchain/src/platform/database"
	"github.com/zikwall/blogchain/src/platform/log"
	"github.com/zikwall/blogchain/src/platform/maxmind"
	"github.com/zikwall/blogchain/src/platform/service"
	"github.com/zikwall/fsclient"
	"net"
	"os"
	"strings"
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
				Value:    ListenerTCP,
				EnvVars:  []string{"LISTENER"},
			},

			// database
			&cli.StringFlag{
				Name:     "database-host",
				Required: true,
				Usage:    "Database host",
				EnvVars:  []string{"DATABASE_HOST"},
			},
			&cli.StringFlag{
				Name:     "database-user",
				Required: true,
				Usage:    "Database user",
				EnvVars:  []string{"DATABASE_USER"},
			},
			&cli.StringFlag{
				Name:     "database-password",
				Required: true,
				Usage:    "Database password",
				EnvVars:  []string{"DATABASE_PASSWORD"},
			},
			&cli.StringFlag{
				Name:     "database-name",
				Required: true,
				Usage:    "Database name",
				EnvVars:  []string{"DATABASE_NAME"},
			},
			&cli.StringFlag{
				Name:     "database-dialect",
				Required: true,
				Usage:    "Database dialect: mysql, postgres, sqlite3, sqlserver etc",
				EnvVars:  []string{"DATABASE_DIALECT"},
			},
			&cli.StringFlag{
				Name:     "rsa-public-key",
				Required: false,
				Usage:    "Container secret key for JWT, and etc.",
				EnvVars:  []string{"RSA_PUBLIC_KEY"},
			},
			&cli.StringFlag{
				Name:     "rsa-private-key",
				Required: false,
				Usage:    "Container secret key for JWT, and etc.",
				EnvVars:  []string{"RSA_PRIVATE_KEY"},
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

			//cdn
			&cli.StringFlag{
				Name:    "cdn-host",
				Usage:   "",
				Value:   "localhost:1337",
				EnvVars: []string{"CDN_HOST"},
			},
			&cli.StringFlag{
				Name:    "cdn-user",
				Usage:   "",
				EnvVars: []string{"CDN_USER"},
			},
			&cli.StringFlag{
				Name:    "cdn-password",
				Usage:   "",
				EnvVars: []string{"CDN_PASSWORD"},
			},

			// dev
			&cli.BoolFlag{
				Name:    "debug",
				Usage:   "Debug mode - details the stages of operation of the service, also in this mode, all logs are sent to stdout",
				EnvVars: []string{"DEBUG"},
			},
		},
	}

	application.Action = func(c *cli.Context) error {
		blogchain, err := service.CreateService(
			context.Background(),
			service.Configuration{
				DatabaseConfiguration: database.Configuration{
					Host:     c.String("database-host"),
					User:     c.String("database-user"),
					Password: c.String("database-password"),
					Name:     c.String("database-name"),
					Dialect:  c.String("database-dialect"),
					Debug:    c.Bool("debug"),
				},
				Container: container.Configuration{},
				ClickhouseConfiguration: clickhouse.Configuration{
					Address:  c.String("clickhouse-address"),
					User:     c.String("clickhouse-user"),
					Password: c.String("clickhouse-password"),
					Database: c.String("clickhouse-database"),
					AltHosts: c.String("clickhouse-alt-hosts"),
					IsDebug:  c.Bool("debug"),
				},
				FinderConfig: maxmind.FinderConfig{
					Path: c.String("maxmind-mmdb"),
				},
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

		app.Use(
			middlewares.WithBlogchainCORSPolicy(service.HttpAccessControl{
				AllowOrigins:     "*",
				AllowMethods:     "*",
				AllowHeaders:     "*",
				AllowCredentials: false,
				ExposeHeaders:    "",
				MaxAge:           0,
			}),
			middlewares.WithBlogchainXHeaderPolicy(),
			middlewares.UseBlogchainRealIp,
		)

		rsa := container.NewRSAContainer(
			container.TestPublicKey, container.TestPrivateKey,
		)

		fsClient, err := fsclient.WithCopyFsClient(fsclient.FsClient{
			Uri:        c.String("cdn-host"),
			SecureType: 1,
			User:       c.String("cdn-user"),
			Password:   c.String("cdn-password"),
		})

		if err != nil {
			return err
		}

		httpController := actions.CreateHttpControllerWithCopy(blogchain.Context, actions.HttpController{
			RSA:        &rsa,
			Db:         blogchain.GetDatabaseConnection(),
			Clickhouse: blogchain.Clickhouse,
			Finder:     blogchain.Finder,
			Uploader:   upload.NewFileUploader(fsClient),
		})

		api := app.Group("/api",
			middlewares.WithBlogchainJWTAuthorization(&rsa),
			middlewares.WithBlogchainUserIdentity(blogchain),
		)
		{
			api.Get("/healthcheck", actions.HealthCheck)
			api.Get("/runtime", actions.RuntimeStatistic(
				blogchain.Container.GetStartedAt(),
			))

			v1 := api.Group("/v1")
			{
				v1.Get("/profile/:username", httpController.Profile)
				v1.Get("/content/:id", httpController.Content)
				v1.Get("/contents/:page?", httpController.Contents)
				v1.Get("/tag/:tag/:page?", httpController.Contents)
				v1.Get("/tags", httpController.Tags)
				v1.Get("/contents/user/:id/:page?", httpController.ContentsUser)
			}

			withAccessControlPolicy := api.Use(
				middlewares.UseBlogchainAccessControlPolicy,
			)

			editor := withAccessControlPolicy.Group("/editor")
			{
				editor.Get("/content/:id", httpController.ContentInformation)
				editor.Post("/content/add", httpController.ContentCreate)
				editor.Post("/content/update/:id", httpController.ContentUpdate)
			}
		}

		// authorization & authentication endpoints
		auth := app.Group("/auth", middlewares.UseBlogchainSignPolicy)
		{
			auth.Post("/register", httpController.Register)
			auth.Post("/login", httpController.Login)
			auth.Post("/logout", httpController.Logout)
		}

		// statistic endpoints
		stats := app.Group("/statistic")
		{
			stats.Post("/post/push", httpController.PushPostStats)
		}

		await, stop := notifier(conf{
			onSignal: func() {
				log.Info("Received a system signal to shut down the application, start the shutdown process..")
			},
		})

		go func() {
			ln, err := resolveListener(
				blogchain.Context,
				c.Int("listener"),
				c.String("bind-socket"),
				c.String("bind-address"),
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

	if err := application.Run(os.Args); err != nil {
		log.Error(err)
	}
}

func resolveListener(context context.Context, listener int, uds, tcp string) (net.Listener, error) {
	if listener == ListenerUDS {
		defer maybeChmodSocket(context, uds)
		ln, err := listenToUnix(uds)

		return ln, err
	}

	if !strings.Contains(tcp, ":") {
		tcp = ":" + tcp
	}

	ln, err := net.Listen(fiber.NetworkTCP4, tcp)

	if err != nil {
		return nil, err
	}

	return ln, nil
}
