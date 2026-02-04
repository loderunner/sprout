// Code generated from Sprout.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type SproutLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var SproutLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func sproutlexerLexerInit() {
	staticData := &SproutLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'('", "':'", "')'", "'let'", "'in'", "'if'", "'then'", "'else'",
		"'true'", "'false'", "'fun'", "'->'", "'='",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "LET", "IN", "IF", "THEN", "ELSE", "TRUE", "FALSE",
		"FUN", "ARROW", "EQUALS", "INT", "IDENT", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "LET", "IN", "IF", "THEN", "ELSE", "TRUE", "FALSE",
		"FUN", "ARROW", "EQUALS", "INT", "IDENT", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 16, 98, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4,
		1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1,
		12, 1, 13, 4, 13, 81, 8, 13, 11, 13, 12, 13, 82, 1, 14, 1, 14, 5, 14, 87,
		8, 14, 10, 14, 12, 14, 90, 9, 14, 1, 15, 4, 15, 93, 8, 15, 11, 15, 12,
		15, 94, 1, 15, 1, 15, 0, 0, 16, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13,
		7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16,
		1, 0, 4, 1, 0, 48, 57, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65,
		90, 95, 95, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 100, 0, 1, 1, 0, 0, 0,
		0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0,
		0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0,
		0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0,
		0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 1, 33, 1,
		0, 0, 0, 3, 35, 1, 0, 0, 0, 5, 37, 1, 0, 0, 0, 7, 39, 1, 0, 0, 0, 9, 43,
		1, 0, 0, 0, 11, 46, 1, 0, 0, 0, 13, 49, 1, 0, 0, 0, 15, 54, 1, 0, 0, 0,
		17, 59, 1, 0, 0, 0, 19, 64, 1, 0, 0, 0, 21, 70, 1, 0, 0, 0, 23, 74, 1,
		0, 0, 0, 25, 77, 1, 0, 0, 0, 27, 80, 1, 0, 0, 0, 29, 84, 1, 0, 0, 0, 31,
		92, 1, 0, 0, 0, 33, 34, 5, 40, 0, 0, 34, 2, 1, 0, 0, 0, 35, 36, 5, 58,
		0, 0, 36, 4, 1, 0, 0, 0, 37, 38, 5, 41, 0, 0, 38, 6, 1, 0, 0, 0, 39, 40,
		5, 108, 0, 0, 40, 41, 5, 101, 0, 0, 41, 42, 5, 116, 0, 0, 42, 8, 1, 0,
		0, 0, 43, 44, 5, 105, 0, 0, 44, 45, 5, 110, 0, 0, 45, 10, 1, 0, 0, 0, 46,
		47, 5, 105, 0, 0, 47, 48, 5, 102, 0, 0, 48, 12, 1, 0, 0, 0, 49, 50, 5,
		116, 0, 0, 50, 51, 5, 104, 0, 0, 51, 52, 5, 101, 0, 0, 52, 53, 5, 110,
		0, 0, 53, 14, 1, 0, 0, 0, 54, 55, 5, 101, 0, 0, 55, 56, 5, 108, 0, 0, 56,
		57, 5, 115, 0, 0, 57, 58, 5, 101, 0, 0, 58, 16, 1, 0, 0, 0, 59, 60, 5,
		116, 0, 0, 60, 61, 5, 114, 0, 0, 61, 62, 5, 117, 0, 0, 62, 63, 5, 101,
		0, 0, 63, 18, 1, 0, 0, 0, 64, 65, 5, 102, 0, 0, 65, 66, 5, 97, 0, 0, 66,
		67, 5, 108, 0, 0, 67, 68, 5, 115, 0, 0, 68, 69, 5, 101, 0, 0, 69, 20, 1,
		0, 0, 0, 70, 71, 5, 102, 0, 0, 71, 72, 5, 117, 0, 0, 72, 73, 5, 110, 0,
		0, 73, 22, 1, 0, 0, 0, 74, 75, 5, 45, 0, 0, 75, 76, 5, 62, 0, 0, 76, 24,
		1, 0, 0, 0, 77, 78, 5, 61, 0, 0, 78, 26, 1, 0, 0, 0, 79, 81, 7, 0, 0, 0,
		80, 79, 1, 0, 0, 0, 81, 82, 1, 0, 0, 0, 82, 80, 1, 0, 0, 0, 82, 83, 1,
		0, 0, 0, 83, 28, 1, 0, 0, 0, 84, 88, 7, 1, 0, 0, 85, 87, 7, 2, 0, 0, 86,
		85, 1, 0, 0, 0, 87, 90, 1, 0, 0, 0, 88, 86, 1, 0, 0, 0, 88, 89, 1, 0, 0,
		0, 89, 30, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0, 91, 93, 7, 3, 0, 0, 92, 91,
		1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 92, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0,
		95, 96, 1, 0, 0, 0, 96, 97, 6, 15, 0, 0, 97, 32, 1, 0, 0, 0, 4, 0, 82,
		88, 94, 1, 6, 0, 0,
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

// SproutLexerInit initializes any static state used to implement SproutLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewSproutLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func SproutLexerInit() {
	staticData := &SproutLexerLexerStaticData
	staticData.once.Do(sproutlexerLexerInit)
}

// NewSproutLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewSproutLexer(input antlr.CharStream) *SproutLexer {
	SproutLexerInit()
	l := new(SproutLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &SproutLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Sprout.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SproutLexer tokens.
const (
	SproutLexerT__0   = 1
	SproutLexerT__1   = 2
	SproutLexerT__2   = 3
	SproutLexerLET    = 4
	SproutLexerIN     = 5
	SproutLexerIF     = 6
	SproutLexerTHEN   = 7
	SproutLexerELSE   = 8
	SproutLexerTRUE   = 9
	SproutLexerFALSE  = 10
	SproutLexerFUN    = 11
	SproutLexerARROW  = 12
	SproutLexerEQUALS = 13
	SproutLexerINT    = 14
	SproutLexerIDENT  = 15
	SproutLexerWS     = 16
)
