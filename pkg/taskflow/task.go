/*
 * @Copyright @ Shopee
 * @Author: Peidong Xu
 * @Email: peidong.xu@shopee.com
 * @Date: 2023/7/17 4:35 下午
 * @Version: kit
 */

package taskflow

import (
	"context"
	"errors"
	"time"
)

type Task interface {
	Execute(ctx context.Context, input map[string]interface{}) (interface{}, error)
	HandleError(context.Context, Runner, error)
}

type MockAPI struct {
	request  map[string]struct{}
	response int
}

func (api *MockAPI) Execute(_ context.Context, input map[string]interface{}) (interface{}, error) {

	for k := range api.request {
		if v, ok := input[k]; !ok || v == nil {
			return nil, errors.New("missing param '" + k + "'")
		}
	}

	// mock time cost
	time.Sleep(500 * time.Millisecond)
	return api.response, nil
}

func (api *MockAPI) HandleError(_ context.Context, flow Runner, err error) {
	print("mock api ", err)
	flow.Done()
}

type CalculateTask struct {
	left  string
	right string
	op    string
}

func (c *CalculateTask) HandleError(_ context.Context, flow Runner, err error) {
	print("calculate ", err)
	flow.Done()
}

func (c *CalculateTask) Execute(_ context.Context, input map[string]interface{}) (interface{}, error) {
	left := input[c.left]

	right := input[c.right]
	switch c.op {
	case "+":
		return toInt(left) + toInt(right), nil

	default:
		return 0, nil
	}
}

func toInt(v interface{}) int {
	if i, ok := v.(int); ok {
		return i
	}
	return 0
}
