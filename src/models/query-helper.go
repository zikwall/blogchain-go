package models

import (
	dbx "github.com/go-ozzo/ozzo-dbx"
	"math"
)

func QueryCount(query *dbx.SelectQuery, pageSize int64) (float64, error) {
	var count int64
	var countPages float64
	var err error
	cloneQuery := *query

	if err := cloneQuery.Select("COUNT(*) as ctn").Row(&count); err == nil {
		countPages = math.Ceil(float64(count) / float64(pageSize))
	}

	return countPages, err
}
