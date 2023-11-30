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
	"strings"
)

var (
	ArgumentTypeError = errors.New("argument type error")
	ArgumentsError    = errors.New("arguments number error")
	DivideZeroError   = errors.New("divide zero")
	NoNeedEvaluate    = errors.New("no need evaluate")
	PlanStageError    = errors.New("plan stage error")
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
