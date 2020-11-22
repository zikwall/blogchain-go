package models

import (
	"github.com/doug-martin/goqu/v9"
	"math"
)

func QueryCount(query *goqu.SelectDataset, pageSize uint) (float64, error) {
	var count int64
	var countPages float64
	var err error
	cloneQuery := *query

	if count, err = cloneQuery.Count(); err == nil {
		countPages = math.Ceil(float64(count) / float64(pageSize))
	}

	return countPages, err
}
