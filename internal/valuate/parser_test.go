/*
 * @Copyright @ Shopee
 * @Author: Peidong Xu
 * @Email: peidong.xu@shopee.com
 * @Date: 2023/11/14 4:59 PM
 * @Version: kit
 */

package valuate

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "kit/internal/valuate/ast"
	"log"
	"os"
	"testing"
)

func TestAST(t *testing.T) {
	testsData := []string{
		"a+b",
		"{a.b}",
	}

	listener := &errorListener{
		logg: log.New(os.Stdout, "test", log.LstdFlags),
	}

	for _, td := range testsData {
		input := antlr.NewInputStream(td)
		lexer := parser.NewvaluateLexer(input)
		lexer.AddErrorListener(listener)

		tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
		p := parser.NewvaluateParser(tokenStream)
		p.AddErrorListener(listener)
		p.Plan()
	}
}
