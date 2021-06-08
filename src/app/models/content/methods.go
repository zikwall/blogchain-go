package content

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	builder "github.com/doug-martin/goqu/v9"
	"github.com/zikwall/blogchain/src/app/exceptions"
	"github.com/zikwall/blogchain/src/app/models"
	"github.com/zikwall/blogchain/src/app/models/content/forms"
	"github.com/zikwall/blogchain/src/app/models/tag"
	"github.com/zikwall/blogchain/src/platform/database"
	"time"
)

func (content *Content) WithTags(tags []tag.Tag) {
	if tags != nil {
		content.Tags = tags
	} else {
		// ToDo: Пофиксить эту шляпу на сторону клиента
		content.Tags = []tag.Tag{}
	}
}

func (content Content) GetTags(context context.Context, conn *database.Instance) ([]tag.Tag, error) {
	return tag.ContextConnection(context, conn).ContentTags(content.Id)
}

func (self Model) Find() *builder.SelectDataset {
	return self.
		Connection().
		Builder().
		Select("content.*").
		From("content")
}

func (c Model) WithFullUser() *builder.SelectDataset {
	query := c.Find()
	query = c.WithUser(query)
	query = c.WithUserProfile(query)

	return query
}

func (self Model) WithUser(query *builder.SelectDataset) *builder.SelectDataset {
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

func (self Model) WithUserProfile(query *builder.SelectDataset) *builder.SelectDataset {
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

func (self Model) UserContent(contentId int64, id int64) (Content, error) {
	var content Content

	query := self.
		WithUser(self.Find()).
		Where(
			builder.And(
				builder.I("content.id").Eq(contentId),
				builder.I("user.id").Eq(id),
			),
		)

	found, err := query.ScanStructContext(self.context, &content)

	if err != nil {
		return Content{}, exceptions.NewErrDatabaseAccess(err)
	} else if !found {
		return Content{}, exceptions.NewErrApplicationLogic(errors.New("user content was not found"))
	}

	tags, err := content.GetTags(self.Context(), self.Connection())

	if err != nil {
		return Content{}, err
	}

	content.WithTags(tags)

	return content, nil
}

func (self Model) CreateContent(f *forms.ContentForm) (Content, error) {
	content := Content{}
	content.Content = f.Content
	content.Title = f.Title
	content.UserId = f.UserId
	content.Annotation = f.Annotation
	content.Uuid = f.UUID

	insert := self.
		Connection().
		Builder().
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
		).
		Executor()

	status, err := insert.ExecContext(self.Context())

	if err != nil {
		return Content{}, err
	}

	if content.Id, err = status.LastInsertId(); err != nil {
		return Content{}, err
	}

	if err = self.UpsertTags(content, f, false); err != nil {
		return Content{}, err
	}

	return content, err
}

func (self Model) UpdateContent(content Content, f *forms.ContentForm) error {
	update := self.
		Connection().
		Builder().
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

	if _, err := update.ExecContext(self.Context()); err != nil {
		return exceptions.NewErrDatabaseAccess(err)
	}

	if err := self.UpsertTags(content, f, true); err != nil {
		return err
	}

	return nil
}

func (self Model) UpsertTags(content Content, f *forms.ContentForm, update bool) error {
	if f.Tags != "" {
		tags := []int{}

		if err := json.Unmarshal([]byte(f.Tags), &tags); err != nil {
			return err
		}

		if update {
			executor := self.
				Connection().
				Builder().
				Delete("content_tag").Where(
				builder.C("content_id").Eq(content.Id),
			).
				Executor()

			if _, err := executor.ExecContext(self.Context()); err != nil {
				return exceptions.NewErrDatabaseAccess(err)
			}
		}

		records := make([]builder.Record, 0, len(tags))
		for _, v := range tags {
			records = append(records, builder.Record{
				"content_id": content.Id,
				"tag_id":     v,
			})
		}

		executor := self.
			Connection().
			Builder().
			Insert("content_tag").
			Rows(records).
			Executor()

		if _, err := executor.ExecContext(self.Context()); err != nil {
			return exceptions.NewErrDatabaseAccess(err)
		}
	}

	return nil
}

func (self Model) FindAllByUser(userid int64, page int64) ([]PublicContent, error, float64) {
	var content []Content

	query := self.
		WithFullUser().
		Where(
			builder.I("user.id").Eq(userid),
		)

	query, count := models.WithPagination(self.Context(), query, uint(page), 4)

	if err := query.ScanStructsContext(self.Context(), &content); err != nil {
		return nil, exceptions.NewErrDatabaseAccess(err), 0
	}

	response, err := self.withMutableResponse(content)

	if err != nil {
		return nil, err, 0
	}

	return response, nil, count
}

func (self Model) FindContentByIdAndUser(id int64, userid int64) (Content, error) {
	content := Content{}

	query := self.
		WithFullUser().
		Where(
			builder.And(
				builder.I("content.id").Eq(id),
				builder.I("user.id").Eq(userid),
			),
		)

	if ok, err := query.ScanStructContext(self.Context(), &content); err != nil {
		return Content{}, exceptions.NewErrDatabaseAccess(err)
	} else if !ok {
		return content, exceptions.NewErrApplicationLogic(errors.New("content with the required ID was not found"))
	}

	if tags, err := content.GetTags(self.Context(), self.Connection()); err == nil {
		content.WithTags(tags)
	}

	return content, nil
}

func (self Model) FindContentById(id int64) (Content, error) {
	content := Content{}
	query := self.
		WithFullUser().
		Where(
			builder.I("content.id").Eq(id),
		)

	if ok, err := query.ScanStructContext(self.Context(), &content); err != nil {
		return content, exceptions.NewErrDatabaseAccess(err)
	} else if !ok {
		return content, exceptions.NewErrApplicationLogic(errors.New("content with the required ID was not found"))
	}

	if tags, err := content.GetTags(self.Context(), self.Connection()); err == nil {
		content.WithTags(tags)
	}

	return content, nil
}

func (self Model) FindAllContent(label string, page int64) ([]PublicContent, error, float64) {
	var content []Content

	query := self.WithFullUser()

	if label != "" {
		query = query.
			LeftJoin(
				builder.T("content_tag"),
				builder.On(
					builder.I("content_tag.content_id").Eq(builder.I("content.id")),
				),
			).
			LeftJoin(
				builder.T("tags"),
				builder.On(
					builder.I("content_tag.tag_id").Eq(builder.I("tags.id")),
				),
			).
			Where(
				builder.I("tags.label").Eq(label),
			)
	}

	query, count := models.WithPagination(self.Context(), query, uint(page), 4)

	if err := query.ScanStructsContext(self.Context(), &content); err != nil {
		return nil, exceptions.NewErrDatabaseAccess(err), 0
	}

	response, err := self.withMutableResponse(content)

	if err != nil {
		return nil, err, 0
	}

	return response, nil, count
}

func (self Model) withMutableResponse(contents []Content) ([]PublicContent, error) {
	idx := make([]interface{}, 0, len(contents))
	contentMap := make(map[int64]*Content, len(contents))

	for _, content := range contents {
		c := content
		contentMap[content.Id] = &c
		idx = append(idx, fmt.Sprintf("%v", content.Id))
	}

	tags, err := tag.
		ContextConnection(self.Context(), self.Connection()).
		ContentGroupedTags(idx...)

	if err != nil {
		return nil, err
	}

	for id, value := range tags {
		if _, ok := contentMap[id]; ok {
			contentMap[id].WithTags(value)
		}
	}

	pc := make([]PublicContent, 0, len(contentMap))
	for _, content := range contentMap {
		pc = append(pc, content.Response())
	}

	return pc, nil
}
