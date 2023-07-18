/*
 * @Copyright @ Shopee
 * @Author: Peidong Xu
 * @Email: peidong.xu@shopee.com
 * @Date: 2023/7/17 5:04 下午
 * @Version: kit
 */

package taskflow

type NodeKey string

type Node struct {
	flow  *FlowRunner // 指向自己所在的flow
	key   NodeKey
	needs []NodeKey // 当前节点依赖的所有节点的Key

	inst   Task        // instance
	result interface{} // task execute result

	stat State
}

// 对于一个节点
// 所有的依赖节点执行完成后
// 把自己设置为Ready
func (node *Node) Ready() bool {
	for _, k := range node.needs {
		f := node.flow
		idx := f.Keys[k]
		if f.Nodes[idx].stat != Terminated {
			return false
		}
	}

	return true
}

// 在这里可以进行参数构造
// 可以复用已有的api proxy 框架
// 先简单实现一种
// 后续可自定义取值逻辑
func (node *Node) buildParams() map[string]interface{} {
	res := map[string]interface{}{}
	for _, k := range node.needs {
		f := node.flow
		res[string(k)] = f.params[k]
	}
	return res
}

// A State indicates the state of a Task.
//
// The following state diagram indicates the possible state transitions:
//
//	       Ready
//	    ↗︎        ↘︎
//	Waiting  ←  Running
//	    ↘︎        ↙︎
//	    Terminated
//
// A Task may move from Waiting to Terminating if one of
// the tasks on which it depends fails.
//
// NOTE: transitions from Running to Waiting are currently not supported. In
// the future this may be possible if a task depends on continuously running
// tasks that send updates.
type State int

const (
	// Waiting indicates a task is blocked on input from another task.
	//
	// NOTE: although this is currently not implemented, a task could
	// theoretically move from the Running to Waiting state.
	Waiting State = iota

	// Ready means a tasks is ready to run, but currently not running.
	Ready

	// Running indicates a goroutine is currently active for a task and that
	// it is not Waiting.
	Running

	// Terminated means a task has stopped running either because it terminated
	// while Running or was aborted by task on which it depends. The error
	// value of a Task indicates the reason for the termination.
	Terminated
)

var stateStrings = map[State]string{
	Waiting:    "Waiting",
	Ready:      "Ready",
	Running:    "Running",
	Terminated: "Terminated",
}

// String reports a human readable string of status s.
func (s State) String() string {
	return stateStrings[s]
}
