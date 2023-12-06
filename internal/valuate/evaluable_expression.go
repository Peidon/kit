/*
 * @Copyright @ Shopee
 * @Author: Peidong Xu
 * @Email: peidong.xu@shopee.com
 * @Date: 2023/11/10 5:30 PM
 * @Version: kit
 */

package valuate

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"kit/internal/valuate/parser"
)

type EvaluableExpression struct {
	antlr.BaseParseTreeVisitor

	text      string
	stage     *evaluationStage
	functions map[string]ExpressionFunction

	errorTrack *StageError
}

func NewExpression(input string) (ee *EvaluableExpression, err error) {

	// lexer
	stream := antlr.NewInputStream(input)
	lexer := parser.NewValuateLexer(stream)

	// parser
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewValuateParser(tokenStream)
	plan := p.Plan()

	// new Evaluable
	ee = &EvaluableExpression{
		text:       input,
		errorTrack: &StageError{},
	}

	// plan stages
	stageInf := plan.Accept(ee)
	if stage, ok := stageInf.(evaluationStage); ok {
		ee.stage = &stage
	}
	if ee.stage == nil {
		err = ee.errorTrack
	}

	return
}

func (eval *EvaluableExpression) String() string {
	return eval.text
}

func (eval *EvaluableExpression) VisitPlan(ctx *parser.PlanContext) interface{} {
	if expr := ctx.Expression(); expr != nil {
		return expr.Accept(eval)
	}
	return nil
}

func (eval *EvaluableExpression) VisitExpression(ctx *parser.ExpressionContext) interface{} {
	var (
		stages []evaluationStage
	)

	if expr := ctx.PrimaryExpr(); expr != nil {
		return expr.Accept(eval)
	}

	for _, expr := range ctx.AllExpression() {
		stageInf := expr.Accept(eval)
		if stageInf == nil {
			eval.errorTrack.Append(expr.GetText())
			return nil
		}

		if stage, ok := stageInf.(evaluationStage); ok {
			stages = append(stages, stage)
		}
	}
	if op := ctx.DIV(); op != nil {
		return evaluationStage{
			symbol:    DIVIDE,
			opType:    binaryOp,
			operator:  divideStage,
			depends:   stages,
			typeCheck: numberTypeCheck,
		}
	}
	if op := ctx.STAR(); op != nil {
		return evaluationStage{
			symbol:    MULTIPLY,
			opType:    binaryOp,
			operator:  multipleStage,
			depends:   stages,
			typeCheck: numberTypeCheck,
		}
	}
	if op := ctx.MODULUS(); op != nil {
		return evaluationStage{
			symbol:    MODULUS,
			opType:    binaryOp,
			operator:  modulusStage,
			depends:   stages,
			typeCheck: numberTypeCheck,
		}
	}
	if op := ctx.PLUS(); op != nil {
		return evaluationStage{
			symbol:    PLUS,
			opType:    binaryOp,
			operator:  addStage,
			depends:   stages,
			typeCheck: numberTypeCheck,
		}
	}
	if op := ctx.MINUS(); op != nil {
		return evaluationStage{
			symbol:    MINUS,
			opType:    binaryOp,
			operator:  subtractStage,
			depends:   stages,
			typeCheck: numberTypeCheck,
		}
	}

	if op := ctx.EQUALS(); op != nil {
		return evaluationStage{
			symbol:    EQ,
			opType:    binaryOp,
			operator:  equalStage,
			depends:   stages,
			typeCheck: comparableCheck,
		}
	}
	if op := ctx.NOT_EQUALS(); op != nil {
		return evaluationStage{
			symbol:    NEQ,
			opType:    binaryOp,
			operator:  notEqualStage,
			depends:   stages,
			typeCheck: comparableCheck,
		}
	}
	if op := ctx.GREATER(); op != nil {
		return evaluationStage{
			symbol:    GT,
			opType:    binaryOp,
			operator:  gtStage,
			depends:   stages,
			typeCheck: comparableCheck,
		}
	}
	if op := ctx.LESS(); op != nil {
		return evaluationStage{
			symbol:    LT,
			opType:    binaryOp,
			operator:  ltStage,
			depends:   stages,
			typeCheck: comparableCheck,
		}
	}
	if op := ctx.GREATER_OR_EQUALS(); op != nil {
		return evaluationStage{
			symbol:    GTE,
			opType:    binaryOp,
			operator:  gteStage,
			depends:   stages,
			typeCheck: comparableCheck,
		}
	}
	if op := ctx.LESS_OR_EQUALS(); op != nil {
		return evaluationStage{
			symbol:    LTE,
			opType:    binaryOp,
			operator:  lteStage,
			depends:   stages,
			typeCheck: comparableCheck,
		}
	}
	if op := ctx.LOGICAL_AND(); op != nil {
		return evaluationStage{
			symbol:    AND,
			opType:    binaryOp,
			operator:  andStage,
			depends:   stages,
			typeCheck: boolTypeCheck,
		}
	}
	if op := ctx.LOGICAL_OR(); op != nil {
		return evaluationStage{
			symbol:    OR,
			opType:    binaryOp,
			operator:  orStage,
			depends:   stages,
			typeCheck: boolTypeCheck,
		}
	}
	return nil
}

