package models

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/zikwall/blogchain/src/app/exceptions"
	"math"
)

func QueryCount(query *goqu.SelectDataset, pageSize uint) (float64, error) {
	var count int64
	var countPages float64
	var err error
	cloneQuery := *query

	if count, err = cloneQuery.Count(); err == nil {
		countPages = math.Ceil(float64(count) / float64(pageSize))
	}

	return countPages, exceptions.NewErrDatabaseAccess(err)
}

func WithPagination(query *goqu.SelectDataset, page, size uint) (*goqu.SelectDataset, float64) {
	countPages, _ := QueryCount(query, size)
	query = query.Offset(page * size).Limit(size)

	return query, countPages
}
