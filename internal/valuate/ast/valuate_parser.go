// Code generated from valuate.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // valuate

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type valuateParser struct {
	*antlr.BaseParser
}

var valuateParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func valuateParserInit() {
	staticData := &valuateParserStaticData
	staticData.literalNames = []string{
		"", "'!'", "'true'", "'false'", "'nil'", "'('", "')'", "'{'", "'}'",
		"'['", "']'", "','", "'.'", "'||'", "'&&'", "'=='", "'!='", "'<'", "'<='",
		"'>'", "'>='", "'/'", "'+'", "'-'", "'*'", "'%'",
	}
	staticData.symbolicNames = []string{
		"", "", "TRUE", "FALSE", "NIL_LIT", "LP", "RP", "L_CURLY", "R_CURLY",
		"L_BRACKET", "R_BRACKET", "COMMA", "DOT", "LOGICAL_OR", "LOGICAL_AND",
		"EQUALS", "NOT_EQUALS", "LESS", "LESS_OR_EQUALS", "GREATER", "GREATER_OR_EQUALS",
		"DIV", "PLUS", "MINUS", "STAR", "MODULUS", "SKIP_", "WHITESPACE", "IDENTIFIER",
		"VAR_ID", "STRING", "STRING_LITERAL", "BYTES_LITERAL", "INT", "DECIMAL_LIT",
		"FLOAT_NUMBER",
	}
	staticData.ruleNames = []string{
		"plan", "expression", "primaryExpr", "unaryExpr", "arguments", "expressionList",
		"operand", "basicLit", "operandName", "qualifiedIdent", "index",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 35, 105, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 3, 1, 29, 8, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 5, 1, 46, 8, 1, 10, 1, 12, 1, 49, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 3, 2, 59, 8, 2, 5, 2, 61, 8, 2, 10, 2, 12, 2, 64, 9,
		2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 3, 4, 71, 8, 4, 1, 4, 1, 4, 1, 5, 1, 5,
		1, 5, 5, 5, 78, 8, 5, 10, 5, 12, 5, 81, 9, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 3, 6, 89, 8, 6, 1, 7, 1, 7, 1, 8, 1, 8, 3, 8, 95, 8, 8, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 0, 2, 2, 4, 11, 0,
		2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 0, 5, 2, 0, 21, 21, 24, 25, 1, 0, 22,
		23, 1, 0, 15, 20, 2, 0, 1, 1, 23, 23, 4, 0, 2, 4, 29, 30, 33, 33, 35, 35,
		107, 0, 22, 1, 0, 0, 0, 2, 28, 1, 0, 0, 0, 4, 50, 1, 0, 0, 0, 6, 65, 1,
		0, 0, 0, 8, 68, 1, 0, 0, 0, 10, 74, 1, 0, 0, 0, 12, 88, 1, 0, 0, 0, 14,
		90, 1, 0, 0, 0, 16, 94, 1, 0, 0, 0, 18, 96, 1, 0, 0, 0, 20, 100, 1, 0,
		0, 0, 22, 23, 3, 2, 1, 0, 23, 24, 5, 0, 0, 1, 24, 1, 1, 0, 0, 0, 25, 26,
		6, 1, -1, 0, 26, 29, 3, 4, 2, 0, 27, 29, 3, 6, 3, 0, 28, 25, 1, 0, 0, 0,
		28, 27, 1, 0, 0, 0, 29, 47, 1, 0, 0, 0, 30, 31, 10, 5, 0, 0, 31, 32, 7,
		0, 0, 0, 32, 46, 3, 2, 1, 6, 33, 34, 10, 4, 0, 0, 34, 35, 7, 1, 0, 0, 35,
		46, 3, 2, 1, 5, 36, 37, 10, 3, 0, 0, 37, 38, 7, 2, 0, 0, 38, 46, 3, 2,
		1, 4, 39, 40, 10, 2, 0, 0, 40, 41, 5, 14, 0, 0, 41, 46, 3, 2, 1, 3, 42,
		43, 10, 1, 0, 0, 43, 44, 5, 13, 0, 0, 44, 46, 3, 2, 1, 2, 45, 30, 1, 0,
		0, 0, 45, 33, 1, 0, 0, 0, 45, 36, 1, 0, 0, 0, 45, 39, 1, 0, 0, 0, 45, 42,
		1, 0, 0, 0, 46, 49, 1, 0, 0, 0, 47, 45, 1, 0, 0, 0, 47, 48, 1, 0, 0, 0,
		48, 3, 1, 0, 0, 0, 49, 47, 1, 0, 0, 0, 50, 51, 6, 2, -1, 0, 51, 52, 3,
		12, 6, 0, 52, 62, 1, 0, 0, 0, 53, 58, 10, 1, 0, 0, 54, 55, 5, 12, 0, 0,
		55, 59, 5, 28, 0, 0, 56, 59, 3, 20, 10, 0, 57, 59, 3, 8, 4, 0, 58, 54,
		1, 0, 0, 0, 58, 56, 1, 0, 0, 0, 58, 57, 1, 0, 0, 0, 59, 61, 1, 0, 0, 0,
		60, 53, 1, 0, 0, 0, 61, 64, 1, 0, 0, 0, 62, 60, 1, 0, 0, 0, 62, 63, 1,
		0, 0, 0, 63, 5, 1, 0, 0, 0, 64, 62, 1, 0, 0, 0, 65, 66, 7, 3, 0, 0, 66,
		67, 3, 2, 1, 0, 67, 7, 1, 0, 0, 0, 68, 70, 5, 5, 0, 0, 69, 71, 3, 10, 5,
		0, 70, 69, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 72, 1, 0, 0, 0, 72, 73,
		5, 6, 0, 0, 73, 9, 1, 0, 0, 0, 74, 79, 3, 2, 1, 0, 75, 76, 5, 11, 0, 0,
		76, 78, 3, 2, 1, 0, 77, 75, 1, 0, 0, 0, 78, 81, 1, 0, 0, 0, 79, 77, 1,
		0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 11, 1, 0, 0, 0, 81, 79, 1, 0, 0, 0, 82,
		89, 3, 14, 7, 0, 83, 89, 3, 16, 8, 0, 84, 85, 5, 5, 0, 0, 85, 86, 3, 2,
		1, 0, 86, 87, 5, 6, 0, 0, 87, 89, 1, 0, 0, 0, 88, 82, 1, 0, 0, 0, 88, 83,
		1, 0, 0, 0, 88, 84, 1, 0, 0, 0, 89, 13, 1, 0, 0, 0, 90, 91, 7, 4, 0, 0,
		91, 15, 1, 0, 0, 0, 92, 95, 5, 28, 0, 0, 93, 95, 3, 18, 9, 0, 94, 92, 1,
		0, 0, 0, 94, 93, 1, 0, 0, 0, 95, 17, 1, 0, 0, 0, 96, 97, 5, 28, 0, 0, 97,
		98, 5, 11, 0, 0, 98, 99, 5, 28, 0, 0, 99, 19, 1, 0, 0, 0, 100, 101, 5,
		9, 0, 0, 101, 102, 3, 2, 1, 0, 102, 103, 5, 10, 0, 0, 103, 21, 1, 0, 0,
		0, 9, 28, 45, 47, 58, 62, 70, 79, 88, 94,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// valuateParserInit initializes any static state used to implement valuateParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewvaluateParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ValuateParserInit() {
	staticData := &valuateParserStaticData
	staticData.once.Do(valuateParserInit)
}

// NewvaluateParser produces a new parser instance for the optional input antlr.TokenStream.
func NewvaluateParser(input antlr.TokenStream) *valuateParser {
	ValuateParserInit()
	this := new(valuateParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &valuateParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "valuate.g4"

	return this
}

// valuateParser tokens.
const (
	valuateParserEOF               = antlr.TokenEOF
	valuateParserT__0              = 1
	valuateParserTRUE              = 2
	valuateParserFALSE             = 3
	valuateParserNIL_LIT           = 4
	valuateParserLP                = 5
	valuateParserRP                = 6
	valuateParserL_CURLY           = 7
	valuateParserR_CURLY           = 8
	valuateParserL_BRACKET         = 9
	valuateParserR_BRACKET         = 10
	valuateParserCOMMA             = 11
	valuateParserDOT               = 12
	valuateParserLOGICAL_OR        = 13
	valuateParserLOGICAL_AND       = 14
	valuateParserEQUALS            = 15
	valuateParserNOT_EQUALS        = 16
	valuateParserLESS              = 17
	valuateParserLESS_OR_EQUALS    = 18
	valuateParserGREATER           = 19
	valuateParserGREATER_OR_EQUALS = 20
	valuateParserDIV               = 21
	valuateParserPLUS              = 22
	valuateParserMINUS             = 23
	valuateParserSTAR              = 24
	valuateParserMODULUS           = 25
	valuateParserSKIP_             = 26
	valuateParserWHITESPACE        = 27
	valuateParserIDENTIFIER        = 28
	valuateParserVAR_ID            = 29
	valuateParserSTRING            = 30
	valuateParserSTRING_LITERAL    = 31
	valuateParserBYTES_LITERAL     = 32
	valuateParserINT               = 33
	valuateParserDECIMAL_LIT       = 34
	valuateParserFLOAT_NUMBER      = 35
)

// valuateParser rules.
const (
	valuateParserRULE_plan           = 0
	valuateParserRULE_expression     = 1
	valuateParserRULE_primaryExpr    = 2
	valuateParserRULE_unaryExpr      = 3
	valuateParserRULE_arguments      = 4
	valuateParserRULE_expressionList = 5
	valuateParserRULE_operand        = 6
	valuateParserRULE_basicLit       = 7
	valuateParserRULE_operandName    = 8
	valuateParserRULE_qualifiedIdent = 9
	valuateParserRULE_index          = 10
)

// IPlanContext is an interface to support dynamic dispatch.
type IPlanContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPlanContext differentiates from other interfaces.
	IsPlanContext()
}

type PlanContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPlanContext() *PlanContext {
	var p = new(PlanContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = valuateParserRULE_plan
	return p
}

func (*PlanContext) IsPlanContext() {}

func NewPlanContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PlanContext {
	var p = new(PlanContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = valuateParserRULE_plan

	return p
}

func (s *PlanContext) GetParser() antlr.Parser { return s.parser }

func (s *PlanContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PlanContext) EOF() antlr.TerminalNode {
	return s.GetToken(valuateParserEOF, 0)
}

func (s *PlanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PlanContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PlanContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case valuateVisitor:
		return t.VisitPlan(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *valuateParser) Plan() (localctx IPlanContext) {
	this := p
	_ = this

	localctx = NewPlanContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, valuateParserRULE_plan)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(22)
		p.expression(0)
	}
	{
		p.SetState(23)
		p.Match(valuateParserEOF)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = valuateParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = valuateParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) PrimaryExpr() IPrimaryExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExprContext)
}

func (s *ExpressionContext) UnaryExpr() IUnaryExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryExprContext)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) STAR() antlr.TerminalNode {
	return s.GetToken(valuateParserSTAR, 0)
}

