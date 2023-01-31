package store

import (
	"fmt"

	"github.com/ralexstokes/relay-monitor/pkg/types"
)

func BuildCategoryFilterClause(query string, filter *types.AnalysisQueryFilter) string {
	if filter == nil {
		return query
	}

	return query + ` AND category ` + filter.Comparator + ` '` + fmt.Sprintf("%d", filter.Category) + `'`
}
