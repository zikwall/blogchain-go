package service

import (
	"github.com/zikwall/blogchain/src/platform/container"
	"github.com/zikwall/blogchain/src/platform/database"
)

type (
	ServiceInstance struct {
		Notify
		HttpAccessControls BlogchainHttpAccessControl
		Container          *container.BlogchainServiceContainer
		database           *database.BlogchainDatabaseInstance
	}
	ServiceConfiguration struct {
		BlogchainDatabaseConfiguration database.BlogchainDatabaseConfiguration
		BlogchainHttpAccessControl     BlogchainHttpAccessControl
		BlogchainContainer             container.BlogchainServiceContainerConfiguration
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

func CreateService(c ServiceConfiguration) (*ServiceInstance, error) {
	b := new(ServiceInstance)
	b.HttpAccessControls = c.BlogchainHttpAccessControl
	b.Container = container.NewBlogchainServiceContainer(c.BlogchainContainer)

	db, err := database.NewBlogchainDatabaseInstance(c.BlogchainDatabaseConfiguration)

	if err != nil {
		return nil, err
	}

	b.database = db

	b.AddNotifiers(
		b.database,
	)

	return b, nil
}

func (b *ServiceInstance) GetBlogchainDatabaseInstance() *database.BlogchainDatabaseInstance {
	return b.database
}
