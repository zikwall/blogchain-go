package service

type (
	BlogchainServiceContainer struct {
		secret []byte
	}
	BlogchainServiceContainerConfiguration struct {
		Secret string
	}
)

func NewBlogchainServiceContainer(c BlogchainServiceContainerConfiguration) *BlogchainServiceContainer {
	return &BlogchainServiceContainer{
		secret: []byte(c.Secret),
	}
}

func (c BlogchainServiceContainer) GetContainerSecret() []byte {
	return c.secret
}
