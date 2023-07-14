package parser

import (
	"fmt"
)

type GolangExecutor struct {
	*BaseGoParserVisitor

	ParamStack []interface{}
}

var (
	golangTreeVisitor = new(BaseGoParserVisitor)
)

func (ge *GolangExecutor) VisitStatementList(ctx *StatementListContext) interface{} {

	stats := ctx.GetChildren()

	for _, stat := range stats {

		sc, ok := stat.(*StatementContext)
		if !ok {
			continue
		}

		sc.Accept(ge)
	}

	return nil
}

func (ge *GolangExecutor) VisitStatement(ctx *StatementContext) interface{} {

	for _, ste := range ctx.GetChildren() {
		if fsc, ok := ste.(*ForStmtContext); ok {
			fmt.Println(fsc.GetText())

			fsc.Accept(ge)
		}

		if ifs, ok := ste.(*IfStmtContext); ok {
			ifs.Accept(ge)
		}
	}
	return nil
}

func (ge *GolangExecutor) VisitIfStmtContext(ctx *IfStmtContext) interface{} {

	//for _, ste := range ctx.GetChildren() {
	//
	//}

	return nil
}

func (ge *GolangExecutor) VisitForStmt(ctx *ForStmtContext) interface{} {

	rangeClause := ctx.RangeClause()

	for _, id := range rangeClause.GetChildren() {
		fmt.Println(id)
	}

	ctx.Block().Accept(ge)

	return nil
}

func (ge *GolangExecutor) VisitBlock(ctx *BlockContext) interface{} {

	for _, s := range ctx.GetChildren() {
		if statements, ok := s.(*StatementListContext); ok {
			statements.Accept(ge)
		}
	}
	return nil
}

func (ge *GolangExecutor) VisitFunctionDecl(ctx *FunctionDeclContext) interface{} {
	block := ctx.Block()
	return block.Accept(ge)
}

func (ge *GolangExecutor) VisitSourceFile(ctx *SourceFileContext) interface{} {
	fmt.Println(ctx.GetText())

	if ge.BaseGoParserVisitor == nil {
		ge.BaseGoParserVisitor = golangTreeVisitor
	}

	trees := ctx.GetChildren()
	for _, tree := range trees {
		if fn, ok := tree.(*FunctionDeclContext); ok {

			fmt.Println(fn.IDENTIFIER().GetText())

			fn.Accept(ge)
		}
	}

	return ge.BaseGoParserVisitor.VisitChildren(ctx)
}
