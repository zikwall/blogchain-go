package tag

import "github.com/zikwall/blogchain/src/models"

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

func NewTagModel() TagModel {
	return TagModel{}
}
