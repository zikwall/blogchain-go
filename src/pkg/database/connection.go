package database

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zikwall/blogchain/src/pkg/log"

	builder "github.com/doug-martin/goqu/v9"
	// nolint:golint // reason
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	_ "github.com/go-sql-driver/mysql"
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

func (cfg *Configuration) checkBeforeInitializing() {
	if cfg.Dialect == "" {
		cfg.Dialect = "mysql"
	}

	if cfg.Host == "" {
		cfg.Host = "@"
	}

	if strings.EqualFold(cfg.Host, "") {
		cfg.Host = "@"
	} else if !strings.Contains(cfg.Host, "@") {
		cfg.Host = fmt.Sprintf("@tcp(%s)", cfg.Host)
	}
}

func NewConnection(ctx context.Context, cfg *Configuration) (*Connection, error) {
	cfg.checkBeforeInitializing()

	db, err := sql.Open(cfg.Dialect, buildConnectionString(cfg))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(25)
	db.SetMaxOpenConns(26)
	db.SetConnMaxLifetime(5 * time.Minute)

	dialect := builder.Dialect(cfg.Dialect)
	connection := &Connection{
		db: dialect.DB(db),
	}

	if cfg.Debug {
		dbLogger := &Logger{}
		dbLogger.SetCallback(func(format string, v ...interface{}) {
			log.Info(v)
		})
		connection.SetLogger(dbLogger)
	}

	return connection, nil
}

func (c *Connection) SetLogger(logger builder.Logger) {
	c.db.Logger(logger)
}

func (c *Connection) Builder() *builder.Database {
	return c.db
}

func buildConnectionString(cfg *Configuration) string {
	return fmt.Sprintf("%s:%s%s/%s", cfg.User, cfg.Password, cfg.Host, cfg.Name)
}

// Drop close not implemented in database
func (c *Connection) Drop() error {
	return nil
}

func (c *Connection) DropMsg() string {
	return "close database: this is not implemented"
}
