package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"testing"
)

const (
	sourceFile = "tests/function.test"
)

func TestNewGoParser(t *testing.T) {

	inputStream, err := antlr.NewFileStream(sourceFile)
	if err != nil {
		t.Error(err)
		return
	}

	el := new(antlr.ConsoleErrorListener)
	lexer := NewGoLexer(inputStream)
	lexer.AddErrorListener(el)

	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := NewGoParser(tokenStream)

	p.BuildParseTrees = true

	executor := new(GolangExecutor)
	p.SourceFile().Accept(executor)
}
