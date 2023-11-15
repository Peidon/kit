/*
 * @Copyright @ Shopee
 * @Author: Peidong Xu
 * @Email: peidong.xu@shopee.com
 * @Date: 2023/11/14 3:56 PM
 * @Version: kit
 */

package valuate

import (
	"errors"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"go.uber.org/zap"
	"log"
	"os"
	"strings"
)

const (
	logicalErrorFormat    string = "Value '%v' cannot be used with the logical operator '%s', it is not a bool"
	modifierErrorFormat   string = "Value '%v' cannot be used with the modifier '%s'"
	comparatorErrorFormat string = "Value '%v' cannot be used with the comparator '%s'"
)

var (
	ArgumentTypeError = errors.New("argument type error")
	ArgumentsError    = errors.New("arguments number error")
	DivideZeroError   = errors.New("divide zero")

	// DefaultListener Default listener
	DefaultListener = &StdListener{
		logg: log.New(os.Stdout, "[debug]", log.LstdFlags),
	}
)

type StdListener struct {
	logg   *log.Logger
	Errors int
}

func (l *StdListener) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {
	l.logg.Println("receive syntax error", zap.Any("errMsg", msg), zap.Any("line", line), zap.Any("column", column))
	l.Errors++
}

func (l *StdListener) ReportAmbiguity(_ antlr.Parser, _ *antlr.DFA, startIndex, stopIndex int, _ bool, _ *antlr.BitSet, _ antlr.ATNConfigSet) {
	l.logg.Println("Ambiguity problems", zap.Any("startIndex", startIndex), zap.Any("stopIndex", stopIndex))
	l.Errors++
}
func (l *StdListener) ReportAttemptingFullContext(p antlr.Parser, _ *antlr.DFA, startIndex, stopIndex int, _ *antlr.BitSet, _ antlr.ATNConfigSet) {
	text := p.GetParserRuleContext().GetText()
	l.logg.Println("AttemptingFullContext problems", zap.Any("startIndex", startIndex), zap.Any("stopIndex", stopIndex), zap.String("text", text))
	l.Errors++
}
func (l *StdListener) ReportContextSensitivity(_ antlr.Parser, _ *antlr.DFA, startIndex, stopIndex, prediction int, _ antlr.ATNConfigSet) {
	l.logg.Println("AttemptingFullContext problems", zap.Any("startIndex", startIndex), zap.Any("stopIndex", stopIndex), zap.Any("prediction", prediction))
	l.Errors++
}

type ExprError struct {
	Msg    string
	tokens []string
}

func (the *ExprError) Error() string {
	return the.Msg + ":" + strings.Join(the.tokens, "->")
}

func (the *ExprError) Append(token string) {
	the.tokens = append(the.tokens, token)
}
