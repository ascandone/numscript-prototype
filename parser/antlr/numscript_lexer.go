// Code generated from Numscript.g4 by ANTLR 4.13.1. DO NOT EDIT.

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

type NumscriptLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var NumscriptLexerLexerStaticData struct {
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

func numscriptlexerLexerInit() {
	staticData := &NumscriptLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "", "", "", "", "'source'", "'destination'", "'send'", "'('", "')'",
		"'['", "']'", "'{'", "'}'", "'='",
	}
	staticData.SymbolicNames = []string{
		"", "WS", "NEWLINE", "MULTILINE_COMMENT", "LINE_COMMENT", "SOURCE",
		"DESTINATION", "SEND", "LPARENS", "RPARENS", "LBRACKET", "RBRACKET",
		"LBRACE", "RBRACE", "EQ", "NUMBER", "VARIABLE_NAME", "ACCOUNT", "ASSET",
	}
	staticData.RuleNames = []string{
		"WS", "NEWLINE", "MULTILINE_COMMENT", "LINE_COMMENT", "SOURCE", "DESTINATION",
		"SEND", "LPARENS", "RPARENS", "LBRACKET", "RBRACKET", "LBRACE", "RBRACE",
		"EQ", "NUMBER", "VARIABLE_NAME", "ACCOUNT", "ASSET",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 18, 154, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 1, 0, 4, 0, 39, 8, 0, 11, 0, 12, 0,
		40, 1, 0, 1, 0, 1, 1, 4, 1, 46, 8, 1, 11, 1, 12, 1, 47, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 5, 2, 55, 8, 2, 10, 2, 12, 2, 58, 9, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 69, 8, 3, 10, 3, 12, 3, 72, 9,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1,
		10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 4, 14, 117, 8, 14,
		11, 14, 12, 14, 118, 1, 15, 1, 15, 4, 15, 123, 8, 15, 11, 15, 12, 15, 124,
		1, 15, 5, 15, 128, 8, 15, 10, 15, 12, 15, 131, 9, 15, 1, 16, 1, 16, 4,
		16, 135, 8, 16, 11, 16, 12, 16, 136, 1, 16, 1, 16, 4, 16, 141, 8, 16, 11,
		16, 12, 16, 142, 5, 16, 145, 8, 16, 10, 16, 12, 16, 148, 9, 16, 1, 17,
		4, 17, 151, 8, 17, 11, 17, 12, 17, 152, 2, 56, 70, 0, 18, 1, 1, 3, 2, 5,
		3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25,
		13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 1, 0, 7, 3, 0, 9, 10, 13, 13,
		32, 32, 2, 0, 10, 10, 13, 13, 1, 0, 48, 57, 2, 0, 95, 95, 97, 122, 3, 0,
		48, 57, 95, 95, 97, 122, 5, 0, 45, 45, 48, 57, 65, 90, 95, 95, 97, 122,
		2, 0, 47, 57, 65, 90, 165, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1,
		0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13,
		1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0,
		21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0,
		0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0,
		0, 1, 38, 1, 0, 0, 0, 3, 45, 1, 0, 0, 0, 5, 49, 1, 0, 0, 0, 7, 64, 1, 0,
		0, 0, 9, 77, 1, 0, 0, 0, 11, 84, 1, 0, 0, 0, 13, 96, 1, 0, 0, 0, 15, 101,
		1, 0, 0, 0, 17, 103, 1, 0, 0, 0, 19, 105, 1, 0, 0, 0, 21, 107, 1, 0, 0,
		0, 23, 109, 1, 0, 0, 0, 25, 111, 1, 0, 0, 0, 27, 113, 1, 0, 0, 0, 29, 116,
		1, 0, 0, 0, 31, 120, 1, 0, 0, 0, 33, 132, 1, 0, 0, 0, 35, 150, 1, 0, 0,
		0, 37, 39, 7, 0, 0, 0, 38, 37, 1, 0, 0, 0, 39, 40, 1, 0, 0, 0, 40, 38,
		1, 0, 0, 0, 40, 41, 1, 0, 0, 0, 41, 42, 1, 0, 0, 0, 42, 43, 6, 0, 0, 0,
		43, 2, 1, 0, 0, 0, 44, 46, 7, 1, 0, 0, 45, 44, 1, 0, 0, 0, 46, 47, 1, 0,
		0, 0, 47, 45, 1, 0, 0, 0, 47, 48, 1, 0, 0, 0, 48, 4, 1, 0, 0, 0, 49, 50,
		5, 47, 0, 0, 50, 51, 5, 42, 0, 0, 51, 56, 1, 0, 0, 0, 52, 55, 3, 5, 2,
		0, 53, 55, 9, 0, 0, 0, 54, 52, 1, 0, 0, 0, 54, 53, 1, 0, 0, 0, 55, 58,
		1, 0, 0, 0, 56, 57, 1, 0, 0, 0, 56, 54, 1, 0, 0, 0, 57, 59, 1, 0, 0, 0,
		58, 56, 1, 0, 0, 0, 59, 60, 5, 42, 0, 0, 60, 61, 5, 47, 0, 0, 61, 62, 1,
		0, 0, 0, 62, 63, 6, 2, 0, 0, 63, 6, 1, 0, 0, 0, 64, 65, 5, 47, 0, 0, 65,
		66, 5, 47, 0, 0, 66, 70, 1, 0, 0, 0, 67, 69, 9, 0, 0, 0, 68, 67, 1, 0,
		0, 0, 69, 72, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 70, 68, 1, 0, 0, 0, 71, 73,
		1, 0, 0, 0, 72, 70, 1, 0, 0, 0, 73, 74, 3, 3, 1, 0, 74, 75, 1, 0, 0, 0,
		75, 76, 6, 3, 0, 0, 76, 8, 1, 0, 0, 0, 77, 78, 5, 115, 0, 0, 78, 79, 5,
		111, 0, 0, 79, 80, 5, 117, 0, 0, 80, 81, 5, 114, 0, 0, 81, 82, 5, 99, 0,
		0, 82, 83, 5, 101, 0, 0, 83, 10, 1, 0, 0, 0, 84, 85, 5, 100, 0, 0, 85,
		86, 5, 101, 0, 0, 86, 87, 5, 115, 0, 0, 87, 88, 5, 116, 0, 0, 88, 89, 5,
		105, 0, 0, 89, 90, 5, 110, 0, 0, 90, 91, 5, 97, 0, 0, 91, 92, 5, 116, 0,
		0, 92, 93, 5, 105, 0, 0, 93, 94, 5, 111, 0, 0, 94, 95, 5, 110, 0, 0, 95,
		12, 1, 0, 0, 0, 96, 97, 5, 115, 0, 0, 97, 98, 5, 101, 0, 0, 98, 99, 5,
		110, 0, 0, 99, 100, 5, 100, 0, 0, 100, 14, 1, 0, 0, 0, 101, 102, 5, 40,
		0, 0, 102, 16, 1, 0, 0, 0, 103, 104, 5, 41, 0, 0, 104, 18, 1, 0, 0, 0,
		105, 106, 5, 91, 0, 0, 106, 20, 1, 0, 0, 0, 107, 108, 5, 93, 0, 0, 108,
		22, 1, 0, 0, 0, 109, 110, 5, 123, 0, 0, 110, 24, 1, 0, 0, 0, 111, 112,
		5, 125, 0, 0, 112, 26, 1, 0, 0, 0, 113, 114, 5, 61, 0, 0, 114, 28, 1, 0,
		0, 0, 115, 117, 7, 2, 0, 0, 116, 115, 1, 0, 0, 0, 117, 118, 1, 0, 0, 0,
		118, 116, 1, 0, 0, 0, 118, 119, 1, 0, 0, 0, 119, 30, 1, 0, 0, 0, 120, 122,
		5, 36, 0, 0, 121, 123, 7, 3, 0, 0, 122, 121, 1, 0, 0, 0, 123, 124, 1, 0,
		0, 0, 124, 122, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 129, 1, 0, 0, 0,
		126, 128, 7, 4, 0, 0, 127, 126, 1, 0, 0, 0, 128, 131, 1, 0, 0, 0, 129,
		127, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130, 32, 1, 0, 0, 0, 131, 129, 1,
		0, 0, 0, 132, 134, 5, 64, 0, 0, 133, 135, 7, 5, 0, 0, 134, 133, 1, 0, 0,
		0, 135, 136, 1, 0, 0, 0, 136, 134, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 137,
		146, 1, 0, 0, 0, 138, 140, 5, 58, 0, 0, 139, 141, 7, 5, 0, 0, 140, 139,
		1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 140, 1, 0, 0, 0, 142, 143, 1, 0,
		0, 0, 143, 145, 1, 0, 0, 0, 144, 138, 1, 0, 0, 0, 145, 148, 1, 0, 0, 0,
		146, 144, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 34, 1, 0, 0, 0, 148, 146,
		1, 0, 0, 0, 149, 151, 7, 6, 0, 0, 150, 149, 1, 0, 0, 0, 151, 152, 1, 0,
		0, 0, 152, 150, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 36, 1, 0, 0, 0,
		13, 0, 40, 47, 54, 56, 70, 118, 124, 129, 136, 142, 146, 152, 1, 6, 0,
		0,
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

// NumscriptLexerInit initializes any static state used to implement NumscriptLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewNumscriptLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func NumscriptLexerInit() {
	staticData := &NumscriptLexerLexerStaticData
	staticData.once.Do(numscriptlexerLexerInit)
}

// NewNumscriptLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewNumscriptLexer(input antlr.CharStream) *NumscriptLexer {
	NumscriptLexerInit()
	l := new(NumscriptLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &NumscriptLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Numscript.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// NumscriptLexer tokens.
const (
	NumscriptLexerWS                = 1
	NumscriptLexerNEWLINE           = 2
	NumscriptLexerMULTILINE_COMMENT = 3
	NumscriptLexerLINE_COMMENT      = 4
	NumscriptLexerSOURCE            = 5
	NumscriptLexerDESTINATION       = 6
	NumscriptLexerSEND              = 7
	NumscriptLexerLPARENS           = 8
	NumscriptLexerRPARENS           = 9
	NumscriptLexerLBRACKET          = 10
	NumscriptLexerRBRACKET          = 11
	NumscriptLexerLBRACE            = 12
	NumscriptLexerRBRACE            = 13
	NumscriptLexerEQ                = 14
	NumscriptLexerNUMBER            = 15
	NumscriptLexerVARIABLE_NAME     = 16
	NumscriptLexerACCOUNT           = 17
	NumscriptLexerASSET             = 18
)
