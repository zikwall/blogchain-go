package content

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	builder "github.com/doug-martin/goqu/v9"
	"github.com/zikwall/blogchain/src/app/exceptions"
	"github.com/zikwall/blogchain/src/app/models"
	"github.com/zikwall/blogchain/src/app/models/common"
	"github.com/zikwall/blogchain/src/app/models/content/forms"
	"github.com/zikwall/blogchain/src/app/models/tag"
	"github.com/zikwall/blogchain/src/platform/database"
	"strings"
	"time"
)

func (m Model) find() *builder.SelectDataset {
	query := m.Connection().Builder()
	return query.Select("content.*").From("content")
}

func (m Model) withFullUserProfile() *builder.SelectDataset {
	query := m.find()
	query = common.JoinWith(query, withUserQuery, withProfileQuery)
	return query
}

func withUserQuery(query *builder.SelectDataset) *builder.SelectDataset {
	query = query.SelectAppend(
		builder.I("user.id").As(builder.C("user.id")),
		builder.I("user.username").As(builder.C("user.username")),
	)
	query = query.LeftJoin(
		builder.T("user"),
		builder.On(
			builder.I("user.id").Eq(builder.I("content.user_id")),
		),
	)
	return query
}

func withProfileQuery(query *builder.SelectDataset) *builder.SelectDataset {
	query = query.SelectAppend(
		builder.I("profile.name").As(builder.C("user.profile.name")),
		builder.I("profile.public_email").As(builder.C("user.profile.public_email")),
		builder.I("profile.avatar").As(builder.C("user.profile.avatar")),
	)
	query = query.LeftJoin(
		builder.T("profile"),
		builder.On(
			builder.I("profile.user_id").Eq(builder.I("user.id")),
		),
	)
	return query
}

func (m Model) UserContent(contentId int64, id int64) (Content, error) {
	var content Content

	query := m.find()
	query = withUserQuery(query).Where(
		builder.And(
			builder.I("content.id").Eq(contentId),
			builder.I("user.id").Eq(id),
		),
	)

	found, err := query.ScanStructContext(m.Context(), &content)
	if err != nil {
		return Content{}, exceptions.NewErrDatabaseAccess(err)
	} else if !found {
		return Content{}, exceptions.NewErrApplicationLogic(errors.New("user content was not found"))
	}

	tags, err := content.getTags(m.Context(), m.Connection())
	if err != nil {
		return Content{}, err
	}

	content.withTags(tags)
	return content, nil
}

func (m Model) CreateContent(form *forms.ContentForm) (Content, error) {
	content := Content{}
	content.Content = form.Content
	content.Title = form.Title
	content.UserId = form.UserId
	content.Annotation = form.Annotation
	content.Uuid = form.UUID

	record := builder.Record{
		"uuid":       content.Uuid,
		"user_id":    content.UserId,
		"title":      content.Title,
		"content":    content.Content,
		"annotation": form.Annotation,
		"image":      content.Image.String,
		"created_at": time.Now().Unix(),
	}

	status, err := m.Connection().
		Builder().
		Insert("content").
		Rows(record).
		Executor().
		ExecContext(m.Context())

	if err != nil {
		return Content{}, err
	}

	if content.Id, err = status.LastInsertId(); err != nil {
		return Content{}, err
	}

	if err = m.upsertTags(content, form, false); err != nil {
		return Content{}, err
	}

	return content, err
}

func (m Model) UpdateContent(content Content, form *forms.ContentForm) error {
	record := builder.Record{
		"title":      form.Title,
		"content":    form.Content,
		"annotation": form.Annotation,
		"image":      content.Image.String,
		"updated_at": time.Now().Unix(),
	}

	_, err := m.Connection().
		Builder().
		Update("content").
		Set(record).
		Where(
			builder.C("id").Eq(content.Id),
		).
		Executor().
		ExecContext(m.Context())

	if err != nil {
		return exceptions.NewErrDatabaseAccess(err)
	}

	if err := m.upsertTags(content, form, true); err != nil {
		return err
	}

	return nil
}

