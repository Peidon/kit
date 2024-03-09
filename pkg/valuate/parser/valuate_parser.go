// Code generated from ValuateParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ValuateParser

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

type ValuateParser struct {
	*antlr.BaseParser
}

var valuateparserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func valuateparserParserInit() {
	staticData := &valuateparserParserStaticData
	staticData.literalNames = []string{
		"", "'true'", "'false'", "'nil'", "'('", "')'", "'{'", "", "'['", "']'",
		"','", "':'", "'.'", "'='", "'||'", "'&&'", "'=='", "'!='", "'<'", "'<='",
		"'>'", "'>='", "'!'", "'/'", "'+'", "'-'", "'*'", "'%'", "", "", "",
		"'${'",
	}
	staticData.symbolicNames = []string{
		"", "TRUE", "FALSE", "NIL_LIT", "LP", "RP", "L_CURLY", "R_CURLY", "L_BRACKET",
		"R_BRACKET", "COMMA", "COLON", "DOT", "ASSIGN", "LOGICAL_OR", "LOGICAL_AND",
		"EQUALS", "NOT_EQUALS", "LESS", "LESS_OR_EQUALS", "GREATER", "GREATER_OR_EQUALS",
		"EXCLAMATION", "DIV", "PLUS", "MINUS", "STAR", "MODULUS", "SKIP_", "WHITESPACE",
		"IDENTIFIER", "VARKEY_OPEN", "STRING", "STRING_LITERAL", "BYTES_LITERAL",
		"INT", "DECIMAL_LIT", "FLOAT_NUMBER", "VARID", "VARKEY_CLOSE",
	}
	staticData.ruleNames = []string{
		"plan", "assignment", "expression", "primaryExpr", "unaryExpr", "arguments",
		"expressionList", "variate", "operand", "basicLit", "obj", "pair", "arr",
		"index",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 39, 165, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 1, 0, 1, 0, 1, 0, 1, 1, 1,
		1, 1, 1, 5, 1, 35, 8, 1, 10, 1, 12, 1, 38, 9, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 2, 1, 2, 1, 2, 3, 2, 47, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 64, 8, 2, 10,
		2, 12, 2, 67, 9, 2, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 73, 8, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 3, 3, 79, 8, 3, 5, 3, 81, 8, 3, 10, 3, 12, 3, 84, 9, 3, 1,
		4, 1, 4, 1, 4, 1, 5, 1, 5, 3, 5, 91, 8, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6,
		5, 6, 98, 8, 6, 10, 6, 12, 6, 101, 9, 6, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7,
		107, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 115, 8, 8, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 125, 8, 9, 1, 10, 1, 10, 1,
		10, 1, 10, 5, 10, 131, 8, 10, 10, 10, 12, 10, 134, 9, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 3, 10, 140, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1,
		12, 1, 12, 1, 12, 5, 12, 150, 8, 12, 10, 12, 12, 12, 153, 9, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 3, 12, 159, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 0, 2, 4, 6, 14, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26,
		0, 4, 2, 0, 23, 23, 26, 27, 1, 0, 24, 25, 1, 0, 16, 21, 2, 0, 22, 22, 25,
		25, 176, 0, 28, 1, 0, 0, 0, 2, 31, 1, 0, 0, 0, 4, 46, 1, 0, 0, 0, 6, 72,
		1, 0, 0, 0, 8, 85, 1, 0, 0, 0, 10, 88, 1, 0, 0, 0, 12, 94, 1, 0, 0, 0,
		14, 106, 1, 0, 0, 0, 16, 114, 1, 0, 0, 0, 18, 124, 1, 0, 0, 0, 20, 139,
		1, 0, 0, 0, 22, 141, 1, 0, 0, 0, 24, 158, 1, 0, 0, 0, 26, 160, 1, 0, 0,
		0, 28, 29, 3, 4, 2, 0, 29, 30, 5, 0, 0, 1, 30, 1, 1, 0, 0, 0, 31, 36, 3,
		14, 7, 0, 32, 33, 5, 12, 0, 0, 33, 35, 5, 30, 0, 0, 34, 32, 1, 0, 0, 0,
		35, 38, 1, 0, 0, 0, 36, 34, 1, 0, 0, 0, 36, 37, 1, 0, 0, 0, 37, 39, 1,
		0, 0, 0, 38, 36, 1, 0, 0, 0, 39, 40, 5, 13, 0, 0, 40, 41, 3, 4, 2, 0, 41,
		42, 5, 0, 0, 1, 42, 3, 1, 0, 0, 0, 43, 44, 6, 2, -1, 0, 44, 47, 3, 6, 3,
		0, 45, 47, 3, 8, 4, 0, 46, 43, 1, 0, 0, 0, 46, 45, 1, 0, 0, 0, 47, 65,
		1, 0, 0, 0, 48, 49, 10, 5, 0, 0, 49, 50, 7, 0, 0, 0, 50, 64, 3, 4, 2, 6,
		51, 52, 10, 4, 0, 0, 52, 53, 7, 1, 0, 0, 53, 64, 3, 4, 2, 5, 54, 55, 10,
		3, 0, 0, 55, 56, 7, 2, 0, 0, 56, 64, 3, 4, 2, 4, 57, 58, 10, 2, 0, 0, 58,
		59, 5, 15, 0, 0, 59, 64, 3, 4, 2, 3, 60, 61, 10, 1, 0, 0, 61, 62, 5, 14,
		0, 0, 62, 64, 3, 4, 2, 2, 63, 48, 1, 0, 0, 0, 63, 51, 1, 0, 0, 0, 63, 54,
		1, 0, 0, 0, 63, 57, 1, 0, 0, 0, 63, 60, 1, 0, 0, 0, 64, 67, 1, 0, 0, 0,
		65, 63, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 5, 1, 0, 0, 0, 67, 65, 1, 0,
		0, 0, 68, 69, 6, 3, -1, 0, 69, 73, 3, 16, 8, 0, 70, 71, 5, 30, 0, 0, 71,
		73, 3, 10, 5, 0, 72, 68, 1, 0, 0, 0, 72, 70, 1, 0, 0, 0, 73, 82, 1, 0,
		0, 0, 74, 78, 10, 2, 0, 0, 75, 76, 5, 12, 0, 0, 76, 79, 5, 30, 0, 0, 77,
		79, 3, 26, 13, 0, 78, 75, 1, 0, 0, 0, 78, 77, 1, 0, 0, 0, 79, 81, 1, 0,
		0, 0, 80, 74, 1, 0, 0, 0, 81, 84, 1, 0, 0, 0, 82, 80, 1, 0, 0, 0, 82, 83,
		1, 0, 0, 0, 83, 7, 1, 0, 0, 0, 84, 82, 1, 0, 0, 0, 85, 86, 7, 3, 0, 0,
		86, 87, 3, 4, 2, 0, 87, 9, 1, 0, 0, 0, 88, 90, 5, 4, 0, 0, 89, 91, 3, 12,
		6, 0, 90, 89, 1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 93,
		5, 5, 0, 0, 93, 11, 1, 0, 0, 0, 94, 99, 3, 4, 2, 0, 95, 96, 5, 10, 0, 0,
		96, 98, 3, 4, 2, 0, 97, 95, 1, 0, 0, 0, 98, 101, 1, 0, 0, 0, 99, 97, 1,
		0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 13, 1, 0, 0, 0, 101, 99, 1, 0, 0, 0,
		102, 103, 5, 31, 0, 0, 103, 104, 5, 38, 0, 0, 104, 107, 5, 39, 0, 0, 105,
		107, 5, 30, 0, 0, 106, 102, 1, 0, 0, 0, 106, 105, 1, 0, 0, 0, 107, 15,
		1, 0, 0, 0, 108, 115, 3, 18, 9, 0, 109, 115, 3, 14, 7, 0, 110, 111, 5,
		4, 0, 0, 111, 112, 3, 4, 2, 0, 112, 113, 5, 5, 0, 0, 113, 115, 1, 0, 0,
		0, 114, 108, 1, 0, 0, 0, 114, 109, 1, 0, 0, 0, 114, 110, 1, 0, 0, 0, 115,
		17, 1, 0, 0, 0, 116, 125, 5, 32, 0, 0, 117, 125, 5, 35, 0, 0, 118, 125,
		5, 37, 0, 0, 119, 125, 3, 20, 10, 0, 120, 125, 3, 24, 12, 0, 121, 125,
		5, 1, 0, 0, 122, 125, 5, 2, 0, 0, 123, 125, 5, 3, 0, 0, 124, 116, 1, 0,
		0, 0, 124, 117, 1, 0, 0, 0, 124, 118, 1, 0, 0, 0, 124, 119, 1, 0, 0, 0,
		124, 120, 1, 0, 0, 0, 124, 121, 1, 0, 0, 0, 124, 122, 1, 0, 0, 0, 124,
		123, 1, 0, 0, 0, 125, 19, 1, 0, 0, 0, 126, 127, 5, 6, 0, 0, 127, 132, 3,
		22, 11, 0, 128, 129, 5, 10, 0, 0, 129, 131, 3, 22, 11, 0, 130, 128, 1,
		0, 0, 0, 131, 134, 1, 0, 0, 0, 132, 130, 1, 0, 0, 0, 132, 133, 1, 0, 0,
		0, 133, 135, 1, 0, 0, 0, 134, 132, 1, 0, 0, 0, 135, 136, 5, 7, 0, 0, 136,
		140, 1, 0, 0, 0, 137, 138, 5, 6, 0, 0, 138, 140, 5, 7, 0, 0, 139, 126,
		1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 140, 21, 1, 0, 0, 0, 141, 142, 5, 32,
		0, 0, 142, 143, 5, 11, 0, 0, 143, 144, 3, 18, 9, 0, 144, 23, 1, 0, 0, 0,
		145, 146, 5, 8, 0, 0, 146, 151, 3, 18, 9, 0, 147, 148, 5, 10, 0, 0, 148,
		150, 3, 18, 9, 0, 149, 147, 1, 0, 0, 0, 150, 153, 1, 0, 0, 0, 151, 149,
		1, 0, 0, 0, 151, 152, 1, 0, 0, 0, 152, 154, 1, 0, 0, 0, 153, 151, 1, 0,
		0, 0, 154, 155, 5, 9, 0, 0, 155, 159, 1, 0, 0, 0, 156, 157, 5, 8, 0, 0,
		157, 159, 5, 9, 0, 0, 158, 145, 1, 0, 0, 0, 158, 156, 1, 0, 0, 0, 159,
		25, 1, 0, 0, 0, 160, 161, 5, 8, 0, 0, 161, 162, 3, 4, 2, 0, 162, 163, 5,
		9, 0, 0, 163, 27, 1, 0, 0, 0, 16, 36, 46, 63, 65, 72, 78, 82, 90, 99, 106,
		114, 124, 132, 139, 151, 158,
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

// ValuateParserInit initializes any static state used to implement ValuateParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewValuateParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ValuateParserInit() {
	staticData := &valuateparserParserStaticData
	staticData.once.Do(valuateparserParserInit)
}

// NewValuateParser produces a new parser instance for the optional input antlr.TokenStream.
func NewValuateParser(input antlr.TokenStream) *ValuateParser {
	ValuateParserInit()
	this := new(ValuateParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &valuateparserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "ValuateParser.g4"

	return this
}

// ValuateParser tokens.
const (
	ValuateParserEOF               = antlr.TokenEOF
	ValuateParserTRUE              = 1
	ValuateParserFALSE             = 2
	ValuateParserNIL_LIT           = 3
	ValuateParserLP                = 4
	ValuateParserRP                = 5
	ValuateParserL_CURLY           = 6
	ValuateParserR_CURLY           = 7
	ValuateParserL_BRACKET         = 8
	ValuateParserR_BRACKET         = 9
	ValuateParserCOMMA             = 10
	ValuateParserCOLON             = 11
	ValuateParserDOT               = 12
	ValuateParserASSIGN            = 13
	ValuateParserLOGICAL_OR        = 14
	ValuateParserLOGICAL_AND       = 15
	ValuateParserEQUALS            = 16
	ValuateParserNOT_EQUALS        = 17
	ValuateParserLESS              = 18
	ValuateParserLESS_OR_EQUALS    = 19
	ValuateParserGREATER           = 20
	ValuateParserGREATER_OR_EQUALS = 21
	ValuateParserEXCLAMATION       = 22
	ValuateParserDIV               = 23
	ValuateParserPLUS              = 24
	ValuateParserMINUS             = 25
	ValuateParserSTAR              = 26
	ValuateParserMODULUS           = 27
	ValuateParserSKIP_             = 28
	ValuateParserWHITESPACE        = 29
	ValuateParserIDENTIFIER        = 30
	ValuateParserVARKEY_OPEN       = 31
	ValuateParserSTRING            = 32
	ValuateParserSTRING_LITERAL    = 33
	ValuateParserBYTES_LITERAL     = 34
	ValuateParserINT               = 35
	ValuateParserDECIMAL_LIT       = 36
	ValuateParserFLOAT_NUMBER      = 37
	ValuateParserVARID             = 38
	ValuateParserVARKEY_CLOSE      = 39
)

// ValuateParser rules.
const (
	ValuateParserRULE_plan           = 0
	ValuateParserRULE_assignment     = 1
	ValuateParserRULE_expression     = 2
	ValuateParserRULE_primaryExpr    = 3
	ValuateParserRULE_unaryExpr      = 4
	ValuateParserRULE_arguments      = 5
	ValuateParserRULE_expressionList = 6
	ValuateParserRULE_variate        = 7
	ValuateParserRULE_operand        = 8
	ValuateParserRULE_basicLit       = 9
	ValuateParserRULE_obj            = 10
	ValuateParserRULE_pair           = 11
	ValuateParserRULE_arr            = 12
	ValuateParserRULE_index          = 13
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
	p.RuleIndex = ValuateParserRULE_plan
	return p
}

func (*PlanContext) IsPlanContext() {}

func NewPlanContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PlanContext {
	var p = new(PlanContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ValuateParserRULE_plan

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
	return s.GetToken(ValuateParserEOF, 0)
}

func (s *PlanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PlanContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PlanContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ValuateParserVisitor:
		return t.VisitPlan(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ValuateParser) Plan() (localctx IPlanContext) {
	this := p
	_ = this

	localctx = NewPlanContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ValuateParserRULE_plan)

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
		p.SetState(28)
		p.expression(0)
	}
	{
		p.SetState(29)
		p.Match(ValuateParserEOF)
	}

	return localctx
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ValuateParserRULE_assignment
	return p
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ValuateParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) Variate() IVariateContext {
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

func (s *AssignmentContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(ValuateParserASSIGN, 0)
}

func (s *AssignmentContext) Expression() IExpressionContext {
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

func (s *AssignmentContext) EOF() antlr.TerminalNode {
	return s.GetToken(ValuateParserEOF, 0)
}

func (s *AssignmentContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(ValuateParserDOT)
}

func (s *AssignmentContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(ValuateParserDOT, i)
}

func (s *AssignmentContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(ValuateParserIDENTIFIER)
}

func (s *AssignmentContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(ValuateParserIDENTIFIER, i)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ValuateParserVisitor:
		return t.VisitAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ValuateParser) Assignment() (localctx IAssignmentContext) {
	this := p
	_ = this

	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ValuateParserRULE_assignment)
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
		p.SetState(31)
		p.Variate()
	}
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ValuateParserDOT {
		{
			p.SetState(32)
			p.Match(ValuateParserDOT)
		}
		{
			p.SetState(33)
			p.Match(ValuateParserIDENTIFIER)
		}

		p.SetState(38)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(39)
		p.Match(ValuateParserASSIGN)
	}
	{
		p.SetState(40)
		p.expression(0)
	}
	{
		p.SetState(41)
		p.Match(ValuateParserEOF)
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
	p.RuleIndex = ValuateParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ValuateParserRULE_expression

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
	return s.GetToken(ValuateParserSTAR, 0)
}

func (s *ExpressionContext) DIV() antlr.TerminalNode {
	return s.GetToken(ValuateParserDIV, 0)
}

func (s *ExpressionContext) MODULUS() antlr.TerminalNode {
	return s.GetToken(ValuateParserMODULUS, 0)
}

func (s *ExpressionContext) PLUS() antlr.TerminalNode {
	return s.GetToken(ValuateParserPLUS, 0)
}

func (s *ExpressionContext) MINUS() antlr.TerminalNode {
	return s.GetToken(ValuateParserMINUS, 0)
}

func (s *ExpressionContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(ValuateParserEQUALS, 0)
}

func (s *ExpressionContext) NOT_EQUALS() antlr.TerminalNode {
	return s.GetToken(ValuateParserNOT_EQUALS, 0)
}

func (s *ExpressionContext) LESS() antlr.TerminalNode {
	return s.GetToken(ValuateParserLESS, 0)
}

func (s *ExpressionContext) LESS_OR_EQUALS() antlr.TerminalNode {
	return s.GetToken(ValuateParserLESS_OR_EQUALS, 0)
}

func (s *ExpressionContext) GREATER() antlr.TerminalNode {
	return s.GetToken(ValuateParserGREATER, 0)
}

func (s *ExpressionContext) GREATER_OR_EQUALS() antlr.TerminalNode {
	return s.GetToken(ValuateParserGREATER_OR_EQUALS, 0)
}

func (s *ExpressionContext) LOGICAL_AND() antlr.TerminalNode {
	return s.GetToken(ValuateParserLOGICAL_AND, 0)
}

func (s *ExpressionContext) LOGICAL_OR() antlr.TerminalNode {
	return s.GetToken(ValuateParserLOGICAL_OR, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ValuateParserVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ValuateParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *ValuateParser) expression(_p int) (localctx IExpressionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, ValuateParserRULE_expression, _p)
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
	p.SetState(46)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ValuateParserTRUE, ValuateParserFALSE, ValuateParserNIL_LIT, ValuateParserLP, ValuateParserL_CURLY, ValuateParserL_BRACKET, ValuateParserIDENTIFIER, ValuateParserVARKEY_OPEN, ValuateParserSTRING, ValuateParserINT, ValuateParserFLOAT_NUMBER:
		{
			p.SetState(44)
			p.primaryExpr(0)
		}

	case ValuateParserEXCLAMATION, ValuateParserMINUS:
		{
			p.SetState(45)
			p.UnaryExpr()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(63)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ValuateParserRULE_expression)
				p.SetState(48)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(49)
					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ValuateParserDIV)|(1<<ValuateParserSTAR)|(1<<ValuateParserMODULUS))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(50)
					p.expression(6)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ValuateParserRULE_expression)
				p.SetState(51)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(52)
					_la = p.GetTokenStream().LA(1)

					if !(_la == ValuateParserPLUS || _la == ValuateParserMINUS) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(53)
					p.expression(5)
				}

			case 3:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ValuateParserRULE_expression)
				p.SetState(54)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(55)
					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ValuateParserEQUALS)|(1<<ValuateParserNOT_EQUALS)|(1<<ValuateParserLESS)|(1<<ValuateParserLESS_OR_EQUALS)|(1<<ValuateParserGREATER)|(1<<ValuateParserGREATER_OR_EQUALS))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(56)
					p.expression(4)
				}

			case 4:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ValuateParserRULE_expression)
				p.SetState(57)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(58)
					p.Match(ValuateParserLOGICAL_AND)
				}
				{
					p.SetState(59)
					p.expression(3)
				}

			case 5:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ValuateParserRULE_expression)
				p.SetState(60)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(61)
					p.Match(ValuateParserLOGICAL_OR)
				}
				{
					p.SetState(62)
					p.expression(2)
				}

			}

		}
		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
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
	p.RuleIndex = ValuateParserRULE_primaryExpr
	return p
}

