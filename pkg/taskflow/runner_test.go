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
	"github.com/stretchr/testify/assert"
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
	flow := BuildFlow([]*Node{&nodeA, &nodeB, &nodeC})

	err := flow.Run(context.Background())
	if err.Empty() {
		t.Log(flow.Result("C"))
		return
	}
	t.Error(err.Error())
}

func TestDep(t *testing.T) {
	/*
	 E -> D -> A
	 \
	  --> C -> B

	* */

	// api A
	a := MockAPI{
		response: 1,
	}
	// api B
	b := MockAPI{
		response: 2,
	}

	// api C depend B
	c := MockAPI{
		request: map[string]struct{}{
			"B": {},
		},
		response: 3,
	}

	// api D depend A
	d := MockAPI{
		request: map[string]struct{}{
			"A": {},
		},
		response: 4,
	}

	e := MockAPI{
		request: map[string]struct{}{
			"C": {},
			"D": {},
		},
		response: 200,
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
		needs: []NodeKey{NodeKey("B")},
		inst:  &c,
	}

	nodeD := Node{
		key:   "D",
		needs: []NodeKey{NodeKey("A")},
		inst:  &d,
	}

	nodeE := Node{
		key:   "E",
		needs: []NodeKey{NodeKey("C"), NodeKey("D")},
		inst:  &e,
	}

	// run flow
	ctx := context.Background()
	flow := BuildFlow([]*Node{&nodeA, &nodeB, &nodeC, &nodeD, &nodeE})
	err := flow.Run(ctx)
	assert.True(t, err.Empty())
	cR, _ := flow.Result("C")
	t.Log("C=", cR)

	dR, _ := flow.Result("D")
	t.Log("D=", dR)

	eR, _ := flow.Result("E")
	t.Log("E=", eR)
}

func TestRunError(t *testing.T) {
	/*
	 E -> D -> A(missing)
	 \
	  --> C -> B

	* */

	// api B
	b := MockAPI{
		response: 2,
	}

	// api C depend B
	c := MockAPI{
		request: map[string]struct{}{
			"B": {},
		},
		response: 3,
	}

	// api D depend A
	d := MockAPI{
		request: map[string]struct{}{
			"A": {},
		},
		response: 4,
	}

	e := MockAPI{
		request: map[string]struct{}{
			"C": {},
			"D": {},
		},
		response: 200,
	}

	// config flow
	nodeB := Node{
		key:  "B",
		inst: &b,
	}

	nodeC := Node{
		key:   "C",
		needs: []NodeKey{NodeKey("B")},
		inst:  &c,
	}

	nodeD := Node{
		key:   "D",
		needs: []NodeKey{NodeKey("A")}, // missing
		inst:  &d,
	}

	nodeE := Node{
		key:   "E",
		needs: []NodeKey{NodeKey("C"), NodeKey("D")},
		inst:  &e,
	}

	// run flow
	ctx := context.Background()

	//missing A
	//D, E error
	flow := BuildFlow([]*Node{&nodeB, &nodeC, &nodeD, &nodeE})
	err := flow.Run(ctx)
	assert.False(t, err.Empty())
	t.Log("flow err=", flow.errs.Error())
}
