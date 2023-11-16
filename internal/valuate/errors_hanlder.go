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
	"strings"
)

var (
	ArgumentTypeError = errors.New("argument type error")
	ArgumentsError    = errors.New("arguments number error")
	DivideZeroError   = errors.New("divide zero")
	NoNeedEvaluate    = errors.New("no evaluate")

	// DefaultListener Default listener
	//DefaultListener = &StdListener{
	//	logg: log.New(os.Stdout, "[debug]", log.LstdFlags),
	//}
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

// StageError track errors when plan stages
type StageError struct {
	tokens []string
}

func (the *StageError) Error() string {
	return "plan stages failed: '" + strings.Join(the.tokens, "' -> '") + "'"
}

func (the *StageError) Append(token string) {
	the.tokens = append(the.tokens, token)
}

// TypeError operator arguments type error
type TypeError struct {
	operate    string
	valueToken string
}

func (ter TypeError) Error() string {
	return "Value '" + ter.valueToken + "' cannot be used with the operator '" + ter.operate + "'"
}
