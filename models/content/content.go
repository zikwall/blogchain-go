package content

import (
	"github.com/zikwall/blogchain/models/user"
)

type Content struct {
	Id         int64
	UserId     int
	Title      string
	Annotation string
	Content    string

	User user.User
}

type PublicContent struct {
	Title      string  `json:"title"`
	Annotation string  `json:"annotation"`
	Content    string  `json:"content"`
	Related    Related `json:"related"`
}

type Related struct {
	Publisher user.PublicUser `json:"publisher"`
}

func (c *Content) ToJSONAPI() PublicContent {
	return PublicContent{
		Title:      c.Title,
		Annotation: c.Annotation,
		Content:    c.Content,
		Related: Related{
			Publisher: user.PublicUser{
				Id:       c.User.Id,
				Username: c.User.Username,
				Profile: user.PublicProfile{
					Name:        c.User.Profile.Name,
					Email:       c.User.Profile.Name,
					Avatar:      c.User.Profile.Avatar.String,
					Location:    c.User.Profile.Location.String,
					Status:      c.User.Profile.Status.String,
					Description: c.User.Profile.Description.String,
				},
			},
		},
	}
}
