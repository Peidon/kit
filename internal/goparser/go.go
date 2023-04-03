package parser

import (
	"fmt"
)

type GolangExecutor struct {
	*BaseGoParserVisitor
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
			fmt.Println(fn.Block().GetText())
		}
	}

	return ge.BaseGoParserVisitor.VisitChildren(ctx)
}