func (*PrimaryExprContext) IsPrimaryExprContext() {}

func NewPrimaryExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExprContext {
	var p = new(PrimaryExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ValuateParserRULE_primaryExpr

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

func (s *PrimaryExprContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ValuateParserIDENTIFIER, 0)
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
	return s.GetToken(ValuateParserDOT, 0)
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

func (s *PrimaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ValuateParserVisitor:
		return t.VisitPrimaryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ValuateParser) PrimaryExpr() (localctx IPrimaryExprContext) {
	return p.primaryExpr(0)
}

func (p *ValuateParser) primaryExpr(_p int) (localctx IPrimaryExprContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewPrimaryExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPrimaryExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 6
	p.EnterRecursionRule(localctx, 6, ValuateParserRULE_primaryExpr, _p)

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
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(69)
			p.Operand()
		}

	case 2:
		{
			p.SetState(70)
			p.Match(ValuateParserIDENTIFIER)
		}
		{
			p.SetState(71)
			p.Arguments()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewPrimaryExprContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ValuateParserRULE_primaryExpr)
			p.SetState(74)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			p.SetState(78)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case ValuateParserDOT:
				{
					p.SetState(75)
					p.Match(ValuateParserDOT)
				}
				{
					p.SetState(76)
					p.Match(ValuateParserIDENTIFIER)
				}

			case ValuateParserL_BRACKET:
				{
					p.SetState(77)
					p.Index()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		}
		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
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
	p.RuleIndex = ValuateParserRULE_unaryExpr
	return p
}

func (*UnaryExprContext) IsUnaryExprContext() {}

func NewUnaryExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryExprContext {
	var p = new(UnaryExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ValuateParserRULE_unaryExpr

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
	return s.GetToken(ValuateParserMINUS, 0)
}

func (s *UnaryExprContext) EXCLAMATION() antlr.TerminalNode {
	return s.GetToken(ValuateParserEXCLAMATION, 0)
}

func (s *UnaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ValuateParserVisitor:
		return t.VisitUnaryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ValuateParser) UnaryExpr() (localctx IUnaryExprContext) {
	this := p
	_ = this

	localctx = NewUnaryExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ValuateParserRULE_unaryExpr)
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
		p.SetState(85)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ValuateParserEXCLAMATION || _la == ValuateParserMINUS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(86)
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
	p.RuleIndex = ValuateParserRULE_arguments
	return p
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ValuateParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) LP() antlr.TerminalNode {
	return s.GetToken(ValuateParserLP, 0)
}

func (s *ArgumentsContext) RP() antlr.TerminalNode {
	return s.GetToken(ValuateParserRP, 0)
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
	case ValuateParserVisitor:
		return t.VisitArguments(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ValuateParser) Arguments() (localctx IArgumentsContext) {
	this := p
	_ = this

	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ValuateParserRULE_arguments)
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
		p.SetState(88)
		p.Match(ValuateParserLP)
	}
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ValuateParserTRUE)|(1<<ValuateParserFALSE)|(1<<ValuateParserNIL_LIT)|(1<<ValuateParserLP)|(1<<ValuateParserL_CURLY)|(1<<ValuateParserL_BRACKET)|(1<<ValuateParserEXCLAMATION)|(1<<ValuateParserMINUS)|(1<<ValuateParserIDENTIFIER)|(1<<ValuateParserVARKEY_OPEN))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(ValuateParserSTRING-32))|(1<<(ValuateParserINT-32))|(1<<(ValuateParserFLOAT_NUMBER-32)))) != 0) {
		{
			p.SetState(89)
			p.ExpressionList()
		}

	}
	{
		p.SetState(92)
		p.Match(ValuateParserRP)
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
	p.RuleIndex = ValuateParserRULE_expressionList
	return p
}

