/*
 * @Copyright @ Shopee
 * @Author: Peidong Xu
 * @Email: peidong.xu@shopee.com
 * @Date: 2023/11/13 3:55 PM
 * @Version: kit
 */

package valuate

import (
	"context"
	"errors"
	"reflect"
	"strconv"
	"strings"
)

type evaluationStage struct {
	symbol OperatorSymbol

	// sub stages
	depends []evaluationStage

	// the operation that will be used to evaluate this stage (such as adding [left] to [right] and return the result)
	operator evaluationOperator

	// operator type
	opType OperatorType

	// ensures values type are appropriate for this stage
	typeCheck stageCombinedTypeCheck
}

func (stage evaluationStage) argsTypeCheck(args ...Any) (err error) {
	if stage.typeCheck == nil {
		return
	}

	if err = stage.typeCheck(args...); err != nil {
		tp := "nil"
		for _, a := range args {

			if a == nil {
				break
			}

			reType := reflect.TypeOf(a)
			tp = reType.String()
			break
		}

		return TypeError{
			operateSym: stage.symbol,
			typeName:   tp,
		}
	}
	return
}

// type checking depends on Stage Operator
type stageCombinedTypeCheck func(arguments ...Any) error

// array operator check
func arrayTypeCheck(arguments ...Any) error {
	arrInf := arguments[left]
	idxInf := arguments[right]

	arrType := reflect.TypeOf(arrInf)
	if arrType.Kind() != reflect.Slice && arrType.Kind() != reflect.Array && !isInt(idxInf) {
		return ArgumentTypeError
	}

	return nil
}

// number operators check
func numberTypeCheck(arguments ...Any) error {
	for _, arg := range arguments {
		if !isNumber(arg) {
			return ArgumentTypeError
		}
	}
	return nil
}

// bool type check
func boolTypeCheck(arguments ...Any) error {
	for _, arg := range arguments {
		if !isBool(arg) {
			return ArgumentTypeError
		}
	}
	return nil
}

// isComparable
func isComparable(obj Any) bool {
	switch obj.(type) {
	case Comparable:
		return true
	}
	return false
}

// isModifier
func isModifier(obj Any) bool {
	switch obj.(type) {
	case Modifier:
		return true
	}
	return false
}

