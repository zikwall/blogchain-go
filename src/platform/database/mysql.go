package database

import (
	"database/sql"
	"fmt"
	builder "github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zikwall/blogchain/src/platform/log"
	"strings"
)

type (
	BlogchainDatabaseInstance struct {
		db *builder.Database
	}
	BlogchainDatabaseConfiguration struct {
		Host     string
		User     string
		Password string
		Port     string
		Name     string
		Dialect  string
		Debug    bool
	}
	DatabaseLogger struct {
		callback func(format string, v ...interface{})
	}
)

func (logger *DatabaseLogger) SetCallback(callbak func(format string, v ...interface{})) {
	logger.callback = callbak
}

func (logger DatabaseLogger) Printf(format string, v ...interface{}) {
	logger.callback(format, v)
}

func NewBlogchainDatabaseInstance(c BlogchainDatabaseConfiguration) (*BlogchainDatabaseInstance, error) {
	d := new(BlogchainDatabaseInstance)

	if c.Dialect == "" {
		c.Dialect = "mysql"
	}

	if c.Host == "" {
		c.Host = "@"
	}

	if strings.EqualFold(c.Host, "") {
		c.Host = "@"
	} else {
		if !strings.Contains(c.Host, "@") {
			c.Host = fmt.Sprintf("@tcp(%s)", c.Host)
		}
	}

	db, err := sql.Open(c.Dialect, makeBlogchainDatabaseConnectionString(c))

	if err != nil {
		return nil, err
	}

	dialect := builder.Dialect(c.Dialect)
	d.db = dialect.DB(db)

	if c.Debug {
		dbLogger := DatabaseLogger{}
		dbLogger.SetCallback(func(format string, v ...interface{}) {
			log.Info(v)
		})

		d.SetLogger(dbLogger)
	}

	return d, nil
}

func (d *BlogchainDatabaseInstance) SetLogger(logger builder.Logger) {
	d.db.Logger(logger)
}

func (d *BlogchainDatabaseInstance) Builder() *builder.Database {
	return d.db
}

func makeBlogchainDatabaseConnectionString(c BlogchainDatabaseConfiguration) string {
	return fmt.Sprintf("%s:%s%s/%s", c.User, c.Password, c.Host, c.Name)
}

// not implemented
func (d BlogchainDatabaseInstance) Close() error {
	return nil
}

func (d BlogchainDatabaseInstance) CloseMessage() string {
	return "Close database: this is not implemented"
}
