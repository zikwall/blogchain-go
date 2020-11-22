package models

import (
	"github.com/doug-martin/goqu/v9"
	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/zikwall/blogchain/src/service"
)

type BlogchainModel struct{}

func (b *BlogchainModel) Query() *dbx.DB {
	return service.GetBlogchainServiceInstance().GetBlogchainDatabaseInstance().Query()
}

func QueryBuilder() *goqu.Database {
	return builder.Builder()
}
