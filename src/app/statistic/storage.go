package statistic

import (
	builder "github.com/doug-martin/goqu/v9"
	"github.com/zikwall/blogchain/src/platform/clickhouse"
)

type Viewers struct {
	PostId int64  `db:"post_id"`
	Views  uint64 `db:"views"`
}

func viewersQuery() *builder.SelectDataset {
	return builder.
		Select(
			builder.C("post_id"),
			builder.L("count(*) as views"),
		).
		From("post_stats").
		GroupBy(builder.C("post_id"))
}

func GetPostViewersCount(ch *clickhouse.Clickhouse, post, owner int64) (uint64, error) {
	var count []Viewers

	query := viewersQuery()
	query = query.Where(
		builder.And(
			builder.C("post_id").Eq(post),
			builder.C("owner_id").Eq(owner),
		),
	)

	rawQuery, _, _ := query.ToSQL()

	if err := ch.Query().Select(&count, rawQuery); err != nil {
		return 0, err
	}

	if len(count) == 0 {
		return 0, nil
	}

	return count[0].Views, nil
}

func GetPostViewersCounts(ch *clickhouse.Clickhouse, posts ...int64) (map[int64]uint64, error) {
	var views []Viewers
	counts := map[int64]uint64{}

	query := viewersQuery()
	query = query.Where(
		builder.And(
			builder.C("post_id").In(posts),
		),
	)

	rawQuery, _, _ := query.ToSQL()

	if err := ch.Query().Select(&views, rawQuery); err != nil {
		return counts, err
	}

	for _, view := range views {
		counts[view.PostId] = view.Views
	}

	return counts, nil
}
