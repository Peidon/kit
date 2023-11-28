/*
 * @Copyright @ Shopee
 * @Author: Peidong Xu
 * @Email: peidong.xu@shopee.com
 * @Date: 2023/11/28 10:41 AM
 * @Version: kit
 */

package valuate

import "reflect"

func toInt(v interface{}) int {
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
	}
	return -1
}

func toFloat64(v interface{}) float64 {
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
	case float32:
		return float64(s)
	case float64:
		return s
	}
	return -1
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
	case int64, int32, int16, int8, int:
		return true
	}
	return false
}

func isInt(value interface{}) bool {
	switch value.(type) {
	case int64, int32, int16, int8, int:
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
