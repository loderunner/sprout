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
		"", "':'", "'='", "'('", "')'", "'||'", "'&&'", "'+'", "'-'", "'*'",
		"'/'", "'!'", "'let'", "'in'", "'if'", "'then'", "'else'", "'true'",
		"'false'", "'fun'", "'->'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "LET", "IN", "IF", "THEN",
		"ELSE", "TRUE", "FALSE", "FUN", "ARROW", "INT", "IDENT", "COMPOP", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "LET", "IN", "IF", "THEN", "ELSE", "TRUE", "FALSE",
		"FUN", "ARROW", "INT", "IDENT", "COMPOP", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 24, 146, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1,
		7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 5, 20, 115, 8, 20, 10, 20, 12,
		20, 118, 9, 20, 3, 20, 120, 8, 20, 1, 21, 1, 21, 5, 21, 124, 8, 21, 10,
		21, 12, 21, 127, 9, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		1, 22, 1, 22, 3, 22, 138, 8, 22, 1, 23, 4, 23, 141, 8, 23, 11, 23, 12,
		23, 142, 1, 23, 1, 23, 0, 0, 24, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13,
		7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16,
		33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 1, 0, 6,
		1, 0, 49, 57, 1, 0, 48, 57, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57,
		65, 90, 95, 95, 97, 122, 2, 0, 60, 60, 62, 62, 3, 0, 9, 10, 13, 13, 32,
		32, 153, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1,
		0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15,
		1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0,
		23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0,
		0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0,
		0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0,
		0, 0, 0, 47, 1, 0, 0, 0, 1, 49, 1, 0, 0, 0, 3, 51, 1, 0, 0, 0, 5, 53, 1,
		0, 0, 0, 7, 55, 1, 0, 0, 0, 9, 57, 1, 0, 0, 0, 11, 60, 1, 0, 0, 0, 13,
		63, 1, 0, 0, 0, 15, 65, 1, 0, 0, 0, 17, 67, 1, 0, 0, 0, 19, 69, 1, 0, 0,
		0, 21, 71, 1, 0, 0, 0, 23, 73, 1, 0, 0, 0, 25, 77, 1, 0, 0, 0, 27, 80,
		1, 0, 0, 0, 29, 83, 1, 0, 0, 0, 31, 88, 1, 0, 0, 0, 33, 93, 1, 0, 0, 0,
		35, 98, 1, 0, 0, 0, 37, 104, 1, 0, 0, 0, 39, 108, 1, 0, 0, 0, 41, 119,
		1, 0, 0, 0, 43, 121, 1, 0, 0, 0, 45, 137, 1, 0, 0, 0, 47, 140, 1, 0, 0,
		0, 49, 50, 5, 58, 0, 0, 50, 2, 1, 0, 0, 0, 51, 52, 5, 61, 0, 0, 52, 4,
		1, 0, 0, 0, 53, 54, 5, 40, 0, 0, 54, 6, 1, 0, 0, 0, 55, 56, 5, 41, 0, 0,
		56, 8, 1, 0, 0, 0, 57, 58, 5, 124, 0, 0, 58, 59, 5, 124, 0, 0, 59, 10,
		1, 0, 0, 0, 60, 61, 5, 38, 0, 0, 61, 62, 5, 38, 0, 0, 62, 12, 1, 0, 0,
		0, 63, 64, 5, 43, 0, 0, 64, 14, 1, 0, 0, 0, 65, 66, 5, 45, 0, 0, 66, 16,
		1, 0, 0, 0, 67, 68, 5, 42, 0, 0, 68, 18, 1, 0, 0, 0, 69, 70, 5, 47, 0,
		0, 70, 20, 1, 0, 0, 0, 71, 72, 5, 33, 0, 0, 72, 22, 1, 0, 0, 0, 73, 74,
		5, 108, 0, 0, 74, 75, 5, 101, 0, 0, 75, 76, 5, 116, 0, 0, 76, 24, 1, 0,
		0, 0, 77, 78, 5, 105, 0, 0, 78, 79, 5, 110, 0, 0, 79, 26, 1, 0, 0, 0, 80,
		81, 5, 105, 0, 0, 81, 82, 5, 102, 0, 0, 82, 28, 1, 0, 0, 0, 83, 84, 5,
		116, 0, 0, 84, 85, 5, 104, 0, 0, 85, 86, 5, 101, 0, 0, 86, 87, 5, 110,
		0, 0, 87, 30, 1, 0, 0, 0, 88, 89, 5, 101, 0, 0, 89, 90, 5, 108, 0, 0, 90,
		91, 5, 115, 0, 0, 91, 92, 5, 101, 0, 0, 92, 32, 1, 0, 0, 0, 93, 94, 5,
		116, 0, 0, 94, 95, 5, 114, 0, 0, 95, 96, 5, 117, 0, 0, 96, 97, 5, 101,
		0, 0, 97, 34, 1, 0, 0, 0, 98, 99, 5, 102, 0, 0, 99, 100, 5, 97, 0, 0, 100,
		101, 5, 108, 0, 0, 101, 102, 5, 115, 0, 0, 102, 103, 5, 101, 0, 0, 103,
		36, 1, 0, 0, 0, 104, 105, 5, 102, 0, 0, 105, 106, 5, 117, 0, 0, 106, 107,
		5, 110, 0, 0, 107, 38, 1, 0, 0, 0, 108, 109, 5, 45, 0, 0, 109, 110, 5,
		62, 0, 0, 110, 40, 1, 0, 0, 0, 111, 120, 5, 48, 0, 0, 112, 116, 7, 0, 0,
		0, 113, 115, 7, 1, 0, 0, 114, 113, 1, 0, 0, 0, 115, 118, 1, 0, 0, 0, 116,
		114, 1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117, 120, 1, 0, 0, 0, 118, 116,
		1, 0, 0, 0, 119, 111, 1, 0, 0, 0, 119, 112, 1, 0, 0, 0, 120, 42, 1, 0,
		0, 0, 121, 125, 7, 2, 0, 0, 122, 124, 7, 3, 0, 0, 123, 122, 1, 0, 0, 0,
		124, 127, 1, 0, 0, 0, 125, 123, 1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126,
		44, 1, 0, 0, 0, 127, 125, 1, 0, 0, 0, 128, 129, 5, 61, 0, 0, 129, 138,
		5, 61, 0, 0, 130, 131, 5, 33, 0, 0, 131, 138, 5, 61, 0, 0, 132, 133, 5,
		60, 0, 0, 133, 138, 5, 61, 0, 0, 134, 135, 5, 62, 0, 0, 135, 138, 5, 61,
		0, 0, 136, 138, 7, 4, 0, 0, 137, 128, 1, 0, 0, 0, 137, 130, 1, 0, 0, 0,
		137, 132, 1, 0, 0, 0, 137, 134, 1, 0, 0, 0, 137, 136, 1, 0, 0, 0, 138,
		46, 1, 0, 0, 0, 139, 141, 7, 5, 0, 0, 140, 139, 1, 0, 0, 0, 141, 142, 1,
		0, 0, 0, 142, 140, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 144, 1, 0, 0,
		0, 144, 145, 6, 23, 0, 0, 145, 48, 1, 0, 0, 0, 6, 0, 116, 119, 125, 137,
		142, 1, 6, 0, 0,
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
	SproutLexerT__3   = 4
	SproutLexerT__4   = 5
	SproutLexerT__5   = 6
	SproutLexerT__6   = 7
	SproutLexerT__7   = 8
	SproutLexerT__8   = 9
	SproutLexerT__9   = 10
	SproutLexerT__10  = 11
	SproutLexerLET    = 12
	SproutLexerIN     = 13
	SproutLexerIF     = 14
	SproutLexerTHEN   = 15
	SproutLexerELSE   = 16
	SproutLexerTRUE   = 17
	SproutLexerFALSE  = 18
	SproutLexerFUN    = 19
	SproutLexerARROW  = 20
	SproutLexerINT    = 21
	SproutLexerIDENT  = 22
	SproutLexerCOMPOP = 23
	SproutLexerWS     = 24
)
