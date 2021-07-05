package statistic

import (
	"context"
	builder "github.com/doug-martin/goqu/v9"
	"github.com/zikwall/blogchain/src/platform/clickhouse"
)

type Viewers struct {
	PostId int64  `db:"post_id"`
	Views  uint64 `db:"views"`
}

func viewersAggregateQuery() *builder.SelectDataset {
	return builder.
		Select(
			builder.C("post_id"),
			builder.L("sum(views) as views"),
		).
		From("post_stats_views").
		GroupBy(
			builder.C("post_id"),
		)
}

func GetPostViewersCount(ctx context.Context, ch *clickhouse.Clickhouse, post int64) (uint64, error) {
	var count uint64
	var postId int64

	rawQuery, _, _ := viewersAggregateQuery().
		Where(
			builder.And(
				builder.C("post_id").Eq(post),
			),
		).
		ToSQL()

	if err := ch.Query().QueryRowContext(ctx, rawQuery).Scan(&postId, &count); err != nil {
		return 0, err
	}

	// current
	count++

	return count, nil
}

func GetPostsViewersCount(ctx context.Context, ch *clickhouse.Clickhouse, posts ...int64) (map[int64]uint64, error) {
	var views []Viewers
	counts := map[int64]uint64{}

	rawQuery, _, _ := viewersAggregateQuery().
		Where(
			builder.And(
				builder.C("post_id").In(posts),
			),
		).ToSQL()

	if err := ch.Query().SelectContext(ctx, &views, rawQuery); err != nil {
		return counts, err
	}

	for _, view := range views {
		counts[view.PostId] = view.Views
	}

	return counts, nil
}
