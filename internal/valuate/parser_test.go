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
		"1+2.0",
		"acb.b.a",
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
