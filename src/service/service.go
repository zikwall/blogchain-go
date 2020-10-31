package service

var service *BlogchainServiceInstance

type (
	BlogchainServiceInstance struct {
		database *BlogchainDatabaseInstance
	}
	BlogchainServiceConfiguration struct {
		BloghainDatabaseConfiguration BloghainDatabaseConfiguration
	}
)

func NewBlogchainServiceInstance(c BlogchainServiceConfiguration) (*BlogchainServiceInstance, error) {
	b := new(BlogchainServiceInstance)

	database, err := NewBlogchainDatabaseInstance(c.BloghainDatabaseConfiguration)

	if err != nil {
		return nil, err
	}

	b.database = database

	service = b

	return b, nil
}

func GetBlogchainServiceInstance() *BlogchainServiceInstance {
	return service
}

func (b *BlogchainServiceInstance) GetBlogchainDatabaseInstance() *BlogchainDatabaseInstance {
	return b.database
}
