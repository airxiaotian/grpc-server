package ifu

import (
	"reflect"

	"github.com/iancoleman/strcase"
)

//GORM failed to bulk insert with Model Slices, so insert with Map
//create a map to insert to db
func ToGormCreateMap(v interface{}) map[string]interface{} {
	t := reflect.ValueOf(v)
	if t.Kind() != reflect.Struct {
		return nil
	}
	fieldsLength := t.NumField()
	GormCreateMap := make(map[string]interface{}, fieldsLength)
	for i := 0; i < fieldsLength; i++ {
		valueField := t.Field(i)
		fieldName := strcase.ToSnake(t.Type().Field(i).Name)
		GormCreateMap[fieldName] = valueField.Interface()
	}
	return GormCreateMap
}
