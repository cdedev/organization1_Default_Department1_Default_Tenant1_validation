// Code generated from /usr/local/lib/Go/212cf23d-32c6-47b2-a754-9f6fa59aba2f/1709693485583/Rule2.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type Rule2Lexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var rule2lexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func rule2lexerLexerInit() {
	staticData := &rule2lexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "EDUCATION", "FOR", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"EDUCATION", "FOR", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 5, 51, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 1, 0, 1, 0, 1, 0, 4, 0, 15, 8, 0, 11, 0, 12, 0, 16, 1, 1, 1, 1,
		1, 1, 1, 1, 4, 1, 23, 8, 1, 11, 1, 12, 1, 24, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 4, 4, 4, 46, 8, 4, 11, 4, 12, 4, 47, 1, 4, 1, 4, 0, 0, 5, 1,
		1, 3, 2, 5, 3, 7, 4, 9, 5, 1, 0, 3, 5, 0, 32, 32, 46, 46, 48, 57, 65, 90,
		97, 122, 5, 0, 32, 32, 46, 46, 49, 57, 65, 90, 97, 122, 3, 0, 9, 10, 13,
		13, 32, 32, 53, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0,
		7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 1, 11, 1, 0, 0, 0, 3, 18, 1, 0, 0, 0,
		5, 26, 1, 0, 0, 0, 7, 30, 1, 0, 0, 0, 9, 45, 1, 0, 0, 0, 11, 12, 3, 7,
		3, 0, 12, 14, 3, 5, 2, 0, 13, 15, 7, 0, 0, 0, 14, 13, 1, 0, 0, 0, 15, 16,
		1, 0, 0, 0, 16, 14, 1, 0, 0, 0, 16, 17, 1, 0, 0, 0, 17, 2, 1, 0, 0, 0,
		18, 19, 3, 7, 3, 0, 19, 20, 3, 5, 2, 0, 20, 22, 3, 5, 2, 0, 21, 23, 7,
		1, 0, 0, 22, 21, 1, 0, 0, 0, 23, 24, 1, 0, 0, 0, 24, 22, 1, 0, 0, 0, 24,
		25, 1, 0, 0, 0, 25, 4, 1, 0, 0, 0, 26, 27, 5, 36, 0, 0, 27, 28, 5, 126,
		0, 0, 28, 29, 5, 36, 0, 0, 29, 6, 1, 0, 0, 0, 30, 31, 5, 32, 0, 0, 31,
		32, 5, 33, 0, 0, 32, 33, 5, 36, 0, 0, 33, 34, 5, 126, 0, 0, 34, 35, 5,
		36, 0, 0, 35, 36, 5, 35, 0, 0, 36, 37, 5, 37, 0, 0, 37, 38, 5, 35, 0, 0,
		38, 39, 5, 38, 0, 0, 39, 40, 5, 36, 0, 0, 40, 41, 5, 38, 0, 0, 41, 42,
		5, 33, 0, 0, 42, 43, 5, 32, 0, 0, 43, 8, 1, 0, 0, 0, 44, 46, 7, 2, 0, 0,
		45, 44, 1, 0, 0, 0, 46, 47, 1, 0, 0, 0, 47, 45, 1, 0, 0, 0, 47, 48, 1,
		0, 0, 0, 48, 49, 1, 0, 0, 0, 49, 50, 6, 4, 0, 0, 50, 10, 1, 0, 0, 0, 4,
		0, 16, 24, 47, 1, 6, 0, 0,
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

// Rule2LexerInit initializes any static state used to implement Rule2Lexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewRule2Lexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func Rule2LexerInit() {
	staticData := &rule2lexerLexerStaticData
	staticData.once.Do(rule2lexerLexerInit)
}

// NewRule2Lexer produces a new lexer instance for the optional input antlr.CharStream.
func NewRule2Lexer(input antlr.CharStream) *Rule2Lexer {
	Rule2LexerInit()
	l := new(Rule2Lexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &rule2lexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Rule2.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// Rule2Lexer tokens.
const (
	Rule2LexerEDUCATION = 1
	Rule2LexerFOR       = 2
	Rule2LexerOWNKEY    = 3
	Rule2LexerSPLITKEY  = 4
	Rule2LexerWS        = 5
)
