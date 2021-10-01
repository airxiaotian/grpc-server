package infra

import (
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/jinzhu/gorm"
)

func AddWhereIfNotEmptyString(db *gorm.DB, query, value string) *gorm.DB {
	if len(value) == 0 {
		return db
	}
	return db.Where(query, value)
}

func NullExprIfEmptyString(value string) interface{} {
	if len(value) == 0 {
		return gorm.Expr("NULL")
	}
	return value
}

func NullExprIfEmptyTime(value time.Time) interface{} {
	if value.IsZero() {
		return gorm.Expr("NULL")
	}
	return value
}

func ToTime(t *timestamp.Timestamp) time.Time {
	if t == nil {
		return time.Time{}
	}
	return t.AsTime()
}