func (s *ExpressionContext) DIV() antlr.TerminalNode {
	return s.GetToken(valuateParserDIV, 0)
}

func (s *ExpressionContext) MODULUS() antlr.TerminalNode {
	return s.GetToken(valuateParserMODULUS, 0)
}

func (s *ExpressionContext) PLUS() antlr.TerminalNode {
	return s.GetToken(valuateParserPLUS, 0)
}

func (s *ExpressionContext) MINUS() antlr.TerminalNode {
	return s.GetToken(valuateParserMINUS, 0)
}

func (s *ExpressionContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(valuateParserEQUALS, 0)
}

func (s *ExpressionContext) NOT_EQUALS() antlr.TerminalNode {
	return s.GetToken(valuateParserNOT_EQUALS, 0)
}

func (s *ExpressionContext) LESS() antlr.TerminalNode {
	return s.GetToken(valuateParserLESS, 0)
}

func (s *ExpressionContext) LESS_OR_EQUALS() antlr.TerminalNode {
	return s.GetToken(valuateParserLESS_OR_EQUALS, 0)
}

func (s *ExpressionContext) GREATER() antlr.TerminalNode {
	return s.GetToken(valuateParserGREATER, 0)
}

func (s *ExpressionContext) GREATER_OR_EQUALS() antlr.TerminalNode {
	return s.GetToken(valuateParserGREATER_OR_EQUALS, 0)
}

