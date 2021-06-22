package statistic

import (
	"github.com/zikwall/clickhouse-buffer/src/types"
)

type PostStats struct {
	PostId   uint64 `json:"post_id"`
	OwnerId  uint64 `json:"owner_id"`
	Os       string `json:"-"`
	Browser  string `json:"-"`
	Platform string `json:"-"`
	Ip       string `json:"-"`
	Country  string `json:"-"`
	Region   string `json:"-"`
	InsertTs string `json:"-"`
	Date     string `json:"-"`
}

func (b PostStats) Row() types.RowSlice {
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
