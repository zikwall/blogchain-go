package statistic

import "github.com/zikwall/blogchain/src/platform/clickhouse"

var postStatsTable = clickhouse.Table{Name: "post_stats", Columns: []string{
	"post_id",
	"owner_id",
	"hostname",
	"os",
	"browser",
	"platform",
	"ip",
	"country",
	"region",
	"insert_ts",
	"date",
}}

type PostStats struct {
	PostId   uint64 `json:"post_id"`
	OwnerId  uint64 `json:"owner_id"`
	Os       string `json:"os"`
	Browser  string `json:"browser"`
	Platform string `json:"platform"`
	Ip       string `json:"-"`
	Country  string `json:"-"`
	Region   string `json:"-"`
	InsertTs string `json:"-"`
	Date     string `json:"-"`
}

func (b *PostStats) flatten() []interface{} {
	return []interface{}{
		b.PostId,
		b.OwnerId,
		b.Os,
		b.Browser,
		b.Platform,
		b.Ip,
		b.Country,
		b.Region,
		b.InsertTs,
		b.Date,
	}
}
