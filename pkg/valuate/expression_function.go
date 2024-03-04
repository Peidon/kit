/*
 * @Copyright @ Shopee
 * @Author: Peidong Xu
 * @Email: peidong.xu@shopee.com
 * @Date: 2023/11/10 5:29 PM
 * @Version: kit
 */

package valuate

import (
	"context"
	"encoding/json"
	"reflect"
	"time"
)

type ExpressionFunction func(ctx context.Context, arguments ...Any) (Any, error)

var (
	// builtin functions
	builtin = map[string]ExpressionFunction{
		"json_marshal": func(_ context.Context, args ...interface{}) (interface{}, error) {
			if len(args) == 0 || len(args) > 1 {
				return nil, FnArgsNumberError
			}
			b, e := json.Marshal(args[0])
			if e != nil {
				return nil, e
			}
			return string(b), nil
		},
		"json_unmarshal": func(_ context.Context, args ...interface{}) (interface{}, error) {
			if len(args) != 2 {
				return nil, FnArgsNumberError
			}
			a := args[0]
			v, ok := a.(string)
			if !ok {
				return nil, FnArgTypeError
			}
			b := []byte(v)
			err := json.Unmarshal(b, args[1])
			return args[1], err
		},
		"unix_timestamp": func(_ context.Context, args ...interface{}) (interface{}, error) {
			if len(args) != 1 {
				return nil, FnArgsNumberError
			}
			a := args[0]
			switch val := a.(type) {
			case time.Time:
				return val.Unix(), nil
			case *time.Time:
				return val.Unix(), nil
			}

			return nil, FnArgTypeError
		},
		"len": length,
	}
)

func length(_ context.Context, args ...Any) (interface{}, error) {
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
