package container

import "time"

type (
	Container struct {
		startedAt time.Time
	}
	Configuration struct{}
)

func NewBlogchainServiceContainer() *Container {
	return &Container{
		startedAt: time.Now(),
	}
}

func (c Container) GetStartedAt() time.Time {
	return c.startedAt
}
