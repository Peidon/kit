/*
 * @Copyright @ Shopee
 * @Author: Peidong Xu
 * @Email: peidong.xu@shopee.com
 * @Date: 2023/11/13 3:57 PM
 * @Version: kit
 */

package valuate

type OperatorSymbol int8

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

var (
	symbolMap = map[OperatorSymbol]string{
		EQ:       "==",
		NEQ:      "!=",
		LT:       "<",
		GT:       ">",
		LTE:      "<=",
		GTE:      ">=",
		AND:      "&&",
		OR:       "||",
		PLUS:     "+",
		MINUS:    "subtract -",
		MULTIPLY: "multiply *",
		DIVIDE:   "divide /",
		MODULUS:  "modulus %",
		INDEX:    "array index []",
		ACCESS:   "accessor .",
		INVERT:   "!",
		NEGATE:   "negate -",
		NOOP:     "noop",

		LITERAL:    "literal",
		VALUE:      "value",
		FUNCTIONAL: "functional",
	}
)

// OperatorArgument /*
/*
Operator Argument Index
*/
type OperatorArgument int8

const (
	// binary arguments index
	left OperatorArgument = iota
	right

	// unary arguments index
	unary OperatorArgument = 0
)

type OperatorType int8

const (
	unaryOp OperatorType = iota + 1
	binaryOp
)

type Any = interface{}

type TokenSymbol string

const (
	NIL    TokenSymbol = "nil"
	Bool   TokenSymbol = "bool"
	String TokenSymbol = "str"
	Int    TokenSymbol = "int"
	Float  TokenSymbol = "float"
)
