package repositories

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	builder "github.com/doug-martin/goqu/v9"
	"github.com/zikwall/blogchain/src/app/exceptions"
	"github.com/zikwall/blogchain/src/app/forms"
	"github.com/zikwall/blogchain/src/platform/database"
	"strings"
	"time"
)

type ContentRepository struct {
	Repository
}

func UseContentRepository(ctx context.Context, conn *database.Connection) ContentRepository {
	return ContentRepository{
		Repository{connection: conn, context: ctx},
	}
}

func (cr ContentRepository) find() *builder.SelectDataset {
	query := cr.Connection().Builder()
	return query.Select("content.*").From("content")
}

func (cr ContentRepository) withFullUserProfile() *builder.SelectDataset {
	query := cr.find()
	query = joinWith(query, withUserQuery, withProfileQuery)
	return query
}

func (cr ContentRepository) UserContent(contentID, id int64) (Content, error) {
	var content Content

	query := cr.find()
	query = withUserQuery(query).Where(
		builder.And(
			builder.I("content.id").Eq(contentID),
			builder.I("user.id").Eq(id),
		),
	)

	found, err := query.ScanStructContext(cr.Context(), &content)
	if err != nil {
		return Content{}, exceptions.ThrowPrivateError(err)
	} else if !found {
		return Content{}, exceptions.ThrowPublicError(errors.New("user content was not found"))
	}

	tags, err := fetchContentTags(cr.Context(), cr.Connection(), content.ID)
	if err != nil {
		return Content{}, err
	}

	content.withTags(tags)
	return content, nil
}

func (cr ContentRepository) CreateContent(form *forms.ContentForm) (Content, error) {
	content := Content{}
	content.Content = form.Content
	content.Title = form.Title
	content.UserID = form.UserID
	content.Annotation = form.Annotation
	content.UUID = form.UUID

	record := builder.Record{
		"uuid":       content.UUID,
		"user_id":    content.UserID,
		"title":      content.Title,
		"content":    content.Content,
		"annotation": form.Annotation,
		"image":      content.Image.String,
		"created_at": time.Now().Unix(),
	}

	status, err := cr.Connection().
		Builder().
		Insert("content").
		Rows(record).
		Executor().
		ExecContext(cr.Context())

	if err != nil {
		return Content{}, err
	}

	if content.ID, err = status.LastInsertId(); err != nil {
		return Content{}, err
	}

	if err := cr.upsertTags(&content, form, false); err != nil {
		return Content{}, err
	}

	return content, err
}

func (cr ContentRepository) UpdateContent(content *Content, form *forms.ContentForm) error {
	record := builder.Record{
		"title":      form.Title,
		"content":    form.Content,
		"annotation": form.Annotation,
		"image":      content.Image.String,
		"updated_at": time.Now().Unix(),
	}

	_, err := cr.Connection().
		Builder().
		Update("content").
		Set(record).
		Where(
			builder.C("id").Eq(content.ID),
		).
		Executor().
		ExecContext(cr.Context())

	if err != nil {
		return exceptions.ThrowPrivateError(err)
	}

	if err := cr.upsertTags(content, form, true); err != nil {
		return err
	}

	return nil
}

func (cr ContentRepository) upsertTags(content *Content, form *forms.ContentForm, update bool) error {
	if form.Tags != "" {
		tags := []int{}

		if err := json.Unmarshal([]byte(form.Tags), &tags); err != nil {
			return err
		}

		if update {
			_, err := cr.Connection().
				Builder().
				Delete("content_tag").
				Where(
					builder.C("content_id").Eq(content.ID),
				).
				Executor().
				ExecContext(cr.Context())

			if err != nil {
				return exceptions.ThrowPrivateError(err)
			}
		}

		records := make([]builder.Record, 0, len(tags))
		for _, v := range tags {
			records = append(records, builder.Record{
				"content_id": content.ID,
				"tag_id":     v,
			})
		}

		_, err := cr.Connection().
			Builder().
			Insert("content_tag").
			Rows(records).
			Executor().
			ExecContext(cr.Context())

		if err != nil {
			return exceptions.ThrowPrivateError(err)
		}
	}

	return nil
}

func (cr ContentRepository) FindAllByUser(userid, page int64) ([]PublicContent, float64, error) {
	query := cr.withFullUserProfile().Where(
		builder.I("user.id").Eq(userid),
	)
	query, count := withPagination(cr.Context(), query, uint(page), 4)

	var content []Content
	if err := query.ScanStructsContext(cr.Context(), &content); err != nil {
		return nil, 0, exceptions.ThrowPrivateError(err)
	}

	response, err := withMutableResponse(cr.Context(), cr.Connection(), content)

	if err != nil {
		return nil, 0, err
	}

	return response, count, nil
}

func (cr ContentRepository) FindContentByIDAndUser(id, userid int64) (Content, error) {
	query := cr.withFullUserProfile().Where(
		builder.And(
			builder.I("content.id").Eq(id),
			builder.I("user.id").Eq(userid),
		),
	)

	content := Content{}
	if ok, err := query.ScanStructContext(cr.Context(), &content); err != nil {
		return content, exceptions.ThrowPrivateError(err)
	} else if !ok {
		return content, exceptions.ThrowPublicError(errors.New("content with the required ID was not found"))
	}

	if tags, err := fetchContentTags(cr.Context(), cr.Connection(), id); err == nil {
		content.withTags(tags)
	}

	return content, nil
}

func (cr ContentRepository) FindContentByID(id int64) (Content, error) {
	query := cr.withFullUserProfile().Where(
		builder.I("content.id").Eq(id),
	)

	content := Content{}
	if ok, err := query.ScanStructContext(cr.Context(), &content); err != nil {
		return content, exceptions.ThrowPrivateError(err)
	} else if !ok {
		return content, exceptions.ThrowPublicError(errors.New("content with the required ID was not found"))
	}

	if tags, err := fetchContentTags(cr.Context(), cr.Connection(), id); err == nil {
		content.withTags(tags)
	}

	return content, nil
}

func (cr ContentRepository) FindAllContent(label string, page int64) ([]PublicContent, float64, error) {
	query := cr.withFullUserProfile()

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

	query, count := withPagination(cr.Context(), query, uint(page), 4)

	var content []Content
	if err := query.ScanStructsContext(cr.Context(), &content); err != nil {
		return nil, 0, exceptions.ThrowPrivateError(err)
	}

	response, err := withMutableResponse(cr.Context(), cr.Connection(), content)

	if err != nil {
		return nil, 0, err
	}

	return response, count, nil
}

// избавиться
func (content *Content) withTags(tags []Tag) {
	if tags != nil {
		content.Tags = tags
	} else {
		// ToDo: Пофиксить эту шляпу на сторону клиента
		content.Tags = []Tag{}
	}
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

func withMutableResponse(ctx context.Context, conn *database.Connection, contents []Content) ([]PublicContent, error) {
	idx := make([]interface{}, 0, len(contents))
	contentMap := make(map[int64]*Content, len(contents))

	for i := range contents {
		contentMap[contents[i].ID] = &contents[i]
		idx = append(idx, fmt.Sprintf("%v", contents[i].ID))
	}

	tags, err := UseTagRepository(ctx, conn).ContentGroupedTags(idx...)

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
