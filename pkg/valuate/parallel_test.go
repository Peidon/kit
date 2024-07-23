/*
 * @Copyright @ Shopee
 * @Author: Peidong Xu
 * @Email: peidong.xu@shopee.com
 * @Date: 2024/7/18 8:19 PM
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

func TestParallelAccess(t *testing.T) {

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
		expr.Parallel()
		expr.CatchError(func(err error) (interface{}, error) {
			t.Log("[info]", err)
			return nil, err
		})
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
				"got type", reflect.TypeOf(got), "\n"+
				"want: ", td.want)
		}
	}
}
