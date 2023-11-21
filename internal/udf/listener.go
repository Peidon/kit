/*
 * @Copyright @ Shopee
 * @Author: Peidong Xu
 * @Email: peidong.xu@shopee.com
 * @Date: 2023/11/21 10:55 AM
 * @Version: kit
 */

package udf

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"go.uber.org/zap"
	"log"
)

var (
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
