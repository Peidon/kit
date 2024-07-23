/*
 * @Copyright @ Shopee
 * @Author: Peidong Xu
 * @Email: peidong.xu@shopee.com
 * @Date: 2024/7/18 3:22 PM
 * @Version: kit
 */

package valuate

import (
	"context"
	"errors"
	"kit/pkg/taskflow"
	"reflect"
	"strconv"
)

const (
	parallelContextParameters = "context_parameters"
)

func (eval *EvaluableExpression) Parallel() *EvaluableExpression {
	eval.isParallel = true
	return eval
}

func (stage evaluationStage) Execute(ctx context.Context, input map[string]interface{}) (interface{}, error) {
	if err := checkOperatorType(&stage); err != nil {
		return nil, err
	}

	args := make([]Any, 0)

	for _, dep := range stage.depends {

		arg := input[dep.id]

		if stage.symbol == AND && isBool(arg) && !arg.(bool) {
			return false, nil
		}
		if stage.symbol == OR && isBool(arg) && arg.(bool) {
			return true, nil
		}
		args = append(args, arg)
	}

	// 参数类型检查
	if err := stage.argsTypeCheck(args...); err != nil {
		return nil, err
	}

	cpi := ctx.Value(parallelContextParameters)
	cp, valid := cpi.(MapParameters)
	if !valid {
		return nil, errors.New("[parallel] parameters error, context params = " + reflect.TypeOf(cpi).String())
	}

	return stage.operator(cp, args...)
}

func (stage evaluationStage) dependsKey() []taskflow.NodeKey {
	needs := make([]taskflow.NodeKey, 0)
	for _, sta := range stage.depends {
		needs = append(needs, taskflow.NodeKey(sta.id))
	}
	return needs
}

// 并行执行
func (eval *EvaluableExpression) stageParallel(stage *evaluationStage, parameters Parameters) (interface{}, error) {
	// build the flow
	nods := initStageNodes(stage)
	flow := taskflow.BuildFlow(nods)

	// run flow
	if eval.ctx == nil {
		eval.ctx = context.Background()
	}
	ctx := context.WithValue(eval.ctx, parallelContextParameters, parameters)
	flow.SetErrorHandler(func(pCtx context.Context, flowErr error) {
		_, _ = eval.errorHandler(flowErr)
	})
	runErr := flow.Run(ctx)
	if runErr != nil {
		return eval.errorHandler(runErr)
	}

	return flow.GetValue(stage.id)
}

func initStageID(stage *evaluationStage) {
	stack := &stageStack{
		elem: []*evaluationStage{stage},
	}

	for !stack.IsEmpty() {
		peek := stack.Pop()
		peek.id = stageID(peek)

		stack.PushBatch(peek.depends)
	}
}

func stageID(stage *evaluationStage) string {
	pr := reflect.ValueOf(stage)

	if pr.CanAddr() {
		return pr.Addr().String()
	}

	pi := uint64(pr.Pointer())
	return strconv.FormatUint(pi, 10)
}

// 将stage 作为task 放置于flow node之中
func initStageNodes(stage *evaluationStage) []*taskflow.Node {

	initStageID(stage)

	nodes := make([]*taskflow.Node, 0)

	stack := &stageStack{
		elem: []*evaluationStage{stage},
	}

	for !stack.IsEmpty() {
		peek := stack.Pop()
		nodes = append(nodes, taskflow.NewNode(peek.id, peek, peek.dependsKey()))

		stack.PushBatch(peek.depends)
	}
	return nodes
}

type stageStack struct {
	peek int
	elem []*evaluationStage
}

func (stack *stageStack) IsEmpty() bool {
	return stack.peek < 0
}

func (stack *stageStack) PushBatch(stages []evaluationStage) {
	for i := range stages {
		stack.Push(&stages[i])
	}
}

func (stack *stageStack) Push(stage *evaluationStage) {
	stack.peek++
	if stack.peek < len(stack.elem) {
		stack.elem[stack.peek] = stage
		return
	}
	stack.elem = append(stack.elem, stage)
}

func (stack *stageStack) Pop() *evaluationStage {
	if stack.peek < 0 {
		return nil
	}
	if len(stack.elem) < 0 {
		return nil
	}
	peek := stack.elem[stack.peek]
	stack.peek--
	return peek
}
