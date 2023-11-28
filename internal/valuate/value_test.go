/*
 * @Copyright @ Shopee
 * @Author: Peidong Xu
 * @Email: peidong.xu@shopee.com
 * @Date: 2023/11/28 5:44 PM
 * @Version: kit
 */

package valuate

import "testing"

func TestAnyValue(t *testing.T) {
	b := []bool{true, false}
	v := AnyValue(b)
	t.Log(v)
}
