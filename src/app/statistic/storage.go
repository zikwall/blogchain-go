package statistic

import (
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

func GetPostViewersCount(ch *clickhouse.Clickhouse, post int64) (uint64, error) {
	var count []Viewers

	rawQuery, _, _ := viewersAggregateQuery().Where(
		builder.And(
			builder.C("post_id").Eq(post),
		),
	).ToSQL()

	if err := ch.Query().Select(&count, rawQuery); err != nil {
		return 0, err
	}

	if len(count) == 0 {
		return 0, nil
	}

	return count[0].Views, nil
}

func GetPostsViewersCount(ch *clickhouse.Clickhouse, posts ...int64) (map[int64]uint64, error) {
	var views []Viewers
	counts := map[int64]uint64{}

	rawQuery, _, _ := viewersAggregateQuery().Where(
		builder.And(
			builder.C("post_id").In(posts),
		),
	).ToSQL()

	if err := ch.Query().Select(&views, rawQuery); err != nil {
		return counts, err
	}

	for _, view := range views {
		counts[view.PostId] = view.Views
	}

	return counts, nil
}
