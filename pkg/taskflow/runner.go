/*
 * @Copyright @ Shopee
 * @Author: Peidong Xu
 * @Email: peidong.xu@shopee.com
 * @Date: 2023/7/17 5:04 下午
 * @Version: kit
 */

package taskflow

import (
	"context"
	"sync"
)

type Runner interface {
	Run(context.Context) error
	Done()
	Terminated() bool
}

type FlowRunner struct {
	ctx        context.Context
	cancelFunc context.CancelFunc

	// use for Parallelization
	// update node status
	nodeChan chan *Node
	done     bool

	// record each node execute output
	// transfer in the whole flow
	params map[NodeKey]interface{}
	parMut sync.RWMutex

	Nodes map[NodeKey]*Node
}

func BuildFlow(nodes []*Node) *FlowRunner {

	flow := &FlowRunner{
		Nodes:    map[NodeKey]*Node{},
		params:   map[NodeKey]interface{}{},
		nodeChan: make(chan *Node, len(nodes)),
	}
	for _, n := range nodes {
		flow.Nodes[n.key] = n
		n.flow = flow
	}
	flow.markReady()
	return flow
}

func (flow *FlowRunner) Run(ctx context.Context) error {
	flow.ctx, flow.cancelFunc = context.WithCancel(ctx)
	defer flow.cancelFunc()
	return flow.runLoop(ctx)
}

func (flow *FlowRunner) executeNode(ctx context.Context, node *Node, params map[string]interface{}) {
	var (
		err error
	)
	withRecover(ctx, func(subCtx context.Context) {
		node.result, err = node.inst.Execute(flow.ctx, params)
		if err != nil {
			node.inst.HandleError(ctx, flow, err)
		}
		flow.nodeChan <- node
	})
}

func (flow *FlowRunner) runLoop(ctx context.Context) error {

	for flow.Running() {
		waiting := false
		running := false

		for _, node := range flow.Nodes {
			switch node.stat {
			case Waiting:
				waiting = true

			case Ready:
				running = true
				node.stat = Running

				// 此处构造输入参数
				// 后续可根据配置自定义构造方式
				params := node.buildParams()

				go flow.executeNode(ctx, node, params)

			case Running:
				running = true

			case Terminated:
				// log info
			}

		}

		// 所有节点都是waiting状态
		if !running && waiting {
			return DeadLockError
		}

		select {
		case <-flow.ctx.Done():
			return TimeoutWithContext

		case node := <-flow.nodeChan:
			node.stat = Terminated

			// 这里可以插入取值逻辑
			flow.UpdateValue(node.key, node.result)

			if flow.Terminated() {
				flow.Done()
				return nil
			}

			flow.markReady()
		}

	}

	return nil
}

func (flow *FlowRunner) markReady() {
	for _, node := range flow.Nodes {
		if node.stat == Waiting && node.Ready() {
			node.stat = Ready
		}
	}
}

func (flow *FlowRunner) Running() bool {
	return !flow.done
}

func (flow *FlowRunner) Done() {
	flow.done = true
}

func (flow *FlowRunner) Terminated() bool {
	for _, node := range flow.Nodes {
		if node.stat != Terminated {
			return false
		}
	}
	return true
}

func (flow *FlowRunner) ValueOf(k NodeKey) (interface{}, bool) {
	flow.parMut.Lock()
	defer flow.parMut.Unlock()
	v, exists := flow.params[k]
	return v, exists
}

func (flow *FlowRunner) UpdateValue(k NodeKey, v interface{}) {
	flow.parMut.Lock()
	defer flow.parMut.Unlock()
	flow.params[k] = v
}
