// Code generated from valuate.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // valuate

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by valuateParser.
type valuateVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by valuateParser#plan.
	VisitPlan(ctx *PlanContext) interface{}

	// Visit a parse tree produced by valuateParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by valuateParser#primaryExpr.
	VisitPrimaryExpr(ctx *PrimaryExprContext) interface{}

	// Visit a parse tree produced by valuateParser#unaryExpr.
	VisitUnaryExpr(ctx *UnaryExprContext) interface{}

	// Visit a parse tree produced by valuateParser#arguments.
	VisitArguments(ctx *ArgumentsContext) interface{}

	// Visit a parse tree produced by valuateParser#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) interface{}

	// Visit a parse tree produced by valuateParser#operand.
	VisitOperand(ctx *OperandContext) interface{}

	// Visit a parse tree produced by valuateParser#basicLit.
	VisitBasicLit(ctx *BasicLitContext) interface{}

	// Visit a parse tree produced by valuateParser#variate.
	VisitVariate(ctx *VariateContext) interface{}

	// Visit a parse tree produced by valuateParser#index.
	VisitIndex(ctx *IndexContext) interface{}
}