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
		"'['", "']'", "','", "':'", "'.'", "'||'", "'&&'", "'=='", "'!='", "'<'",
		"'<='", "'>'", "'>='", "'/'", "'+'", "'-'", "'*'", "'%'",
	}
	staticData.symbolicNames = []string{
		"", "", "TRUE", "FALSE", "NIL_LIT", "LP", "RP", "L_CURLY", "R_CURLY",
		"L_BRACKET", "R_BRACKET", "COMMA", "COLON", "DOT", "LOGICAL_OR", "LOGICAL_AND",
		"EQUALS", "NOT_EQUALS", "LESS", "LESS_OR_EQUALS", "GREATER", "GREATER_OR_EQUALS",
		"DIV", "PLUS", "MINUS", "STAR", "MODULUS", "SKIP_", "WHITESPACE", "IDENTIFIER",
		"STRING_LITERAL", "BYTES_LITERAL", "INT", "DECIMAL_LIT", "FLOAT_NUMBER",
		"STRING",
	}
	staticData.ruleNames = []string{
		"plan", "expression", "primaryExpr", "unaryExpr", "arguments", "expressionList",
		"operand", "basicLit", "obj", "pair", "arr", "variate", "index",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 35, 157, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 3,
		1, 33, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 50, 8, 1, 10, 1, 12, 1, 53, 9, 1, 1,
		2, 1, 2, 1, 2, 3, 2, 58, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 65,
		8, 2, 5, 2, 67, 8, 2, 10, 2, 12, 2, 70, 9, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1,
		4, 3, 4, 77, 8, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 5, 5, 84, 8, 5, 10, 5,
		12, 5, 87, 9, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 94, 8, 6, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 104, 8, 7, 1, 8, 1, 8, 1,
		8, 1, 8, 5, 8, 110, 8, 8, 10, 8, 12, 8, 113, 9, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 3, 8, 119, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10,
		5, 10, 129, 8, 10, 10, 10, 12, 10, 132, 9, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 3, 10, 138, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 145, 8,
		11, 10, 11, 12, 11, 148, 9, 11, 1, 11, 3, 11, 151, 8, 11, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 0, 2, 2, 4, 13, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18,
		20, 22, 24, 0, 4, 2, 0, 22, 22, 25, 26, 1, 0, 23, 24, 1, 0, 16, 21, 2,
		0, 1, 1, 24, 24, 169, 0, 26, 1, 0, 0, 0, 2, 32, 1, 0, 0, 0, 4, 57, 1, 0,
		0, 0, 6, 71, 1, 0, 0, 0, 8, 74, 1, 0, 0, 0, 10, 80, 1, 0, 0, 0, 12, 93,
		1, 0, 0, 0, 14, 103, 1, 0, 0, 0, 16, 118, 1, 0, 0, 0, 18, 120, 1, 0, 0,
		0, 20, 137, 1, 0, 0, 0, 22, 150, 1, 0, 0, 0, 24, 152, 1, 0, 0, 0, 26, 27,
		3, 2, 1, 0, 27, 28, 5, 0, 0, 1, 28, 1, 1, 0, 0, 0, 29, 30, 6, 1, -1, 0,
		30, 33, 3, 4, 2, 0, 31, 33, 3, 6, 3, 0, 32, 29, 1, 0, 0, 0, 32, 31, 1,
		0, 0, 0, 33, 51, 1, 0, 0, 0, 34, 35, 10, 5, 0, 0, 35, 36, 7, 0, 0, 0, 36,
		50, 3, 2, 1, 6, 37, 38, 10, 4, 0, 0, 38, 39, 7, 1, 0, 0, 39, 50, 3, 2,
		1, 5, 40, 41, 10, 3, 0, 0, 41, 42, 7, 2, 0, 0, 42, 50, 3, 2, 1, 4, 43,
		44, 10, 2, 0, 0, 44, 45, 5, 15, 0, 0, 45, 50, 3, 2, 1, 3, 46, 47, 10, 1,
		0, 0, 47, 48, 5, 14, 0, 0, 48, 50, 3, 2, 1, 2, 49, 34, 1, 0, 0, 0, 49,
		37, 1, 0, 0, 0, 49, 40, 1, 0, 0, 0, 49, 43, 1, 0, 0, 0, 49, 46, 1, 0, 0,
		0, 50, 53, 1, 0, 0, 0, 51, 49, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52, 3, 1,
		0, 0, 0, 53, 51, 1, 0, 0, 0, 54, 55, 6, 2, -1, 0, 55, 58, 3, 12, 6, 0,
		56, 58, 3, 22, 11, 0, 57, 54, 1, 0, 0, 0, 57, 56, 1, 0, 0, 0, 58, 68, 1,
		0, 0, 0, 59, 64, 10, 1, 0, 0, 60, 61, 5, 13, 0, 0, 61, 65, 5, 29, 0, 0,
		62, 65, 3, 24, 12, 0, 63, 65, 3, 8, 4, 0, 64, 60, 1, 0, 0, 0, 64, 62, 1,
		0, 0, 0, 64, 63, 1, 0, 0, 0, 65, 67, 1, 0, 0, 0, 66, 59, 1, 0, 0, 0, 67,
		70, 1, 0, 0, 0, 68, 66, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 5, 1, 0, 0,
		0, 70, 68, 1, 0, 0, 0, 71, 72, 7, 3, 0, 0, 72, 73, 3, 2, 1, 0, 73, 7, 1,
		0, 0, 0, 74, 76, 5, 5, 0, 0, 75, 77, 3, 10, 5, 0, 76, 75, 1, 0, 0, 0, 76,
		77, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 79, 5, 6, 0, 0, 79, 9, 1, 0, 0,
		0, 80, 85, 3, 2, 1, 0, 81, 82, 5, 11, 0, 0, 82, 84, 3, 2, 1, 0, 83, 81,
		1, 0, 0, 0, 84, 87, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0, 85, 86, 1, 0, 0, 0,
		86, 11, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 88, 94, 3, 14, 7, 0, 89, 90, 5,
		5, 0, 0, 90, 91, 3, 2, 1, 0, 91, 92, 5, 6, 0, 0, 92, 94, 1, 0, 0, 0, 93,
		88, 1, 0, 0, 0, 93, 89, 1, 0, 0, 0, 94, 13, 1, 0, 0, 0, 95, 104, 5, 35,
		0, 0, 96, 104, 5, 32, 0, 0, 97, 104, 5, 34, 0, 0, 98, 104, 3, 16, 8, 0,
		99, 104, 3, 20, 10, 0, 100, 104, 5, 2, 0, 0, 101, 104, 5, 3, 0, 0, 102,
		104, 5, 4, 0, 0, 103, 95, 1, 0, 0, 0, 103, 96, 1, 0, 0, 0, 103, 97, 1,
		0, 0, 0, 103, 98, 1, 0, 0, 0, 103, 99, 1, 0, 0, 0, 103, 100, 1, 0, 0, 0,
		103, 101, 1, 0, 0, 0, 103, 102, 1, 0, 0, 0, 104, 15, 1, 0, 0, 0, 105, 106,
		5, 7, 0, 0, 106, 111, 3, 18, 9, 0, 107, 108, 5, 11, 0, 0, 108, 110, 3,
		18, 9, 0, 109, 107, 1, 0, 0, 0, 110, 113, 1, 0, 0, 0, 111, 109, 1, 0, 0,
		0, 111, 112, 1, 0, 0, 0, 112, 114, 1, 0, 0, 0, 113, 111, 1, 0, 0, 0, 114,
		115, 5, 8, 0, 0, 115, 119, 1, 0, 0, 0, 116, 117, 5, 7, 0, 0, 117, 119,
		5, 8, 0, 0, 118, 105, 1, 0, 0, 0, 118, 116, 1, 0, 0, 0, 119, 17, 1, 0,
		0, 0, 120, 121, 5, 30, 0, 0, 121, 122, 5, 12, 0, 0, 122, 123, 3, 14, 7,
		0, 123, 19, 1, 0, 0, 0, 124, 125, 5, 9, 0, 0, 125, 130, 3, 14, 7, 0, 126,
		127, 5, 11, 0, 0, 127, 129, 3, 14, 7, 0, 128, 126, 1, 0, 0, 0, 129, 132,
		1, 0, 0, 0, 130, 128, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131, 133, 1, 0,
		0, 0, 132, 130, 1, 0, 0, 0, 133, 134, 5, 10, 0, 0, 134, 138, 1, 0, 0, 0,
		135, 136, 5, 9, 0, 0, 136, 138, 5, 10, 0, 0, 137, 124, 1, 0, 0, 0, 137,
		135, 1, 0, 0, 0, 138, 21, 1, 0, 0, 0, 139, 151, 5, 29, 0, 0, 140, 141,
		5, 7, 0, 0, 141, 146, 5, 29, 0, 0, 142, 143, 5, 13, 0, 0, 143, 145, 5,
		29, 0, 0, 144, 142, 1, 0, 0, 0, 145, 148, 1, 0, 0, 0, 146, 144, 1, 0, 0,
		0, 146, 147, 1, 0, 0, 0, 147, 149, 1, 0, 0, 0, 148, 146, 1, 0, 0, 0, 149,
		151, 5, 8, 0, 0, 150, 139, 1, 0, 0, 0, 150, 140, 1, 0, 0, 0, 151, 23, 1,
		0, 0, 0, 152, 153, 5, 9, 0, 0, 153, 154, 3, 2, 1, 0, 154, 155, 5, 10, 0,
		0, 155, 25, 1, 0, 0, 0, 16, 32, 49, 51, 57, 64, 68, 76, 85, 93, 103, 111,
		118, 130, 137, 146, 150,
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
	valuateParserCOLON             = 12
	valuateParserDOT               = 13
	valuateParserLOGICAL_OR        = 14
	valuateParserLOGICAL_AND       = 15
	valuateParserEQUALS            = 16
	valuateParserNOT_EQUALS        = 17
	valuateParserLESS              = 18
	valuateParserLESS_OR_EQUALS    = 19
	valuateParserGREATER           = 20
	valuateParserGREATER_OR_EQUALS = 21
	valuateParserDIV               = 22
	valuateParserPLUS              = 23
	valuateParserMINUS             = 24
	valuateParserSTAR              = 25
	valuateParserMODULUS           = 26
	valuateParserSKIP_             = 27
	valuateParserWHITESPACE        = 28
	valuateParserIDENTIFIER        = 29
	valuateParserSTRING_LITERAL    = 30
	valuateParserBYTES_LITERAL     = 31
	valuateParserINT               = 32
	valuateParserDECIMAL_LIT       = 33
	valuateParserFLOAT_NUMBER      = 34
	valuateParserSTRING            = 35
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
	valuateParserRULE_obj            = 8
	valuateParserRULE_pair           = 9
	valuateParserRULE_arr            = 10
	valuateParserRULE_variate        = 11
	valuateParserRULE_index          = 12
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
		p.SetState(26)
		p.expression(0)
	}
	{
		p.SetState(27)
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
	p.SetState(32)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case valuateParserTRUE, valuateParserFALSE, valuateParserNIL_LIT, valuateParserLP, valuateParserL_CURLY, valuateParserL_BRACKET, valuateParserIDENTIFIER, valuateParserINT, valuateParserFLOAT_NUMBER, valuateParserSTRING:
		{
			p.SetState(30)
			p.primaryExpr(0)
		}

	case valuateParserT__0, valuateParserMINUS:
		{
			p.SetState(31)
			p.UnaryExpr()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(49)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, valuateParserRULE_expression)
				p.SetState(34)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(35)
					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<valuateParserDIV)|(1<<valuateParserSTAR)|(1<<valuateParserMODULUS))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(36)
					p.expression(6)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, valuateParserRULE_expression)
				p.SetState(37)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(38)
					_la = p.GetTokenStream().LA(1)

					if !(_la == valuateParserPLUS || _la == valuateParserMINUS) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(39)
					p.expression(5)
				}

			case 3:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, valuateParserRULE_expression)
				p.SetState(40)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(41)
					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<valuateParserEQUALS)|(1<<valuateParserNOT_EQUALS)|(1<<valuateParserLESS)|(1<<valuateParserLESS_OR_EQUALS)|(1<<valuateParserGREATER)|(1<<valuateParserGREATER_OR_EQUALS))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(42)
					p.expression(4)
				}

			case 4:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, valuateParserRULE_expression)
				p.SetState(43)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(44)
					p.Match(valuateParserLOGICAL_AND)
				}
				{
					p.SetState(45)
					p.expression(3)
				}

			case 5:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, valuateParserRULE_expression)
				p.SetState(46)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(47)
					p.Match(valuateParserLOGICAL_OR)
				}
				{
					p.SetState(48)
					p.expression(2)
				}

			}

		}
		p.SetState(53)
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

