package service

import (
	"context"
	"fmt"
	"github.com/zikwall/blogchain/src/platform/clickhouse"
	"github.com/zikwall/blogchain/src/platform/container"
	"github.com/zikwall/blogchain/src/platform/database"
	"github.com/zikwall/blogchain/src/platform/log"
	"github.com/zikwall/blogchain/src/platform/maxmind"
	"runtime"
	"strconv"
	"time"
)

type (
	Instance struct {
		notify     Notify
		Container  *container.Container
		Clickhouse *clickhouse.Clickhouse
		Finder     *maxmind.Finder
		Context    context.Context
		cancelFunc context.CancelFunc
		database   *database.Instance
	}
	Configuration struct {
		DatabaseConfiguration   database.Configuration
		Container               container.Configuration
		ClickhouseConfiguration clickhouse.Configuration
		FinderConfig            maxmind.FinderConfig
		IsDebug                 bool
	}
	HttpAccessControl struct {
		AllowOrigins     string
		AllowMethods     string
		AllowHeaders     string
		AllowCredentials bool
		ExposeHeaders    string
		MaxAge           int
	}
)

func CreateService(ctx context.Context, c Configuration) (*Instance, error) {
	b := new(Instance)
	b.Container = container.NewBlogchainServiceContainer(c.Container)
	b.Context, b.cancelFunc = context.WithCancel(ctx)

	finder, err := maxmind.CreateFinder(c.FinderConfig)

	if err != nil {
		return nil, err
	}

	b.Finder = finder

	db, err := database.NewInstance(b.Context, c.DatabaseConfiguration)

	if err != nil {
		return nil, err
	}

	b.database = db

	ch, err := clickhouse.NewClickhouse(b.Context, c.ClickhouseConfiguration)

	if err != nil {
		return nil, err
	}

	b.Clickhouse = ch

	b.notify.AddNotifiers(
		b.database,
		b.Clickhouse,
		b.Finder,
	)

	return b, nil
}

func (b *Instance) GetDatabaseInstance() *database.Instance {
	return b.database
}

func (s Instance) Shutdown(onError func(error)) {
	log.Info("Shutdown Blogchain Service via System signal")

	// cancel root context
	s.cancelFunc()

	for _, notifier := range s.notify.notifiers {
		log.Info(notifier.CloseMessage())

		if err := notifier.Close(); err != nil {
			onError(err)
		}
	}
}

func (s Instance) Stacktrace() {
	log.Info("Waiting for the server completion report to be generated")

	<-time.After(time.Second * 2)

	memory := runtime.MemStats{}
	runtime.ReadMemStats(&memory)

	colored := func(category, context string) string {
		return fmt.Sprintf("%s: %s", log.Colored(category, log.Cyan), log.Colored(context, log.Green))
	}

	fmt.Printf(
		"%s \n \t - %s \n \t - %s \n",
		log.Colored("REPORT", log.Green),
		colored("Number of remaining goroutines:", strconv.Itoa(runtime.NumGoroutine())),
		colored("Number of operations of the garbage collector:", strconv.Itoa(int(memory.NumGC))),
	)
}
