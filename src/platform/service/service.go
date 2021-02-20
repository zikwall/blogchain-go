package service

import "github.com/zikwall/blogchain/src/platform/log"

type (
	BlogchainServiceInstance struct {
		Notify
		HttpAccessControls BlogchainHttpAccessControl
		Container          *BlogchainServiceContainer
		database           *BlogchainDatabaseInstance
	}
	BlogchainServiceConfiguration struct {
		BlogchainDatabaseConfiguration BlogchainDatabaseConfiguration
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

	database, err := NewBlogchainDatabaseInstance(c.BlogchainDatabaseConfiguration)

	if err != nil {
		return nil, err
	}

	dbLogger := BlogchainDatabaseLogger{}
	dbLogger.SetCallback(func(format string, v ...interface{}) {
		log.Info(v)
	})

	database.SetLogger(dbLogger)

	b.database = database

	b.AddNotifiers(
		b.database,
	)

	return b, nil
}

func (b *BlogchainServiceInstance) GetBlogchainDatabaseInstance() *BlogchainDatabaseInstance {
	return b.database
}