func (s *ExpressionContext) LOGICAL_AND() antlr.TerminalNode {
	return s.GetToken(valuateParserLOGICAL_AND, 0)
}

func (s *ExpressionContext) LOGICAL_OR() antlr.TerminalNode {
	return s.GetToken(valuateParserLOGICAL_OR, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case valuateVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *valuateParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *valuateParser) expression(_p int) (localctx IExpressionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, valuateParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(28)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case valuateParserTRUE, valuateParserFALSE, valuateParserNIL_LIT, valuateParserLP, valuateParserIDENTIFIER, valuateParserVAR_ID, valuateParserSTRING, valuateParserINT, valuateParserFLOAT_NUMBER:
		{
			p.SetState(26)
			p.primaryExpr(0)
		}

	case valuateParserT__0, valuateParserMINUS:
		{
			p.SetState(27)
			p.UnaryExpr()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(45)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, valuateParserRULE_expression)
				p.SetState(30)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(31)
					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<valuateParserDIV)|(1<<valuateParserSTAR)|(1<<valuateParserMODULUS))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(32)
					p.expression(6)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, valuateParserRULE_expression)
				p.SetState(33)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(34)
					_la = p.GetTokenStream().LA(1)

					if !(_la == valuateParserPLUS || _la == valuateParserMINUS) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(35)
					p.expression(5)
				}

			case 3:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, valuateParserRULE_expression)
				p.SetState(36)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(37)
					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<valuateParserEQUALS)|(1<<valuateParserNOT_EQUALS)|(1<<valuateParserLESS)|(1<<valuateParserLESS_OR_EQUALS)|(1<<valuateParserGREATER)|(1<<valuateParserGREATER_OR_EQUALS))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(38)
					p.expression(4)
				}

			case 4:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, valuateParserRULE_expression)
				p.SetState(39)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(40)
					p.Match(valuateParserLOGICAL_AND)
				}
				{
					p.SetState(41)
					p.expression(3)
				}

			case 5:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, valuateParserRULE_expression)
				p.SetState(42)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(43)
					p.Match(valuateParserLOGICAL_OR)
				}
				{
					p.SetState(44)
					p.expression(2)
				}

			}

		}
		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

