// Code generated from /usr/local/lib/Go/212cf23d-32c6-47b2-a754-9f6fa59aba2f/1709267325888/AppBasedRule.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // AppBasedRule

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

type AppBasedRuleParser struct {
	*antlr.BaseParser
}

var appbasedruleParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func appbasedruleParserInit() {
	staticData := &appbasedruleParserStaticData
	staticData.literalNames = []string{
		"", "", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "GLUCOSE", "GLUCOSE_ACOSH", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "glucose", "glucose_acosh",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 5, 15, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 0, 0, 3, 0, 2, 4, 0, 0, 11, 0, 6, 1, 0, 0,
		0, 2, 10, 1, 0, 0, 0, 4, 12, 1, 0, 0, 0, 6, 7, 3, 2, 1, 0, 7, 8, 3, 4,
		2, 0, 8, 9, 5, 0, 0, 1, 9, 1, 1, 0, 0, 0, 10, 11, 5, 1, 0, 0, 11, 3, 1,
		0, 0, 0, 12, 13, 5, 2, 0, 0, 13, 5, 1, 0, 0, 0, 0,
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

// AppBasedRuleParserInit initializes any static state used to implement AppBasedRuleParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAppBasedRuleParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AppBasedRuleParserInit() {
	staticData := &appbasedruleParserStaticData
	staticData.once.Do(appbasedruleParserInit)
}

// NewAppBasedRuleParser produces a new parser instance for the optional input antlr.TokenStream.
func NewAppBasedRuleParser(input antlr.TokenStream) *AppBasedRuleParser {
	AppBasedRuleParserInit()
	this := new(AppBasedRuleParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &appbasedruleParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "AppBasedRule.g4"

	return this
}

// AppBasedRuleParser tokens.
const (
	AppBasedRuleParserEOF           = antlr.TokenEOF
	AppBasedRuleParserGLUCOSE       = 1
	AppBasedRuleParserGLUCOSE_ACOSH = 2
	AppBasedRuleParserOWNKEY        = 3
	AppBasedRuleParserSPLITKEY      = 4
	AppBasedRuleParserWS            = 5
)

// AppBasedRuleParser rules.
const (
	AppBasedRuleParserRULE_expression    = 0
	AppBasedRuleParserRULE_glucose       = 1
	AppBasedRuleParserRULE_glucose_acosh = 2
)

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
	p.RuleIndex = AppBasedRuleParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AppBasedRuleParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Glucose() IGlucoseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGlucoseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGlucoseContext)
}

func (s *ExpressionContext) Glucose_acosh() IGlucose_acoshContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGlucose_acoshContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGlucose_acoshContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(AppBasedRuleParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AppBasedRuleListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AppBasedRuleListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *AppBasedRuleParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AppBasedRuleParserRULE_expression)

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
		p.SetState(6)
		p.Glucose()
	}
	{
		p.SetState(7)
		p.Glucose_acosh()
	}
	{
		p.SetState(8)
		p.Match(AppBasedRuleParserEOF)
	}

	return localctx
}

// IGlucoseContext is an interface to support dynamic dispatch.
type IGlucoseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGlucoseContext differentiates from other interfaces.
	IsGlucoseContext()
}

type GlucoseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGlucoseContext() *GlucoseContext {
	var p = new(GlucoseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AppBasedRuleParserRULE_glucose
	return p
}

func (*GlucoseContext) IsGlucoseContext() {}

func NewGlucoseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GlucoseContext {
	var p = new(GlucoseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AppBasedRuleParserRULE_glucose

	return p
}

func (s *GlucoseContext) GetParser() antlr.Parser { return s.parser }

func (s *GlucoseContext) GLUCOSE() antlr.TerminalNode {
	return s.GetToken(AppBasedRuleParserGLUCOSE, 0)
}

func (s *GlucoseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GlucoseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GlucoseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AppBasedRuleListener); ok {
		listenerT.EnterGlucose(s)
	}
}

func (s *GlucoseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AppBasedRuleListener); ok {
		listenerT.ExitGlucose(s)
	}
}

func (p *AppBasedRuleParser) Glucose() (localctx IGlucoseContext) {
	this := p
	_ = this

	localctx = NewGlucoseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AppBasedRuleParserRULE_glucose)

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
		p.SetState(10)
		p.Match(AppBasedRuleParserGLUCOSE)
	}

	return localctx
}

// IGlucose_acoshContext is an interface to support dynamic dispatch.
type IGlucose_acoshContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGlucose_acoshContext differentiates from other interfaces.
	IsGlucose_acoshContext()
}

type Glucose_acoshContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGlucose_acoshContext() *Glucose_acoshContext {
	var p = new(Glucose_acoshContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AppBasedRuleParserRULE_glucose_acosh
	return p
}

func (*Glucose_acoshContext) IsGlucose_acoshContext() {}

func NewGlucose_acoshContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Glucose_acoshContext {
	var p = new(Glucose_acoshContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AppBasedRuleParserRULE_glucose_acosh

	return p
}

func (s *Glucose_acoshContext) GetParser() antlr.Parser { return s.parser }

func (s *Glucose_acoshContext) GLUCOSE_ACOSH() antlr.TerminalNode {
	return s.GetToken(AppBasedRuleParserGLUCOSE_ACOSH, 0)
}

func (s *Glucose_acoshContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Glucose_acoshContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Glucose_acoshContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AppBasedRuleListener); ok {
		listenerT.EnterGlucose_acosh(s)
	}
}

func (s *Glucose_acoshContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AppBasedRuleListener); ok {
		listenerT.ExitGlucose_acosh(s)
	}
}

func (p *AppBasedRuleParser) Glucose_acosh() (localctx IGlucose_acoshContext) {
	this := p
	_ = this

	localctx = NewGlucose_acoshContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, AppBasedRuleParserRULE_glucose_acosh)

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
		p.SetState(12)
		p.Match(AppBasedRuleParserGLUCOSE_ACOSH)
	}

	return localctx
}
