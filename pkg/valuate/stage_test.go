/*
 * @Copyright @ Shopee
 * @Author: Peidong Xu
 * @Email: peidong.xu@shopee.com
 * @Date: 2023/11/21 11:38 AM
 * @Version: kit
 */

package valuate

import "testing"

func TestDiv(t *testing.T) {
	res, err := modulusStage(nil, 2, 3)
	t.Log(res, err)
}
