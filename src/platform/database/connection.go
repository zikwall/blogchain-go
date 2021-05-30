package database

import (
	"context"
	"database/sql"
	"fmt"
	builder "github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zikwall/blogchain/src/platform/log"
	"strings"
	"time"
)

type (
	Instance struct {
		db *builder.Database
	}
	Configuration struct {
		Host     string
		User     string
		Password string
		Port     string
		Name     string
		Dialect  string
		Debug    bool
	}
	Logger struct {
		callback func(format string, v ...interface{})
	}
)

func (logger *Logger) SetCallback(callbak func(format string, v ...interface{})) {
	logger.callback = callbak
}

func (logger Logger) Printf(format string, v ...interface{}) {
	logger.callback(format, v)
}

func NewInstance(c context.Context, conf Configuration) (*Instance, error) {
	d := new(Instance)

	if conf.Dialect == "" {
		conf.Dialect = "mysql"
	}

	if conf.Host == "" {
		conf.Host = "@"
	}

	if strings.EqualFold(conf.Host, "") {
		conf.Host = "@"
	} else {
		if !strings.Contains(conf.Host, "@") {
			conf.Host = fmt.Sprintf("@tcp(%s)", conf.Host)
		}
	}

	db, err := sql.Open(conf.Dialect, makeConnectionString(conf))

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(c, 10*time.Second)

	defer func() {
		cancel()
	}()

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(25)
	db.SetMaxOpenConns(26)
	db.SetConnMaxLifetime(5 * time.Minute)

	dialect := builder.Dialect(conf.Dialect)
	d.db = dialect.DB(db)

	if conf.Debug {
		dbLogger := Logger{}
		dbLogger.SetCallback(func(format string, v ...interface{}) {
			log.Info(v)
		})

		d.SetLogger(dbLogger)
	}

	return d, nil
}

func (d *Instance) SetLogger(logger builder.Logger) {
	d.db.Logger(logger)
}

func (d *Instance) Builder() *builder.Database {
	return d.db
}

func makeConnectionString(c Configuration) string {
	return fmt.Sprintf("%s:%s%s/%s", c.User, c.Password, c.Host, c.Name)
}

// Close not implemented
func (d Instance) Close() error {
	return nil
}

func (d Instance) CloseMessage() string {
	return "Close database: this is not implemented"
}
