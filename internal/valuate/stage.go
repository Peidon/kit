/*
 * @Copyright @ Shopee
 * @Author: Peidong Xu
 * @Email: peidong.xu@shopee.com
 * @Date: 2023/11/13 3:55 PM
 * @Version: kit
 */

package valuate

import (
	"reflect"
)

type evaluationStage struct {
	symbol OperatorSymbol

	// 这里基于ANTLR4语法树后序遍历
	// 否则不能这么设计, 因为不确定顺序则不能确定符号左右的值
	// Operator 中的arguments 同理
	depends []evaluationStage

	// the operation that will be used to evaluate this stage (such as adding [left] to [right] and return the result)
	operator evaluationOperator

	// operator type
	opType OperatorType

	// if specified, will override whatever is used in "leftTypeCheck" and "rightTypeCheck".
	// primarily used for specific operators that don't care which side a given type is on, but still requires one side to be of a given type
	// (like string concat)
	typeCheck stageCombinedTypeCheck

	// regardless of which type check is used, this string format will be used as the error message for type errors
	typeErrorFormat string
}

// type checking depends on Stage Operator
type stageCombinedTypeCheck func(arguments ...Any) bool

// Operator in Stage
type evaluationOperator func(parameters Parameters, arguments ...Any) (Any, error)

// equalStage symbol ==
// @param arguments left right
func equalStage(_ Parameters, arguments ...Any) (Any, error) {
	l := arguments[left]
	r := arguments[right]
	return boolInterface(reflect.DeepEqual(l, r)), nil
}

// notEqualStage symbol !=
// @param arguments left right
func notEqualStage(_ Parameters, arguments ...Any) (Any, error) {
	l := arguments[left]
	r := arguments[right]
	return boolInterface(!reflect.DeepEqual(l, r)), nil
}

// modify op '+'
func addStage(_ Parameters, arguments ...Any) (Any, error) {
	l := arguments[left]
	r := arguments[right]

	if isString(l) && isString(r) {
		return addString(l.(string), r.(string))
	}

	if isNumber(l) && isNumber(r) {
		return addNumber(l, r)
	}
	return nil, ArgumentTypeError
}

// modify op '/'
func divideStage(_ Parameters, arguments ...Any) (Any, error) {
	l := arguments[left]
	r := arguments[right]

	if !isNumber(l) || !isNumber(r) {
		return nil, ArgumentTypeError
	}

	lv := toFloat64(l)
	rv := toFloat64(r)

	if rv == 0.0 {
		return nil, DivideZeroError
	}

	return lv / rv, nil
}

// modify op '*'
func multipleStage(_ Parameters, arguments ...Any) (Any, error) {
	l := arguments[left]
	r := arguments[right]

	if !isNumber(l) || !isNumber(r) {
		return nil, ArgumentTypeError
	}

	lv := toFloat64(l)
	rv := toFloat64(r)

	return lv * rv, nil
}

// modify op '%'
func modulusStage(_ Parameters, arguments ...Any) (Any, error) {
	l := arguments[left]
	r := arguments[right]

	if !isInt(l) || !isInt(r) {
		return nil, ArgumentTypeError
	}

	lv := toInt64(l)
	rv := toInt64(r)

	if rv == 0 {
		return nil, DivideZeroError
	}

	return lv % rv, nil
}

// modify op '-'
func subtractStage(_ Parameters, arguments ...Any) (Any, error) {
	l := arguments[left]
	r := arguments[right]

	if isFloat(l) || isFloat(r) {
		return l.(float64) - r.(float64), nil
	}
	return l.(int64) - r.(int64), nil
}

func addNumber(left Any, right Any) (Any, error) {
	if isFloat(left) || isFloat(right) {
		return left.(float64) + right.(float64), nil
	}

	return left.(int64) + right.(int64), nil
}

func addString(left, right string) (string, error) {
	return left + right, nil
}

func isString(value interface{}) bool {
	switch value.(type) {
	case string:
		return true
	}
	return false
}

func toInt64(v interface{}) int64 {
	v = indirect(v)

	switch s := v.(type) {
	case int64:
		return s
	case int32:
		return int64(s)
	case int16:
		return int64(s)
	case int8:
		return int64(s)
	case int:
		return int64(s)
	}
	return 0
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
	return 0.0
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

func isFloat(value interface{}) bool {
	switch value.(type) {
	case float64, float32:
		return true
	}
	return false
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
