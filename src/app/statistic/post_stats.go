package statistic

import (
	"github.com/zikwall/clickhouse-buffer/src/types"
)

type PostStats struct {
	PostId   uint64 `json:"post_id"`
	OwnerId  uint64 `json:"owner_id"`
	Os       string `json:"os"`
	Browser  string `json:"browser"`
	Platform string `json:"platform"`
	Ip       string `json:"ip"`
	Country  string `json:"country"`
	Region   string `json:"region"`
	InsertTs string `json:"insert_ts"`
	Date     string `json:"date"`
}

func (b *PostStats) Row() types.RowSlice {
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
