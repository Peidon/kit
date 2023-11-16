/*
 * @Copyright @ Shopee
 * @Author: Peidong Xu
 * @Email: peidong.xu@shopee.com
 * @Date: 2023/11/13 4:08 PM
 * @Version: kit
 */

package valuate

import "errors"

type Parameters interface {
	Get(name string) (Any, error)
}

type MapParameters map[string]Any

func (p MapParameters) Get(name string) (Any, error) {

	value, found := p[name]

	if !found {
		errorMessage := "No parameter '" + name + "' found."
		return nil, errors.New(errorMessage)
	}

	return value, nil
}

var DUMMY_PARAMETERS = MapParameters(map[string]interface{}{})
