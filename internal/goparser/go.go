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

func (ge *GolangExecutor) VisitSourceFile(ctx *SourceFileContext) interface{} {
	fmt.Println(ctx.GetText())

	if ge.BaseGoParserVisitor == nil {
		ge.BaseGoParserVisitor = golangTreeVisitor
	}

	trees := ctx.GetChildren()
	for _, tree := range trees {
		if fn, ok := tree.(*FunctionDeclContext); ok {

			fmt.Println(fn.IDENTIFIER().GetText())

			block := fn.Block()

			st := block.GetChildren()
			for _, s := range st {
				if statements, ok := s.(*StatementListContext); ok {
					stat := statements.GetChild(0)

					if sc, ok := stat.(*StatementContext); ok {
						fmt.Println(sc.GetText())
					}
				}
			}
		}
	}

	return ge.BaseGoParserVisitor.VisitChildren(ctx)
}
