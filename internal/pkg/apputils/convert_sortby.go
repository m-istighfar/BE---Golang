package apputils

import (
	"fmt"
	"strings"
)

func ConvertSortByToSQL(sortBy, defaultSortField, defaultSortOrder string) string {
	if sortBy == "" {
		return fmt.Sprintf("%s %s", defaultSortField, defaultSortOrder)
	}

	order := "ASC"
	if strings.HasPrefix(sortBy, "-") {
		order = "DESC"
		sortBy = sortBy[1:]
	}

	sortableFields := map[string]string{
		"name":       "name",
		"price":      "price",
		"stock":      "stock",
		"created_at": "created_at",
		"updated_at": "updated_at",
	}

	if field, ok := sortableFields[sortBy]; ok {
		return fmt.Sprintf("%s %s", field, order)
	}

	return fmt.Sprintf("%s %s", defaultSortField, defaultSortOrder)
}
