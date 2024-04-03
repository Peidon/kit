/*
 * @Copyright @ Shopee
 * @Author: Peidong Xu
 * @Email: peidong.xu@shopee.com
 * @Date: 2023/11/16 3:38 PM
 * @Version: kit
 */

package valuate

import (
	"context"
	"encoding/json"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestEvaluableExpression_Evaluate(t *testing.T) {
	testData := []struct {
		input  string
		want   interface{}
		params MapParameters

		hasError bool

		OmitError bool
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
			input:  "!a",
			params: MapParameters(map[string]Any{"a": false}),
			want:   true,
		},
		{
			input:  "-a",
			params: MapParameters(map[string]Any{"a": 1}),
			want:   -1,
		},
		{
			input:     "a == 0",
			params:    MapParameters(map[string]Any{"abc": 3}),
			want:      true, // error has been caught
			OmitError: true,
		},
		{
			input:    "a == nil",
			params:   MapParameters(map[string]Any{"abc": 3}),
			hasError: true,
		},
		{
			input:  "abc != nil",
			params: MapParameters(map[string]Any{"abc": &abc{}}),
			want:   true,
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
			input: "[]",
			want:  []interface{}{},
		},
		{
			input: "[1,2,3]",
			want:  []interface{}{1, 2, 3},
		},
		{
			input: `["abc", "d"]`,
			want:  []interface{}{"abc", "d"},
		},
		{
			input:  "a[1]",
			want:   int64(2),
			params: MapParameters(map[string]Any{"a": []int{1, 2, 3}}),
		},
		{
			input:  "a[idx]",
			want:   int64(2),
			params: MapParameters(map[string]Any{"a": []int{1, 2, 3}, "idx": 1}),
		},
	}

	for _, td := range testData {
		expr, err := Expression(td.input)
		if err != nil {
			t.Error(err)
			return
		}
		if td.OmitError {
			expr.CatchError(omitError)
		}
		got, er := expr.Evaluate(td.params)
		if er != nil && !td.hasError {
			t.Error(er)
			return
		}
		if er != nil {
			t.Log("[info]", er)
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

func omitError(_ error) (interface{}, error) {
	return 0, nil
}

type param struct {
	IntValue int
	StrValue string
	BooValue bool
	IsEmpty  bool
}

func (p *param) Equal(other interface{}) bool {
	if o, ok := other.(*param); ok {
		return p.IntValue == o.IntValue && p.StrValue == o.StrValue && p.BooValue == o.BooValue
	}
	if o, ok := other.(int); ok {
		return p.IntValue == o
	}
	return false
}

func (p *param) Greater(other interface{}) bool {
	if o, ok := other.(*param); ok {
		return p.IntValue > o.IntValue
	}
	if o, ok := other.(int); ok {
		return p.IntValue > o
	}
	return false
}

func (p *param) ToFloat() float64 {
	return float64(p.IntValue)
}

type emptyParam struct{}

func (empty emptyParam) Modify(_ context.Context, op string, _ interface{}) (interface{}, error) {
	if op == "+" {
		return 0, nil
	}
	if op == "==" {
		return false, nil
	}

	return nil, nil
}

func emptyError(err error) (interface{}, error) {
	switch err.(type) {
	case ParameterNotFound:
		return emptyParam{}, nil
	}
	return nil, err
}

func TestOverwrite(t *testing.T) {

	testData := []struct {
		input  string
		want   interface{}
		params MapParameters
	}{
		{
			input: "a < b",
			want:  true,
			params: MapParameters{
				"a": &param{
					IntValue: 1,
				},
				"b": &param{
					IntValue: 2,
				},
			},
		},
		{
			input: "a < b",
			want:  true,
			params: MapParameters{
				"a": &param{
					IntValue: 1,
				},
				"b": 2,
			},
		},
		{
			input: "a == b",
			want:  true,
			params: MapParameters{
				"a": &param{
					StrValue: "abc",
				},
				"b": &param{
					StrValue: "abc",
				},
			},
		},
		{
			input: "a > b",
			want:  true,
			params: MapParameters{
				"a": 3,
				"b": &param{
					IntValue: 2,
				},
			},
		},
		{
			input: "a + b",
			want:  5.0,
			params: MapParameters{
				"a": 3,
				"b": &param{
					IntValue: 2,
				},
			},
		},
		{
			input: "a - b",
			want:  1.0,
			params: MapParameters{
				"a": 3,
				"b": &param{
					IntValue: 2,
				},
			},
		},
		{
			input: "a + b",
			want:  0, // catch empty error
			params: MapParameters{
				"b": &param{
					IntValue: 2,
				},
			},
		},
		{
			input: "a * b",
			want:  6,
			params: MapParameters{
				"a": &param{
					IntValue: 3,
				},
				"b": &param{
					IntValue: 2,
				},
			},
		},
	}

	for _, td := range testData {
		expr, err := Expression(td.input)
		if err != nil {
			t.Error(err)
			return
		}

		got, er := expr.CatchError(emptyError).Evaluate(td.params)
		if er != nil {
			t.Log("[info] ", er)
		}

		if got != td.want {
			t.Error("eval failed: <"+td.input+"> got: ", got, "\n"+
				"got type", reflect.TypeOf(got).String(), "\n"+
				"want: ", td.want)
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

		G: GG{
			h: 10,
			I: []byte{'g'},
			J: ti,
		},

		GgList: []*GG{
			{
				h: 1,
				I: []byte{'i'},
				J: ti,
			},
			{
				h: 2,
				I: []byte{'h', 'i'},
			},
		},
	}

	stu := struct {
		A json.Number `json:"field_a"`
		B string
		C float64
	}{
		A: json.Number("567"),
		B: "abc",
		C: 2.5,
	}
	mma := map[string]interface{}{
		"big_number":  json.Number("1234"),
		"some_struct": stu,
	}

	fs := map[string]ExpressionFunction{
		"context_value": func(c context.Context, args ...interface{}) (interface{}, error) {
			if len(args) == 0 || len(args) > 1 {
				return nil, FnArgsNumberError
			}
			return c.Value(args[0]), nil
		},
		"json_marshal": func(_ context.Context, args ...interface{}) (interface{}, error) {
			if len(args) == 0 || len(args) > 1 {
				return nil, FnArgsNumberError
			}
			return json.Marshal(args[0])
		},
		"bytes_to_string": func(_ context.Context, args ...interface{}) (interface{}, error) {
			if len(args) == 0 || len(args) > 1 {
				return nil, FnArgsNumberError
			}
			a := args[0]
			v, ok := a.([]byte)
			if !ok {
				return nil, FnArgTypeError
			}

			return string(v), nil
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
		"time_stamp": func(_ context.Context, args ...interface{}) (interface{}, error) {
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
	}

	testData := []struct {
		input    string
		want     interface{}
		params   MapParameters
		hasError bool

		functions map[string]ExpressionFunction
	}{
		{
			input:  "m.some_struct",
			params: map[string]Any{"m": mma},
			want:   stu,
		},
		{
			input:  "m.big_number",
			params: map[string]Any{"m": mma},
			want:   json.Number("1234"),
		},
		{
			input:  "!in(${abc}, \"ab\")",
			params: MapParameters(map[string]Any{"abc": "abc"}),
			want:   false,
			functions: map[string]ExpressionFunction{
				"in": func(ctx context.Context, args ...Any) (Any, error) {
					if len(args) != 2 {
						return nil, FnArgsNumberError
					}
					a := args[0]
					b := args[1]
					v, ok := a.(string)
					if !ok {
						return nil, FnArgTypeError
					}
					k, pass := b.(string)
					if !pass {
						return nil, FnArgTypeError
					}

					return strings.Contains(v, k), nil
				},
			},
		},
		{
			input:  "s.abc.b == efg",
			want:   true,
			params: MapParameters(map[string]Any{"s": s, "efg": s.E.B}),
		},
		{
			input:  "s.abc.b == \"ef\" && s.ag == 10",
			want:   false,
			params: MapParameters(map[string]Any{"s": s, "efg": s.E.B}),
		},
		{
			input:  "s.abc.b == efg || s.ag == 10",
			want:   true,
			params: MapParameters(map[string]Any{"s": s, "efg": s.E.B}),
		},
		{
			input:  "s.G.h",
			want:   uint64(10),
			params: MapParameters(map[string]Any{"s": s}),
		},
		{
			input:  "time_stamp(s.G.J) == time_stamp(ti)",
			want:   true,
			params: MapParameters(map[string]Any{"s": s, "ti": ti}),

			functions: fs,
		},
		{
			input:  "s.GgList[0].I",
			want:   []byte{'i'},
			params: MapParameters(map[string]Any{"s": s}),
		},
		{
			input:  "len(s.GgList) == 2",
			want:   true,
			params: MapParameters(map[string]Any{"s": s}),
		},
		{
			input:  "len(s.GgList) + len(s.G.I) == n",
			want:   true,
			params: MapParameters(map[string]Any{"s": s, "n": len(s.GgList) + len(s.G.I)}),
		},
		{
			input:    "len(s)",
			hasError: true,
			params:   MapParameters(map[string]Any{"s": s}),
		},
		{
			input:     "context_value(s)",
			want:      "a",
			params:    MapParameters(map[string]Any{"s": "a"}),
			functions: fs,
		},
		{
			input:     "len(json_marshal(s.abc)) > 0",
			params:    MapParameters(map[string]Any{"s": s}),
			want:      true,
			functions: fs,
		},
		{
			input:     "bytes_to_string(json_marshal(s.abc))",
			params:    MapParameters(map[string]Any{"s": s}),
			want:      `{"G":{"I":null,"J":"0001-01-01T00:00:00Z"},"GgList":null,"a":3,"abc":null,"b":"efg","c":false,"d":null,"f":0}`,
			functions: fs,
		},
		{
			input: "json_unmarshal(a, b)",
			params: MapParameters(map[string]Any{
				"a": `{"G":{"I":[],"J":"0001-01-01T00:00:00Z","h":0},"GgList":[],"a":3,"abc":null,"b":"efg","c":false,"d":null,"f":0} `,
				"b": &s,
			}),
			want:      &s,
			functions: fs,
		},
	}

	for _, td := range testData {
		expr, err := ExpressionWithFunctions(td.input, td.functions)
		if err != nil {
			t.Error(err)
			return
		}

		ctx := context.WithValue(context.Background(), "a", "a")
		got, er := expr.WithContext(ctx).Evaluate(td.params)
		if er != nil && !td.hasError {
			t.Error(er)
			return
		}
		if er != nil {
			t.Log("[info] ", er)
		}

		if isArray(td.want) {
			w := AnyValue(td.want)
			g := AnyValue(got)

			if !g.Equal(w) {
				t.Error("eval failed: <"+td.input+">  got: ", got, "\nwant: ", td.want)
				continue
			}
			continue
		}

		w := AnyValue(td.want)
		g := AnyValue(got)

		if !g.Equal(w) {
			t.Error("eval failed: <"+td.input+"> got: ", got, "\n"+
				"got type", reflect.TypeOf(got).String(), "\n"+
				"want: ", td.want)
		}
	}
}
