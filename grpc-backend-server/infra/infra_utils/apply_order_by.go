package ifu

import (
	"github.com/jinzhu/gorm"
)

func ApplyOrderBys(db *gorm.DB, orderBy interface{}, defaultOrderByField string) *gorm.DB {
	appliedByParams := false
	if orderBy != nil {
		if orderByString := GetOrderByString(orderBy); len(orderByString) > 0 {
			db = db.Order(orderByString)
			appliedByParams = true
		}
	}
	if !appliedByParams {
		db = db.Order(defaultOrderByField)
	}
	return db
}
