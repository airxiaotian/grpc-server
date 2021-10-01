package ifu

import (
	"reflect"
	"strings"

	"git.paylabo.com/c002/harp/backend-purchase/app/usecase/utils"
	"github.com/iancoleman/strcase"
)

func GetOrderByString(v interface{}) string {
	orderBys := make(map[string]string, 20)
	t := reflect.ValueOf(v).Elem()
	if t.Kind() == reflect.Struct {
		for i := 0; i < t.NumField(); i++ {
			valueField := t.Field(i)
			fieldName := strcase.ToSnake(t.Type().Field(i).Name)
			//tag := typeField.Tag
			if valueField.Kind() == reflect.Int32 {
				if e, ok := valueField.Interface().(utils.SortEnum); ok {
					switch e {
					case utils.SortEnumASC:
						orderBys[fieldName] = "asc"
					case utils.SortEnumDESC:
						orderBys[fieldName] = "desc"
					}
				}
			}
			//fmt.Printf("Field Name: %s,\t Field Value: %v,\t Tag Value: %s\n", typeField.Name, valueField.Interface(), tag.Get("tag_name"))
		}
		var sb strings.Builder
		isFirst := true
		for k, v := range orderBys {
			if !isFirst {
				sb.WriteString(",")
			} else {
				isFirst = false
			}
			sb.WriteString(k)
			sb.WriteString(" ")
			sb.WriteString(v)
		}
		return sb.String()
	}
	return ""
}
