// Code generated from ValuateParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ValuateParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by ValuateParser.
type ValuateParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ValuateParser#plan.
	VisitPlan(ctx *PlanContext) interface{}

	// Visit a parse tree produced by ValuateParser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by ValuateParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by ValuateParser#primaryExpr.
	VisitPrimaryExpr(ctx *PrimaryExprContext) interface{}

	// Visit a parse tree produced by ValuateParser#unaryExpr.
	VisitUnaryExpr(ctx *UnaryExprContext) interface{}

	// Visit a parse tree produced by ValuateParser#arguments.
	VisitArguments(ctx *ArgumentsContext) interface{}

	// Visit a parse tree produced by ValuateParser#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) interface{}

	// Visit a parse tree produced by ValuateParser#variate.
	VisitVariate(ctx *VariateContext) interface{}

	// Visit a parse tree produced by ValuateParser#operand.
	VisitOperand(ctx *OperandContext) interface{}

	// Visit a parse tree produced by ValuateParser#basicLit.
	VisitBasicLit(ctx *BasicLitContext) interface{}

	// Visit a parse tree produced by ValuateParser#obj.
	VisitObj(ctx *ObjContext) interface{}

	// Visit a parse tree produced by ValuateParser#pair.
	VisitPair(ctx *PairContext) interface{}

	// Visit a parse tree produced by ValuateParser#arr.
	VisitArr(ctx *ArrContext) interface{}

	// Visit a parse tree produced by ValuateParser#index.
	VisitIndex(ctx *IndexContext) interface{}
}
