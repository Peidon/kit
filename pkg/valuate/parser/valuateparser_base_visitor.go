// Code generated from ValuateParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ValuateParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseValuateParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseValuateParserVisitor) VisitPlan(ctx *PlanContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseValuateParserVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseValuateParserVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseValuateParserVisitor) VisitPrimaryExpr(ctx *PrimaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseValuateParserVisitor) VisitUnaryExpr(ctx *UnaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseValuateParserVisitor) VisitArguments(ctx *ArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseValuateParserVisitor) VisitExpressionList(ctx *ExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseValuateParserVisitor) VisitVariate(ctx *VariateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseValuateParserVisitor) VisitOperand(ctx *OperandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseValuateParserVisitor) VisitBasicLit(ctx *BasicLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseValuateParserVisitor) VisitObj(ctx *ObjContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseValuateParserVisitor) VisitPair(ctx *PairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseValuateParserVisitor) VisitArr(ctx *ArrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseValuateParserVisitor) VisitIndex(ctx *IndexContext) interface{} {
	return v.VisitChildren(ctx)
}
