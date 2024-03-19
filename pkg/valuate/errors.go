/*
 * @Copyright @ Shopee
 * @Author: Peidong Xu
 * @Email: peidong.xu@shopee.com
 * @Date: 2023/11/14 3:56 PM
 * @Version: kit
 */

package valuate

import (
	"errors"
	"strconv"
	"strings"
)

var (
	ArgumentTypeError = errors.New("argument type error")
	ArgumentsError    = errors.New("arguments number error")
	DivideZeroError   = errors.New("divide zero")
	NoNeedEvaluate    = errors.New("no need evaluate")
	PlanStageError    = errors.New("plan stage error")
	FnArgsNumberError = errors.New("function arguments number error")
	FnArgTypeError    = errors.New("function argument type error")
)

// StageError track errors when plan stages
type StageError struct {
	tokens []string
}

func (the *StageError) Error() string {
	return "plan stages failed: '" + strings.Join(the.tokens, "' -> '") + "'"
}

func (the *StageError) Append(token string) {
	the.tokens = append(the.tokens, token)
}

// TypeError operator arguments type error
type TypeError struct {
	operateSym OperatorSymbol
	typeName   string
}

func (ter TypeError) Error() string {
	return "Type '" + ter.typeName + "' cannot be used with the operator '" + symbolMap[ter.operateSym] + "'"
}

type AccessError struct {
	structType string
	fieldName  string
}

func (aer AccessError) Error() string {
	return "Field " + aer.fieldName + " not exists in " + aer.structType
}

type ArrayIndexError struct {
	ArrayLen int
	Index    int
}

func (err ArrayIndexError) Error() string {
	return "index " + strconv.Itoa(err.Index) + " out of range, array length is " + strconv.Itoa(err.ArrayLen)
}

type ParameterNotFound struct {
	paramName string
}

func (pa ParameterNotFound) Error() string {
	return "No parameter '" + pa.paramName + "' found."
}

func (pa ParameterNotFound) KeyName() string {
	return pa.paramName
}

type ErrorStrategy func(err error) (interface{}, error)

func defaultStrategy(err error) (interface{}, error) {
	return nil, err
}
