/*
 * @Copyright @ Shopee
 * @Author: Peidong Xu
 * @Email: peidong.xu@shopee.com
 * @Date: 2023/11/16 3:38 PM
 * @Version: kit
 */

package valuate

import "testing"

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
			t.Error("got: ", got, "\nwant: ", td.want)
		}
	}
}

type abc struct{}

func TestAccess(t *testing.T) {
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

		G: FG{},
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
			t.Error("got: ", got, "\nwant: ", td.want)
		}
	}
}
