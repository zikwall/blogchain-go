package tag

import (
	"context"
	"github.com/zikwall/blogchain/src/platform/database"
)

type (
	Model struct {
		context    context.Context
		connection *database.Instance
	}
	Tag struct {
		Id    int64  `json:"id"`
		Name  string `json:"name"`
		Label string `json:"label"`
	}
	ContentTag struct {
		Id        int64 `json:"id"`
		ContentId int64 `json:"content_id"`
		TagId     int64 `json:"tag_id"`
	}
)

func ContextConnection(context context.Context, connection *database.Instance) Model {
	return Model{
		connection: connection,
		context:    context,
	}
}

func (t Model) Connection() *database.Instance {
	return t.connection
}
