package actions

import (
	"github.com/zikwall/blogchain/src/platform/container"
	"github.com/zikwall/blogchain/src/platform/database"
)

type (
	BlogchainActionProvider struct {
		rsa container.RSA
		db  *database.BlogchainDatabaseInstance
	}
	ActionsRequiredInstances struct {
		RSA container.RSA
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
