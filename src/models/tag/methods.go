package tag

import (
	dbx "github.com/go-ozzo/ozzo-dbx"
)

func (t TagModel) GetTags() ([]Tag, error) {
	tags := []Tag{}

	err := t.Query().
		Select("*").
		From("tags").
		All(&tags)

	return tags, err
}

func (t TagModel) GetTagsByContent(id int64) ([]Tag, error) {
	tags := []Tag{}

	err := t.Query().
		Select("tags.*").
		From("tags").
		LeftJoin("content_tag", dbx.NewExp("content_tag.tag_id=tags.id")).
		Where(dbx.HashExp{"content_tag.content_id": id}).
		All(&tags)

	return tags, err
}

func AttachTagQuery(query *dbx.SelectQuery, tag string) {
	query.LeftJoin("content_tag", dbx.NewExp("content_tag.content_id=content.id")).
		LeftJoin("tags", dbx.NewExp("content_tag.tag_id=tags.id")).
		Where(dbx.HashExp{"tags.label": tag})
}
