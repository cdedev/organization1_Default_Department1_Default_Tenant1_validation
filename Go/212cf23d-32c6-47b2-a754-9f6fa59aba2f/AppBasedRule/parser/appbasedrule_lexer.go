// Code generated from /usr/local/lib/Go/212cf23d-32c6-47b2-a754-9f6fa59aba2f/1709267325888/AppBasedRule.g4 by ANTLR 4.10.1. DO NOT EDIT.

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

type AppBasedRuleLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var appbasedrulelexerLexerStaticData struct {
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

func appbasedrulelexerLexerInit() {
	staticData := &appbasedrulelexerLexerStaticData
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
		"", "GLUCOSE", "GLUCOSE_ACOSH", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"GLUCOSE", "GLUCOSE_ACOSH", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 5, 53, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 1, 0, 1, 0, 1, 0, 5, 0, 15, 8, 0, 10, 0, 12, 0, 18, 9, 0, 1, 1,
		1, 1, 1, 1, 1, 1, 5, 1, 24, 8, 1, 10, 1, 12, 1, 27, 9, 1, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 4, 4, 4, 48, 8, 4, 11, 4, 12, 4, 49, 1, 4, 1, 4,
		0, 0, 5, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 1, 0, 2, 1, 0, 48, 57, 3, 0, 9,
		10, 13, 13, 32, 32, 55, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0,
		0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 1, 11, 1, 0, 0, 0, 3, 19, 1,
		0, 0, 0, 5, 28, 1, 0, 0, 0, 7, 32, 1, 0, 0, 0, 9, 47, 1, 0, 0, 0, 11, 12,
		3, 7, 3, 0, 12, 16, 3, 5, 2, 0, 13, 15, 7, 0, 0, 0, 14, 13, 1, 0, 0, 0,
		15, 18, 1, 0, 0, 0, 16, 14, 1, 0, 0, 0, 16, 17, 1, 0, 0, 0, 17, 2, 1, 0,
		0, 0, 18, 16, 1, 0, 0, 0, 19, 20, 3, 7, 3, 0, 20, 21, 3, 5, 2, 0, 21, 25,
		3, 5, 2, 0, 22, 24, 7, 0, 0, 0, 23, 22, 1, 0, 0, 0, 24, 27, 1, 0, 0, 0,
		25, 23, 1, 0, 0, 0, 25, 26, 1, 0, 0, 0, 26, 4, 1, 0, 0, 0, 27, 25, 1, 0,
		0, 0, 28, 29, 5, 36, 0, 0, 29, 30, 5, 126, 0, 0, 30, 31, 5, 36, 0, 0, 31,
		6, 1, 0, 0, 0, 32, 33, 5, 32, 0, 0, 33, 34, 5, 33, 0, 0, 34, 35, 5, 36,
		0, 0, 35, 36, 5, 126, 0, 0, 36, 37, 5, 36, 0, 0, 37, 38, 5, 35, 0, 0, 38,
		39, 5, 37, 0, 0, 39, 40, 5, 35, 0, 0, 40, 41, 5, 38, 0, 0, 41, 42, 5, 36,
		0, 0, 42, 43, 5, 38, 0, 0, 43, 44, 5, 33, 0, 0, 44, 45, 5, 32, 0, 0, 45,
		8, 1, 0, 0, 0, 46, 48, 7, 1, 0, 0, 47, 46, 1, 0, 0, 0, 48, 49, 1, 0, 0,
		0, 49, 47, 1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 51, 52,
		6, 4, 0, 0, 52, 10, 1, 0, 0, 0, 4, 0, 16, 25, 49, 1, 6, 0, 0,
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

// AppBasedRuleLexerInit initializes any static state used to implement AppBasedRuleLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewAppBasedRuleLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func AppBasedRuleLexerInit() {
	staticData := &appbasedrulelexerLexerStaticData
	staticData.once.Do(appbasedrulelexerLexerInit)
}

// NewAppBasedRuleLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewAppBasedRuleLexer(input antlr.CharStream) *AppBasedRuleLexer {
	AppBasedRuleLexerInit()
	l := new(AppBasedRuleLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &appbasedrulelexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "AppBasedRule.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// AppBasedRuleLexer tokens.
const (
	AppBasedRuleLexerGLUCOSE       = 1
	AppBasedRuleLexerGLUCOSE_ACOSH = 2
	AppBasedRuleLexerOWNKEY        = 3
	AppBasedRuleLexerSPLITKEY      = 4
	AppBasedRuleLexerWS            = 5
)
