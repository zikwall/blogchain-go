package clickhouse

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/ClickHouse/clickhouse-go"
	"github.com/jmoiron/sqlx"
)

type Connection struct {
	db      *sqlx.DB
	context context.Context
}

type Configuration struct {
	Address  string
	Password string
	User     string
	Database string
	AltHosts string
	IsDebug  bool
}

type Table struct {
	Name    string
	Columns []string
}

func NewConnection(ctx context.Context, cfg *Configuration) (*Connection, error) {
	connect, err := sqlx.Open("clickhouse", buildConnectionString(cfg))
	if err != nil {
		return nil, err
	}

	connect.SetMaxIdleConns(20)
	connect.SetMaxOpenConns(21)
	connect.SetConnMaxLifetime(5 * time.Minute)

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := connect.PingContext(ctx); err != nil {
		var e *clickhouse.Exception
		if errors.As(err, &e) {
			return nil, fmt.Errorf("[%d] %s \n%s", e.Code, e.Message, e.StackTrace)
		}
		return nil, err
	}

	connection := &Connection{
		db:      connect,
		context: ctx,
	}

	return connection, nil
}

func (c *Connection) Query() *sqlx.DB {
	return c.db
}

func (c *Connection) Drop() error {
	return c.db.Close()
}

func (c Connection) DropMsg() string {
	return "close clickhouse connection pool"
}

func buildConnectionString(cfg *Configuration) string {
	u := url.URL{
		Scheme: "tcp",
		Host:   cfg.Address + ":9000",
	}
	debug := "false"
	if cfg.IsDebug {
		debug = "true"
	}
	q := u.Query()
	q.Set("debug", debug)
	q.Set("username", cfg.User)
	q.Set("password", cfg.Password)
	q.Set("database", cfg.Database)
	if len(cfg.AltHosts) > 0 {
		q.Set("alt_hosts", cfg.AltHosts)
	}
	u.RawQuery = q.Encode()
	return u.String()
}
