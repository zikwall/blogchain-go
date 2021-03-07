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

func (a BlogchainActionProvider) _common(status uint8, message string) BlogchainMessageResponse {
	return BlogchainMessageResponse{
		BlogchainCommonResponseAttributes: BlogchainCommonResponseAttributes{
			Status: status,
		},
		Message: message,
	}
}

func (a BlogchainActionProvider) message(message string) BlogchainMessageResponse {
	return a._common(200, message)
}

func (a BlogchainActionProvider) error(err error) BlogchainMessageResponse {
	return a._common(100, err.Error())
}