// IPrimaryExprContext is an interface to support dynamic dispatch.
type IPrimaryExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryExprContext differentiates from other interfaces.
	IsPrimaryExprContext()
}

type PrimaryExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExprContext() *PrimaryExprContext {
	var p = new(PrimaryExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = valuateParserRULE_primaryExpr
	return p
}

func (*PrimaryExprContext) IsPrimaryExprContext() {}

func NewPrimaryExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExprContext {
	var p = new(PrimaryExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = valuateParserRULE_primaryExpr

	return p
}

func (s *PrimaryExprContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExprContext) Operand() IOperandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *PrimaryExprContext) PrimaryExpr() IPrimaryExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExprContext)
}

func (s *PrimaryExprContext) DOT() antlr.TerminalNode {
	return s.GetToken(valuateParserDOT, 0)
}

func (s *PrimaryExprContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(valuateParserIDENTIFIER, 0)
}

func (s *PrimaryExprContext) Index() IIndexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexContext)
}

func (s *PrimaryExprContext) Arguments() IArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *PrimaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case valuateVisitor:
		return t.VisitPrimaryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *valuateParser) PrimaryExpr() (localctx IPrimaryExprContext) {
	return p.primaryExpr(0)
}

func (p *valuateParser) primaryExpr(_p int) (localctx IPrimaryExprContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewPrimaryExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPrimaryExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, valuateParserRULE_primaryExpr, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(51)
		p.Operand()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewPrimaryExprContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, valuateParserRULE_primaryExpr)
			p.SetState(53)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			p.SetState(58)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case valuateParserDOT:
				{
					p.SetState(54)
					p.Match(valuateParserDOT)
				}
				{
					p.SetState(55)
					p.Match(valuateParserIDENTIFIER)
				}

			case valuateParserL_BRACKET:
				{
					p.SetState(56)
					p.Index()
				}

			case valuateParserLP:
				{
					p.SetState(57)
					p.Arguments()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		}
		p.SetState(64)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}

	return localctx
}

