package repositories

import (
	"context"
	"github.com/zikwall/blogchain/src/platform/database"
)

type Repository struct {
	connection *database.Connection
	context    context.Context
}

func (r *Repository) Context() context.Context {
	if r.context == nil {
		return context.Background()
	}

	return r.context
}

func (r *Repository) Connection() *database.Connection {
	return r.connection
}
