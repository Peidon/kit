/*
 * @Copyright @ Shopee
 * @Author: Peidong Xu
 * @Email: peidong.xu@shopee.com
 * @Date: 2023/11/22 11:52 AM
 * @Version: kit
 */

package valuate

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
	Array
	Struct
)

func GetInt(v Value) int64 {
	return v.Integer
}

func GetString(v Value) string {
	return v.string
}

func GetBool(v Value) bool {
	return v.Integer == 1
}

type Comparable interface {
	Equal(other Comparable) bool
	Greater(other Comparable) bool
}
