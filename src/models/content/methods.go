package content

import (
	"encoding/json"
	"fmt"
	builder "github.com/doug-martin/goqu/v9"
	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
	"github.com/zikwall/blogchain/src/models"
	"github.com/zikwall/blogchain/src/models/content/forms"
	"github.com/zikwall/blogchain/src/models/tag"
	"github.com/zikwall/blogchain/src/models/user"
	"github.com/zikwall/blogchain/src/utils"
	"time"
)

type Contents struct {
	Content      `json:"content"`
	user.User    `json:"user"`
	user.Profile `json:"profile"`
}

func (contents Contents) Response() PublicContent {
	return PublicContent{
		Id:         contents.Content.Id,
		Uuid:       contents.Uuid,
		Title:      contents.Title,
		Annotation: contents.Annotation,
		Content:    contents.Content.Content,
		CreatedAt:  contents.Content.CreatedAt.Int64,
		UpdatedAt:  contents.Content.UpdatedAt.Int64,
		Image:      contents.Image.String,
		Related: Related{
			Publisher: user.PublicUser{
				Id:       contents.UserId,
				Username: contents.User.Username,
				Profile: user.PublicProfile{
					Name:        contents.Profile.Name,
					Email:       contents.Profile.PublicEmail,
					Avatar:      contents.Profile.Avatar.String,
					Location:    contents.Profile.Location.String,
					Status:      contents.Profile.Status.String,
					Description: contents.Profile.Description.String,
				},
			},
			Tags: contents.Tags,
		},
	}
}

func (contents *Contents) WithTags(tags []tag.Tag) {
	contents.Tags = tags
}

func (content Content) GetTags() ([]tag.Tag, error) {
	u := tag.NewTagModel()
	tags, err := u.ContentTags(content.Id)

	return tags, err
}

func (self ContentModel) Find() *builder.SelectDataset {
	return models.QueryBuilder().Select("content.*").From("content")
}

func (c ContentModel) FindWith() *builder.SelectDataset {
	query := c.Find()
	query = c.WithUser(query)
	query = c.WithUserProfile(query)

	return query
}

func (self ContentModel) WithUser(query *builder.SelectDataset) *builder.SelectDataset {
	return query.
		SelectAppend(
			"user.username",
		).
		LeftJoin(
			builder.T("user"),
			builder.On(
				builder.I("user.id").Eq(builder.I("content.user_id")),
			),
		)
}

func (self ContentModel) WithUserProfile(query *builder.SelectDataset) *builder.SelectDataset {
	return query.
		SelectAppend(
			"profile.user_id",
			"profile.name",
			"profile.public_email",
			"profile.avatar",
		).
		LeftJoin(
			builder.T("profile"),
			builder.On(
				builder.I("profile.user_id").Eq(builder.I("user.id")),
			),
		)
}

func (self ContentModel) UserContent(contentId int64, id int64) (Contents, error) {
	var content Contents

	query := self.Find()
	query = self.WithUser(query)

	query = query.Where(
		builder.And(
			builder.I("content.id").Eq(contentId),
			builder.I("user.id").Eq(id),
		),
	)

	_, err := query.ScanStruct(&content)

	if err != nil {
		return content, err
	}

	return content, nil
}

func (c ContentModel) CreateContent(f *forms.ContentForm, ctx *fiber.Ctx) (Content, error) {
	content := Content{}
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

	insert := models.QueryBuilder().
		Insert("content").
		Rows(
			builder.Record{
				"uuid":       content.Uuid,
				"user_id":    content.UserId,
				"title":      content.Title,
				"content":    content.Content,
				"annotation": f.Annotation,
				"image":      content.Image.String,
				"created_at": time.Now().Unix(),
			},
		).Executor()

	status, err := insert.Exec()
	content.Id, err = status.LastInsertId()

	if err == nil {
		err = c.UpsertTags(content, f, true)
	}

	return content, err
}

func (c ContentModel) UpdateContent(content Content, f *forms.ContentForm, ctx *fiber.Ctx) error {
	if f.GetImage().Err == nil {
		_ = SaveImage(content, f, ctx)
	}

	update := models.QueryBuilder().
		Update("content").
		Set(
			builder.Record{
				"title":      content.Title,
				"content":    content.Content,
				"annotation": f.Annotation,
				"image":      content.Image.String,
				"updated_at": time.Now().Unix(),
			},
		).
		Where(
			builder.C("id").Eq(content.Id),
		).
		Executor()

	_, err := update.Exec()

	if err == nil {
		err = c.UpsertTags(content, f, true)
	}

	return err
}

