/*
 * @Copyright @ Shopee
 * @Author: Peidong Xu
 * @Email: peidong.xu@shopee.com
 * @Date: 2023/3/29 7:07 下午
 * @Version: kit
 */

package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"testing"
)

const (
	sourceCode = `
func Add(a,b int) int {
	return a+b
}
`
)

func TestNewGoParser(t *testing.T) {

	el := new(antlr.ConsoleErrorListener)
	input := antlr.NewInputStream(sourceCode)
	lexer := NewGoLexer(input)
	lexer.AddErrorListener(el)

	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewGoParser(tokenStream)
	p.AddErrorListener(el)

	t.Log("Parser = ", p)
}
