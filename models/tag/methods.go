package tag

import "github.com/zikwall/blogchain/di"

func GetTags() ([]Tag, error) {
	tags := []Tag{}

	err := di.DI().Database.Query().
		Select("*").
		From("tags").
		All(&tags)

	return tags, err
}
