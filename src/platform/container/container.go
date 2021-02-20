package container

import "time"

type (
	BlogchainServiceContainer struct {
		startedAt time.Time
	}
	BlogchainServiceContainerConfiguration struct{}
)

func NewBlogchainServiceContainer(c BlogchainServiceContainerConfiguration) *BlogchainServiceContainer {
	return &BlogchainServiceContainer{
		startedAt: time.Now(),
	}
}

func (c BlogchainServiceContainer) GetStartedAt() time.Time {
	return c.startedAt
}
