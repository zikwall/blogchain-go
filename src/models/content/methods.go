package content

import (
	"encoding/json"
	"fmt"
	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
	"github.com/zikwall/blogchain/src/models"
	forms2 "github.com/zikwall/blogchain/src/models/content/forms"
	tag2 "github.com/zikwall/blogchain/src/models/tag"
	user2 "github.com/zikwall/blogchain/src/models/user"
	"time"
)

func createImagePath(uuidv4 string) string {
	return fmt.Sprintf("%s.png", uuidv4)
}

func (c ContentModel) CreateContent(f *forms2.ContentForm, ctx *fiber.Ctx) (*Content, error) {
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
		_ = SaveImage(content, f, ctx)
	}

	status, err := c.Query().Insert("content", dbx.Params{
		"uuid":       content.Uuid,
		"user_id":    content.UserId,
		"title":      content.Title,
		"content":    content.Content,
		"annotation": f.Annotation,
		"image":      content.Image.String,
		"created_at": time.Now().Unix(),
	}).Execute()

	content.Id, err = status.LastInsertId()

	if err == nil {
		err = c.UpsertTags(content, f, true)
	}

	return content, err
}

func (c ContentModel) UpdateContent(content *Content, f *forms2.ContentForm, ctx *fiber.Ctx) error {
	if f.GetImage().Err == nil {
		_ = SaveImage(content, f, ctx)
	}

	_, err := c.Query().Update("content", dbx.Params{
		"title":      f.Title,
		"content":    f.Content,
		"annotation": f.Annotation,
		"image":      content.Image.String,
		"updated_at": time.Now().Unix(),
	}, dbx.HashExp{"id": content.Id}).Execute()

	if err == nil {
		err = c.UpsertTags(content, f, true)
	}

	return err
}

func (c ContentModel) UpsertTags(content *Content, f *forms2.ContentForm, update bool) error {
	var err error

	if f.Tags != "" {
		tags := []int{}

		if err := json.Unmarshal([]byte(f.Tags), &tags); err == nil {

			// todo calculate diff from request and existing tags (?)
			if update {
				_, err = c.Query().
					Delete("content_tag", dbx.HashExp{"content_id": content.Id}).
					Execute()
			}

			// todo batch upsert & limited tags
			for _, v := range tags {
				_, err = c.Query().Upsert("content_tag", dbx.Params{
					"content_id": content.Id,
					"tag_id":     v,
				}, "content_id=content_id", "tag_id=tag_id").Execute()
			}
		}
	}

	return err
}

func SaveImage(content *Content, f *forms2.ContentForm, c *fiber.Ctx) error {
	content.Image.String = createImagePath(content.Uuid)
	err := c.SaveFile(f.GetImage().File, fmt.Sprintf("./public/uploads/%s", content.Image.String))

	return err
}

func (c ContentModel) Find() *dbx.SelectQuery {
	query := c.Query().
		Select(
			"content.*",
			"u.id as user.id",
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

func (c ContentModel) FindAllByUser(userid int64, page int64) ([]PublicContent, error, float64) {
	var content []Content

	query := c.Find().
		Where(dbx.HashExp{"u.id": userid})

	var pageSize int64
	pageSize = 4

	countPages, _ := models.QueryCount(query, pageSize)
	query.Offset(page * pageSize).Limit(pageSize)

	err := query.All(&content)

	pc := []PublicContent{}
	for _, v := range content {
		_ = v.WithTags()
		pc = append(pc, v.ToJSONAPI())
	}

	return pc, err, countPages
}

func (c ContentModel) FindContentByIdAndUser(id int64, userid int64) (*Content, error) {
	content := &Content{
		Id:      0,
		UserId:  0,
		Title:   "",
		Content: "",
		User: user2.User{
			Id:       0,
			Username: "",
			Email:    "",
			Profile:  user2.Profile{},
		},
	}

	err := c.Find().
		Where(dbx.HashExp{"content.id": id}).
		AndWhere(dbx.HashExp{"u.id": userid}).
		One(&c)

	if err == nil {
		err = content.WithTags()
	}

	return content, err
}

func (c ContentModel) FindContentById(id int64) (Content, error) {
	content := Content{
		Id:      0,
		UserId:  0,
		Title:   "",
		Content: "",
		User: user2.User{
			Id:       0,
			Username: "",
			Email:    "",
			Profile:  user2.Profile{},
		},
	}

	err := c.Find().
		Where(dbx.HashExp{"content.id": id}).
		One(&content)

	if err == nil {
		err = content.WithTags()
	}

	return content, err
}

func (c ContentModel) FindAllContent(label string, page int64) ([]PublicContent, error, float64) {
	var content []Content

	query := c.Find()

	if label != "" {
		tag2.AttachTagQuery(query, label)
	}

	var pageSize int64
	pageSize = 4

	countPages, _ := models.QueryCount(query, pageSize)
	query.Offset(page * pageSize).Limit(pageSize)

	err := query.All(&content)

	if err != nil {
		return nil, err, 0
	}

	pc := []PublicContent{}
	for _, v := range content {
		_ = v.WithTags()
		pc = append(pc, v.ToJSONAPI())
	}

	return pc, err, countPages
}
