package repositories

import (
	"database/sql"
)

type Content struct {
	ID         int64          `db:"id"`
	UUID       string         `db:"uuid"`
	UserID     int64          `db:"user_id"`
	Title      string         `db:"title"`
	Annotation string         `db:"annotation"`
	Content    string         `db:"content"`
	CreatedAt  sql.NullInt64  `db:"created_at"`
	UpdatedAt  sql.NullInt64  `db:"updated_at"`
	Image      sql.NullString `db:"image"`

	User User `db:"user"`
	Tags []Tag
}

type PublicContent struct {
	ID         int64  `json:"id"`
	UUID       string `json:"uuid"`
	Title      string `json:"title"`
	Annotation string `json:"annotation"`
	Content    string `json:"content"`
	CreatedAt  int64  `json:"created_at"`
	UpdatedAt  int64  `json:"updated_at"`
	Image      string `json:"image"`

	Related Related `json:"related"`
}

type Related struct {
	Publisher PublicUser `json:"publisher"`
	Tags      []Tag      `json:"tags"`
}

func (content *Content) Response() PublicContent {
	return PublicContent{
		ID:         content.ID,
		UUID:       content.UUID,
		Title:      content.Title,
		Annotation: content.Annotation,
		Content:    content.Content,
		CreatedAt:  content.CreatedAt.Int64,
		UpdatedAt:  content.UpdatedAt.Int64,
		Image:      content.Image.String,
		Related: Related{
			Publisher: content.User.Properties(),
			Tags:      content.Tags,
		},
	}
}
