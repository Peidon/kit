/*
 * @Copyright @ Shopee
 * @Author: Peidong Xu
 * @Email: peidong.xu@shopee.com
 * @Date: 2023/11/13 3:55 PM
 * @Version: kit
 */

package valuate

import (
	"errors"
	"reflect"
	"strconv"
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
}

// type checking depends on Stage Operator
type stageCombinedTypeCheck func(arguments ...Any) error

// number operators check
func numberTypeCheck(arguments ...Any) error {
	for _, arg := range arguments {
		if !isNumber(arg) {
			return ArgumentTypeError
		}
	}
	return nil
}

// evaluationOperator
// Operator in Stage
type evaluationOperator func(parameters Parameters, arguments ...Any) (Any, error)

func makeLiteralStage(literal string, tp TokenSymbol) evaluationOperator {
	return func(parameters Parameters, arguments ...Any) (Any, error) {
		switch tp {
		case NIL:
			return nil, nil
		case Bool:
			return strconv.ParseBool(literal)
		case Int:
			return strconv.ParseInt(literal, 10, 64)
		case Float:
			return strconv.ParseFloat(literal, 64)
		case String:
			return strconv.Unquote(literal)
		}

		return nil, errors.New("make literal '" + literal + "' stage error, token type=" + string(tp))
	}
}

// equalStage symbol ==
// @param arguments left right
func equalStage(_ Parameters, arguments ...Any) (Any, error) {
	l, r := arguments[left], arguments[right]
	return boolInterface(reflect.DeepEqual(l, r)), nil
}

// notEqualStage symbol !=
// @param arguments left right
func notEqualStage(_ Parameters, arguments ...Any) (Any, error) {
	l, r := arguments[left], arguments[right]
	return boolInterface(!reflect.DeepEqual(l, r)), nil
}

// modify op '+'
func addStage(_ Parameters, arguments ...Any) (Any, error) {
	l, r := arguments[left], arguments[right]

	if isFloat(l) || isFloat(r) {
		return toFloat64(l) + toFloat64(r), nil
	}
	return toInt(l) + toInt(r), nil
}

// modify op '/'
func divideStage(_ Parameters, arguments ...Any) (Any, error) {
	l, r := arguments[left], arguments[right]

	if isInt(l) && isInt(r) {
		return toInt(l) / toInt(r), nil
	}

	rv := toFloat64(r)
	if rv == 0.0 {
		return nil, DivideZeroError
	}

	return toFloat64(l) / rv, nil
}

// modify op '*'
func multipleStage(_ Parameters, arguments ...Any) (Any, error) {
	l, r := arguments[left], arguments[right]

	if isInt(l) && isInt(r) {
		return toInt(l) * toInt(r), nil
	}

	return toFloat64(l) * toFloat64(r), nil
}

// modify op '%'
func modulusStage(_ Parameters, arguments ...Any) (Any, error) {
	l, r := arguments[left], arguments[right]

	if !isInt(l) || !isInt(r) {
		return nil, ArgumentTypeError
	}

	lv := toInt(l)
	rv := toInt(r)

	if rv == 0 {
		return nil, DivideZeroError
	}

	return lv % rv, nil
}

// modify op '-'
func subtractStage(_ Parameters, arguments ...Any) (Any, error) {
	l, r := arguments[left], arguments[right]

	if isFloat(l) || isFloat(r) {
		return l.(float64) - r.(float64), nil
	}
	return l.(int64) - r.(int64), nil
}

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
