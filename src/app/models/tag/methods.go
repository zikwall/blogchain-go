package tag

import (
	builder "github.com/doug-martin/goqu/v9"
	"github.com/zikwall/blogchain/src/app/exceptions"
)

func (self Model) Find() *builder.SelectDataset {
	return self.Connection().Builder().Select("tags.*").From("tags")
}

func (self Model) All() ([]Tag, error) {
	var tags []Tag

	if err := self.Find().ScanStructsContext(self.Context(), &tags); err != nil {
		return nil, exceptions.NewErrDatabaseAccess(err)
	}

	return tags, nil
}

func (self Model) OnContentCondition(query *builder.SelectDataset, id ...interface{}) *builder.SelectDataset {
	where := builder.I("content_tag.content_id").Eq(id)

	if len(id) > 1 {
		where = builder.I("content_tag.content_id").In(id...)
	}

	return query.
		LeftJoin(
			builder.T("content_tag"),
			builder.On(
				builder.I("content_tag.tag_id").Eq(builder.I("tags.id")),
			),
		).
		Where(where)
}

type TagContent struct {
	Tag
	ContentId int64 `db:"content_id"`
}

func (self Model) ContentGroupedTags(id ...interface{}) (map[int64][]Tag, error) {
	var tags []TagContent

	query := self.Find().SelectAppend(
		"content_tag.content_id",
	)

	if err := self.OnContentCondition(query, id...).ScanStructsContext(self.Context(), &tags); err != nil {
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

func (self Model) ContentTags(id int64) ([]Tag, error) {
	var tags []Tag

	if err := self.OnContentCondition(self.Find(), id).ScanStructsContext(self.Context(), &tags); err != nil {
		return nil, exceptions.NewErrDatabaseAccess(err)
	}

	return tags, nil
}
