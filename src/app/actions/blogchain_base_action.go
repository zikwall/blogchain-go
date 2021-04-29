package actions

import (
	"github.com/zikwall/blogchain/src/app/statistic"
	"github.com/zikwall/blogchain/src/platform/container"
	"github.com/zikwall/blogchain/src/platform/database"
	"github.com/zikwall/blogchain/src/platform/maxmind"
)

type (
	BlogchainActionProvider struct {
		rsa          container.RSA
		db           *database.BlogchainDatabaseInstance
		statsBatcher *statistic.ClickhouseBatcher
		finder       *maxmind.Finder
	}
	ActionsRequiredInstances struct {
		RSA          container.RSA
		Db           *database.BlogchainDatabaseInstance
		StatsBatcher *statistic.ClickhouseBatcher
		Finder       *maxmind.Finder
	}
)

func NewBlogchainActionProvider(conf ActionsRequiredInstances) BlogchainActionProvider {
	a := BlogchainActionProvider{
		rsa:          conf.RSA,
		db:           conf.Db,
		statsBatcher: conf.StatsBatcher,
		finder:       conf.Finder,
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

func (a BlogchainActionProvider) response(response interface{}) BlogchainResponse {
	return BlogchainResponse{
		Response: response,
	}
}

func (a BlogchainActionProvider) message(message string) BlogchainMessageResponse {
	return a._common(200, message)
}

func (a BlogchainActionProvider) error(err error) BlogchainMessageResponse {
	return a._common(100, err.Error())
}
