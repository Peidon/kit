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
	"errors"
)

type Runner interface {
	Run(context.Context) Error
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

	Keys  map[NodeKey]int
	Nodes []*Node

	errs Error
}

func BuildFlow(nodes []*Node) *FlowRunner {

	flow := &FlowRunner{
		Keys:     map[NodeKey]int{},
		params:   map[NodeKey]interface{}{},
		nodeChan: make(chan *Node, len(nodes)),

		Nodes: nodes,
	}
	for i, n := range nodes {
		flow.Keys[n.key] = i
		n.flow = flow
	}
	flow.markReady()
	return flow
}

func (flow *FlowRunner) Run(ctx context.Context) Error {
	flow.ctx, flow.cancelFunc = context.WithCancel(ctx)
	defer flow.cancelFunc()

	flow.runLoop()

	return flow.errs
}

func (flow *FlowRunner) runLoop() {

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

				go func(node *Node, params map[string]interface{}) {
					var err error
					node.result, err = node.inst.Execute(flow.ctx, params)
					if err != nil {
						flow.addErr(ErrorMessage{
							key: node.key,
							err: err,
						})
					}

					flow.nodeChan <- node
				}(node, params)

			case Running:
				running = true

			case Terminated:
				// log info
			}

		}

		if !running {
			if waiting {
				flow.addErr(ErrorMessage{
					err: errors.New("dead lock"),
				})
			}
			break
		}

		select {
		case <-flow.ctx.Done():
			return
		case node := <-flow.nodeChan:
			node.stat = Terminated

			// 这里可以插入取值逻辑
			flow.params[node.key] = node.result

			if flow.Done() {
				flow.done = true
				return
			}

			flow.markReady()
		}

	}
}

func (flow *FlowRunner) markReady() {
	for _, node := range flow.Nodes {
		if node.stat == Waiting && node.Ready() {
			node.stat = Ready
		}
	}
}

func (flow *FlowRunner) Running() bool {
	return !flow.done && flow.errs.Empty()
}

func (flow *FlowRunner) Done() bool {
	for _, node := range flow.Nodes {
		if node.stat != Terminated {
			return false
		}
	}
	return true
}

func (flow *FlowRunner) Result(k NodeKey) (interface{}, bool) {
	v, exists := flow.params[k]
	return v, exists
}

func (flow *FlowRunner) addErr(err ErrorMessage) {
	flow.errs.Append(err)
}

// track each node error
type Error struct {
	errs []ErrorMessage
}

func (err Error) Error() string {
	msg := ""
	for _, m := range err.errs {
		msg = " | " + string(m.key) + m.err.Error()
	}
	return msg
}

func (err Error) Empty() bool {
	return len(err.errs) == 0
}

func (err *Error) Append(msg ErrorMessage) {
	err.errs = append(err.errs, msg)
}

type ErrorMessage struct {
	key NodeKey
	err error
}
