package repository

import (
	"github.com/zikwall/clickhouse-buffer/src/buffer"
)

const PostStatsTable = "blogchain.post_stats"

var PostStatsColumns = []string{
	"post_id", "owner_id", "os", "browser", "platform",
	"ip", "country", "region", "insert_ts", "date",
}

type PostStats struct {
	PostID   uint64 `json:"post_id"`
	OwnerID  uint64 `json:"owner_id"`
	Os       string `json:"os"`
	Browser  string `json:"browser"`
	Platform string `json:"platform"`
	IP       string `json:"ip"`
	Country  string `json:"country"`
	Region   string `json:"region"`
	InsertTS string `json:"insert_ts"`
	Date     string `json:"date"`
}

func (b *PostStats) Row() buffer.RowSlice {
	return []interface{}{
		b.PostID,
		b.OwnerID,
		b.Os,
		b.Browser,
		b.Platform,
		b.IP,
		b.Country,
		b.Region,
		b.InsertTS,
		b.Date,
	}
}
