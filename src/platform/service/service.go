package service

import (
	"context"
	"fmt"
	"github.com/zikwall/blogchain/src/platform/clickhouse"
	"github.com/zikwall/blogchain/src/platform/container"
	"github.com/zikwall/blogchain/src/platform/database"
	"github.com/zikwall/blogchain/src/platform/log"
	"github.com/zikwall/blogchain/src/platform/maxmind"
	"github.com/zikwall/clickhouse-buffer"
	"runtime"
	"strconv"
	"time"
)

// Blogchain is basic structure is the "core" of the entire application and somewhat resembles dependency injection,
// but in a simpler form. It has such important properties as the top-level context and the ability to cancel it.
// It also includes all the main component instances, such as databases, queues and caches,
// and various internal schedulers.
type Blogchain struct {
	notify            Notify
	Container         *container.Container
	Clickhouse        *clickhouse.Connection
	ChBuffer          *clickhouse.BufferAdapter
	Reader            *maxmind.Reader
	Context           context.Context
	cancelRootContext context.CancelFunc
	database          *database.Connection
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

func CreateBlogchainService(ctx context.Context, c *Configuration) (*Blogchain, error) {
	blogchain := new(Blogchain)
	blogchain.Container = container.NewBlogchainServiceContainer()
	blogchain.Context, blogchain.cancelRootContext = context.WithCancel(ctx)

	finder, err := maxmind.CreateReader(c.ReaderConfig)

	if err != nil {
		return nil, err
	}

	blogchain.Reader = finder

	db, err := database.NewConnection(blogchain.Context, c.DatabaseConfiguration)

	if err != nil {
		return nil, err
	}

	blogchain.database = db

	ch, err := clickhouse.NewConnection(blogchain.Context, c.ClickhouseConfiguration)

	if err != nil {
		return nil, err
	}

	blogchain.Clickhouse = ch

	chBuffer, err := clickhousebuffer.NewClickhouseWithSqlx(blogchain.Clickhouse.Query())

	if err != nil {
		return nil, err
	}

	blogchain.ChBuffer = clickhouse.NewClickhouseBufferAdapter(
		clickhousebuffer.NewClientWithOptions(blogchain.Context, chBuffer,
			clickhousebuffer.DefaultOptions().
				SetFlushInterval(2000).
				SetBatchSize(5000),
		),
	)

	blogchain.notify.AddNotifiers(
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

// Shutdown method of implementing the cleaning of resources and connections when the application is shut down.
// Shutdown method addresses all connected listeners (which implement the corresponding interface)
// and calls the methods of this interface in turn.
// Shutdown accepts a single argument - a callback function with an argument of type error for custom error handling
// ```code
// 	Shutdown(func(err error) {
//		log.Warning(err)
//		bugsnag.Notify(err)
//	})
// ```
func (b *Blogchain) Shutdown(onError func(error)) {
	log.Info("shutdown Blogchain Service via System signal")

	// cancel the root context and clear all allocated resources
	b.cancelRootContext()
	for _, notifier := range b.notify.notifiers {
		log.Info(notifier.CloseMessage())

		if err := notifier.Close(); err != nil {
			onError(err)
		}
	}
}

// Stacktrace simple output of debugging information on the operation of the application
// and on the status of released resources (gorutin)
func (b *Blogchain) Stacktrace() {
	log.Info("waiting for the server completion report to be generated")

	<-time.After(time.Second * 2)

	memory := runtime.MemStats{}
	runtime.ReadMemStats(&memory)

	colored := func(category, context string) string {
		return fmt.Sprintf("%s: %s", log.Colored(category, log.Cyan), log.Colored(context, log.Green))
	}

	fmt.Printf(
		"%s \n \t - %s \n \t - %s \n",
		log.Colored("REPORT", log.Green),
		colored("number of remaining goroutines:", strconv.Itoa(runtime.NumGoroutine())),
		colored("number of operations of the garbage collector:", strconv.Itoa(int(memory.NumGC))),
	)
}
