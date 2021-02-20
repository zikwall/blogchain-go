package database

import "github.com/doug-martin/goqu/v9"

type Repository interface {
	Connection() *goqu.Database
}
