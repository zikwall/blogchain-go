package models

import (
	"database/sql"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	"log"
)
import _ "github.com/go-sql-driver/mysql"

type Builder struct {
	db *goqu.Database
}

var builder = NewBuilder()

func NewBuilder() Builder {
	dialect := goqu.Dialect("mysql")

	db, err := sql.Open("mysql", "blogchain:123456@/blogchain")

	if err != nil {
		log.Fatal(err)
	}

	b := Builder{}
	b.db = dialect.DB(db)

	return b
}

func (b *Builder) Builder() *goqu.Database {
	return b.db
}
