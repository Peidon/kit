/*
 * @Copyright @ Shopee
 * @Author: Peidong Xu
 * @Email: peidong.xu@shopee.com
 * @Date: 2023/11/10 5:29 PM
 * @Version: kit
 */

package valuate

type ExpressionFunction func(arguments ...Any) (Any, error)

type Value struct {
	Type      ValueType
	Integer   int64
	string    string
	Interface interface{}
}

type ValueType int

const (
	Unknown ValueType = iota
	IntegerType
	StringType
	BoolType
)
