package common

import builder "github.com/doug-martin/goqu/v9"

type JoinFn func(query *builder.SelectDataset) *builder.SelectDataset

func JoinWith(query *builder.SelectDataset, joins ...JoinFn) *builder.SelectDataset {
	for _, join := range joins {
		query = join(query)
	}

	return query
}
