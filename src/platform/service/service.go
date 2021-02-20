package service

import (
	"github.com/zikwall/blogchain/src/platform/database"
	"github.com/zikwall/blogchain/src/platform/log"
)

type (
	BlogchainServiceInstance struct {
		Notify
		HttpAccessControls BlogchainHttpAccessControl
		Container          *BlogchainServiceContainer
		database           *database.BlogchainDatabaseInstance
	}
	BlogchainServiceConfiguration struct {
		BlogchainDatabaseConfiguration database.BlogchainDatabaseConfiguration
		BlogchainHttpAccessControl     BlogchainHttpAccessControl
		BlogchainContainer             BlogchainServiceContainerConfiguration
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

func NewBlogchainServiceInstance(c BlogchainServiceConfiguration) (*BlogchainServiceInstance, error) {
	b := new(BlogchainServiceInstance)
	b.HttpAccessControls = c.BlogchainHttpAccessControl
	b.Container = NewBlogchainServiceContainer(c.BlogchainContainer)

	db, err := database.NewBlogchainDatabaseInstance(c.BlogchainDatabaseConfiguration)

	if err != nil {
		return nil, err
	}

	dbLogger := database.BlogchainDatabaseLogger{}
	dbLogger.SetCallback(func(format string, v ...interface{}) {
		log.Info(v)
	})

	db.SetLogger(dbLogger)

	b.database = db

	b.AddNotifiers(
		b.database,
	)

	return b, nil
}

func (b *BlogchainServiceInstance) GetBlogchainDatabaseInstance() *database.BlogchainDatabaseInstance {
	return b.database
}
