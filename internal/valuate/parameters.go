/*
 * @Copyright @ Shopee
 * @Author: Peidong Xu
 * @Email: peidong.xu@shopee.com
 * @Date: 2023/11/13 4:08 PM
 * @Version: kit
 */

package valuate

import (
	"errors"
)

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

var DummyParameters = MapParameters(map[string]interface{}{})

type VariableParameters MapParameters

func (p VariableParameters) Get(name string) (Any, error) {
	key := variableKey(name)
	value, found := p[key]

	if !found {
		errorMessage := "No parameter '" + name + "' found."
		return nil, errors.New(errorMessage)
	}

	return value, nil
}

func variableKey(name string) string {
	length := len(name)

	if length > 2 && name[0] == '{' {
		return name[1 : length-1]
	}
	return name
}
