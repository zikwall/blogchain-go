package models

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/zikwall/blogchain/src/service"
)

type BlogchainModel struct{}

func (model BlogchainModel) QueryBuilder() *goqu.Database {
	return service.GetBlogchainServiceInstance().GetBlogchainDatabaseInstance().Builder()
}
