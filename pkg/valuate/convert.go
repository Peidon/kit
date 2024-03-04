/*
 * @Copyright @ Shopee
 * @Author: Peidong Xu
 * @Email: peidong.xu@shopee.com
 * @Date: 2023/11/28 10:41 AM
 * @Version: kit
 */

package valuate

import (
	"reflect"
	"strings"
	"time"
)

func toInt(v interface{}) int {
	switch s := v.(type) {
	case Number:
		f := s.ToFloat()
		return int(f)
	case nil:
		return 0
	}
	v = indirect(v)

	switch s := v.(type) {
	case int64:
		return int(s)
	case int32:
		return int(s)
	case int16:
		return int(s)
	case int8:
		return int(s)
	case int:
		return s
	case uint64:
		return int(s)
	case uint32:
		return int(s)
	case uint16:
		return int(s)
	case uint8:
		return int(s)
	case uint:
		return int(s)
	case Number:
		f := s.ToFloat()
		return int(f)
	}
	return -1
}

func toFloat64(v interface{}) float64 {
	switch s := v.(type) {
	case Number:
		return s.ToFloat()
	case nil:
		return 0
	}

	v = indirect(v)

	switch s := v.(type) {
	case int64:
		return float64(s)
	case int32:
		return float64(s)
	case int16:
		return float64(s)
	case int8:
		return float64(s)
	case int:
		return float64(s)
	case uint64:
		return float64(s)
	case uint32:
		return float64(s)
	case uint16:
		return float64(s)
	case uint8:
		return float64(s)
	case uint:
		return float64(s)
	case float32:
		return float64(s)
	case float64:
		return s
	case Number:
		return s.ToFloat()
	}
	return -1
}

func toTime(v interface{}) time.Time {
	v = indirect(v)

	switch s := v.(type) {
	case time.Time:
		return s
	}
	return noopTime
}

func isTime(value interface{}) bool {
	switch value.(type) {
	case time.Time, *time.Time:
		return true
	}
	return false
}

func isArray(value interface{}) bool {
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		return true
	}
	return false
}

func isPtr(value interface{}) bool {
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Pointer, reflect.UnsafePointer:
		return true
	}
	return false
}

func isBool(value interface{}) bool {
	switch value.(type) {
	case bool:
		return true
	}
	return false
}

func isNumber(value interface{}) bool {
	switch value.(type) {
	case float64, float32:
		return true
	case int64, int32, int16, int8, int, uint64, uint32, uint16, uint8, uint:
		return true
	case Number:
		return true
	case Modifier:
		return true
	}
	return false
}

func isInt(value interface{}) bool {
	switch value.(type) {
	case int64, int32, int16, int8, int, uint64, uint32, uint16, uint8, uint:
		return true
	case Number:
		return true
	}
	return false
}

func isString(value interface{}) bool {
	switch value.(type) {
	case string:
		return true
	}
	return false
}

func isFloat(value interface{}) bool {
	switch value.(type) {
	case float64, float32:
		return true
	case Number:
		return true
	}
	return false
}

func isNil(value interface{}) bool {
	return value == nil
}

var (
	_true  = Any(true)
	_false = Any(false)
)

/*
Converting a boolean to an interface{} requires an allocation.
We can use interned bool to avoid this cost.
*/
func boolInterface(b bool) Any {
	if b {
		return _true
	}
	return _false
}

func indirect(a interface{}) interface{} {
	if a == nil {
		return nil
	}
	if t := reflect.TypeOf(a); t.Kind() != reflect.Ptr {
		return a
	}
	v := reflect.ValueOf(a)
	for v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}

const (
	TagJson      = "json"
	TagIgnore    = "-"
	TagOmitEmpty = "omitempty"
	EmptyStr     = ""
	CommaStr     = ","
)

func getKey(tagStr string) (key string, ignore, omitempty bool) {

	if tagStr == EmptyStr {
		return "", false, false
	}

	key = strings.Split(tagStr, CommaStr)[0]
	if key == TagIgnore {
		ignore = true
		return
	}
	if strings.Contains(tagStr, TagOmitEmpty) {
		omitempty = true
	}

	return
}
