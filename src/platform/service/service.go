package service

var service *BlogchainServiceInstance

type (
	BlogchainServiceInstance struct {
		Notify
		HttpAccessControls BlogchainHttpAccessControl
		Container          *BlogchainServiceContainer
		database           *BlogchainDatabaseInstance
		logger             *BlogchainInternalLogger
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
	b.logger = NewBlogchainInternalLogger(c.IsDebug)

	database, err := NewBlogchainDatabaseInstance(c.BlogchainDatabaseConfiguration)

	if err != nil {
		return nil, err
	}

	dbLogger := BlogchainDatabaseLogger{}
	dbLogger.SetCallback(func(format string, v ...interface{}) {
		b.logger.Info(v)
	})

	database.SetLogger(dbLogger)

	b.database = database

	b.AddNotifiers(
		b.database,
		b.logger,
	)

	service = b

	return b, nil
}

func GetBlogchainServiceInstance() *BlogchainServiceInstance {
	return service
}

func (b *BlogchainServiceInstance) GetBlogchainDatabaseInstance() *BlogchainDatabaseInstance {
	return b.database
}

func (b *BlogchainServiceInstance) GetInternalLogger() *BlogchainInternalLogger {
	return b.logger
}
