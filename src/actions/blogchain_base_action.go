package actions

import (
	"github.com/zikwall/blogchain/src/lib"
	"github.com/zikwall/blogchain/src/service"
)

type BlogchainActionProvider struct {
	rsa lib.RSA
	db  *service.BlogchainDatabaseInstance
}

func NewBlogchainActionProvider(rsa lib.RSA, db *service.BlogchainDatabaseInstance) BlogchainActionProvider {
	a := BlogchainActionProvider{
		rsa: rsa,
		db:  db,
	}

	return a
}
