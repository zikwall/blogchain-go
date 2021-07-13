package database

import (
	"context"
	"database/sql"
	"fmt"
	builder "github.com/doug-martin/goqu/v9"
	// nolint:golint // reason
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zikwall/blogchain/src/platform/log"
	"strings"
	"time"
)

type Connection struct {
	db *builder.Database
}

type Configuration struct {
	Host     string
	User     string
	Password string
	Port     string
	Name     string
	Dialect  string
	Debug    bool
}

func (conf *Configuration) checkBeforeInitializing() {
	if conf.Dialect == "" {
		conf.Dialect = "mysql"
	}

	if conf.Host == "" {
		conf.Host = "@"
	}

	if strings.EqualFold(conf.Host, "") {
		conf.Host = "@"
	} else if !strings.Contains(conf.Host, "@") {
		conf.Host = fmt.Sprintf("@tcp(%s)", conf.Host)
	}
}

func NewConnection(c context.Context, conf *Configuration) (*Connection, error) {
	conf.checkBeforeInitializing()

	db, err := sql.Open(conf.Dialect, buildConnectionString(conf))

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(25)
	db.SetMaxOpenConns(26)
	db.SetConnMaxLifetime(5 * time.Minute)

	dialect := builder.Dialect(conf.Dialect)
	connection := new(Connection)
	connection.db = dialect.DB(db)

	if conf.Debug {
		dbLogger := Logger{}
		dbLogger.SetCallback(func(format string, v ...interface{}) {
			log.Info(v)
		})

		connection.SetLogger(dbLogger)
	}

	return connection, nil
}

func (conn *Connection) SetLogger(logger builder.Logger) {
	conn.db.Logger(logger)
}

func (conn *Connection) Builder() *builder.Database {
	return conn.db
}

func buildConnectionString(c *Configuration) string {
	return fmt.Sprintf("%s:%s%s/%s", c.User, c.Password, c.Host, c.Name)
}

// Close not implemented
func (conn Connection) Close() error {
	return nil
}

func (conn Connection) CloseMessage() string {
	return "close database: this is not implemented"
}
