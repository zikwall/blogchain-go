package tag

import (
	builder "github.com/doug-martin/goqu/v9"
)

func (self TagModel) Find() *builder.SelectDataset {
	return self.QueryBuilder().Select("tags.*").From("tags")
}

func (self TagModel) All() ([]Tag, error) {
	var tags []Tag

	query := self.Find()

	if err := query.ScanStructs(&tags); err != nil {
		return nil, err
	}

	return tags, nil
}

func (self TagModel) OnContentCondition(query *builder.SelectDataset, id int64) *builder.SelectDataset {
	return query.
		LeftJoin(
			builder.T("content_tag"),
			builder.On(
				builder.I("content_tag.tag_id").Eq(builder.I("tags.id")),
			),
		).
		Where(
			builder.I("content_tag.content_id").Eq(id),
		)
}

func (self TagModel) ContentTags(id int64) ([]Tag, error) {
	var tags []Tag

	query := self.Find()
	query = self.OnContentCondition(query, id)

	if err := query.ScanStructs(&tags); err != nil {
		return nil, err
	}

	return tags, nil
}