func (m Model) upsertTags(content Content, form *forms.ContentForm, update bool) error {
	if form.Tags != "" {
		tags := []int{}

		if err := json.Unmarshal([]byte(form.Tags), &tags); err != nil {
			return err
		}

		if update {
			_, err := m.Connection().
				Builder().
				Delete("content_tag").
				Where(
					builder.C("content_id").Eq(content.Id),
				).
				Executor().
				ExecContext(m.Context())

			if err != nil {
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

		_, err := m.Connection().
			Builder().
			Insert("content_tag").
			Rows(records).
			Executor().
			ExecContext(m.Context())

		if err != nil {
			return exceptions.NewErrDatabaseAccess(err)
		}
	}

	return nil
}

func (m Model) FindAllByUser(userid int64, page int64) ([]PublicContent, error, float64) {
	query := m.withFullUserProfile().Where(
		builder.I("user.id").Eq(userid),
	)
	query, count := models.WithPagination(m.Context(), query, uint(page), 4)

	var content []Content
	if err := query.ScanStructsContext(m.Context(), &content); err != nil {
		return nil, exceptions.NewErrDatabaseAccess(err), 0
	}

	response, err := m.withMutableResponse(content)

	if err != nil {
		return nil, err, 0
	}

	return response, nil, count
}

func (m Model) FindContentByIdAndUser(id int64, userid int64) (Content, error) {
	query := m.withFullUserProfile().Where(
		builder.And(
			builder.I("content.id").Eq(id),
			builder.I("user.id").Eq(userid),
		),
	)

	content := Content{}
	if ok, err := query.ScanStructContext(m.Context(), &content); err != nil {
		return Content{}, exceptions.NewErrDatabaseAccess(err)
	} else if !ok {
		return content, exceptions.NewErrApplicationLogic(errors.New("content with the required ID was not found"))
	}

	if tags, err := content.getTags(m.Context(), m.Connection()); err == nil {
		content.withTags(tags)
	}

	return content, nil
}

func (m Model) FindContentById(id int64) (Content, error) {
	query := m.withFullUserProfile().Where(
		builder.I("content.id").Eq(id),
	)

	content := Content{}
	if ok, err := query.ScanStructContext(m.Context(), &content); err != nil {
		return content, exceptions.NewErrDatabaseAccess(err)
	} else if !ok {
		return content, exceptions.NewErrApplicationLogic(errors.New("content with the required ID was not found"))
	}

	if tags, err := content.getTags(m.Context(), m.Connection()); err == nil {
		content.withTags(tags)
	}

	return content, nil
}

func (m Model) FindAllContent(label string, page int64) ([]PublicContent, error, float64) {
	query := m.withFullUserProfile()

	if !strings.EqualFold(label, "") {
		query = query.LeftJoin(
			builder.T("content_tag"),
			builder.On(
				builder.I("content_tag.content_id").Eq(builder.I("content.id")),
			),
		)
		query = query.LeftJoin(
			builder.T("tags"),
			builder.On(
				builder.I("content_tag.tag_id").Eq(builder.I("tags.id")),
			),
		)
		query = query.Where(
			builder.I("tags.label").Eq(label),
		)
	}

	query, count := models.WithPagination(m.Context(), query, uint(page), 4)

	var content []Content
	if err := query.ScanStructsContext(m.Context(), &content); err != nil {
		return nil, exceptions.NewErrDatabaseAccess(err), 0
	}

	response, err := m.withMutableResponse(content)

	if err != nil {
		return nil, err, 0
	}

	return response, nil, count
}

func (m Model) withMutableResponse(contents []Content) ([]PublicContent, error) {
	idx := make([]interface{}, 0, len(contents))
	contentMap := make(map[int64]*Content, len(contents))

	for _, content := range contents {
		c := content
		contentMap[content.Id] = &c
		idx = append(idx, fmt.Sprintf("%v", content.Id))
	}

	tags, err := tag.ContextConnection(m.Context(), m.Connection()).ContentGroupedTags(idx...)

	if err != nil {
		return nil, err
	}

	for id, value := range tags {
		if _, ok := contentMap[id]; ok {
			contentMap[id].withTags(value)
		}
	}

	pc := make([]PublicContent, 0, len(contentMap))
	for _, content := range contentMap {
		pc = append(pc, content.Response())
	}

	return pc, nil
}

func (content *Content) withTags(tags []tag.Tag) {
	if tags != nil {
		content.Tags = tags
	} else {
		// ToDo: Пофиксить эту шляпу на сторону клиента
		content.Tags = []tag.Tag{}
	}
}

func (content Content) getTags(context context.Context, conn *database.Instance) ([]tag.Tag, error) {
	tags, err := tag.ContextConnection(context, conn).ContentTags(content.Id)
	return tags, err
}