// IUnaryExprContext is an interface to support dynamic dispatch.
type IUnaryExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnaryExprContext differentiates from other interfaces.
	IsUnaryExprContext()
}

type UnaryExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryExprContext() *UnaryExprContext {
	var p = new(UnaryExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = valuateParserRULE_unaryExpr
	return p
}

func (*UnaryExprContext) IsUnaryExprContext() {}

func NewUnaryExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryExprContext {
	var p = new(UnaryExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = valuateParserRULE_unaryExpr

	return p
}

func (s *UnaryExprContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryExprContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *UnaryExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(valuateParserMINUS, 0)
}

func (s *UnaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case valuateVisitor:
		return t.VisitUnaryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *valuateParser) UnaryExpr() (localctx IUnaryExprContext) {
	this := p
	_ = this

	localctx = NewUnaryExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, valuateParserRULE_unaryExpr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(65)
		_la = p.GetTokenStream().LA(1)

		if !(_la == valuateParserT__0 || _la == valuateParserMINUS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(66)
		p.expression(0)
	}

	return localctx
}

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = valuateParserRULE_arguments
	return p
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = valuateParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) LP() antlr.TerminalNode {
	return s.GetToken(valuateParserLP, 0)
}

func (s *ArgumentsContext) RP() antlr.TerminalNode {
	return s.GetToken(valuateParserRP, 0)
}

func (s *ArgumentsContext) ExpressionList() IExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case valuateVisitor:
		return t.VisitArguments(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *valuateParser) Arguments() (localctx IArgumentsContext) {
	this := p
	_ = this

	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, valuateParserRULE_arguments)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(68)
		p.Match(valuateParserLP)
	}
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<valuateParserT__0)|(1<<valuateParserTRUE)|(1<<valuateParserFALSE)|(1<<valuateParserNIL_LIT)|(1<<valuateParserLP)|(1<<valuateParserMINUS)|(1<<valuateParserIDENTIFIER)|(1<<valuateParserVAR_ID)|(1<<valuateParserSTRING))) != 0) || _la == valuateParserINT || _la == valuateParserFLOAT_NUMBER {
		{
			p.SetState(69)
			p.ExpressionList()
		}

	}
	{
		p.SetState(72)
		p.Match(valuateParserRP)
	}

	return localctx
}

// IExpressionListContext is an interface to support dynamic dispatch.
type IExpressionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionListContext differentiates from other interfaces.
	IsExpressionListContext()
}

type ExpressionListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionListContext() *ExpressionListContext {
	var p = new(ExpressionListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = valuateParserRULE_expressionList
	return p
}

func (*ExpressionListContext) IsExpressionListContext() {}

func NewExpressionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionListContext {
	var p = new(ExpressionListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = valuateParserRULE_expressionList

	return p
}

func (s *ExpressionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionListContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionListContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(valuateParserCOMMA)
}

func (s *ExpressionListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(valuateParserCOMMA, i)
}

func (s *ExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case valuateVisitor:
		return t.VisitExpressionList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *valuateParser) ExpressionList() (localctx IExpressionListContext) {
	this := p
	_ = this

	localctx = NewExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, valuateParserRULE_expressionList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(74)
		p.expression(0)
	}
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == valuateParserCOMMA {
		{
			p.SetState(75)
			p.Match(valuateParserCOMMA)
		}
		{
			p.SetState(76)
			p.expression(0)
		}

		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IOperandContext is an interface to support dynamic dispatch.
type IOperandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperandContext differentiates from other interfaces.
	IsOperandContext()
}

type OperandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperandContext() *OperandContext {
	var p = new(OperandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = valuateParserRULE_operand
	return p
}

func (*OperandContext) IsOperandContext() {}

func NewOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperandContext {
	var p = new(OperandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = valuateParserRULE_operand

	return p
}

func (s *OperandContext) GetParser() antlr.Parser { return s.parser }

func (s *OperandContext) BasicLit() IBasicLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBasicLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBasicLitContext)
}

func (s *OperandContext) OperandName() IOperandNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperandNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperandNameContext)
}

