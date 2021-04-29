package service

import (
	"context"
	"github.com/zikwall/blogchain/src/platform/clickhouse"
	"github.com/zikwall/blogchain/src/platform/container"
	"github.com/zikwall/blogchain/src/platform/database"
	"github.com/zikwall/blogchain/src/platform/maxmind"
)

type (
	ServiceInstance struct {
		Notify
		HttpAccessControls BlogchainHttpAccessControl
		Container          *container.BlogchainServiceContainer
		Clickhouse         *clickhouse.Clickhouse
		Finder             *maxmind.Finder
		Context            context.Context
		cancelFunc         context.CancelFunc
		database           *database.BlogchainDatabaseInstance
	}
	ServiceConfiguration struct {
		BlogchainDatabaseConfiguration database.BlogchainDatabaseConfiguration
		BlogchainHttpAccessControl     BlogchainHttpAccessControl
		BlogchainContainer             container.BlogchainServiceContainerConfiguration
		ClickhouseConfiguration        clickhouse.ClickhouseConfiguration
		FinderConfig                   maxmind.FinderConfig
		IsDebug                        bool
	}
	BlogchainHttpAccessControl struct {
		AllowOrigins     string
		AllowMethods     string
		AllowHeaders     string
		AllowCredentials bool
		ExposeHeaders    string
		MaxAge           int
	}
)

func CreateService(ctx context.Context, c ServiceConfiguration) (*ServiceInstance, error) {
	b := new(ServiceInstance)
	b.HttpAccessControls = c.BlogchainHttpAccessControl
	b.Container = container.NewBlogchainServiceContainer(c.BlogchainContainer)
	b.Context, b.cancelFunc = context.WithCancel(ctx)

	finder, err := maxmind.CreateFinder(c.FinderConfig)

	if err != nil {
		return nil, err
	}

	b.Finder = finder

	db, err := database.NewBlogchainDatabaseInstance(c.BlogchainDatabaseConfiguration)

	if err != nil {
		return nil, err
	}

	b.database = db

	ch, err := clickhouse.NewClickhouse(c.ClickhouseConfiguration)

	if err != nil {
		return nil, err
	}

	b.Clickhouse = ch

	b.AddNotifiers(
		b.database,
		b.Clickhouse,
		b.Finder,
	)

	return b, nil
}

func (b *ServiceInstance) GetBlogchainDatabaseInstance() *database.BlogchainDatabaseInstance {
	return b.database
}
