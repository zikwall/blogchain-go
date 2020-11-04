package service

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

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
	b.logger = NewBlogchainInternalLogger(c.IsDebug)

	database, err := NewBlogchainDatabaseInstance(c.BloghainDatabaseConfiguration)

	if err != nil {
		return nil, err
	}

	database.SetExecuteLoggerFunction(func(ctx context.Context, t time.Duration, sql string, result sql.Result, err error) {
		b.logger.Info(fmt.Sprintf("[%.2fms] Execute SQL: %v", float64(t.Milliseconds()), sql))
	})

	database.SetQueryLoggerFunction(func(ctx context.Context, t time.Duration, sql string, rows *sql.Rows, err error) {
		b.logger.Info(fmt.Sprintf("[%.2fms] Query SQL: %v", float64(t.Milliseconds()), sql))
	})

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
