package repositories

import (
	"context"
	"github.com/zikwall/blogchain/src/platform/database"
)

type Repository struct {
	connection *database.Instance
	context    context.Context
}

func (r *Repository) Context() context.Context {
	if r.context == nil {
		return context.Background()
	}

	return r.context
}

func (r *Repository) Connection() *database.Instance {
	return r.connection
}
