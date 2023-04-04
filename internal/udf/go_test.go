package udf

import (
	"kit/pkg/slice"
	"reflect"
	"testing"
)

func TestExecute(t *testing.T) {
	ss := []string{"a", "b"}
	s := "a"

	c := slice.Contains(ss, s)
	t.Log(c)

	fn := reflect.ValueOf(slice.Contains)

	argValues := make([]reflect.Value, 2)
	for i, arg := range []interface{}{ss, s} {
		argValues[i] = reflect.ValueOf(arg)
	}

	vs := fn.Call(argValues)
	for _, v := range vs {
		t.Log(v.Interface())
	}
}
