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

func queryCount(ctx context.Context, query *builder.SelectDataset, pageSize uint) (float64, error) {
	var count int64
	var countPages float64
	var err error
	cloneQuery := *query

	if count, err = cloneQuery.CountContext(ctx); err == nil {
		countPages = math.Ceil(float64(count) / float64(pageSize))
	}

	return countPages, exceptions.ThrowPrivateError(err)
}

func withPagination(ctx context.Context, query *builder.SelectDataset, page, size uint) (
	pquery *builder.SelectDataset, count float64,
) {
	count, _ = queryCount(ctx, query, size)
	pquery = query.Offset(page * size).Limit(size)

	return pquery, count
}
