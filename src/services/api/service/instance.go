package service

import (
	"context"

	"github.com/zikwall/blogchain/src/pkg/clickhouse"
	"github.com/zikwall/blogchain/src/pkg/container"
	"github.com/zikwall/blogchain/src/pkg/database"
	"github.com/zikwall/blogchain/src/pkg/drop"
	"github.com/zikwall/blogchain/src/pkg/maxmind"

	"github.com/zikwall/clickhouse-buffer"
)

// Blogchain is basic structure is the "core" of the entire application and somewhat resembles dependency injection,
// but in a simpler form. It has such important properties as the top-level context and the ability to cancel it.
// It also includes all the main component instances, such as databases, queues and caches,
// and various internal schedulers.
type Blogchain struct {
	Container  *container.Container
	Clickhouse *clickhouse.Connection
	ChBuffer   *clickhouse.BufferAdapter
	Reader     *maxmind.Reader
	database   *database.Connection
	*drop.Impl
}

type Configuration struct {
	DatabaseConfiguration   *database.Configuration
	Container               container.Configuration
	ClickhouseConfiguration *clickhouse.Configuration
	ReaderConfig            maxmind.ReaderConfig
	IsDebug                 bool
}

type HTTPAccessControl struct {
	AllowOrigins     string
	AllowMethods     string
	AllowHeaders     string
	AllowCredentials bool
	ExposeHeaders    string
	MaxAge           int
}

func New(ctx context.Context, c *Configuration) (*Blogchain, error) {
	var err error
	blogchain := &Blogchain{
		Container: container.NewBlogchainServiceContainer(),
	}
	blogchain.Impl = drop.NewContext(ctx)

	blogchain.Reader, err = maxmind.CreateReader(c.ReaderConfig)
	if err != nil {
		return nil, err
	}

	blogchain.database, err = database.NewConnection(blogchain.Context(), c.DatabaseConfiguration)
	if err != nil {
		return nil, err
	}

	blogchain.Clickhouse, err = clickhouse.NewConnection(blogchain.Context(), c.ClickhouseConfiguration)
	if err != nil {
		return nil, err
	}

	chBuffer, err := clickhousebuffer.NewClickhouseWithSqlx(blogchain.Clickhouse.Query())
	if err != nil {
		return nil, err
	}

	blogchain.ChBuffer = clickhouse.NewClickhouseBufferAdapter(
		clickhousebuffer.NewClientWithOptions(blogchain.Context(), chBuffer,
			clickhousebuffer.DefaultOptions().
				SetFlushInterval(2000).
				SetBatchSize(5000),
		),
	)

	blogchain.AddDroppers(
		blogchain.database,
		blogchain.Clickhouse,
		blogchain.Reader,
		blogchain.ChBuffer,
	)

	return blogchain, nil
}

func (b *Blogchain) Database() *database.Connection {
	return b.database
}
