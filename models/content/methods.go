package content

import (
	"fmt"
	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/gofiber/fiber"
	uuid "github.com/satori/go.uuid"
	"github.com/zikwall/blogchain/di"
	"github.com/zikwall/blogchain/models/content/forms"
	"github.com/zikwall/blogchain/models/tag"
	"github.com/zikwall/blogchain/models/user"
	"time"
)

func createImagePath(uuidv4 string) string {
	return fmt.Sprintf("%s.png", uuidv4)
}

func CreateContent(f *forms.ContentForm, c *fiber.Ctx) (*Content, error) {
	content := &Content{
		Id:      0,
		UserId:  0,
		Title:   "",
		Content: "",
	}

	content.Content = f.Content
	content.Title = f.Title
	content.UserId = f.UserId
	content.Annotation = f.Annotation

	var err error
	uv4 := uuid.Must(uuid.NewV4(), err)
	content.Uuid = uv4.String()

	if f.GetImage().Err == nil {
		content.Image.String = createImagePath(content.Uuid)
		_ = c.SaveFile(f.GetImage().File, fmt.Sprintf("./public/uploads/%s", content.Image.String))
	}

	status, err := di.DI().Database.Query().Insert("content", dbx.Params{
		"uuid":       content.Uuid,
		"user_id":    content.UserId,
		"title":      content.Title,
		"content":    content.Content,
		"annotation": f.Annotation,
		"image":      content.Image.String,
		"created_at": time.Now().Unix(),
	}).Execute()

	content.Id, err = status.LastInsertId()

	return content, err
}

func UpdateContent(content *Content, f *forms.ContentForm, c *fiber.Ctx) error {
	if f.GetImage().Err == nil {
		content.Image.String = createImagePath(content.Uuid)
		err := c.SaveFile(f.GetImage().File, fmt.Sprintf("./public/uploads/%s", content.Image.String))

		if err != nil {
			panic(err)
		}
	}

	_, err := di.DI().Database.Query().Update("content", dbx.Params{
		"title":      f.Title,
		"content":    f.Content,
		"annotation": f.Annotation,
		"image":      content.Image.String,
		"updated_at": time.Now().Unix(),
	}, dbx.HashExp{"id": content.Id}).Execute()

	return err
}

func Find() *dbx.SelectQuery {
	query :=
		di.DI().Database.Query().
			Select(
				"content.*",
				"u.username as user.username",
				"p.name as user.profile.name",
				"p.public_email as user.profile.public_email",
				"p.avatar as user.profile.avatar",
			).
			From("content").
			LeftJoin("user u", dbx.NewExp("u.id=content.user_id")).
			LeftJoin("profile p", dbx.NewExp("p.user_id=u.id"))

	return query
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

	err := Find().
		Where(dbx.HashExp{"content.id": id}).
		AndWhere(dbx.HashExp{"u.id": userid}).
		One(&c)

	if err == nil {
		err = c.WithTags()
	}

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

	err := Find().
		Where(dbx.HashExp{"content.id": id}).
		One(&c)

	if err == nil {
		err = c.WithTags()
	}

	return c, err
}

func FindAllContent(label string) ([]PublicContent, error) {
	var c []Content

	query := Find()

	if label != "" {
		tag.AttachTagQuery(query, label)
	}

	err := query.All(&c)

	if err != nil {
		return nil, err
	}

	pc := []PublicContent{}
	for _, v := range c {
		_ = v.WithTags()
		pc = append(pc, v.ToJSONAPI())
	}

	return pc, err
}