func (eval *EvaluableExpression) VisitPrimaryExpr(ctx *parser.PrimaryExprContext) interface{} {
	if varID := ctx.Variate(); varID != nil {
		return varID.Accept(eval)
	}
	if operand := ctx.Operand(); operand != nil {
		return operand.Accept(eval)
	}

	primary := ctx.PrimaryExpr()
	if primary == nil {
		eval.errorTrack.Append(ctx.GetText())
		return nil
	}

	stageInf := primary.Accept(eval)
	if primaryStage, ok := stageInf.(evaluationStage); ok {
		identify := ctx.IDENTIFIER()

		if ctx.DOT() != nil && identify != nil {
			// accessors op
			field := identify.GetText()
			op := makeAccessorStage(field)

			return evaluationStage{
				symbol:   ACCESS,
				operator: op,
				depends:  []evaluationStage{primaryStage},
			}
		}

		if idx := ctx.Index(); idx != nil {
			// array index op
			idxInf := idx.Accept(eval)

			if idxStage, ok := idxInf.(evaluationStage); ok {
				return evaluationStage{
					opType:    binaryOp,
					symbol:    INDEX,
					operator:  arrayIndexStage,
					typeCheck: arrayTypeCheck,
					depends:   []evaluationStage{primaryStage, idxStage},
				}
			}
		}

	}

	// functional
	if identifier, args := ctx.IDENTIFIER(), ctx.Arguments(); identifier != nil && args != nil {
		functionName := identifier.GetText()
		argsInf := args.Accept(eval)
		if argumentStage, ok := argsInf.(evaluationStage); ok {
			op := makeFuncStage(functionName, eval.functions)
			return evaluationStage{
				opType:   unaryOp,
				symbol:   FUNCTIONAL,
				operator: op,
				depends:  []evaluationStage{argumentStage},
			}
		}
	}

	eval.errorTrack.Append(ctx.GetText())
	return nil
}

func (eval *EvaluableExpression) VisitUnaryExpr(ctx *parser.UnaryExprContext) interface{} {

	var dep []evaluationStage

	if expr := ctx.Expression(); expr != nil {
		stageInf := expr.Accept(eval)

		if stageInf == nil {
			eval.errorTrack.Append(expr.GetText())
			return nil
		}

		if stage, ok := stageInf.(evaluationStage); ok {
			dep = append(dep, stage)
		}
	}

	if minus := ctx.MINUS(); minus != nil {
		return evaluationStage{
			symbol:    NEGATE,
			opType:    unaryOp,
			operator:  negateStage,
			typeCheck: numberTypeCheck,
			depends:   dep,
		}
	}

	if invert := ctx.EXCLAMATION(); invert != nil {
		return evaluationStage{
			symbol:    INVERT,
			opType:    unaryOp,
			operator:  invertStage,
			typeCheck: boolTypeCheck,
			depends:   dep,
		}
	}

	eval.errorTrack.Append(ctx.GetText())
	return nil
}

func (eval *EvaluableExpression) VisitArguments(ctx *parser.ArgumentsContext) interface{} {
	if expr := ctx.ExpressionList(); expr != nil {
		return expr.Accept(eval)
	}

	eval.errorTrack.Append(ctx.GetText())
	return nil
}

func (eval *EvaluableExpression) VisitExpressionList(ctx *parser.ExpressionListContext) interface{} {

	var dep []evaluationStage

	for _, expr := range ctx.AllExpression() {

		stageInf := expr.Accept(eval)

		if stageInf == nil {
			eval.errorTrack.Append(expr.GetText())
			return nil
		}

		if stage, ok := stageInf.(evaluationStage); ok {
			dep = append(dep, stage)
		}
	}

	return evaluationStage{
		symbol:   VALUE,
		depends:  dep,
		operator: exprListStage,
	}
}

func (eval *EvaluableExpression) VisitOperand(ctx *parser.OperandContext) interface{} {
	if basicLit := ctx.BasicLit(); basicLit != nil {
		return basicLit.Accept(eval)
	}
	if expr := ctx.Expression(); expr != nil {
		return expr.Accept(eval)
	}

	eval.errorTrack.Append(ctx.GetText())
	return nil
}

