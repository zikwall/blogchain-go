package database

import (
	"context"
	"github.com/doug-martin/goqu/v9"
)

type Repository interface {
	Connection() *goqu.Database
}

type Connector interface {
	Connection() *Connection
	Context() context.Context
}
