/*
 * @Copyright @ Shopee
 * @Author: Peidong Xu
 * @Email: peidong.xu@shopee.com
 * @Date: 2023/11/14 4:59 PM
 * @Version: kit
 */

package valuate

import (
	"testing"
)

func TestAST(t *testing.T) {
	testsData := []string{
		"a+b-c",
		"{a.b}",
		"abc(x)",
		"",
	}
	for _, td := range testsData {
		eval, err := NewExpression(td)
		t.Log(eval)
		if err != nil {
			//t.Error(err)
			//return
		}
	}
}

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