func (eval *EvaluableExpression) VisitBasicLit(ctx *parser.BasicLitContext) interface{} {
	if nilLit := ctx.NIL_LIT(); nilLit != nil {
		op := makeLiteralStage(nilLit.GetText(), NIL)
		return evaluationStage{
			symbol:   LITERAL,
			operator: op,
		}
	}
	if lit := ctx.TRUE(); lit != nil {
		op := makeLiteralStage(lit.GetText(), Bool)
		return evaluationStage{
			symbol:   LITERAL,
			operator: op,
		}
	}
	if lit := ctx.FALSE(); lit != nil {
		op := makeLiteralStage(lit.GetText(), Bool)
		return evaluationStage{
			symbol:   LITERAL,
			operator: op,
		}
	}
	if lit := ctx.INT(); lit != nil {
		op := makeLiteralStage(lit.GetText(), Int)
		return evaluationStage{
			symbol:   LITERAL,
			operator: op,
		}
	}
	if lit := ctx.STRING(); lit != nil {
		op := makeLiteralStage(lit.GetText(), String)
		return evaluationStage{
			symbol:   LITERAL,
			operator: op,
		}
	}
	if lit := ctx.FLOAT_NUMBER(); lit != nil {
		op := makeLiteralStage(lit.GetText(), Float)
		return evaluationStage{
			symbol:   LITERAL,
			operator: op,
		}
	}
	if lit := ctx.Arr(); lit != nil {
		return lit.Accept(eval)
	}

	eval.errorTrack.Append(ctx.GetText())
	return nil
}

func (eval *EvaluableExpression) VisitArr(ctx *parser.ArrContext) interface{} {

	var ifs []interface{}

	for _, lit := range ctx.AllBasicLit() {
		ifs = append(ifs, lit.Accept(eval))
	}

	dep, err := stagesInf(ifs)
	if err != nil {
		eval.errorTrack.Append(ctx.GetText())
		return nil
	}

	return evaluationStage{
		symbol:   LITERAL,
		operator: arrayValueStage,
		depends:  dep,
	}
}

func (eval *EvaluableExpression) VisitObj(ctx *parser.ObjContext) interface{} {
	return eval.VisitChildren(ctx)
}

func (eval *EvaluableExpression) VisitPair(ctx *parser.PairContext) interface{} {
	return eval.VisitChildren(ctx)
}

func (eval *EvaluableExpression) VisitVariate(ctx *parser.VariateContext) interface{} {

	if varID := ctx.VARID(); varID != nil {
		k := varID.GetText()

		if len(k) == 0 {
			eval.errorTrack.Append(ctx.GetText())
			return nil
		}

		op := makeParameterStage(k)
		return evaluationStage{
			symbol:   VALUE,
			operator: op,
		}
	}

	if varID := ctx.IDENTIFIER(); varID != nil {
		k := varID.GetText()
		op := makeParameterStage(k)
		return evaluationStage{
			symbol:   VALUE,
			operator: op,
		}
	}

	eval.errorTrack.Append(ctx.GetText())
	return nil
}

func (eval *EvaluableExpression) VisitIndex(ctx *parser.IndexContext) interface{} {

	var ifs []interface{}

	if idx := ctx.Expression(); idx != nil {
		stageInf := idx.Accept(eval)
		ifs = append(ifs, stageInf)
	}

	dep, err := stagesInf(ifs)
	if err != nil {
		eval.errorTrack.Append(ctx.GetText())
		return nil
	}

	return evaluationStage{
		symbol:    INDEX,
		opType:    unaryOp,
		operator:  indexStage,
		depends:   dep,
		typeCheck: numberTypeCheck,
	}
}

func (eval *EvaluableExpression) Evaluate(parameters map[string]interface{}) (interface{}, error) {
	if parameters == nil {
		return eval.Eval(nil)
	}

	return eval.Eval(MapParameters(parameters))
}

func (eval *EvaluableExpression) Eval(parameters Parameters) (interface{}, error) {
	if eval.stage == nil {
		return nil, NoNeedEvaluate
	}

	if parameters == nil {
		parameters = DummyParameters
	}

	return eval.evaluateStage(eval.stage, parameters)
}

func (eval *EvaluableExpression) evaluateStage(stage *evaluationStage, parameters Parameters) (interface{}, error) {

	if err := checkOperatorType(stage); err != nil {
		return nil, err
	}

	args := make([]Any, 0)

	// 这里简单使用递归实现
	// 可以使用并行化或者调度算法进行性能优化
	// 前提是很多stage 的耗时确实到了需要优化的地步
	for _, dep := range stage.depends {
		arg, err := eval.evaluateStage(&dep, parameters)
		if err != nil {
			return nil, err
		}
		args = append(args, arg)
	}

	// 参数类型检查
	if err := stage.argsTypeCheck(args...); err != nil {
		return nil, err
	}

	return stage.operator(parameters, args...)
}

func checkOperatorType(stage *evaluationStage) error {
	if len(stage.depends) < int(stage.opType) {
		return ArgumentsError
	}
	return nil
}

func stagesInf(ifs []interface{}) (stages []evaluationStage, err error) {
	for _, inf := range ifs {
		if inf == nil {
			err = PlanStageError
			return
		}

		if stage, ok := inf.(evaluationStage); ok {
			stages = append(stages, stage)
		}
	}
	return
}
