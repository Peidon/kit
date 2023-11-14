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

	children []evaluationStage

	// the operation that will be used to evaluate this stage (such as adding [left] to [right] and return the result)
	operator evaluationOperator

	// ensures that both left and right values are appropriate for this stage. Returns an error if they aren't operable.
	leftTypeCheck  stageTypeCheck
	rightTypeCheck stageTypeCheck

	// if specified, will override whatever is used in "leftTypeCheck" and "rightTypeCheck".
	// primarily used for specific operators that don't care which side a given type is on, but still requires one side to be of a given type
	// (like string concat)
	typeCheck stageCombinedTypeCheck

	// regardless of which type check is used, this string format will be used as the error message for type errors
	typeErrorFormat string
}

type evaluationOperator func(parameters Parameters, arguments ...Any) (Any, error)

type stageTypeCheck func(value Any) bool
type stageCombinedTypeCheck func(left Any, right Any) bool

func equalStage(_ Parameters, arguments ...Any) (Any, error) {
	l := arguments[left]
	r := arguments[right]
	return boolInterface(reflect.DeepEqual(l, r)), nil
}

func notEqualStage(_ Parameters, arguments ...Any) (Any, error) {
	l := arguments[left]
	r := arguments[right]
	return boolInterface(!reflect.DeepEqual(l, r)), nil
}

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

func subtractStage(_ Parameters, arguments ...Any) (interface{}, error) {
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
