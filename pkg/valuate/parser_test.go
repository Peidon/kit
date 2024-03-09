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
		"${a.b}",
		"abc(x)",
		"1+2.0",
		"acb.b[0].x",
		"!abc",
		"-1",
		"0",
		"true",
		"false",
		"nil",
		`!in(${abc}, "YY")`,
	}
	for _, td := range testsData {
		eval, err := Expression(td)
		t.Log(eval)
		if err != nil {
			t.Error(err)
			return
		}
	}
}
