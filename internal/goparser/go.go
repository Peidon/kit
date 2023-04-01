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
	fmt.Print(ctx.GetText())

	if ge.BaseGoParserVisitor == nil {
		ge.BaseGoParserVisitor = golangTreeVisitor
	}
	return ge.BaseGoParserVisitor.VisitChildren(ctx)
}