// comparable operators check
func comparableCheck(arguments ...Any) error {
	for _, arg := range arguments {
		if !isNumber(arg) &&
			!isString(arg) &&
			!isBool(arg) &&
			!isTime(arg) &&
			!isArray(arg) &&
			!isComparable(arg) &&
			!isModifier(arg) &&
			!isNil(arg) &&
			!isPtr(arg) {
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

func makeParameterStage(parameterName string) evaluationOperator {

	return func(parameters Parameters, arguments ...Any) (Any, error) {
		value, err := parameters.Get(parameterName)
		if err != nil {
			return nil, err
		}

		return value, nil
	}
}

// equalStage symbol ==
// @param arguments left right
func equalStage(p Parameters, arguments ...Any) (Any, error) {
	arguments = getAnyList(arguments)
	l, r := arguments[left], arguments[right]

	ctx := p.GetContext()

	if isModifier(l) {
		return l.(Modifier).Modify(ctx, "==", r)
	}
	if isModifier(r) {
		return r.(Modifier).Modify(ctx, "==", l)
	}

	if isNumber(l) && isNumber(r) {
		return toFloat64(l) == toFloat64(r), nil
	}
	if isString(l) && isString(r) {
		return l.(string) == r.(string), nil
	}
	if isBool(l) && isBool(r) {
		return l.(bool) == r.(bool), nil
	}

	if isTime(l) && isTime(r) {
		lv, rv := toTime(l), toTime(r)
		return lv.Sub(rv) == 0, nil
	}

	if isComparable(l) {
		return l.(Comparable).Equal(r), nil
	}
	if isComparable(r) {
		return r.(Comparable).Equal(l), nil
	}
	return boolInterface(reflect.DeepEqual(l, r)), nil
}

// notEqualStage symbol !=
// @param arguments left right
func notEqualStage(p Parameters, arguments ...Any) (Any, error) {
	arguments = getAnyList(arguments)
	l, r := arguments[left], arguments[right]

	ctx := p.GetContext()

	if isModifier(l) {
		return l.(Modifier).Modify(ctx, "!=", r)
	}
	if isModifier(r) {
		return r.(Modifier).Modify(ctx, "!=", l)
	}

	if isNumber(l) && isNumber(r) {
		return toFloat64(l) != toFloat64(r), nil
	}
	if isString(l) && isString(r) {
		return l.(string) != r.(string), nil
	}
	if isBool(l) && isBool(r) {
		return l.(bool) != r.(bool), nil
	}
	if isTime(l) && isTime(r) {
		lv, rv := toTime(l), toTime(r)
		return lv.Sub(rv) != 0, nil
	}

	if isComparable(l) {
		a := l.(Comparable)
		return !a.Equal(r), nil
	}
	if isComparable(r) {
		b := r.(Comparable)
		return !b.Equal(l), nil
	}
	return boolInterface(!reflect.DeepEqual(l, r)), nil
}

// gtStage symbol '>'
// greater than
func gtStage(p Parameters, arguments ...Any) (Any, error) {
	arguments = getAnyList(arguments)
	l, r := arguments[left], arguments[right]

	ctx := p.GetContext()

	if isModifier(l) {
		return l.(Modifier).Modify(ctx, ">", r)
	}
	if isModifier(r) {
		return r.(Modifier).Modify(ctx, "<", l)
	}

	if isNumber(l) && isNumber(r) {
		return toFloat64(l) > toFloat64(r), nil
	}
	if isTime(l) && isTime(r) {
		lv, rv := toTime(l), toTime(r)
		return lv.After(rv), nil
	}
	if isComparable(l) {
		a := l.(Comparable)
		return a.Greater(r), nil
	}
	if isComparable(r) {
		b := r.(Comparable)
		return !b.Greater(l) && !b.Equal(l), nil
	}
	return nil, ArgumentTypeError
}

// ltStage symbol '<'
// less than
func ltStage(p Parameters, arguments ...Any) (Any, error) {
	arguments = getAnyList(arguments)
	l, r := arguments[left], arguments[right]

	ctx := p.GetContext()

	if isModifier(l) {
		return l.(Modifier).Modify(ctx, "<", r)
	}
	if isModifier(r) {
		return r.(Modifier).Modify(ctx, ">", l)
	}

	if isNumber(l) && isNumber(r) {
		a := toFloat64(l)
		b := toFloat64(r)
		return a < b, nil
	}
	if isTime(l) && isTime(r) {
		lv, rv := toTime(l), toTime(r)
		return lv.Before(rv), nil
	}
	if isComparable(l) {
		a := l.(Comparable)
		return !a.Greater(r) && !a.Equal(r), nil
	}
	if isComparable(r) {
		b := r.(Comparable)
		return b.Greater(l), nil
	}
	return nil, ArgumentTypeError
}

// gteStage symbol '>='
// greater or equal
func gteStage(p Parameters, arguments ...Any) (Any, error) {
	arguments = getAnyList(arguments)
	l, r := arguments[left], arguments[right]

	ctx := p.GetContext()

	if isModifier(l) {
		return l.(Modifier).Modify(ctx, ">=", r)
	}
	if isModifier(r) {
		return r.(Modifier).Modify(ctx, "<=", l)
	}

	if isNumber(l) && isNumber(r) {
		return toFloat64(l) >= toFloat64(r), nil
	}
	if isTime(l) && isTime(r) {
		lv, rv := toTime(l), toTime(r)
		return !lv.Before(rv), nil
	}
	if isComparable(l) {
		a := l.(Comparable)
		return a.Greater(r) || a.Equal(r), nil
	}
	if isComparable(r) {
		b := r.(Comparable)
		return !b.Greater(l), nil
	}
	return nil, ArgumentTypeError
}

// lteStage symbol '<='
// less or equal
func lteStage(p Parameters, arguments ...Any) (Any, error) {
	arguments = getAnyList(arguments)
	l, r := arguments[left], arguments[right]

	ctx := p.GetContext()

	if isModifier(l) {
		return l.(Modifier).Modify(ctx, "<=", r)
	}
	if isModifier(r) {
		return r.(Modifier).Modify(ctx, ">=", l)
	}

	if isNumber(l) && isNumber(r) {
		return toFloat64(l) <= toFloat64(r), nil
	}
	if isTime(l) && isTime(r) {
		lv, rv := toTime(l), toTime(r)
		return !lv.After(rv), nil
	}
	if isComparable(l) {
		a := l.(Comparable)
		return !a.Greater(r), nil
	}
	if isComparable(r) {
		b := r.(Comparable)
		return b.Greater(l) || b.Equal(l), nil
	}
	return nil, ArgumentTypeError
}

// andStage symbol '&&'
func andStage(_ Parameters, arguments ...Any) (Any, error) {
	arguments = getAnyList(arguments)
	l, r := arguments[left], arguments[right]
	if isBool(l) && isBool(r) {
		return l.(bool) && r.(bool), nil
	}
	return nil, ArgumentTypeError
}

// orStage symbol '||'
func orStage(_ Parameters, arguments ...Any) (Any, error) {
	arguments = getAnyList(arguments)
	l, r := arguments[left], arguments[right]
	if isBool(l) && isBool(r) {
		return l.(bool) || r.(bool), nil
	}
	return nil, ArgumentTypeError
}

// modify op '+'
func addStage(p Parameters, arguments ...Any) (Any, error) {
	arguments = getAnyList(arguments)
	l, r := arguments[left], arguments[right]

	ctx := p.GetContext()

	if isModifier(l) {
		return l.(Modifier).Modify(ctx, "+", r)
	}
	if isModifier(r) {
		return r.(Modifier).Modify(ctx, "+", l)
	}

	if isFloat(l) || isFloat(r) {
		a := toFloat64(l)
		b := toFloat64(r)
		return a + b, nil
	}
	return toInt(l) + toInt(r), nil
}

// modify op '/'
func divideStage(p Parameters, arguments ...Any) (Any, error) {
	arguments = getAnyList(arguments)
	l, r := arguments[left], arguments[right]

	ctx := p.GetContext()

	if isModifier(l) {
		return l.(Modifier).Modify(ctx, "/", r)
	}
	if isModifier(r) {
		return r.(Modifier).Modify(ctx, "1/", l)
	}

	if isInt(l) && isInt(r) {
		rv := toInt(r)
		if rv == 0 {
			return nil, DivideZeroError
		}
		return toInt(l) / rv, nil
	}

	rv := toFloat64(r)
	if rv == 0.0 {
		return nil, DivideZeroError
	}

	return toFloat64(l) / rv, nil
}

// modify op '*'
func multipleStage(p Parameters, arguments ...Any) (Any, error) {
	arguments = getAnyList(arguments)
	l, r := arguments[left], arguments[right]

	ctx := p.GetContext()

	if isModifier(l) {
		return l.(Modifier).Modify(ctx, "*", r)
	}
	if isModifier(r) {
		return r.(Modifier).Modify(ctx, "*", l)
	}

	if isInt(l) && isInt(r) {
		return toInt(l) * toInt(r), nil
	}

	return toFloat64(l) * toFloat64(r), nil
}

// modify op '%'
func modulusStage(_ Parameters, arguments ...Any) (Any, error) {
	arguments = getAnyList(arguments)
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
func subtractStage(p Parameters, arguments ...Any) (Any, error) {
	arguments = getAnyList(arguments)
	l, r := arguments[left], arguments[right]

	ctx := p.GetContext()

	if isModifier(l) {
		return l.(Modifier).Modify(ctx, "-", r)
	}
	if isModifier(r) {
		return r.(Modifier).Modify(ctx, "1-", l)
	}

	if isFloat(l) || isFloat(r) {
		return toFloat64(l) - toFloat64(r), nil
	}
	return toInt(l) - toInt(r), nil
}

// negate op '-'
func negateStage(_ Parameters, arguments ...Any) (Any, error) {
	arguments = getAnyList(arguments)
	u := arguments[unary]
	if isInt(u) {
		return -toInt(u), nil
	}
	if isFloat(u) || isNumber(u) {
		return -toFloat64(u), nil
	}
	return nil, ArgumentTypeError
}

// invert op '!'
func invertStage(_ Parameters, arguments ...Any) (Any, error) {
	u := getAny(arguments[unary])
	if isBool(u) {
		return !u.(bool), nil
	}
	return nil, ArgumentTypeError
}

// functional arguments list
func exprListStage(_ Parameters, arguments ...Any) (Any, error) {
	return getAnyList(arguments), nil
}

// index :  '[' expression ']'
func indexStage(_ Parameters, arguments ...Any) (Any, error) {
	arg := getAny(arguments[unary])
	if isInt(arg) {
		return toInt(arg), nil
	}
	return nil, ArgumentTypeError
}

// array value literal
// [expr, ... , expr]
func arrayValueStage(_ Parameters, arguments ...Any) (Any, error) {
	return getAnyList(arguments), nil
}

// array index
// primary_expr '[' index ']'
func arrayIndexStage(_ Parameters, arguments ...Any) (Any, error) {
	arguments = getAnyList(arguments)
	arrInf := arguments[left]  // primary expr
	idxInf := arguments[right] // index

	arrVal := AnyValue(arrInf)
	idx := AnyValue(idxInf)

	arr := arrVal.GetArray()
	i := int(idx.GetInt())

	if i < 0 || i >= len(arr) {
		return nil, ArrayIndexError{
			Index:    i,
			ArrayLen: len(arr),
		}
	}

	return arr[i], nil
}

// functional
func makeFuncStage(funcName string, fs map[string]ExpressionFunction) evaluationOperator {
	return func(parameters Parameters, arguments ...Any) (ret Any, err error) {

		lis := arguments[unary]

		args, ok := lis.([]interface{})
		if !ok {
			argVal := AnyValue(lis)
			vars := argVal.GetArray()

			for _, v := range vars {
				args = append(args, v.Get())
			}
		}

		var fn ExpressionFunction

		// customize
		if fn, ok = fs[funcName]; !ok {

			// builtin
			fn, ok = builtin[funcName]
		}
		if !ok {
			return nil, errors.New("function " + funcName + " not exists")
		}

		ctx := context.Background()
		ctxPart, existsErr := parameters.Get(contextParam)
		if c, valid := ctxPart.(context.Context); valid && existsErr == nil {
			ctx = c
		}

		ret, err = fn(ctx, args...)
		switch err {
		case FnArgTypeError:
			ts := argTypes(args...)
			return nil, errors.New("function '" + funcName + "' arguments type error : " + strings.Join(ts, ","))
		case FnArgsNumberError:
			num := len(args)
			return nil, errors.New("function '" + funcName + "' arguments number error, number = " + strconv.Itoa(num))
		case nil:
			return
		default:
			return nil, errors.New("function '" + funcName + "' error, " + err.Error())
		}
	}
}

// op symbol '.'
func makeAccessorStage(field string) evaluationOperator {
	return func(parameters Parameters, arguments ...Any) (Any, error) {
		if len(arguments) == 0 {
			return nil, TypeError{
				operateSym: ACCESS,
				typeName:   "void",
			}
		}
		obj := arguments[0]
		v := AnyValue(obj)

		structVal := v.GetStruct()
		if structVal == nil {
			return nil, TypeError{
				operateSym: ACCESS,
				typeName:   v.GetString(),
			}
		}

		structTyp := v.GetString()
		v = structVal.Field(field)
		if v.Type == VoidType {
			// 字段不存在
			return nil, AccessError{
				structType: structTyp,
				fieldName:  field,
			}
		}

		return v, nil
	}
}
