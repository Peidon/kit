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
	"time"
)

type Task interface {
	Execute(ctx context.Context, input map[string]interface{}) (interface{}, error)
}

type MockAPI struct {
	response int
}

func (api *MockAPI) Execute(ctx context.Context, input map[string]interface{}) (interface{}, error) {
	time.Sleep(500 * time.Millisecond)
	return api.response, nil
}

type CalculateTask struct {
	left  string
	right string
	op    string
}

func (c *CalculateTask) Execute(ctx context.Context, input map[string]interface{}) (interface{}, error) {
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
