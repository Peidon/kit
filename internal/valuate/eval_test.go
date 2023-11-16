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
		input string
		want  interface{}
	}{
		{
			input: "1+2.0",
			want:  interface{}(3.0),
		},
	}

	for _, td := range testData {
		expr, err := NewExpression(td.input)
		if err != nil {
			t.Error(err)
			return
		}
		got, er := expr.Evaluate(nil)
		if er != nil {
			t.Error(er)
			return
		}
		if got != td.want {
			t.Error("got: ", got, "\nwant: ", td.want)
		}
	}
}