func (s *PrimaryExprContext) Variate() IVariateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariateContext)
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
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(55)
			p.Operand()
		}

	case 2:
		{
			p.SetState(56)
			p.Variate()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewPrimaryExprContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, valuateParserRULE_primaryExpr)
			p.SetState(59)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			p.SetState(64)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case valuateParserDOT:
				{
					p.SetState(60)
					p.Match(valuateParserDOT)
				}
				{
					p.SetState(61)
					p.Match(valuateParserIDENTIFIER)
				}

			case valuateParserL_BRACKET:
				{
					p.SetState(62)
					p.Index()
				}

			case valuateParserLP:
				{
					p.SetState(63)
					p.Arguments()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		}
		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
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
		p.SetState(71)
		_la = p.GetTokenStream().LA(1)

		if !(_la == valuateParserT__0 || _la == valuateParserMINUS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(72)
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
		p.SetState(74)
		p.Match(valuateParserLP)
	}
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<valuateParserT__0)|(1<<valuateParserTRUE)|(1<<valuateParserFALSE)|(1<<valuateParserNIL_LIT)|(1<<valuateParserLP)|(1<<valuateParserL_CURLY)|(1<<valuateParserL_BRACKET)|(1<<valuateParserMINUS)|(1<<valuateParserIDENTIFIER))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(valuateParserINT-32))|(1<<(valuateParserFLOAT_NUMBER-32))|(1<<(valuateParserSTRING-32)))) != 0) {
		{
			p.SetState(75)
			p.ExpressionList()
		}

	}
	{
		p.SetState(78)
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
		p.SetState(80)
		p.expression(0)
	}
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == valuateParserCOMMA {
		{
			p.SetState(81)
			p.Match(valuateParserCOMMA)
		}
		{
			p.SetState(82)
			p.expression(0)
		}

		p.SetState(87)
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

	p.SetState(93)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case valuateParserTRUE, valuateParserFALSE, valuateParserNIL_LIT, valuateParserL_CURLY, valuateParserL_BRACKET, valuateParserINT, valuateParserFLOAT_NUMBER, valuateParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(88)
			p.BasicLit()
		}

	case valuateParserLP:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(89)
			p.Match(valuateParserLP)
		}
		{
			p.SetState(90)
			p.expression(0)
		}
		{
			p.SetState(91)
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

func (s *BasicLitContext) STRING() antlr.TerminalNode {
	return s.GetToken(valuateParserSTRING, 0)
}

func (s *BasicLitContext) INT() antlr.TerminalNode {
	return s.GetToken(valuateParserINT, 0)
}

func (s *BasicLitContext) FLOAT_NUMBER() antlr.TerminalNode {
	return s.GetToken(valuateParserFLOAT_NUMBER, 0)
}

func (s *BasicLitContext) Obj() IObjContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjContext)
}

