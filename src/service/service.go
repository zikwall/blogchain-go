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
		BloghainDatabaseConfiguration BloghainDatabaseConfiguration
		BlogchainHttpAccessControl    BlogchainHttpAccessControl
		BlogchainContainer            BlogchainServiceContainerConfiguration
		IsDebug                       bool
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

	database, err := NewBlogchainDatabaseInstance(c.BloghainDatabaseConfiguration)

	if err != nil {
		return nil, err
	}

	b.database = database
	b.logger = NewBlogchainInternalLogger(c.IsDebug)

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
