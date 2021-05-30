package tag

import (
	builder "github.com/doug-martin/goqu/v9"
	"github.com/zikwall/blogchain/src/app/exceptions"
)

func (self TagModel) Find() *builder.SelectDataset {
	return self.Connection().Builder().Select("tags.*").From("tags")
}

func (self TagModel) All() ([]Tag, error) {
	var tags []Tag

	query := self.Find()

	if err := query.ScanStructs(&tags); err != nil {
		return nil, exceptions.NewErrDatabaseAccess(err)
	}

	return tags, nil
}

func (self TagModel) OnContentCondition(query *builder.SelectDataset, id ...interface{}) *builder.SelectDataset {
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

func (self TagModel) ContentGroupedTags(id ...interface{}) (map[int64][]Tag, error) {
	var tags []TagContent

	withContent := func(query *builder.SelectDataset) *builder.SelectDataset {
		return query.SelectAppend(
			"content_tag.content_id",
		)
	}

	query := self.Find()
	query = self.OnContentCondition(query, id...)
	query = withContent(query)

	if err := query.ScanStructs(&tags); err != nil {
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

func (self TagModel) ContentTags(id int64) ([]Tag, error) {
	var tags []Tag

	query := self.Find()
	query = self.OnContentCondition(query, id)

	if err := query.ScanStructs(&tags); err != nil {
		return nil, exceptions.NewErrDatabaseAccess(err)
	}

	return tags, nil
}
