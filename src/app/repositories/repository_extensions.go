package repositories

import (
	"context"
	builder "github.com/doug-martin/goqu/v9"
	"github.com/zikwall/blogchain/src/app/exceptions"
	"math"
)

type joinFn func(query *builder.SelectDataset) *builder.SelectDataset

func joinWith(query *builder.SelectDataset, joins ...joinFn) *builder.SelectDataset {
	for _, join := range joins {
		query = join(query)
	}

	return query
}

func queryCount(context context.Context, query *builder.SelectDataset, pageSize uint) (float64, error) {
	var count int64
	var countPages float64
	var err error
	cloneQuery := *query

	if count, err = cloneQuery.CountContext(context); err == nil {
		countPages = math.Ceil(float64(count) / float64(pageSize))
	}

	return countPages, exceptions.NewErrDatabaseAccess(err)
}

func withPagination(context context.Context, query *builder.SelectDataset, page, size uint) (*builder.SelectDataset, float64) {
	countPages, _ := queryCount(context, query, size)
	query = query.Offset(page * size).Limit(size)

	return query, countPages
}