func (*ExpressionListContext) IsExpressionListContext() {}

func NewExpressionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionListContext {
	var p = new(ExpressionListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ValuateParserRULE_expressionList

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
	return s.GetTokens(ValuateParserCOMMA)
}

func (s *ExpressionListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ValuateParserCOMMA, i)
}

func (s *ExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ValuateParserVisitor:
		return t.VisitExpressionList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ValuateParser) ExpressionList() (localctx IExpressionListContext) {
	this := p
	_ = this

	localctx = NewExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ValuateParserRULE_expressionList)
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
		p.SetState(94)
		p.expression(0)
	}
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ValuateParserCOMMA {
		{
			p.SetState(95)
			p.Match(ValuateParserCOMMA)
		}
		{
			p.SetState(96)
			p.expression(0)
		}

		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.RuleIndex = ValuateParserRULE_variate
	return p
}

func (*VariateContext) IsVariateContext() {}

func NewVariateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariateContext {
	var p = new(VariateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ValuateParserRULE_variate

	return p
}

func (s *VariateContext) GetParser() antlr.Parser { return s.parser }

func (s *VariateContext) VARKEY_OPEN() antlr.TerminalNode {
	return s.GetToken(ValuateParserVARKEY_OPEN, 0)
}

func (s *VariateContext) VARID() antlr.TerminalNode {
	return s.GetToken(ValuateParserVARID, 0)
}

func (s *VariateContext) VARKEY_CLOSE() antlr.TerminalNode {
	return s.GetToken(ValuateParserVARKEY_CLOSE, 0)
}

func (s *VariateContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ValuateParserIDENTIFIER, 0)
}

