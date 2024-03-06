// Code generated from /usr/local/lib/Go/212cf23d-32c6-47b2-a754-9f6fa59aba2f/1709694127334/Rule3.g4 by ANTLR 4.10.1. DO NOT EDIT.

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

type Rule3Lexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var rule3lexerLexerStaticData struct {
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

func rule3lexerLexerInit() {
	staticData := &rule3lexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "EDUCATION", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"EDUCATION", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 4, 48, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1,
		0, 1, 0, 1, 0, 5, 0, 13, 8, 0, 10, 0, 12, 0, 16, 9, 0, 1, 0, 5, 0, 19,
		8, 0, 10, 0, 12, 0, 22, 9, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3,
		4, 3, 43, 8, 3, 11, 3, 12, 3, 44, 1, 3, 1, 3, 0, 0, 4, 1, 1, 3, 2, 5, 3,
		7, 4, 1, 0, 3, 1, 0, 46, 46, 4, 0, 46, 46, 48, 57, 65, 90, 97, 122, 3,
		0, 9, 10, 13, 13, 32, 32, 50, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5,
		1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 1, 9, 1, 0, 0, 0, 3, 23, 1, 0, 0, 0, 5, 27,
		1, 0, 0, 0, 7, 42, 1, 0, 0, 0, 9, 10, 3, 5, 2, 0, 10, 14, 3, 3, 1, 0, 11,
		13, 7, 0, 0, 0, 12, 11, 1, 0, 0, 0, 13, 16, 1, 0, 0, 0, 14, 12, 1, 0, 0,
		0, 14, 15, 1, 0, 0, 0, 15, 20, 1, 0, 0, 0, 16, 14, 1, 0, 0, 0, 17, 19,
		7, 1, 0, 0, 18, 17, 1, 0, 0, 0, 19, 22, 1, 0, 0, 0, 20, 18, 1, 0, 0, 0,
		20, 21, 1, 0, 0, 0, 21, 2, 1, 0, 0, 0, 22, 20, 1, 0, 0, 0, 23, 24, 5, 36,
		0, 0, 24, 25, 5, 126, 0, 0, 25, 26, 5, 36, 0, 0, 26, 4, 1, 0, 0, 0, 27,
		28, 5, 32, 0, 0, 28, 29, 5, 33, 0, 0, 29, 30, 5, 36, 0, 0, 30, 31, 5, 126,
		0, 0, 31, 32, 5, 36, 0, 0, 32, 33, 5, 35, 0, 0, 33, 34, 5, 37, 0, 0, 34,
		35, 5, 35, 0, 0, 35, 36, 5, 38, 0, 0, 36, 37, 5, 36, 0, 0, 37, 38, 5, 38,
		0, 0, 38, 39, 5, 33, 0, 0, 39, 40, 5, 32, 0, 0, 40, 6, 1, 0, 0, 0, 41,
		43, 7, 2, 0, 0, 42, 41, 1, 0, 0, 0, 43, 44, 1, 0, 0, 0, 44, 42, 1, 0, 0,
		0, 44, 45, 1, 0, 0, 0, 45, 46, 1, 0, 0, 0, 46, 47, 6, 3, 0, 0, 47, 8, 1,
		0, 0, 0, 4, 0, 14, 20, 44, 1, 6, 0, 0,
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

// Rule3LexerInit initializes any static state used to implement Rule3Lexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewRule3Lexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func Rule3LexerInit() {
	staticData := &rule3lexerLexerStaticData
	staticData.once.Do(rule3lexerLexerInit)
}

// NewRule3Lexer produces a new lexer instance for the optional input antlr.CharStream.
func NewRule3Lexer(input antlr.CharStream) *Rule3Lexer {
	Rule3LexerInit()
	l := new(Rule3Lexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &rule3lexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Rule3.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// Rule3Lexer tokens.
const (
	Rule3LexerEDUCATION = 1
	Rule3LexerOWNKEY    = 2
	Rule3LexerSPLITKEY  = 3
	Rule3LexerWS        = 4
)
