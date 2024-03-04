/*
 * @Copyright @ Shopee
 * @Author: Peidong Xu
 * @Email: peidong.xu@shopee.com
 * @Date: 2023/11/13 4:08 PM
 * @Version: kit
 */

package valuate

import (
	"context"
)

type Parameters interface {
	Get(name string) (Any, error)
	GetContext() context.Context
}

type MapParameters map[string]Any

func (p MapParameters) Get(name string) (Any, error) {

	value, found := p[name]

	if !found {
		return nil, ParameterNotFound{
			paramName: name,
		}
	}

	return value, nil
}

func (p MapParameters) GetContext() context.Context {
	ctxPart, existsErr := p.Get(contextParam)
	if c, valid := ctxPart.(context.Context); valid && existsErr == nil {
		return c
	}
	return context.Background()
}

var DummyParameters = MapParameters(map[string]Any{})
