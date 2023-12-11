/*
 * @Copyright @ Shopee
 * @Author: Peidong Xu
 * @Email: peidong.xu@shopee.com
 * @Date: 2023/11/13 4:08 PM
 * @Version: kit
 */

package valuate

type Parameters interface {
	Get(name string) (Any, error)
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

var DummyParameters = MapParameters(map[string]Any{})
