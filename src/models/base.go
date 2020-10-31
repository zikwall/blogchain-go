package models

import (
	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/zikwall/blogchain/src/service"
)

type BlogchainModel struct{}

func (b *BlogchainModel) Query() *dbx.DB {
	return service.GetBlogchainServiceInstance().GetBlogchainDatabaseInstance().Query()
}
