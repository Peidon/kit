/*
 * @Copyright @ Shopee
 * @Author: Peidong Xu
 * @Email: peidong.xu@shopee.com
 * @Date: 2023/11/13 3:57 PM
 * @Version: kit
 */

package valuate

type OperatorSymbol int

const (
	NOOP OperatorSymbol = iota
	LITERAL
	VALUE
	EQ
	NEQ
	GT
	LT
	GTE
	LTE

	AND
	OR

	PLUS
	MINUS
	MULTIPLY
	DIVIDE
	MODULUS

	NEGATE
	INVERT

	INDEX
	FUNCTIONAL
	ACCESS
)

// OperatorArgument /*
/*
Operator Argument Index
*/
type OperatorArgument int

const (
	// binary arguments index
	left OperatorArgument = iota
	right

	// unary arguments index
	unaryIndex = 0
)

type OperatorType int

const (
	unary OperatorType = iota + 1
	binary
)

type Any = interface{}
