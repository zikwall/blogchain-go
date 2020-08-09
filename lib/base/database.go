package base

import (
	"fmt"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

type DatabaseConnection struct {
	connection *dbx.DB
	config     DatabaseConnectionConfig
}

type DatabaseConnectionConfig struct {
	Host     string
	User     string
	Password string
	Port     string
	Name     string
	Driver   string
}

func NewDatabaseConnection(config DatabaseConnectionConfig) (*DatabaseConnection, error) {
	d := &DatabaseConnection{}
	d.config = config
	if err := d.Open(); err != nil {
		return nil, err
	}

	return d, nil
}

func (d DatabaseConnection) Open() error {
	connection, err := dbx.Open("mysql", _safeConnectionString(d.config))

	if err != nil {
		return err
	}

	d.connection = connection
	return nil
}

func (d DatabaseConnection) Close() error {
	if err := d.connection.Close(); err != nil {
		return err
	}

	return nil
}

func (d DatabaseConnection) BuildQuery() *dbx.DB {
	return d.connection
}

func _safeConnectionString(c DatabaseConnectionConfig) string {
	return fmt.Sprintf("%s:%s%s/%s", c.User, c.Password, c.Host, c.Name)
}
