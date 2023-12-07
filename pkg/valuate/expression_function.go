/*
 * @Copyright @ Shopee
 * @Author: Peidong Xu
 * @Email: peidong.xu@shopee.com
 * @Date: 2023/11/10 5:29 PM
 * @Version: kit
 */

package valuate

import "reflect"

type ExpressionFunction func(arguments ...Any) (Any, error)

var (
	// builtin functions
	builtin = map[string]ExpressionFunction{
		"len": length,
	}
)

func length(args ...Any) (interface{}, error) {
	if len(args) == 0 || len(args) > 1 {
		return nil, FnArgsNumberError
	}
	arr := args[0]
	v := reflect.ValueOf(arr)
	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		return v.Len(), nil
	case reflect.String:
		return len(v.String()), nil
	default:
		return nil, FnArgTypeError
	}
}

func argTypes(args ...Any) []string {
	var ts []string
	for _, a := range args {
		t := reflect.TypeOf(a)
		ts = append(ts, t.String())
	}
	return ts
}