func (s *BasicLitContext) Arr() IArrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrContext)
}

func (s *BasicLitContext) TRUE() antlr.TerminalNode {
	return s.GetToken(valuateParserTRUE, 0)
}

func (s *BasicLitContext) FALSE() antlr.TerminalNode {
	return s.GetToken(valuateParserFALSE, 0)
}

func (s *BasicLitContext) NIL_LIT() antlr.TerminalNode {
	return s.GetToken(valuateParserNIL_LIT, 0)
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

	p.SetState(103)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case valuateParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(95)
			p.Match(valuateParserSTRING)
		}

	case valuateParserINT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(96)
			p.Match(valuateParserINT)
		}

	case valuateParserFLOAT_NUMBER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(97)
			p.Match(valuateParserFLOAT_NUMBER)
		}

	case valuateParserL_CURLY:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(98)
			p.Obj()
		}

	case valuateParserL_BRACKET:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(99)
			p.Arr()
		}

	case valuateParserTRUE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(100)
			p.Match(valuateParserTRUE)
		}

	case valuateParserFALSE:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(101)
			p.Match(valuateParserFALSE)
		}

	case valuateParserNIL_LIT:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(102)
			p.Match(valuateParserNIL_LIT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IObjContext is an interface to support dynamic dispatch.
type IObjContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjContext differentiates from other interfaces.
	IsObjContext()
}

type ObjContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjContext() *ObjContext {
	var p = new(ObjContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = valuateParserRULE_obj
	return p
}

func (*ObjContext) IsObjContext() {}

func NewObjContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjContext {
	var p = new(ObjContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = valuateParserRULE_obj

	return p
}

func (s *ObjContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(valuateParserL_CURLY, 0)
}

func (s *ObjContext) AllPair() []IPairContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPairContext); ok {
			len++
		}
	}

	tst := make([]IPairContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPairContext); ok {
			tst[i] = t.(IPairContext)
			i++
		}
	}

	return tst
}

func (s *ObjContext) Pair(i int) IPairContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPairContext); ok {
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

	return t.(IPairContext)
}

func (s *ObjContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(valuateParserR_CURLY, 0)
}

func (s *ObjContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(valuateParserCOMMA)
}

func (s *ObjContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(valuateParserCOMMA, i)
}

func (s *ObjContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case valuateVisitor:
		return t.VisitObj(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *valuateParser) Obj() (localctx IObjContext) {
	this := p
	_ = this

	localctx = NewObjContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, valuateParserRULE_obj)
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

	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(105)
			p.Match(valuateParserL_CURLY)
		}
		{
			p.SetState(106)
			p.Pair()
		}
		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == valuateParserCOMMA {
			{
				p.SetState(107)
				p.Match(valuateParserCOMMA)
			}
			{
				p.SetState(108)
				p.Pair()
			}

			p.SetState(113)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(114)
			p.Match(valuateParserR_CURLY)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(116)
			p.Match(valuateParserL_CURLY)
		}
		{
			p.SetState(117)
			p.Match(valuateParserR_CURLY)
		}

	}

	return localctx
}

