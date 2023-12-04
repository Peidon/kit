/*
 * @Copyright @ Shopee
 * @Author: Peidong Xu
 * @Email: peidong.xu@shopee.com
 * @Date: 2023/11/16 3:38 PM
 * @Version: kit
 */

package valuate

import (
	"reflect"
	"testing"
	"time"
)

func TestEvaluableExpression_Evaluate(t *testing.T) {
	testData := []struct {
		input  string
		want   interface{}
		params MapParameters

		hasError bool
	}{
		{
			input: "1+2.1",
			want:  3.1,
		},
		{
			input: "nil",
			want:  nil,
		},
		{
			input:    "a == nil",
			params:   MapParameters(map[string]Any{"abc": 3}),
			hasError: true,
		},
		{
			input:    "abc == nil",
			params:   MapParameters(map[string]Any{"abc": abc{}}),
			hasError: true,
		},
		{
			input:  "abc == nil",
			want:   true,
			params: MapParameters(map[string]Any{"abc": nil}),
		},
		{
			input: "true",
			want:  true,
		},
		{
			input: "false",
			want:  false,
		},
		{
			input: "3 * 3",
			want:  9,
		},
		{
			input: "4 / 2.0",
			want:  2.0,
		},
		{
			input: "5 / 2",
			want:  2,
		},
		{
			input: "5 % 2",
			want:  1,
		},
		{
			input: "5 > 2",
			want:  true,
		},
		{
			input: "5 <= 2",
			want:  false,
		},
		{
			input:  "abc - 2",
			want:   1,
			params: MapParameters(map[string]Any{"abc": 3}),
		},
		{
			input:    "abc && 2",
			params:   MapParameters(map[string]Any{"abc": 3}),
			hasError: true,
		},
		{
			input:  "abc > 2 && ${c} == true",
			want:   true,
			params: MapParameters(map[string]Any{"abc": 3, "c": true}),
		},
		{
			input:  "${ab.c} - 2.0",
			want:   1.0,
			params: MapParameters(map[string]Any{"ab.c": 3}),
		},
		{
			input: "[1,2,3]",
			want:  []interface{}{1, 2, 3},
		},
		{
			input: `["abc", "d"]`,
			want:  []interface{}{"abc", "d"},
		},
	}

	for _, td := range testData {
		expr, err := NewExpression(td.input)
		if err != nil {
			t.Error(err)
			return
		}
		got, er := expr.Evaluate(td.params)
		if er != nil && !td.hasError {
			t.Error(er)
			return
		}
		if er != nil {
			t.Log(er)
		}

		// testing array value
		if arr, ok := td.want.([]interface{}); ok {
			gotArr, ok := got.([]interface{})
			if !ok {
				t.Error("got type not []interface\n got type : ", reflect.TypeOf(got).String())
			}
			for i := range arr {
				a := AnyValue(arr[i])
				b := AnyValue(gotArr[i])
				if !a.Equal(b) {
					t.Error("Array Elem ", arr[i], " and ", gotArr[i], " not equal\n")
					continue
				}

			}

			continue
		}

		if got != td.want {
			t.Error("got: ", got, "\n got type", reflect.TypeOf(got).String(), "\nwant: ", td.want)
		}
	}
}

type abc struct{}

func TestAccess(t *testing.T) {

	ti := time.Now()

	s := ABC{
		a: 12,
		B: "abc",
		c: true,

		D: []int64{1, 2, 3},

		E: &ABC{
			a: 3,
			B: "efg",
		},

		f: 1234.2345,

		G: FG{
			J: ti,
			h: 10,
			g: []byte{'g'},
		},
	}

	testData := []struct {
		input    string
		want     interface{}
		params   MapParameters
		hasError bool
	}{
		{
			input:  "s.abc.b == efg",
			want:   true,
			params: MapParameters(map[string]Any{"s": s, "efg": s.E.B}),
		},
		{
			input:  "s.G.h",
			want:   uint64(10),
			params: MapParameters(map[string]Any{"s": s}),
		},
		{
			input:  "s.G.J == ti",
			want:   true,
			params: MapParameters(map[string]Any{"s": s, "ti": ti}),
		},
	}

	for _, td := range testData {
		expr, err := NewExpression(td.input)
		if err != nil {
			t.Error(err)
			return
		}
		got, er := expr.Evaluate(td.params)
		if er != nil && !td.hasError {
			t.Error(er)
			return
		}
		if er != nil {
			t.Log(er)
		}
		if got != td.want {
			t.Error("got: ", got, "\n got type", reflect.TypeOf(got).String(), "\nwant: ", td.want)
		}
	}
}
