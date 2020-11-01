package service

import "time"

type (
	BlogchainServiceContainer struct {
		secret    []byte
		startedAt time.Time
	}
	BlogchainServiceContainerConfiguration struct {
		Secret string
	}
)

func NewBlogchainServiceContainer(c BlogchainServiceContainerConfiguration) *BlogchainServiceContainer {
	return &BlogchainServiceContainer{
		secret:    []byte(c.Secret),
		startedAt: time.Now(),
	}
}

func (c BlogchainServiceContainer) GetContainerSecret() []byte {
	return c.secret
}

func (c BlogchainServiceContainer) GetStartedAt() time.Time {
	return c.startedAt
}
