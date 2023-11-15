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

	Error *ExprError
}

func NewExpression(input string) (ee *EvaluableExpression, err error) {

	// lexer
	stream := antlr.NewInputStream(input)
	lexer := parser.NewvaluateLexer(stream)
	lexer.AddErrorListener(DefaultListener)

	// parser
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewvaluateParser(tokenStream)
	p.AddErrorListener(DefaultListener)
	plan := p.Plan()

	// new Evaluable
	ee = &EvaluableExpression{
		text:  input,
		Error: &ExprError{},
	}

	// load functions
	// todo

	// plan stages
	stageInf := plan.Accept(ee)
	if stage, ok := stageInf.(evaluationStage); ok {
		ee.stage = &stage
	}
	if ee.stage == nil {
		err = ee.Error
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
			eval.Error.tokens = append(eval.Error.tokens, expr.GetText())
			return nil
		}

		if stage, ok := stageInf.(evaluationStage); ok {
			stages = append(stages, stage)
		}
	}
	if op := ctx.DIV(); op != nil {
		return evaluationStage{
			symbol:   DIVIDE,
			opType:   binary,
			operator: divideStage,
			depends:  stages,
			// types check todo
		}
	}
	if op := ctx.STAR(); op != nil {
		return evaluationStage{
			symbol:   MULTIPLY,
			opType:   binary,
			operator: multipleStage,
			depends:  stages,
		}
	}
	if op := ctx.PLUS(); op != nil {
		return evaluationStage{
			symbol:   PLUS,
			opType:   binary,
			operator: addStage,
			depends:  stages,
		}
	}
	if op := ctx.MINUS(); op != nil {
		return evaluationStage{
			symbol:   MINUS,
			opType:   binary,
			operator: subtractStage,
			depends:  stages,
		}
	}
	if op := ctx.MODULUS(); op != nil {
		return evaluationStage{
			symbol:   MODULUS,
			opType:   binary,
			operator: modulusStage,
			depends:  stages,
		}
	}
	return nil
}

func (eval *EvaluableExpression) VisitPrimaryExpr(ctx *parser.PrimaryExprContext) interface{} {
	text := ctx.GetText()

	println("visit primary expression", text)
	return nil
}

func (eval *EvaluableExpression) VisitUnaryExpr(ctx *parser.UnaryExprContext) interface{} {
	println("visit unary expression")
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
	text := ctx.GetText()
	println("visit operand ", text)
	return nil
}

func (eval *EvaluableExpression) VisitBasicLit(ctx *parser.BasicLitContext) interface{} {
	println("visit basic Lit")
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

func (eval *EvaluableExpression) VisitInteger(ctx *parser.IntegerContext) interface{} {
	return nil
}

func (eval *EvaluableExpression) Evaluate(parameters map[string]interface{}) (interface{}, error) {
	if parameters == nil {

	}
	return nil, nil
}
