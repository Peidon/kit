// Code generated from valuate.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // valuate

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BasevaluateVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasevaluateVisitor) VisitPlan(ctx *PlanContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasevaluateVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasevaluateVisitor) VisitPrimaryExpr(ctx *PrimaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasevaluateVisitor) VisitUnaryExpr(ctx *UnaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasevaluateVisitor) VisitArguments(ctx *ArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasevaluateVisitor) VisitExpressionList(ctx *ExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasevaluateVisitor) VisitOperand(ctx *OperandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasevaluateVisitor) VisitBasicLit(ctx *BasicLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasevaluateVisitor) VisitObj(ctx *ObjContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasevaluateVisitor) VisitPair(ctx *PairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasevaluateVisitor) VisitArr(ctx *ArrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasevaluateVisitor) VisitVariate(ctx *VariateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasevaluateVisitor) VisitIndex(ctx *IndexContext) interface{} {
	return v.VisitChildren(ctx)
}
