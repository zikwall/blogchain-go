package repositories

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/zikwall/blogchain/src/pkg/database"
	"github.com/zikwall/blogchain/src/pkg/exceptions"
	"github.com/zikwall/blogchain/src/services/api/forms"

	builder "github.com/doug-martin/goqu/v9"
)

type ContentRepository struct {
	Repository
}

func UseContentRepository(ctx context.Context, conn *database.Connection) *ContentRepository {
	return &ContentRepository{
		Repository{connection: conn, context: ctx},
	}
}

func (r *ContentRepository) find() *builder.SelectDataset {
	query := r.Connection().Builder()
	return query.Select("content.*").From("content")
}

func (r *ContentRepository) withFullUserProfile() *builder.SelectDataset {
	query := r.find()
	query = database.JoinWith(query, withUserQuery, withProfileQuery)
	return query
}

func (r *ContentRepository) UserContent(contentID, id int64) (Content, error) {
	var content Content

	query := r.find()
	query = withUserQuery(query).Where(
		builder.And(
			builder.I("content.id").Eq(contentID),
			builder.I("user.id").Eq(id),
		),
	)

	found, err := query.ScanStructContext(r.Context(), &content)
	if err != nil {
		return Content{}, exceptions.ThrowPrivateError(err)
	} else if !found {
		return Content{}, exceptions.ThrowPublicError(errors.New("user content was not found"))
	}
	tags, err := fetchContentTags(r.Context(), r.Connection(), content.ID)
	if err != nil {
		return Content{}, err
	}
	content.withTags(tags)
	return content, nil
}

func (r *ContentRepository) CreateContent(form *forms.ContentForm) (Content, error) {
	content := Content{
		Content:    form.Content,
		Title:      form.Title,
		UserID:     form.UserID,
		Annotation: form.Annotation,
		UUID:       form.UUID,
	}

	record := builder.Record{
		"uuid":       content.UUID,
		"user_id":    content.UserID,
		"title":      content.Title,
		"content":    content.Content,
		"annotation": form.Annotation,
		"image":      content.Image.String,
		"created_at": time.Now().Unix(),
	}

	status, err := r.Connection().
		Builder().
		Insert("content").
		Rows(record).
		Executor().
		ExecContext(r.Context())
	if err != nil {
		return Content{}, err
	}
	if content.ID, err = status.LastInsertId(); err != nil {
		return Content{}, err
	}
	if err := r.upsertTags(&content, form, false); err != nil {
		return Content{}, err
	}
	return content, err
}

func (r *ContentRepository) UpdateContent(content *Content, form *forms.ContentForm) error {
	record := builder.Record{
		"title":      form.Title,
		"content":    form.Content,
		"annotation": form.Annotation,
		"image":      content.Image.String,
		"updated_at": time.Now().Unix(),
	}

	_, err := r.Connection().
		Builder().
		Update("content").
		Set(record).
		Where(
			builder.C("id").Eq(content.ID),
		).
		Executor().
		ExecContext(r.Context())
	if err != nil {
		return exceptions.ThrowPrivateError(err)
	}
	if err := r.upsertTags(content, form, true); err != nil {
		return err
	}
	return nil
}

func (r *ContentRepository) upsertTags(content *Content, form *forms.ContentForm, update bool) error {
	if form.Tags != "" {
		tags := []int{}
		if err := json.Unmarshal([]byte(form.Tags), &tags); err != nil {
			return err
		}

		if update {
			_, err := r.Connection().
				Builder().
				Delete("content_tag").
				Where(
					builder.C("content_id").Eq(content.ID),
				).
				Executor().
				ExecContext(r.Context())

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

		_, err := r.Connection().
			Builder().
			Insert("content_tag").
			Rows(records).
			Executor().
			ExecContext(r.Context())

		if err != nil {
			return exceptions.ThrowPrivateError(err)
		}
	}

	return nil
}

func (r *ContentRepository) FindAllByUser(userid, page int64) ([]PublicContent, float64, error) {
	query := r.withFullUserProfile().Where(
		builder.I("user.id").Eq(userid),
	)
	query, count := database.WithPagination(r.Context(), query, uint(page), 4)

	var content []Content
	if err := query.ScanStructsContext(r.Context(), &content); err != nil {
		return nil, 0, exceptions.ThrowPrivateError(err)
	}

	response, err := withMutableResponse(r.Context(), r.Connection(), content)
	if err != nil {
		return nil, 0, err
	}

	return response, count, nil
}

func (r *ContentRepository) FindContentByIDAndUser(id, userid int64) (Content, error) {
	query := r.withFullUserProfile().Where(
		builder.And(
			builder.I("content.id").Eq(id),
			builder.I("user.id").Eq(userid),
		),
	)

	content := Content{}
	if ok, err := query.ScanStructContext(r.Context(), &content); err != nil {
		return content, exceptions.ThrowPrivateError(err)
	} else if !ok {
		return content, exceptions.ThrowPublicError(errors.New("content with the required ID was not found"))
	}
	if tags, err := fetchContentTags(r.Context(), r.Connection(), id); err == nil {
		content.withTags(tags)
	}
	return content, nil
}

func (r *ContentRepository) FindContentByID(id int64) (Content, error) {
	query := r.withFullUserProfile().Where(
		builder.I("content.id").Eq(id),
	)

	content := Content{}
	if ok, err := query.ScanStructContext(r.Context(), &content); err != nil {
		return content, exceptions.ThrowPrivateError(err)
	} else if !ok {
		return content, exceptions.ThrowPublicError(errors.New("content with the required ID was not found"))
	}
	if tags, err := fetchContentTags(r.Context(), r.Connection(), id); err == nil {
		content.withTags(tags)
	}
	return content, nil
}

func (r *ContentRepository) FindAllContent(label string, page int64) ([]PublicContent, float64, error) {
	query := r.withFullUserProfile()
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

	query, count := database.WithPagination(r.Context(), query, uint(page), 4)
	var content []Content
	if err := query.ScanStructsContext(r.Context(), &content); err != nil {
		return nil, 0, exceptions.ThrowPrivateError(err)
	}

	response, err := withMutableResponse(r.Context(), r.Connection(), content)
	if err != nil {
		return nil, 0, err
	}
	return response, count, nil
}

// избавиться
func (c *Content) withTags(tags []Tag) {
	if tags != nil {
		c.Tags = tags
	} else {
		// ToDo: Пофиксить эту шляпу на сторону клиента
		c.Tags = []Tag{}
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
	for i := range contents {
		idx = append(idx, contents[i].ID)
	}
	tags, err := UseTagRepository(ctx, conn).ContentGroupedTags(idx...)
	if err != nil {
		return nil, err
	}
	publicContents := make([]PublicContent, 0, len(contents))
	for i := range contents {
		if value, ok := tags[contents[i].ID]; ok {
			contents[i].withTags(value)
		}
		publicContents = append(publicContents, contents[i].GetPublicContent())
	}
	return publicContents, nil
}
