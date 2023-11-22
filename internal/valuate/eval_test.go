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
			input:  "abc - 2",
			want:   1,
			params: MapParameters(map[string]Any{"abc": 3}),
		},
		{
			input:  "{ab.c} - 2.0",
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
		if er != nil {
			t.Error(er)
			return
		}
		if got != td.want {
			t.Error("got: ", got, "\nwant: ", td.want)
		}
	}
}
