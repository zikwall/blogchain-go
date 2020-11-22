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
	"github.com/zikwall/blogchain/src/utils"
	"time"
)

func (content *Content) WithTags(tags []tag.Tag) {
	if tags != nil && len(tags) > 0 {
		content.Tags = tags
	} else {
		// ToDo: Пофиксить эту шляпу на сторону клиента
		content.Tags = []tag.Tag{}
	}
}

func (content Content) GetTags() ([]tag.Tag, error) {
	u := tag.NewTagModel()
	tags, err := u.ContentTags(content.Id)

	return tags, err
}

func (self ContentModel) Find() *builder.SelectDataset {
	return self.QueryBuilder().Select("content.*").From("content")
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
			builder.I("user.id").As(builder.C("user.id")),
			builder.I("user.username").As(builder.C("user.username")),
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
			builder.I("profile.name").As(builder.C("user.profile.name")),
			builder.I("profile.public_email").As(builder.C("user.profile.public_email")),
			builder.I("profile.avatar").As(builder.C("user.profile.avatar")),
		).
		LeftJoin(
			builder.T("profile"),
			builder.On(
				builder.I("profile.user_id").Eq(builder.I("user.id")),
			),
		)
}

func (self ContentModel) WithPagination(query *builder.SelectDataset, page, size uint) (*builder.SelectDataset, float64) {
	countPages, _ := models.QueryCount(query, size)
	query = query.Offset(page * size).Limit(size)

	return query, countPages
}

func (self ContentModel) WithMutableResponse(contents []Content) ([]PublicContent, error) {
	idx := make([]interface{}, 0, len(contents))
	contentMap := make(map[int64]*Content, len(contents))

	for _, content := range contents {
		c := content
		contentMap[content.Id] = &c
		idx = append(idx, fmt.Sprintf("%v", content.Id))
	}

	t := tag.NewTagModel()

	tags, err := t.ContentGroupedTags(idx...)

	if err != nil {
		return nil, err
	}

	for id, value := range tags {
		if _, ok := contentMap[id]; ok {

			contentMap[id].WithTags(value)
		}
	}

	var pc []PublicContent
	for _, content := range contentMap {
		pc = append(pc, content.Response())
	}

	return pc, nil
}

func (self ContentModel) UserContent(contentId int64, id int64) (Content, error) {
	var content Content

	query := self.Find()
	query = self.WithUser(query)

	query = query.Where(
		builder.And(
			builder.I("content.id").Eq(contentId),
			builder.I("user.id").Eq(id),
		),
	)

	if _, err := query.ScanStruct(&content); err != nil {
		return Content{}, err
	}

	tags, err := content.GetTags()

	if err != nil {
		return Content{}, err
	}

	content.WithTags(tags)

	return content, nil
}

func (self ContentModel) CreateContent(f *forms.ContentForm, ctx *fiber.Ctx) (Content, error) {
	content := Content{}
	content.Content = f.Content
	content.Title = f.Title
	content.UserId = f.UserId
	content.Annotation = f.Annotation

	var err error
	uv4 := uuid.Must(uuid.NewV4(), err)
	content.Uuid = uv4.String()

	// ToDo: проверка ошибок
	if f.GetImage().Err == nil {
		if err = SaveImage(&content, f, ctx); err != nil {
			return Content{}, err
		}
	}

	insert := self.QueryBuilder().
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

	if err != nil {
		return Content{}, err
	}

	if content.Id, err = status.LastInsertId(); err != nil {
		return Content{}, err
	}

	if err = self.UpsertTags(content, f, true); err != nil {
		return Content{}, err
	}

	return content, err
}

func (self ContentModel) UpdateContent(content Content, f *forms.ContentForm, ctx *fiber.Ctx) error {
	// ToDo: проверка ошибок
	if f.GetImage().Err == nil {
		if err := SaveImage(&content, f, ctx); err != nil {
			return err
		}
	}

	update := self.QueryBuilder().
		Update("content").
		Set(
			builder.Record{
				"title":      f.Title,
				"content":    f.Content,
				"annotation": f.Annotation,
				"image":      content.Image.String,
				"updated_at": time.Now().Unix(),
			},
		).
		Where(
			builder.C("id").Eq(content.Id),
		).
		Executor()

	if _, err := update.Exec(); err != nil {
		return err
	}

	if err := self.UpsertTags(content, f, true); err != nil {
		return err
	}

	return nil
}

func (self ContentModel) UpsertTags(content Content, f *forms.ContentForm, update bool) error {
	if f.Tags != "" {
		tags := []int{}

		if err := json.Unmarshal([]byte(f.Tags), &tags); err != nil {
			return err
		}

		if update {
			executor := self.QueryBuilder().Delete("content_tag").Where(
				builder.C("content_id").Eq(content.Id),
			).Executor()

			if _, err := executor.Exec(); err != nil {
				return err
			}
		}

		// ToDo: Batch Insert
		executor := self.QueryBuilder()
		for _, v := range tags {
			insert := executor.Insert("content_tag").Rows(
				builder.Record{
					"content_id": content.Id,
					"tag_id":     v,
				},
			).Executor()

			if _, err := insert.Exec(); err != nil {
				return err
			}
		}
	}

	return nil
}

// ToDo: нажо что-то сделать с этим методом, он ни туда ни сюда...
func SaveImage(content *Content, f *forms.ContentForm, c *fiber.Ctx) error {
	content.Image.String = utils.CreateImagePath(content.Uuid)
	path := fmt.Sprintf("./src/public/uploads/%s", content.Image.String)

	return utils.SaveFile(c, f.GetImage().File, path)
}

func (self ContentModel) FindAllByUser(userid int64, page int64) ([]PublicContent, error, float64) {
	var content []Content

	query := self.FindWith().Where(builder.I("user.id").Eq(userid))
	query, count := self.WithPagination(query, uint(page), 4)

	if err := query.ScanStructs(&content); err != nil {
		return nil, err, 0
	}

	response, err := self.WithMutableResponse(content)

	if err != nil {
		return nil, err, 0
	}

	return response, nil, count
}

func (self ContentModel) FindContentByIdAndUser(id int64, userid int64) (Content, error) {
	content := Content{}

	query := self.FindWith().
		Where(
			builder.And(
				builder.I("content.id").Eq(id),
				builder.I("user.id").Eq(userid),
			),
		)

	if _, err := query.ScanStruct(&content); err != nil {
		return Content{}, err
	}

	if tags, err := content.GetTags(); err == nil {
		content.WithTags(tags)
	}

	return content, nil
}

func (self ContentModel) FindContentById(id int64) (Content, error) {
	content := Content{}
	query := self.FindWith().Where(builder.I("content.id").Eq(id))

	if _, err := query.ScanStruct(&content); err != nil {
		return Content{}, err
	}

	if tags, err := content.GetTags(); err == nil {
		content.WithTags(tags)
	}

	return content, nil
}

func (self ContentModel) FindAllContent(label string, page int64) ([]PublicContent, error, float64) {
	var content []Content

	query := self.FindWith()

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

	query, count := self.WithPagination(query, uint(page), 4)

	if err := query.ScanStructs(&content); err != nil {
		return nil, err, 0
	}

	response, err := self.WithMutableResponse(content)

	if err != nil {
		return nil, err, 0
	}

	return response, nil, count
}
