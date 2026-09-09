package helper

import (
	"strings"

	"gorm.io/gorm"
)

func ApplySmartSearch(query *gorm.DB, column, search string) *gorm.DB {
	if search == "" {
		return query
	}
	words := strings.Fields(search)
	for _, word := range words {
		query = query.Where(column+" ILIKE ?", "%"+word+"%")
	}
	return query
}
