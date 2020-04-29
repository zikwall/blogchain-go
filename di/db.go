package di

import (
	"fmt"
	dbx "github.com/go-ozzo/ozzo-dbx"
	_ "github.com/go-sql-driver/mysql"
)

type DBConfig struct {
	Host string
	User string
	Pass string
	Port string
	Name string
	Driv string
}

type Database struct {
	DB *dbx.DB
}

func (db *Database) Close() {
	_ = db.DB.Close()
}

func (db *Database) Query() *dbx.DB {
	return db.DB
}

func (db *Database) Open(config DBConfig) {
	var err error

	db.DB, err = dbx.Open("mysql", connectionString(config))

	if err != nil {
		panic(err)
	}
}

func connectionString(c DBConfig) string {
	return fmt.Sprintf("%s:%s%s/%s", c.User, c.Pass, c.Host, c.Name)
}
