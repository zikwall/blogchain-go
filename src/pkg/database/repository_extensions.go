package database

import (
	"context"
	"math"

	"github.com/zikwall/blogchain/src/pkg/exceptions"

	builder "github.com/doug-martin/goqu/v9"
)

// JoinFn
// ```code
// 	func withUserQuery(query *builder.SelectDataset) *builder.SelectDataset {
//		query = query.SelectAppend(
//			builder.I(...),
//		)
//		query = query.LeftJoin(
//			builder.T("user"),
//			builder.On(
//				builder.I("user.id").Eq(builder.I("content.user_id")),
//			),
//		)
//		return query
//	}
// ```
type JoinFn func(query *builder.SelectDataset) *builder.SelectDataset

// JoinWith the function is a wrapper for convenient merging of query parts, for example, such as SQL JOIN
//
// ```code
// 	func withFullUserProfile() *builder.SelectDataset {
//		query = ...
//		query = joinWith(query, withUserQuery, withProfileQuery)
//		return query
//	}
// ```
func JoinWith(query *builder.SelectDataset, joins ...JoinFn) *builder.SelectDataset {
	for _, join := range joins {
		query = join(query)
	}

	return query
}

// queryCount function determines the number of "pages" depending on the page size (the larger, the smaller the number of pages)
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

// WithPagination wrapper function for creating page-by-page pagination, automatically tracks the offset and size
func WithPagination(ctx context.Context, query *builder.SelectDataset, page, size uint) (
	paginatedQuery *builder.SelectDataset, count float64,
) {
	count, _ = queryCount(ctx, query, size)
	paginatedQuery = query.Offset(page * size).Limit(size)

	return paginatedQuery, count
}
