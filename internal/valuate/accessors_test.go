/*
 * @Copyright @ Shopee
 * @Author: Peidong Xu
 * @Email: peidong.xu@shopee.com
 * @Date: 2023/11/7 2:51 PM
 * @Version: kit
 */

package valuate

import (
	"github.com/Peidon/govaluate"
	"testing"
)

func TestAccessors(t *testing.T) {
	expression, _ := govaluate.NewEvaluableExpression("abc.A > 0 && [x.y] > 100")
	parameters := make(map[string]interface{}, 8)
	parameters["foo"] = -1

	parameters["abc"] = Abc{
		A: 2,
		B: "s",
	}

	parameters["x.y"] = 123

	result, err := expression.Evaluate(parameters)
	if err != nil {
		t.Error(err)
	}

	var (
		a = float32(1.0)
		b = int64(8)
	)

	c := float64(a) + float64(b)

	t.Log(c)
	t.Log(result)
}

type Abc struct {
	A int
	B string
	c bool
}

func (abc Abc) Echo(hello string) string {
	return hello
}