func (c ContentModel) UpsertTags(content Content, f *forms.ContentForm, update bool) error {
	var err error

	if f.Tags != "" {
		tags := []int{}

		if err := json.Unmarshal([]byte(f.Tags), &tags); err == nil {

			// todo calculate diff from request and existing tags (?)
			if update {
				executor := models.QueryBuilder().Delete("content_tag").Where(
					builder.C("content_id").Eq(content.Id),
				).Executor()

				_, err = executor.Exec()
			}

			// todo batch upsert & limited tags
			executor := models.QueryBuilder()
			for _, v := range tags {
				insert := executor.Insert("content_tag").Rows(
					builder.Record{
						"content_id": content.Id,
						"tag_id":     v,
					},
				).Executor()

				_, err = insert.Exec()
			}
		}
	}

	return err
}

func SaveImage(content Content, f *forms.ContentForm, c *fiber.Ctx) error {
	content.Image.String = utils.CreateImagePath(content.Uuid)
	path := fmt.Sprintf("./public/uploads/%s", content.Image.String)

	return utils.SaveFile(c, f.GetImage().File, path)
}

func (c ContentModel) FindAllByUser(userid int64, page int64) ([]PublicContent, error, float64) {
	var content []Contents

	query := c.FindWith().Where(builder.I("user.id").Eq(userid))

	raw, _, _ := query.ToSQL()

	fmt.Println(raw)

	var pageSize uint
	pageSize = 4

	countPages, _ := models.QueryCount(query, pageSize)
	query.Offset(uint(page) * pageSize).Limit(pageSize)

	err := query.ScanStructs(&content)

	pc := []PublicContent{}
	for _, v := range content {
		if tags, err := v.GetTags(); err == nil {
			v.WithTags(tags)
		}

		pc = append(pc, v.Response())
	}

	return pc, err, countPages
}

func (c ContentModel) FindContentByIdAndUser(id int64, userid int64) (*Contents, error) {
	content := &Contents{}

	_, err := c.FindWith().
		Where(
			builder.And(
				builder.I("content.id").Eq(id),
				builder.I("user.id").Eq(userid),
			),
		).
		ScanStruct(&content)

	if err == nil {
		if tags, err := content.GetTags(); err == nil {
			content.WithTags(tags)
		}
	}

	return content, err
}

func (c ContentModel) FindContentById(id int64) (Contents, error) {
	content := Contents{}
	query := c.FindWith().Where(builder.I("content.id").Eq(id))

	_, err := query.ScanStruct(&content)

	if err == nil {
		if tags, err := content.GetTags(); err == nil {
			content.WithTags(tags)
		}
	}

	return content, err
}

func (c ContentModel) FindAllContent(label string, page int64) ([]PublicContent, error, float64) {
	var content []Contents

	query := c.FindWith()

	if label != "" {
		onTagCondition := func(query *builder.SelectDataset, tag string) *builder.SelectDataset {
			withContentTag := func(query *builder.SelectDataset) *builder.SelectDataset {
				return query.LeftJoin(
					builder.T("content_tag"),
					builder.On(
						builder.I("content_tag.content_id").Eq(builder.I("content.id")),
					),
				)
			}

			withTags := func(query *builder.SelectDataset) *builder.SelectDataset {
				return query.LeftJoin(
					builder.T("tags"),
					builder.On(
						builder.I("content_tag.tag_id").Eq(builder.I("tags.id")),
					),
				)
			}

			query = withContentTag(query)
			query = withTags(query)

			return query.Where(builder.I("tags.label").Eq(tag))
		}

		query = onTagCondition(query, label)
	}

	var pageSize uint
	pageSize = 4

	countPages, _ := models.QueryCount(query, pageSize)
	query = query.Offset(uint(page) * pageSize).Limit(pageSize)

	err := query.ScanStructs(&content)

	if err != nil {
		return nil, err, 0
	}

	pc := []PublicContent{}

	for _, v := range content {
		if tags, err := v.GetTags(); err == nil {
			v.WithTags(tags)
		}

		pc = append(pc, v.Response())
	}

	return pc, err, countPages
}