// IPairContext is an interface to support dynamic dispatch.
type IPairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPairContext differentiates from other interfaces.
	IsPairContext()
}

type PairContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPairContext() *PairContext {
	var p = new(PairContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = valuateParserRULE_pair
	return p
}

func (*PairContext) IsPairContext() {}

func NewPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PairContext {
	var p = new(PairContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = valuateParserRULE_pair

	return p
}

func (s *PairContext) GetParser() antlr.Parser { return s.parser }

func (s *PairContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(valuateParserSTRING_LITERAL, 0)
}

func (s *PairContext) COLON() antlr.TerminalNode {
	return s.GetToken(valuateParserCOLON, 0)
}

func (s *PairContext) BasicLit() IBasicLitContext {
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

func (s *PairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PairContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case valuateVisitor:
		return t.VisitPair(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *valuateParser) Pair() (localctx IPairContext) {
	this := p
	_ = this

	localctx = NewPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, valuateParserRULE_pair)

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
		p.SetState(120)
		p.Match(valuateParserSTRING_LITERAL)
	}
	{
		p.SetState(121)
		p.Match(valuateParserCOLON)
	}
	{
		p.SetState(122)
		p.BasicLit()
	}

	return localctx
}

// IArrContext is an interface to support dynamic dispatch.
type IArrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrContext differentiates from other interfaces.
	IsArrContext()
}

type ArrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrContext() *ArrContext {
	var p = new(ArrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = valuateParserRULE_arr
	return p
}

func (*ArrContext) IsArrContext() {}

func NewArrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrContext {
	var p = new(ArrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = valuateParserRULE_arr

	return p
}

func (s *ArrContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(valuateParserL_BRACKET, 0)
}

func (s *ArrContext) AllBasicLit() []IBasicLitContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBasicLitContext); ok {
			len++
		}
	}

	tst := make([]IBasicLitContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBasicLitContext); ok {
			tst[i] = t.(IBasicLitContext)
			i++
		}
	}

	return tst
}

