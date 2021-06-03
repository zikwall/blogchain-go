package clickhouse

import (
	"context"
	"errors"
	"fmt"
	"github.com/ClickHouse/clickhouse-go"
	"github.com/jmoiron/sqlx"
	"github.com/zikwall/blogchain/src/platform/log"
	"strings"
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

// Insert Currently, the client library does not support the JSONEachRow format, only native byte blocks
// There is no support for user interfaces as well as simple execution of an already prepared request
// The entire batch bid is implemented through so-called "transactions",
// although Clickhouse does not support them - it is only a client solution for preparing requests
func (c *Clickhouse) Insert(ctx context.Context, table Table, rows [][]interface{}) (uint64, error) {
	var affected uint64

	tx, err := c.db.BeginTx(ctx, nil)

	if err != nil {
		return 0, err
	}

	stmt, err := tx.PrepareContext(ctx, insertQuery(table.Name, table.Columns))

	if err != nil {
		// If you do not call the rollback function there will be a memory leak and goroutine
		// Such a leak can occur if there is no access to the table or there is no table itself
		if err = tx.Rollback(); err != nil {
			log.Warning(err)
		}

		return 0, err
	}

	defer func() {
		if err := stmt.Close(); err != nil {
			log.Warning(err)
		}
	}()

	timeoutContext, cancel := context.WithTimeout(ctx, time.Second*15)

	defer cancel()

	for _, row := range rows {
		// row affected is not supported
		if _, err := stmt.ExecContext(timeoutContext, row...); err == nil {
			affected += 1
		} else {
			log.Warning(err)
		}
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return affected, nil
}

// InsertWithMetrics use elastic apm metrics
func (c *Clickhouse) InsertWithMetrics(table Table, rows [][]interface{}) (uint64, error) {
	return c.Insert(c.context, table, rows)
}

func insertQuery(table string, cols []string) string {
	placeholders := make([]string, 0, len(cols))

	for range cols {
		placeholders = append(placeholders, "?")
	}

	prepared := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		table,
		strings.Join(cols, ", "),
		strings.Join(placeholders, ", "),
	)

	return prepared
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
