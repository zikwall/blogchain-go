package clickhouse

import (
	"context"
	"errors"
	"fmt"
	"github.com/ClickHouse/clickhouse-go"
	"github.com/jmoiron/sqlx"
	"time"
)

type (
	Clickhouse struct {
		db      *sqlx.DB
		context context.Context
	}
	Configuration struct {
		Address  string
		Password string
		User     string
		Database string
		AltHosts string
		IsDebug  bool
	}
	Table struct {
		Name    string
		Columns []string
	}
)

func NewClickhouse(c context.Context, conf Configuration) (*Clickhouse, error) {
	connect, err := sqlx.Open("clickhouse", buildConnectionString(conf))

	if err != nil {
		return nil, err
	}

	connect.SetMaxIdleConns(20)
	connect.SetMaxOpenConns(21)
	connect.SetConnMaxLifetime(5 * time.Minute)

	ctx, cancel := context.WithTimeout(c, 10*time.Second)

	defer func() {
		cancel()
	}()

	if err := connect.PingContext(ctx); err != nil {
		var e *clickhouse.Exception
		if errors.As(err, &e) {
			return nil, fmt.Errorf("[%d] %s \n%s\n", e.Code, e.Message, e.StackTrace)
		}

		return nil, err
	}

	ch := new(Clickhouse)
	ch.db = connect
	ch.context = c

	return ch, nil
}

func (c Clickhouse) Query() *sqlx.DB {
	return c.db
}

func buildConnectionString(c Configuration) string {
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

func (c Clickhouse) Close() error {
	return c.db.Close()
}

func (c Clickhouse) CloseMessage() string {
	return "Close Clickhouse connection"
}
