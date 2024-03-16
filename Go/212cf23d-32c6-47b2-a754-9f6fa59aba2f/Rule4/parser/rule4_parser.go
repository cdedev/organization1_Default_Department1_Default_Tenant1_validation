// Code generated from /usr/local/lib/Go/212cf23d-32c6-47b2-a754-9f6fa59aba2f/1710560357049/Rule4.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Rule4

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

type Rule4Parser struct {
	*antlr.BaseParser
}

var rule4ParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func rule4ParserInit() {
	staticData := &rule4ParserStaticData
	staticData.literalNames = []string{
		"", "", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "FIELD1", "FIELD2", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "field1", "field2",
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

// Rule4ParserInit initializes any static state used to implement Rule4Parser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewRule4Parser(). You can call this function if you wish to initialize the static state ahead
// of time.
func Rule4ParserInit() {
	staticData := &rule4ParserStaticData
	staticData.once.Do(rule4ParserInit)
}

// NewRule4Parser produces a new parser instance for the optional input antlr.TokenStream.
func NewRule4Parser(input antlr.TokenStream) *Rule4Parser {
	Rule4ParserInit()
	this := new(Rule4Parser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &rule4ParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Rule4.g4"

	return this
}

// Rule4Parser tokens.
const (
	Rule4ParserEOF      = antlr.TokenEOF
	Rule4ParserFIELD1   = 1
	Rule4ParserFIELD2   = 2
	Rule4ParserOWNKEY   = 3
	Rule4ParserSPLITKEY = 4
	Rule4ParserWS       = 5
)

// Rule4Parser rules.
const (
	Rule4ParserRULE_expression = 0
	Rule4ParserRULE_field1     = 1
	Rule4ParserRULE_field2     = 2
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
	p.RuleIndex = Rule4ParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Rule4ParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Field1() IField1Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IField1Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IField1Context)
}

func (s *ExpressionContext) Field2() IField2Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IField2Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IField2Context)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(Rule4ParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Rule4Listener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Rule4Listener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *Rule4Parser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Rule4ParserRULE_expression)

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
		p.Field1()
	}
	{
		p.SetState(7)
		p.Field2()
	}
	{
		p.SetState(8)
		p.Match(Rule4ParserEOF)
	}

	return localctx
}

// IField1Context is an interface to support dynamic dispatch.
type IField1Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsField1Context differentiates from other interfaces.
	IsField1Context()
}

type Field1Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField1Context() *Field1Context {
	var p = new(Field1Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Rule4ParserRULE_field1
	return p
}

func (*Field1Context) IsField1Context() {}

func NewField1Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field1Context {
	var p = new(Field1Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Rule4ParserRULE_field1

	return p
}

func (s *Field1Context) GetParser() antlr.Parser { return s.parser }

func (s *Field1Context) FIELD1() antlr.TerminalNode {
	return s.GetToken(Rule4ParserFIELD1, 0)
}

func (s *Field1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field1Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Rule4Listener); ok {
		listenerT.EnterField1(s)
	}
}

func (s *Field1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Rule4Listener); ok {
		listenerT.ExitField1(s)
	}
}

func (p *Rule4Parser) Field1() (localctx IField1Context) {
	this := p
	_ = this

	localctx = NewField1Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, Rule4ParserRULE_field1)

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
		p.Match(Rule4ParserFIELD1)
	}

	return localctx
}

// IField2Context is an interface to support dynamic dispatch.
type IField2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsField2Context differentiates from other interfaces.
	IsField2Context()
}

type Field2Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField2Context() *Field2Context {
	var p = new(Field2Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Rule4ParserRULE_field2
	return p
}

func (*Field2Context) IsField2Context() {}

func NewField2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field2Context {
	var p = new(Field2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Rule4ParserRULE_field2

	return p
}

func (s *Field2Context) GetParser() antlr.Parser { return s.parser }

func (s *Field2Context) FIELD2() antlr.TerminalNode {
	return s.GetToken(Rule4ParserFIELD2, 0)
}

func (s *Field2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Rule4Listener); ok {
		listenerT.EnterField2(s)
	}
}

func (s *Field2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Rule4Listener); ok {
		listenerT.ExitField2(s)
	}
}

func (p *Rule4Parser) Field2() (localctx IField2Context) {
	this := p
	_ = this

	localctx = NewField2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, Rule4ParserRULE_field2)

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
		p.Match(Rule4ParserFIELD2)
	}

	return localctx
}
