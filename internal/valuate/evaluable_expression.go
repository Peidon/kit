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
	antlr.ParseTreeVisitor
	input     string
	stage     *evaluationStage
	functions map[string]ExpressionFunction
}

func NewExpression(input string) *EvaluableExpression {
	return &EvaluableExpression{
		input: input,
	}
}

func (v *EvaluableExpression) VisitPlan(ctx *parser.PlanContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *EvaluableExpression) VisitExpression(ctx *parser.ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *EvaluableExpression) VisitPrimaryExpr(ctx *parser.PrimaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *EvaluableExpression) VisitUnaryExpr(ctx *parser.UnaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *EvaluableExpression) VisitArguments(ctx *parser.ArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *EvaluableExpression) VisitExpressionList(ctx *parser.ExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *EvaluableExpression) VisitOperand(ctx *parser.OperandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *EvaluableExpression) VisitBasicLit(ctx *parser.BasicLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *EvaluableExpression) VisitOperandName(ctx *parser.OperandNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *EvaluableExpression) VisitQualifiedIdent(ctx *parser.QualifiedIdentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *EvaluableExpression) VisitIndex(ctx *parser.IndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *EvaluableExpression) VisitInteger(ctx *parser.IntegerContext) interface{} {
	return v.VisitChildren(ctx)
}
