package actions

import (
	"github.com/zikwall/blogchain/src/lib"
	"github.com/zikwall/blogchain/src/service"
)

type (
	BlogchainActionProvider struct {
		rsa lib.RSA
		db  *service.BlogchainDatabaseInstance
	}
	ActionsRequiredInstances struct {
		RSA lib.RSA
		Db  *service.BlogchainDatabaseInstance
	}
)

func NewBlogchainActionProvider(conf ActionsRequiredInstances) BlogchainActionProvider {
	a := BlogchainActionProvider{
		rsa: conf.RSA,
		db:  conf.Db,
	}

	return a
}
