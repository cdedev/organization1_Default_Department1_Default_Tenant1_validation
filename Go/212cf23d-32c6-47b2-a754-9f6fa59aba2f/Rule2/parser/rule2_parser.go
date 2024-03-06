// Code generated from /usr/local/lib/Go/212cf23d-32c6-47b2-a754-9f6fa59aba2f/1709693485583/Rule2.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Rule2

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

type Rule2Parser struct {
	*antlr.BaseParser
}

var rule2ParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func rule2ParserInit() {
	staticData := &rule2ParserStaticData
	staticData.literalNames = []string{
		"", "", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "EDUCATION", "FOR", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "education", "for",
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

// Rule2ParserInit initializes any static state used to implement Rule2Parser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewRule2Parser(). You can call this function if you wish to initialize the static state ahead
// of time.
func Rule2ParserInit() {
	staticData := &rule2ParserStaticData
	staticData.once.Do(rule2ParserInit)
}

// NewRule2Parser produces a new parser instance for the optional input antlr.TokenStream.
func NewRule2Parser(input antlr.TokenStream) *Rule2Parser {
	Rule2ParserInit()
	this := new(Rule2Parser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &rule2ParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Rule2.g4"

	return this
}

// Rule2Parser tokens.
const (
	Rule2ParserEOF       = antlr.TokenEOF
	Rule2ParserEDUCATION = 1
	Rule2ParserFOR       = 2
	Rule2ParserOWNKEY    = 3
	Rule2ParserSPLITKEY  = 4
	Rule2ParserWS        = 5
)

// Rule2Parser rules.
const (
	Rule2ParserRULE_expression = 0
	Rule2ParserRULE_education  = 1
	Rule2ParserRULE_for        = 2
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
	p.RuleIndex = Rule2ParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Rule2ParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Education() IEducationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEducationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEducationContext)
}

func (s *ExpressionContext) For() IForContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(Rule2ParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Rule2Listener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Rule2Listener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *Rule2Parser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Rule2ParserRULE_expression)

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
		p.Education()
	}
	{
		p.SetState(7)
		p.For()
	}
	{
		p.SetState(8)
		p.Match(Rule2ParserEOF)
	}

	return localctx
}

// IEducationContext is an interface to support dynamic dispatch.
type IEducationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEducationContext differentiates from other interfaces.
	IsEducationContext()
}

type EducationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEducationContext() *EducationContext {
	var p = new(EducationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Rule2ParserRULE_education
	return p
}

func (*EducationContext) IsEducationContext() {}

func NewEducationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EducationContext {
	var p = new(EducationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Rule2ParserRULE_education

	return p
}

func (s *EducationContext) GetParser() antlr.Parser { return s.parser }

func (s *EducationContext) EDUCATION() antlr.TerminalNode {
	return s.GetToken(Rule2ParserEDUCATION, 0)
}

func (s *EducationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EducationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EducationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Rule2Listener); ok {
		listenerT.EnterEducation(s)
	}
}

func (s *EducationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Rule2Listener); ok {
		listenerT.ExitEducation(s)
	}
}

func (p *Rule2Parser) Education() (localctx IEducationContext) {
	this := p
	_ = this

	localctx = NewEducationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, Rule2ParserRULE_education)

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
		p.Match(Rule2ParserEDUCATION)
	}

	return localctx
}

// IForContext is an interface to support dynamic dispatch.
type IForContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsForContext differentiates from other interfaces.
	IsForContext()
}

type ForContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForContext() *ForContext {
	var p = new(ForContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Rule2ParserRULE_for
	return p
}

func (*ForContext) IsForContext() {}

func NewForContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForContext {
	var p = new(ForContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Rule2ParserRULE_for

	return p
}

func (s *ForContext) GetParser() antlr.Parser { return s.parser }

func (s *ForContext) FOR() antlr.TerminalNode {
	return s.GetToken(Rule2ParserFOR, 0)
}

func (s *ForContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Rule2Listener); ok {
		listenerT.EnterFor(s)
	}
}

func (s *ForContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Rule2Listener); ok {
		listenerT.ExitFor(s)
	}
}

func (p *Rule2Parser) For() (localctx IForContext) {
	this := p
	_ = this

	localctx = NewForContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, Rule2ParserRULE_for)

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
		p.Match(Rule2ParserFOR)
	}

	return localctx
}
