// Code generated from /usr/local/lib/Go/212cf23d-32c6-47b2-a754-9f6fa59aba2f/1710560357049/Rule4.g4 by ANTLR 4.10.1. DO NOT EDIT.

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

type Rule4Lexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var rule4lexerLexerStaticData struct {
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

func rule4lexerLexerInit() {
	staticData := &rule4lexerLexerStaticData
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
		"", "FIELD1", "FIELD2", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"FIELD1", "FIELD2", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 5, 60, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 31, 8, 1, 10,
		1, 12, 1, 34, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 4, 4, 55,
		8, 4, 11, 4, 12, 4, 56, 1, 4, 1, 4, 0, 0, 5, 1, 1, 3, 2, 5, 3, 7, 4, 9,
		5, 1, 0, 3, 1, 0, 48, 57, 5, 0, 32, 32, 46, 46, 48, 57, 65, 90, 97, 122,
		3, 0, 9, 10, 13, 13, 32, 32, 61, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0,
		5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 1, 11, 1, 0, 0, 0, 3,
		26, 1, 0, 0, 0, 5, 35, 1, 0, 0, 0, 7, 39, 1, 0, 0, 0, 9, 54, 1, 0, 0, 0,
		11, 12, 3, 7, 3, 0, 12, 13, 3, 5, 2, 0, 13, 14, 7, 0, 0, 0, 14, 15, 7,
		0, 0, 0, 15, 16, 7, 0, 0, 0, 16, 17, 7, 0, 0, 0, 17, 18, 7, 0, 0, 0, 18,
		19, 7, 0, 0, 0, 19, 20, 7, 0, 0, 0, 20, 21, 7, 0, 0, 0, 21, 22, 7, 0, 0,
		0, 22, 23, 7, 0, 0, 0, 23, 24, 7, 0, 0, 0, 24, 25, 7, 0, 0, 0, 25, 2, 1,
		0, 0, 0, 26, 27, 3, 7, 3, 0, 27, 28, 3, 5, 2, 0, 28, 32, 3, 5, 2, 0, 29,
		31, 7, 1, 0, 0, 30, 29, 1, 0, 0, 0, 31, 34, 1, 0, 0, 0, 32, 30, 1, 0, 0,
		0, 32, 33, 1, 0, 0, 0, 33, 4, 1, 0, 0, 0, 34, 32, 1, 0, 0, 0, 35, 36, 5,
		36, 0, 0, 36, 37, 5, 126, 0, 0, 37, 38, 5, 36, 0, 0, 38, 6, 1, 0, 0, 0,
		39, 40, 5, 32, 0, 0, 40, 41, 5, 33, 0, 0, 41, 42, 5, 36, 0, 0, 42, 43,
		5, 126, 0, 0, 43, 44, 5, 36, 0, 0, 44, 45, 5, 35, 0, 0, 45, 46, 5, 37,
		0, 0, 46, 47, 5, 35, 0, 0, 47, 48, 5, 38, 0, 0, 48, 49, 5, 36, 0, 0, 49,
		50, 5, 38, 0, 0, 50, 51, 5, 33, 0, 0, 51, 52, 5, 32, 0, 0, 52, 8, 1, 0,
		0, 0, 53, 55, 7, 2, 0, 0, 54, 53, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 54,
		1, 0, 0, 0, 56, 57, 1, 0, 0, 0, 57, 58, 1, 0, 0, 0, 58, 59, 6, 4, 0, 0,
		59, 10, 1, 0, 0, 0, 3, 0, 32, 56, 1, 6, 0, 0,
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

// Rule4LexerInit initializes any static state used to implement Rule4Lexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewRule4Lexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func Rule4LexerInit() {
	staticData := &rule4lexerLexerStaticData
	staticData.once.Do(rule4lexerLexerInit)
}

// NewRule4Lexer produces a new lexer instance for the optional input antlr.CharStream.
func NewRule4Lexer(input antlr.CharStream) *Rule4Lexer {
	Rule4LexerInit()
	l := new(Rule4Lexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &rule4lexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Rule4.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// Rule4Lexer tokens.
const (
	Rule4LexerFIELD1   = 1
	Rule4LexerFIELD2   = 2
	Rule4LexerOWNKEY   = 3
	Rule4LexerSPLITKEY = 4
	Rule4LexerWS       = 5
)
