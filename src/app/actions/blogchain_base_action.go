package actions

import (
	"github.com/zikwall/blogchain/src/app/lib"
	"github.com/zikwall/blogchain/src/platform/service"
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
