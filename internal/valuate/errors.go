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
	operate    string
	valueToken string
}

func (ter TypeError) Error() string {
	return "Value '" + ter.valueToken + "' cannot be used with the operator '" + ter.operate + "'"
}
