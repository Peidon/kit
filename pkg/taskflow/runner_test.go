/*
 * @Copyright @ Shopee
 * @Author: Peidong Xu
 * @Email: peidong.xu@shopee.com
 * @Date: 2023/7/18 11:57 上午
 * @Version: kit
 */

package taskflow

import (
	"context"
	"testing"
)

func TestRun(t *testing.T) {
	// api A
	a := MockAPI{
		response: 1,
	}
	// api B
	b := MockAPI{
		response: 2,
	}

	// Calculate A + B Task
	c := CalculateTask{
		left:  "A",
		right: "B",
		op:    "+",
	}

	// config flow
	nodeA := Node{
		key:  "A",
		inst: &a,
	}
	nodeB := Node{
		key:  "B",
		inst: &b,
	}

	nodeC := Node{
		key:   "C",
		needs: []NodeKey{NodeKey("A"), NodeKey("B")},
		inst:  &c,
	}

	// run flow
	flow := new(FlowRunner)
	flow.Build([]*Node{&nodeA, &nodeB, &nodeC})

	err := flow.Run(context.Background())
	if err.Empty() {
		t.Log(flow.Result("C"))
		return
	}
	t.Error(err.Error())
}
