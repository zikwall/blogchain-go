package clickhouse

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go"
	"github.com/jmoiron/sqlx"
	"github.com/zikwall/blogchain/src/platform/log"
	"strings"
	"time"
)

type (
	Clickhouse struct {
		db *sqlx.DB
	}
	ClickhouseConfiguration struct {
		Address  string
		Password string
		User     string
		Database string
		IsDebug  bool
	}
	Table struct {
		Name    string
		Columns []string
	}
)

func NewClickhouse(conf ClickhouseConfiguration) (*Clickhouse, error) {
	connect, err := sqlx.Open("clickhouse", buildConnectionString(conf))

	if err != nil {
		return nil, err
	}

	if err := connect.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			return nil, fmt.Errorf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		}

		return nil, err
	}

	ch := new(Clickhouse)
	ch.db = connect

	return ch, nil
}

func (c Clickhouse) Query() *sqlx.DB {
	return c.db
}

// Currently, the client library does not support the JSONEachRow format, only native byte blocks
// There is no support for user interfaces as well as simple execution of an already prepared request
// The entire batch bid is implemented through so-called "transactions",
// although Clickhouse does not support them - it is only a client solution for preparing requests
func (c *Clickhouse) Insert(ctx context.Context, table Table, rows [][]interface{}) (uint64, error) {
	var affected uint64

	tx, err := c.db.Begin()

	if err != nil {
		return 0, err
	}

	stmt, err := tx.Prepare(insertQuery(table.Name, table.Columns))

	if err != nil {
		// If you do not call the rollback function there will be a memory leak and goroutine
		// Such a leak can occur if there is no access to the table or there is no table itself
		if err := tx.Rollback(); err != nil {
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

// use elastic apm metrics
func (c *Clickhouse) InsertWithMetrics(table Table, rows [][]interface{}) (uint64, error) {
	return c.Insert(context.Background(), table, rows)
}

func insertQuery(table string, cols []string) string {
	placeholders := []string{}

	for _, _ = range cols {
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

func buildConnectionString(c ClickhouseConfiguration) string {
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

	return build
}

func (c Clickhouse) Close() error {
	return c.db.Close()
}

func (c Clickhouse) CloseMessage() string {
	return "Close Clickhouse connection"
}