func (s *VariateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ValuateParserVisitor:
		return t.VisitVariate(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ValuateParser) Variate() (localctx IVariateContext) {
	this := p
	_ = this

	localctx = NewVariateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ValuateParserRULE_variate)

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

	p.SetState(106)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ValuateParserVARKEY_OPEN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(102)
			p.Match(ValuateParserVARKEY_OPEN)
		}
		{
			p.SetState(103)
			p.Match(ValuateParserVARID)
		}
		{
			p.SetState(104)
			p.Match(ValuateParserVARKEY_CLOSE)
		}

	case ValuateParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(105)
			p.Match(ValuateParserIDENTIFIER)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.RuleIndex = ValuateParserRULE_operand
	return p
}

func (*OperandContext) IsOperandContext() {}

func NewOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperandContext {
	var p = new(OperandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ValuateParserRULE_operand

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

func (s *OperandContext) Variate() IVariateContext {
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

func (s *OperandContext) LP() antlr.TerminalNode {
	return s.GetToken(ValuateParserLP, 0)
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
	return s.GetToken(ValuateParserRP, 0)
}

func (s *OperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperandContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ValuateParserVisitor:
		return t.VisitOperand(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ValuateParser) Operand() (localctx IOperandContext) {
	this := p
	_ = this

	localctx = NewOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ValuateParserRULE_operand)

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

	p.SetState(114)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ValuateParserTRUE, ValuateParserFALSE, ValuateParserNIL_LIT, ValuateParserL_CURLY, ValuateParserL_BRACKET, ValuateParserSTRING, ValuateParserINT, ValuateParserFLOAT_NUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(108)
			p.BasicLit()
		}

	case ValuateParserIDENTIFIER, ValuateParserVARKEY_OPEN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(109)
			p.Variate()
		}

	case ValuateParserLP:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(110)
			p.Match(ValuateParserLP)
		}
		{
			p.SetState(111)
			p.expression(0)
		}
		{
			p.SetState(112)
			p.Match(ValuateParserRP)
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
	p.RuleIndex = ValuateParserRULE_basicLit
	return p
}

func (*BasicLitContext) IsBasicLitContext() {}

func NewBasicLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BasicLitContext {
	var p = new(BasicLitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ValuateParserRULE_basicLit

	return p
}

func (s *BasicLitContext) GetParser() antlr.Parser { return s.parser }

func (s *BasicLitContext) STRING() antlr.TerminalNode {
	return s.GetToken(ValuateParserSTRING, 0)
}

func (s *BasicLitContext) INT() antlr.TerminalNode {
	return s.GetToken(ValuateParserINT, 0)
}

func (s *BasicLitContext) FLOAT_NUMBER() antlr.TerminalNode {
	return s.GetToken(ValuateParserFLOAT_NUMBER, 0)
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
	return s.GetToken(ValuateParserTRUE, 0)
}

func (s *BasicLitContext) FALSE() antlr.TerminalNode {
	return s.GetToken(ValuateParserFALSE, 0)
}

func (s *BasicLitContext) NIL_LIT() antlr.TerminalNode {
	return s.GetToken(ValuateParserNIL_LIT, 0)
}

func (s *BasicLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BasicLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BasicLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ValuateParserVisitor:
		return t.VisitBasicLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ValuateParser) BasicLit() (localctx IBasicLitContext) {
	this := p
	_ = this

	localctx = NewBasicLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ValuateParserRULE_basicLit)

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

	p.SetState(124)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ValuateParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(116)
			p.Match(ValuateParserSTRING)
		}

	case ValuateParserINT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(117)
			p.Match(ValuateParserINT)
		}

	case ValuateParserFLOAT_NUMBER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(118)
			p.Match(ValuateParserFLOAT_NUMBER)
		}

	case ValuateParserL_CURLY:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(119)
			p.Obj()
		}

	case ValuateParserL_BRACKET:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(120)
			p.Arr()
		}

	case ValuateParserTRUE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(121)
			p.Match(ValuateParserTRUE)
		}

	case ValuateParserFALSE:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(122)
			p.Match(ValuateParserFALSE)
		}

	case ValuateParserNIL_LIT:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(123)
			p.Match(ValuateParserNIL_LIT)
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
	p.RuleIndex = ValuateParserRULE_obj
	return p
}

