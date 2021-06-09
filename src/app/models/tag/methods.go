package tag

import (
	builder "github.com/doug-martin/goqu/v9"
	"github.com/zikwall/blogchain/src/app/exceptions"
)

func (m Model) find() *builder.SelectDataset {
	query := m.Connection().Builder()
	return query.Select("tags.*").From("tags")
}

func (m Model) All() ([]Tag, error) {
	var tags []Tag

	if err := m.find().ScanStructsContext(m.Context(), &tags); err != nil {
		return nil, exceptions.NewErrDatabaseAccess(err)
	}

	return tags, nil
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

type TagContent struct {
	Tag
	ContentId int64 `db:"content_id"`
}

func (m Model) ContentGroupedTags(id ...interface{}) (map[int64][]Tag, error) {
	var tags []TagContent

	query := m.find().SelectAppend(
		"content_tag.content_id",
	)

	if err := withContent(query, id...).ScanStructsContext(m.Context(), &tags); err != nil {
		return nil, exceptions.NewErrDatabaseAccess(err)
	}

	grouped := make(map[int64][]Tag, len(tags))

	for _, tag := range tags {
		grouped[tag.ContentId] = append(grouped[tag.ContentId], Tag{
			Id:    tag.Id,
			Name:  tag.Name,
			Label: tag.Label,
		})
	}

	return grouped, nil
}

func (m Model) ContentTags(id int64) ([]Tag, error) {
	var tags []Tag

	if err := withContent(m.find(), id).ScanStructsContext(m.Context(), &tags); err != nil {
		return nil, exceptions.NewErrDatabaseAccess(err)
	}

	return tags, nil
}
