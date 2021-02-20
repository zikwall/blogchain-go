package models

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/zikwall/blogchain/src/platform/service"
)

type (
	QueryInterface interface {
		Builder() *goqu.Database
	}
	BlogchainModel struct {
		Connection *service.BlogchainDatabaseInstance
	}
)

func (model BlogchainModel) Builder() *goqu.Database {
	return model.Connection.Builder()
}