func (*ObjContext) IsObjContext() {}

func NewObjContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjContext {
	var p = new(ObjContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ValuateParserRULE_obj

	return p
}

func (s *ObjContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(ValuateParserL_CURLY, 0)
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
	return s.GetToken(ValuateParserR_CURLY, 0)
}

func (s *ObjContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ValuateParserCOMMA)
}

func (s *ObjContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ValuateParserCOMMA, i)
}

func (s *ObjContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ValuateParserVisitor:
		return t.VisitObj(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ValuateParser) Obj() (localctx IObjContext) {
	this := p
	_ = this

	localctx = NewObjContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ValuateParserRULE_obj)
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

	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(126)
			p.Match(ValuateParserL_CURLY)
		}
		{
			p.SetState(127)
			p.Pair()
		}
		p.SetState(132)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ValuateParserCOMMA {
			{
				p.SetState(128)
				p.Match(ValuateParserCOMMA)
			}
			{
				p.SetState(129)
				p.Pair()
			}

			p.SetState(134)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(135)
			p.Match(ValuateParserR_CURLY)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(137)
			p.Match(ValuateParserL_CURLY)
		}
		{
			p.SetState(138)
			p.Match(ValuateParserR_CURLY)
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
	p.RuleIndex = ValuateParserRULE_pair
	return p
}

