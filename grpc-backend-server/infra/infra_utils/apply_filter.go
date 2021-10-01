package ifu

import (
	"strconv"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/jinzhu/gorm"
)

func ApplyFilter(db *gorm.DB, filter []string, fieldName string) *gorm.DB {

	if len(filter) > 0 {
		if arr, err := SliceAtoi(filter); err == nil {
			return db.Where(fieldName+" IN (?)", arr)
		}
	}
	return db
}

func ApplyStringFilter(db *gorm.DB, filter []string, fieldName string) *gorm.DB {
	if len(filter) > 0 {
		return db.Where(fieldName+" IN (?)", filter)
	}
	return db
}

func ApplyDateRange(db *gorm.DB, field string, from *timestamp.Timestamp, to *timestamp.Timestamp) *gorm.DB {
	if from != nil && to != nil {
		fromTime := from.AsTime()
		toTime := to.AsTime()
		if fromTime.Before(toTime) {
			db = db.Where(field+" BETWEEN ? AND ?", fromTime, toTime)
		}
	}
	return db
}

func SliceAtoi(sa []string) ([]int, error) {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err != nil {
			return si, err
		}
		si = append(si, i)
	}
	return si, nil
}

func SliceItoa(sa []int64) []string {
	si := make([]string, 0, len(sa))
	for _, a := range sa {
		i := strconv.Itoa(int(a))
		si = append(si, i)
	}
	return si
}
