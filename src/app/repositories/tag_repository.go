package repositories

import (
	"context"
	builder "github.com/doug-martin/goqu/v9"
	"github.com/zikwall/blogchain/src/app/exceptions"
	"github.com/zikwall/blogchain/src/platform/database"
)

type TagRepository struct {
	Repository
}

type TagContent struct {
	Tag
	ContentID int64 `db:"content_id"`
}

func UseTagRepository(ctx context.Context, conn *database.Connection) TagRepository {
	return TagRepository{
		Repository{connection: conn, context: ctx},
	}
}

func (tr TagRepository) find() *builder.SelectDataset {
	query := tr.Connection().Builder()
	return query.Select("tags.*").From("tags")
}

func (tr TagRepository) All() ([]Tag, error) {
	var tags []Tag

	if err := tr.find().ScanStructsContext(tr.Context(), &tags); err != nil {
		return nil, exceptions.ThrowPrivateError(err)
	}

	return tags, nil
}

func (tr TagRepository) ContentGroupedTags(id ...interface{}) (map[int64][]Tag, error) {
	var tags []TagContent

	query := tr.find().SelectAppend(
		"content_tag.content_id",
	)

	if err := withContent(query, id...).ScanStructsContext(tr.Context(), &tags); err != nil {
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

func (tr TagRepository) ContentTags(id int64) ([]Tag, error) {
	var tags []Tag

	if err := withContent(tr.find(), id).ScanStructsContext(tr.Context(), &tags); err != nil {
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