func (*PairContext) IsPairContext() {}

func NewPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PairContext {
	var p = new(PairContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ValuateParserRULE_pair

	return p
}

func (s *PairContext) GetParser() antlr.Parser { return s.parser }

func (s *PairContext) STRING() antlr.TerminalNode {
	return s.GetToken(ValuateParserSTRING, 0)
}

func (s *PairContext) COLON() antlr.TerminalNode {
	return s.GetToken(ValuateParserCOLON, 0)
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
	case ValuateParserVisitor:
		return t.VisitPair(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ValuateParser) Pair() (localctx IPairContext) {
	this := p
	_ = this

	localctx = NewPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ValuateParserRULE_pair)

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
		p.SetState(141)
		p.Match(ValuateParserSTRING)
	}
	{
		p.SetState(142)
		p.Match(ValuateParserCOLON)
	}
	{
		p.SetState(143)
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
	p.RuleIndex = ValuateParserRULE_arr
	return p
}

func (*ArrContext) IsArrContext() {}

func NewArrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrContext {
	var p = new(ArrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ValuateParserRULE_arr

	return p
}

func (s *ArrContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(ValuateParserL_BRACKET, 0)
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
	return s.GetToken(ValuateParserR_BRACKET, 0)
}

func (s *ArrContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ValuateParserCOMMA)
}

func (s *ArrContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ValuateParserCOMMA, i)
}

