package content

import (
	"database/sql"
	"github.com/zikwall/blogchain/models/user"
)

type Content struct {
	Id         int64
	Uuid       string
	UserId     int64
	Title      string
	Annotation string
	Content    string
	CreatedAt  sql.NullInt64
	UpdatedAt  sql.NullInt64
	Image      sql.NullString

	User user.User
}

type PublicContent struct {
	Id         int64  `json:"id"`
	Uuid       string `json:"uuid"`
	Title      string `json:"title"`
	Annotation string `json:"annotation"`
	Content    string `json:"content"`
	CreatedAt  int64  `json:"created_at"`
	UpdatedAt  int64  `json:"updated_at"`
	Image      string `json:"image"`

	Related Related `json:"related"`
}

type Related struct {
	Publisher user.PublicUser `json:"publisher"`
}

func (c *Content) ToJSONAPI() PublicContent {
	return PublicContent{
		Id:         c.Id,
		Uuid:       c.Uuid,
		Title:      c.Title,
		Annotation: c.Annotation,
		Content:    c.Content,
		CreatedAt:  c.CreatedAt.Int64,
		UpdatedAt:  c.UpdatedAt.Int64,
		Image:      c.Image.String,
		Related: Related{
			Publisher: c.User.Properties(),
		},
	}
}
