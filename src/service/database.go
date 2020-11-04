package service

import (
	"fmt"
	dbx "github.com/go-ozzo/ozzo-dbx"
	_ "github.com/go-sql-driver/mysql"
)

type (
	BlogchainDatabaseInstance struct {
		db *dbx.DB
	}
	BloghainDatabaseConfiguration struct {
		Host     string
		User     string
		Password string
		Port     string
		Name     string
		Driver   string
	}
)

func NewBlogchainDatabaseInstance(c BloghainDatabaseConfiguration) (*BlogchainDatabaseInstance, error) {
	d := new(BlogchainDatabaseInstance)

	if c.Driver == "" {
		c.Driver = "mysql"
	}

	if c.Host == "" {
		c.Host = "@"
	}

	db, err := dbx.Open(c.Driver, makeBlogchainDatabaseConnectionString(c))

	if err != nil {
		return nil, err
	}

	d.db = db

	return d, nil
}

func (d *BlogchainDatabaseInstance) SetQueryLoggerFunction(f dbx.QueryLogFunc) {
	d.db.QueryLogFunc = f
}

func (d *BlogchainDatabaseInstance) SetExecuteLoggerFunction(f dbx.ExecLogFunc) {
	d.db.ExecLogFunc = f
}

func (d *BlogchainDatabaseInstance) Query() *dbx.DB {
	return d.db
}

func makeBlogchainDatabaseConnectionString(c BloghainDatabaseConfiguration) string {
	return fmt.Sprintf("%s:%s%s/%s", c.User, c.Password, c.Host, c.Name)
}

func (d BlogchainDatabaseInstance) Close() error {
	return d.db.Close()
}

func (d BlogchainDatabaseInstance) CloseMessage() string {
	return "Close database"
}
