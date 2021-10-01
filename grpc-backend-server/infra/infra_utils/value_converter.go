package ifu

import (
	"fmt"
	"strconv"
	"time"

	"git.paylabo.com/c002/harp/backend-purchase/app/usecase/utils"
	pb "git.paylabo.com/c002/harp/backend-purchase/interfaces/proto/git.paylabo.com/c002/harp"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/protobuf/types/known/wrapperspb"

	wrappers "github.com/golang/protobuf/ptypes/wrappers"
)

func ConvertStringPointerToInt64(v *string) int64 {
	if v == nil {
		return 0
	}
	i, _ := strconv.ParseInt(*v, 10, 64)
	return i
}

func ConvertStringPointerToString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

func ConvertStringPointerToFloat64(v *string) float64 {
	if v == nil {
		return 0
	}

	i, err := strconv.ParseFloat(*v, 64)
	if err == nil {
		return i
	}
	return 0
}

func ConvertInt32ValueToString(i *wrappers.Int32Value) *string {
	if i == nil {
		return nil
	}
	v := strconv.Itoa(int(i.Value))
	return &v
}

func ConvertDoubleValueToString(i *wrappers.DoubleValue) *string {
	if i == nil {
		return nil
	}
	v := strconv.FormatFloat(float64(i.Value), 'f', 2, 64)
	return &v
}

func ConvertStringValueToString(i *wrappers.StringValue) *string {
	if i == nil {
		return nil
	}
	v := i.Value
	return &v
}

func ToInt64(v string) int64 {
	i, _ := strconv.ParseInt(v, 10, 64)
	return i
}

func ToFloat64(v string) float64 {
	i, _ := strconv.ParseFloat(v, 64)
	return i
}

func ToInt64Pointer(v string) *int64 {
	if v == "" {
		return nil
	}

	i, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return nil
	}
	ip := new(int64)
	*ip = i
	return ip

}

func ToBoolPointer(v *wrappers.BoolValue) *bool {
	if v == nil {
		return nil
	}
	result := v.Value
	return &result
}

func Int64ToInt32Value(v int64) *wrappers.Int32Value {
	return &wrappers.Int32Value{Value: int32(v)}
}

func Int64PointerToInt32Value(v *int64) *wrappers.Int32Value {
	if v == nil {
		return nil
	}
	return &wrappers.Int32Value{Value: int32(*v)}
}

//internal use only
func stringToInt32Value(s string) *wrappers.Int32Value {

	i, err := strconv.ParseInt(s, 10, 32)
	if err == nil {
		return wrapperspb.Int32(int32(i))
	}
	return nil
}

//internal use only
func stringToDoubleValue(s string) *wrappers.DoubleValue {

	i, err := strconv.ParseFloat(s, 64)
	if err == nil {
		return wrapperspb.Double(float64(i))
	}
	return nil
}

func ToInt32Value(s interface{}) *wrappers.Int32Value {

	if s == nil {
		return nil
	}

	switch v := s.(type) {
	case string:
		if str, ok := s.(string); ok {
			return stringToInt32Value(str)
		}
	case *string:
		if str, ok := s.(*string); ok {
			if str == nil {
				return nil
			}
			return stringToInt32Value(*str)
		}
	default:
		{
			fmt.Println("type is", v)
		}
	}
	return nil
}

func ToDoubleValue(s interface{}) *wrappers.DoubleValue {

	if s == nil {
		return nil
	}

	switch v := s.(type) {
	case string:
		if str, ok := s.(string); ok {
			return stringToDoubleValue(str)
		}
	case *string:
		if str, ok := s.(*string); ok {
			if str == nil {
				return nil
			}
			return stringToDoubleValue(*str)
		}
	default:
		{
			fmt.Println("type is", v)
		}
	}
	return nil
}

func ToTimestamp(t time.Time) *timestamp.Timestamp {
	if t.IsZero() {
		return nil
	}
	ts, err := ptypes.TimestampProto(t)
	if err != nil {
		return nil
	}
	return ts

}

func ToSortEnum(e pb.SortEnum) utils.SortEnum {
	return utils.SortEnum(int(e))
}
