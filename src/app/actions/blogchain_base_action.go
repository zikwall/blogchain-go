package actions

import (
	"github.com/zikwall/blogchain/src/app/lib"
	"github.com/zikwall/blogchain/src/platform/database"
)

type (
	BlogchainActionProvider struct {
		rsa lib.RSA
		db  *database.BlogchainDatabaseInstance
	}
	ActionsRequiredInstances struct {
		RSA lib.RSA
		Db  *database.BlogchainDatabaseInstance
	}
)

func NewBlogchainActionProvider(conf ActionsRequiredInstances) BlogchainActionProvider {
	a := BlogchainActionProvider{
		rsa: conf.RSA,
		db:  conf.Db,
	}

	return a
}
