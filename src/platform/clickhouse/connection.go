package clickhouse

import (
	"context"
	"errors"
	"fmt"
	"github.com/ClickHouse/clickhouse-go"
	"github.com/jmoiron/sqlx"
	"time"
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

func NewConnection(c context.Context, conf *Configuration) (*Connection, error) {
	connect, err := sqlx.Open("clickhouse", buildConnectionString(conf))

	if err != nil {
		return nil, err
	}

	connect.SetMaxIdleConns(20)
	connect.SetMaxOpenConns(21)
	connect.SetConnMaxLifetime(5 * time.Minute)

	ctx, cancel := context.WithTimeout(c, 10*time.Second)

	defer cancel()

	if err := connect.PingContext(ctx); err != nil {
		var e *clickhouse.Exception
		if errors.As(err, &e) {
			return nil, fmt.Errorf("[%d] %s \n%s", e.Code, e.Message, e.StackTrace)
		}

		return nil, err
	}

	ch := new(Connection)
	ch.db = connect
	ch.context = c

	return ch, nil
}

func (conn Connection) Query() *sqlx.DB {
	return conn.db
}

func buildConnectionString(c *Configuration) string {
	debug := "false"

	if c.IsDebug {
		debug = "true"
	}
	build := fmt.Sprintf(
		"tcp://%s:9000?debug=%s&username=%s&password=%s&database=%s",
		c.Address,
		debug,
		c.User,
		c.Password,
		c.Database,
	)
	if len(c.AltHosts) > 0 {
		build = fmt.Sprintf("%s&alt_hosts=%s", build, c.AltHosts)
	}

	return build
}

func (conn Connection) Close() error {
	return conn.db.Close()
}

func (conn Connection) CloseMessage() string {
	return "close Clickhouse connection"
}
