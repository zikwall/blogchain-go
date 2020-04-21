package content

import (
	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/zikwall/blogchain/di"
	"github.com/zikwall/blogchain/models/content/forms"
	"github.com/zikwall/blogchain/models/user"
	"time"
)

func CreateContent(f *forms.ContentForm) (*Content, error) {
	c := &Content{
		Id:      0,
		UserId:  0,
		Title:   "",
		Content: "",
	}

	c.Content = f.Content
	c.Title = f.Title
	c.UserId = f.UserId

	status, err := di.DI().Database.Query().Insert("content", dbx.Params{
		"user_id":    c.UserId,
		"title":      c.Title,
		"content":    c.Content,
		"annotation": f.Annotation,
		"image":      f.Image,
		"created_at": time.Now().Unix(),
	}).Execute()

	c.Id, err = status.LastInsertId()

	return c, err
}

func UpdateContent(c *Content, f *forms.ContentForm) error {
	_, err := di.DI().Database.Query().Update("content", dbx.Params{
		"title":      f.Title,
		"content":    f.Content,
		"annotation": f.Annotation,
		"image":      f.Image,
		"updated_at": time.Now().Unix(),
	}, dbx.HashExp{"id": c.Id}).Execute()

	return err
}

func FindContentByIdAndUser(id int64, userid int64) (*Content, error) {
	c := &Content{
		Id:      0,
		UserId:  0,
		Title:   "",
		Content: "",
		User: user.User{
			Id:       0,
			Username: "",
			Email:    "",
			Profile:  user.Profile{},
		},
	}

	err := di.DI().Database.Query().
		Select(
			"content.*",
			"u.username as user.username",
			"p.name as user.profile.name",
			"p.public_email as user.profile.public_email",
			"p.avatar as user.profile.avatar",
		).
		From("content").
		LeftJoin("user u", dbx.NewExp("u.id=content.user_id")).
		LeftJoin("profile p", dbx.NewExp("p.user_id=u.id")).
		Where(dbx.HashExp{"content.id": id}).
		AndWhere(dbx.HashExp{"u.id": userid}).
		One(&c)

	return c, err
}

func FindContentById(id int64) (*Content, error) {
	c := &Content{
		Id:      0,
		UserId:  0,
		Title:   "",
		Content: "",
		User: user.User{
			Id:       0,
			Username: "",
			Email:    "",
			Profile:  user.Profile{},
		},
	}

	err := di.DI().Database.Query().
		Select(
			"content.*",
			"u.username as user.username",
			"p.name as user.profile.name",
			"p.public_email as user.profile.public_email",
			"p.avatar as user.profile.avatar",
		).
		From("content").
		LeftJoin("user u", dbx.NewExp("u.id=content.user_id")).
		LeftJoin("profile p", dbx.NewExp("p.user_id=u.id")).
		Where(dbx.HashExp{"content.id": id}).
		One(&c)

	return c, err
}

func FindAllContent() ([]PublicContent, error) {
	var c []Content

	err := di.DI().Database.Query().
		Select(
			"content.*",
			"u.username as user.username",
			"p.name as user.profile.name",
			"p.public_email as user.profile.public_email",
			"p.avatar as user.profile.avatar",
		).
		From("content").
		LeftJoin("user u", dbx.NewExp("u.id=content.user_id")).
		LeftJoin("profile p", dbx.NewExp("p.user_id=u.id")).
		All(&c)

	if err != nil {
		return nil, err
	}

	pc := []PublicContent{}
	for _, v := range c {
		pc = append(pc, v.ToJSONAPI())
	}

	return pc, err
}