func (s *OperandContext) LP() antlr.TerminalNode {
	return s.GetToken(valuateParserLP, 0)
}

func (s *OperandContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *OperandContext) RP() antlr.TerminalNode {
	return s.GetToken(valuateParserRP, 0)
}

func (s *OperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperandContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case valuateVisitor:
		return t.VisitOperand(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *valuateParser) Operand() (localctx IOperandContext) {
	this := p
	_ = this

	localctx = NewOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, valuateParserRULE_operand)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(88)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case valuateParserTRUE, valuateParserFALSE, valuateParserNIL_LIT, valuateParserVAR_ID, valuateParserSTRING, valuateParserINT, valuateParserFLOAT_NUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(82)
			p.BasicLit()
		}

	case valuateParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(83)
			p.OperandName()
		}

	case valuateParserLP:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(84)
			p.Match(valuateParserLP)
		}
		{
			p.SetState(85)
			p.expression(0)
		}
		{
			p.SetState(86)
			p.Match(valuateParserRP)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBasicLitContext is an interface to support dynamic dispatch.
type IBasicLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBasicLitContext differentiates from other interfaces.
	IsBasicLitContext()
}

type BasicLitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBasicLitContext() *BasicLitContext {
	var p = new(BasicLitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = valuateParserRULE_basicLit
	return p
}

func (*BasicLitContext) IsBasicLitContext() {}

func NewBasicLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BasicLitContext {
	var p = new(BasicLitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = valuateParserRULE_basicLit

	return p
}

func (s *BasicLitContext) GetParser() antlr.Parser { return s.parser }

func (s *BasicLitContext) NIL_LIT() antlr.TerminalNode {
	return s.GetToken(valuateParserNIL_LIT, 0)
}

func (s *BasicLitContext) TRUE() antlr.TerminalNode {
	return s.GetToken(valuateParserTRUE, 0)
}

func (s *BasicLitContext) FALSE() antlr.TerminalNode {
	return s.GetToken(valuateParserFALSE, 0)
}

func (s *BasicLitContext) INT() antlr.TerminalNode {
	return s.GetToken(valuateParserINT, 0)
}

func (s *BasicLitContext) STRING() antlr.TerminalNode {
	return s.GetToken(valuateParserSTRING, 0)
}

func (s *BasicLitContext) FLOAT_NUMBER() antlr.TerminalNode {
	return s.GetToken(valuateParserFLOAT_NUMBER, 0)
}

func (s *BasicLitContext) VAR_ID() antlr.TerminalNode {
	return s.GetToken(valuateParserVAR_ID, 0)
}

func (s *BasicLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BasicLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BasicLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case valuateVisitor:
		return t.VisitBasicLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *valuateParser) BasicLit() (localctx IBasicLitContext) {
	this := p
	_ = this

	localctx = NewBasicLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, valuateParserRULE_basicLit)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(90)
		_la = p.GetTokenStream().LA(1)

		if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<valuateParserTRUE)|(1<<valuateParserFALSE)|(1<<valuateParserNIL_LIT)|(1<<valuateParserVAR_ID)|(1<<valuateParserSTRING))) != 0) || _la == valuateParserINT || _la == valuateParserFLOAT_NUMBER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IOperandNameContext is an interface to support dynamic dispatch.
type IOperandNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperandNameContext differentiates from other interfaces.
	IsOperandNameContext()
}

type OperandNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperandNameContext() *OperandNameContext {
	var p = new(OperandNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = valuateParserRULE_operandName
	return p
}

func (*OperandNameContext) IsOperandNameContext() {}

func NewOperandNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperandNameContext {
	var p = new(OperandNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = valuateParserRULE_operandName

	return p
}

func (s *OperandNameContext) GetParser() antlr.Parser { return s.parser }

func (s *OperandNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(valuateParserIDENTIFIER, 0)
}

func (s *OperandNameContext) QualifiedIdent() IQualifiedIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQualifiedIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQualifiedIdentContext)
}

func (s *OperandNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperandNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case valuateVisitor:
		return t.VisitOperandName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *valuateParser) OperandName() (localctx IOperandNameContext) {
	this := p
	_ = this

	localctx = NewOperandNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, valuateParserRULE_operandName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(92)
			p.Match(valuateParserIDENTIFIER)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(93)
			p.QualifiedIdent()
		}

	}

	return localctx
}

// IQualifiedIdentContext is an interface to support dynamic dispatch.
type IQualifiedIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQualifiedIdentContext differentiates from other interfaces.
	IsQualifiedIdentContext()
}

type QualifiedIdentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQualifiedIdentContext() *QualifiedIdentContext {
	var p = new(QualifiedIdentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = valuateParserRULE_qualifiedIdent
	return p
}

func (*QualifiedIdentContext) IsQualifiedIdentContext() {}

func NewQualifiedIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QualifiedIdentContext {
	var p = new(QualifiedIdentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = valuateParserRULE_qualifiedIdent

	return p
}

func (s *QualifiedIdentContext) GetParser() antlr.Parser { return s.parser }

func (s *QualifiedIdentContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(valuateParserIDENTIFIER)
}

func (s *QualifiedIdentContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(valuateParserIDENTIFIER, i)
}

func (s *QualifiedIdentContext) COMMA() antlr.TerminalNode {
	return s.GetToken(valuateParserCOMMA, 0)
}

func (s *QualifiedIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QualifiedIdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QualifiedIdentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case valuateVisitor:
		return t.VisitQualifiedIdent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *valuateParser) QualifiedIdent() (localctx IQualifiedIdentContext) {
	this := p
	_ = this

	localctx = NewQualifiedIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, valuateParserRULE_qualifiedIdent)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)
		p.Match(valuateParserIDENTIFIER)
	}
	{
		p.SetState(97)
		p.Match(valuateParserCOMMA)
	}
	{
		p.SetState(98)
		p.Match(valuateParserIDENTIFIER)
	}

	return localctx
}

// IIndexContext is an interface to support dynamic dispatch.
type IIndexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIndexContext differentiates from other interfaces.
	IsIndexContext()
}

type IndexContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexContext() *IndexContext {
	var p = new(IndexContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = valuateParserRULE_index
	return p
}

func (*IndexContext) IsIndexContext() {}

func NewIndexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexContext {
	var p = new(IndexContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = valuateParserRULE_index

	return p
}

func (s *IndexContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(valuateParserL_BRACKET, 0)
}

func (s *IndexContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IndexContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(valuateParserR_BRACKET, 0)
}

func (s *IndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case valuateVisitor:
		return t.VisitIndex(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *valuateParser) Index() (localctx IIndexContext) {
	this := p
	_ = this

	localctx = NewIndexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, valuateParserRULE_index)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(100)
		p.Match(valuateParserL_BRACKET)
	}
	{
		p.SetState(101)
		p.expression(0)
	}
	{
		p.SetState(102)
		p.Match(valuateParserR_BRACKET)
	}

	return localctx
}

func (p *valuateParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 2:
		var t *PrimaryExprContext = nil
		if localctx != nil {
			t = localctx.(*PrimaryExprContext)
		}
		return p.PrimaryExpr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *valuateParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *valuateParser) PrimaryExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 5:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