func (s *ArrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ValuateParserVisitor:
		return t.VisitArr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ValuateParser) Arr() (localctx IArrContext) {
	this := p
	_ = this

	localctx = NewArrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ValuateParserRULE_arr)
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

	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(145)
			p.Match(ValuateParserL_BRACKET)
		}
		{
			p.SetState(146)
			p.BasicLit()
		}
		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ValuateParserCOMMA {
			{
				p.SetState(147)
				p.Match(ValuateParserCOMMA)
			}
			{
				p.SetState(148)
				p.BasicLit()
			}

			p.SetState(153)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(154)
			p.Match(ValuateParserR_BRACKET)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(156)
			p.Match(ValuateParserL_BRACKET)
		}
		{
			p.SetState(157)
			p.Match(ValuateParserR_BRACKET)
		}

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
	p.RuleIndex = ValuateParserRULE_index
	return p
}

func (*IndexContext) IsIndexContext() {}

func NewIndexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexContext {
	var p = new(IndexContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ValuateParserRULE_index

	return p
}

func (s *IndexContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(ValuateParserL_BRACKET, 0)
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
	return s.GetToken(ValuateParserR_BRACKET, 0)
}

func (s *IndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ValuateParserVisitor:
		return t.VisitIndex(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ValuateParser) Index() (localctx IIndexContext) {
	this := p
	_ = this

	localctx = NewIndexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ValuateParserRULE_index)

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
		p.SetState(160)
		p.Match(ValuateParserL_BRACKET)
	}
	{
		p.SetState(161)
		p.expression(0)
	}
	{
		p.SetState(162)
		p.Match(ValuateParserR_BRACKET)
	}

	return localctx
}

func (p *ValuateParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 3:
		var t *PrimaryExprContext = nil
		if localctx != nil {
			t = localctx.(*PrimaryExprContext)
		}
		return p.PrimaryExpr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ValuateParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
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

func (p *ValuateParser) PrimaryExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 5:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
