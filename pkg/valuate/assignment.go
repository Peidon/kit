/*
 * @Copyright @ Shopee
 * @Author: Peidong Xu
 * @Email: peidong.xu@shopee.com
 * @Date: 2023/12/15 4:07 PM
 * @Version: kit
 */

package valuate

import (
	"errors"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"kit/pkg/valuate/parser"
)

type Assignment struct {
	Variable   string
	Accessor   string
	AssignMode int8
	Expression string
	Eval       *EvaluableExpression
}

func ParseAssignment(input string) (*Assignment, error) {
	return AssignWithFunctions(input, nil)
}

func AssignWithFunctions(input string, fs map[string]ExpressionFunction) (*Assignment, error) {
	// lexer
	stream := antlr.NewInputStream(input)
	lexer := parser.NewValuateLexer(stream)

	// parser
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewValuateParser(tokenStream)
	as := p.Assignment()

	// new Evaluable
	ee := &EvaluableExpression{
		text:       input,
		errorTrack: &StageError{},
		functions:  fs,
	}

	// build assignment
	inf := as.Accept(ee)
	if assignment, ok := inf.(*Assignment); ok {
		if assignment.Eval == nil {
			return nil, errors.New("parse expression error, expression: " + assignment.Expression)
		}
		return assignment, nil
	}
	return nil, errors.New("parse assignment error, input: " + input)
}

func (a Assignment) Do(parameters map[string]Any) (map[string]Any, error) {
	k := a.Variable
	old := parameters[k]
	v, err := a.Eval.Evaluate(parameters)
	if err != nil {
		return nil, err
	}
	parameters[k] = v
	return map[string]Any{k: old}, nil
}
