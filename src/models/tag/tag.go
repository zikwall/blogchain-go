package tag

import (
	"github.com/zikwall/blogchain/src/models"
	"github.com/zikwall/blogchain/src/service"
)

type (
	TagModel struct {
		models.BlogchainModel
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

func NewTagModel(conn *service.BlogchainDatabaseInstance) TagModel {
	return TagModel{struct {
		Connection *service.BlogchainDatabaseInstance
	}{Connection: conn}}
}