func (s *ArrContext) BasicLit(i int) IBasicLitContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBasicLitContext); ok {
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

	return t.(IBasicLitContext)
}

func (s *ArrContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(valuateParserR_BRACKET, 0)
}

func (s *ArrContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(valuateParserCOMMA)
}

func (s *ArrContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(valuateParserCOMMA, i)
}

func (s *ArrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case valuateVisitor:
		return t.VisitArr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *valuateParser) Arr() (localctx IArrContext) {
	this := p
	_ = this

	localctx = NewArrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, valuateParserRULE_arr)
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

	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(124)
			p.Match(valuateParserL_BRACKET)
		}
		{
			p.SetState(125)
			p.BasicLit()
		}
		p.SetState(130)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == valuateParserCOMMA {
			{
				p.SetState(126)
				p.Match(valuateParserCOMMA)
			}
			{
				p.SetState(127)
				p.BasicLit()
			}

			p.SetState(132)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(133)
			p.Match(valuateParserR_BRACKET)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(135)
			p.Match(valuateParserL_BRACKET)
		}
		{
			p.SetState(136)
			p.Match(valuateParserR_BRACKET)
		}

	}

	return localctx
}

// IVariateContext is an interface to support dynamic dispatch.
type IVariateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariateContext differentiates from other interfaces.
	IsVariateContext()
}

type VariateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariateContext() *VariateContext {
	var p = new(VariateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = valuateParserRULE_variate
	return p
}

func (*VariateContext) IsVariateContext() {}

func NewVariateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariateContext {
	var p = new(VariateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = valuateParserRULE_variate

	return p
}

func (s *VariateContext) GetParser() antlr.Parser { return s.parser }

func (s *VariateContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(valuateParserIDENTIFIER)
}

func (s *VariateContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(valuateParserIDENTIFIER, i)
}

func (s *VariateContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(valuateParserL_CURLY, 0)
}

func (s *VariateContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(valuateParserR_CURLY, 0)
}

func (s *VariateContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(valuateParserDOT)
}

func (s *VariateContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(valuateParserDOT, i)
}

func (s *VariateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case valuateVisitor:
		return t.VisitVariate(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *valuateParser) Variate() (localctx IVariateContext) {
	this := p
	_ = this

	localctx = NewVariateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, valuateParserRULE_variate)
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

	p.SetState(150)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case valuateParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(139)
			p.Match(valuateParserIDENTIFIER)
		}

	case valuateParserL_CURLY:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(140)
			p.Match(valuateParserL_CURLY)
		}
		{
			p.SetState(141)
			p.Match(valuateParserIDENTIFIER)
		}
		p.SetState(146)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == valuateParserDOT {
			{
				p.SetState(142)
				p.Match(valuateParserDOT)
			}
			{
				p.SetState(143)
				p.Match(valuateParserIDENTIFIER)
			}

			p.SetState(148)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(149)
			p.Match(valuateParserR_CURLY)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.EnterRule(localctx, 24, valuateParserRULE_index)

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
		p.SetState(152)
		p.Match(valuateParserL_BRACKET)
	}
	{
		p.SetState(153)
		p.expression(0)
	}
	{
		p.SetState(154)
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
