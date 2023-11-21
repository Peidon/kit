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
	parser "kit/internal/valuate/ast"
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
	lexer := parser.NewvaluateLexer(stream)

	// parser
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewvaluateParser(tokenStream)
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
			eval.errorTrack.tokens = append(eval.errorTrack.tokens, expr.GetText())
			return nil
		}

		if stage, ok := stageInf.(evaluationStage); ok {
			stages = append(stages, stage)
		}
	}
	if op := ctx.DIV(); op != nil {
		return evaluationStage{
			symbol:    DIVIDE,
			opType:    binary,
			operator:  divideStage,
			depends:   stages,
			typeCheck: numberTypeCheck,
		}
	}
	if op := ctx.STAR(); op != nil {
		return evaluationStage{
			symbol:    MULTIPLY,
			opType:    binary,
			operator:  multipleStage,
			depends:   stages,
			typeCheck: numberTypeCheck,
		}
	}
	if op := ctx.PLUS(); op != nil {
		return evaluationStage{
			symbol:    PLUS,
			opType:    binary,
			operator:  addStage,
			depends:   stages,
			typeCheck: numberTypeCheck,
		}
	}
	if op := ctx.MINUS(); op != nil {
		return evaluationStage{
			symbol:    MINUS,
			opType:    binary,
			operator:  subtractStage,
			depends:   stages,
			typeCheck: numberTypeCheck,
		}
	}
	if op := ctx.MODULUS(); op != nil {
		return evaluationStage{
			symbol:    MODULUS,
			opType:    binary,
			operator:  modulusStage,
			depends:   stages,
			typeCheck: numberTypeCheck,
		}
	}
	return nil
}

func (eval *EvaluableExpression) VisitPrimaryExpr(ctx *parser.PrimaryExprContext) interface{} {
	if operand := ctx.Operand(); operand != nil {
		return operand.Accept(eval)
	}

	return nil
}

func (eval *EvaluableExpression) VisitUnaryExpr(ctx *parser.UnaryExprContext) interface{} {

	if minus := ctx.MINUS(); minus != nil {

	}
	return nil
}

func (eval *EvaluableExpression) VisitArguments(ctx *parser.ArgumentsContext) interface{} {
	println("visit arguments")
	return nil
}

func (eval *EvaluableExpression) VisitExpressionList(ctx *parser.ExpressionListContext) interface{} {
	println("visit expression list")
	return nil
}

func (eval *EvaluableExpression) VisitOperand(ctx *parser.OperandContext) interface{} {
	if basicLit := ctx.BasicLit(); basicLit != nil {
		return basicLit.Accept(eval)
	}
	return eval.VisitChildren(ctx)
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

	if varID := ctx.VAR_ID(); varID != nil {

	}

	return nil
}

func (eval *EvaluableExpression) VisitOperandName(ctx *parser.OperandNameContext) interface{} {
	println("visit Operand Name")
	return nil
}

func (eval *EvaluableExpression) VisitQualifiedIdent(ctx *parser.QualifiedIdentContext) interface{} {
	println("visit qualified identify")
	return nil
}

func (eval *EvaluableExpression) VisitIndex(ctx *parser.IndexContext) interface{} {
	return nil
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

	if stage.typeCheck != nil {
		if err := stage.typeCheck(args...); err != nil {
			return nil, err
		}
	}

	return stage.operator(parameters, args...)
}

func checkOperatorType(stage *evaluationStage) error {
	if len(stage.depends) < int(stage.opType) {
		return ArgumentsError
	}
	return nil
}
