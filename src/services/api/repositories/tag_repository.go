package repositories

import (
	"context"

	"github.com/zikwall/blogchain/src/pkg/database"
	"github.com/zikwall/blogchain/src/pkg/exceptions"

	builder "github.com/doug-martin/goqu/v9"
)

type TagRepository struct {
	Repository
}

type TagContent struct {
	Tag
	ContentID int64 `db:"content_id"`
}

func UseTagRepository(ctx context.Context, conn *database.Connection) *TagRepository {
	return &TagRepository{
		Repository{connection: conn, context: ctx},
	}
}

func (r *TagRepository) find() *builder.SelectDataset {
	query := r.Connection().Builder()
	return query.Select("tags.*").From("tags")
}

func (r *TagRepository) All() ([]Tag, error) {
	var tags []Tag
	if err := r.find().ScanStructsContext(r.Context(), &tags); err != nil {
		return nil, exceptions.ThrowPrivateError(err)
	}
	return tags, nil
}

func (r *TagRepository) ContentGroupedTags(id ...interface{}) (map[int64][]Tag, error) {
	var tags []TagContent
	query := r.find().SelectAppend(
		"content_tag.content_id",
	)
	if err := withContent(query, id...).ScanStructsContext(r.Context(), &tags); err != nil {
		return nil, exceptions.ThrowPrivateError(err)
	}
	grouped := make(map[int64][]Tag, len(tags))
	for _, tag := range tags {
		grouped[tag.ContentID] = append(grouped[tag.ContentID], Tag{
			ID:    tag.ID,
			Name:  tag.Name,
			Label: tag.Label,
		})
	}
	return grouped, nil
}

func (r *TagRepository) ContentTags(id int64) ([]Tag, error) {
	var tags []Tag
	if err := withContent(r.find(), id).ScanStructsContext(r.Context(), &tags); err != nil {
		return nil, exceptions.ThrowPrivateError(err)
	}
	return tags, nil
}

func fetchContentTags(ctx context.Context, conn *database.Connection, id int64) ([]Tag, error) {
	tags, err := UseTagRepository(ctx, conn).ContentTags(id)
	return tags, err
}

func withContent(query *builder.SelectDataset, id ...interface{}) *builder.SelectDataset {
	query = query.LeftJoin(
		builder.T("content_tag"),
		builder.On(
			builder.I("content_tag.tag_id").Eq(builder.I("tags.id")),
		),
	)
	if len(id) > 1 {
		query = query.Where(builder.I("content_tag.content_id").In(id...))
	} else {
		query = query.Where(builder.I("content_tag.content_id").Eq(id))
	}
	return query
}
