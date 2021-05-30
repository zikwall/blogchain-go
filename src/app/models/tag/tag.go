package tag

import (
	"github.com/zikwall/blogchain/src/platform/database"
)

type (
	TagModel struct {
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

func CreateTagConnection(connection *database.Instance) TagModel {
	return TagModel{
		connection: connection,
	}
}

func (t TagModel) Connection() *database.Instance {
	return t.connection
}
