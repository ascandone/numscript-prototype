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
		"'['", "']'", "'='",
	}
	staticData.SymbolicNames = []string{
		"", "WS", "NEWLINE", "MULTILINE_COMMENT", "LINE_COMMENT", "SOURCE",
		"DESTINATION", "SEND", "LPARENS", "RPARENS", "LBRACKET", "RBRACKET",
		"EQ", "NUMBER", "VARIABLE_NAME", "ACCOUNT", "ASSET",
	}
	staticData.RuleNames = []string{
		"WS", "NEWLINE", "MULTILINE_COMMENT", "LINE_COMMENT", "SOURCE", "DESTINATION",
		"SEND", "LPARENS", "RPARENS", "LBRACKET", "RBRACKET", "EQ", "NUMBER",
		"VARIABLE_NAME", "ACCOUNT", "ASSET",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 16, 146, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 1, 0, 4, 0, 35, 8, 0, 11, 0, 12, 0, 36, 1, 0, 1, 0, 1, 1, 4, 1,
		42, 8, 1, 11, 1, 12, 1, 43, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 51, 8,
		2, 10, 2, 12, 2, 54, 9, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1,
		3, 1, 3, 5, 3, 65, 8, 3, 10, 3, 12, 3, 68, 9, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12,
		4, 12, 109, 8, 12, 11, 12, 12, 12, 110, 1, 13, 1, 13, 4, 13, 115, 8, 13,
		11, 13, 12, 13, 116, 1, 13, 5, 13, 120, 8, 13, 10, 13, 12, 13, 123, 9,
		13, 1, 14, 1, 14, 4, 14, 127, 8, 14, 11, 14, 12, 14, 128, 1, 14, 1, 14,
		4, 14, 133, 8, 14, 11, 14, 12, 14, 134, 5, 14, 137, 8, 14, 10, 14, 12,
		14, 140, 9, 14, 1, 15, 4, 15, 143, 8, 15, 11, 15, 12, 15, 144, 2, 52, 66,
		0, 16, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10,
		21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 1, 0, 7, 3, 0, 9, 10, 13,
		13, 32, 32, 2, 0, 10, 10, 13, 13, 1, 0, 48, 57, 2, 0, 95, 95, 97, 122,
		3, 0, 48, 57, 95, 95, 97, 122, 5, 0, 45, 45, 48, 57, 65, 90, 95, 95, 97,
		122, 2, 0, 47, 57, 65, 90, 157, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0,
		5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0,
		13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0,
		0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0,
		0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 1, 34, 1, 0, 0, 0, 3, 41, 1, 0,
		0, 0, 5, 45, 1, 0, 0, 0, 7, 60, 1, 0, 0, 0, 9, 73, 1, 0, 0, 0, 11, 80,
		1, 0, 0, 0, 13, 92, 1, 0, 0, 0, 15, 97, 1, 0, 0, 0, 17, 99, 1, 0, 0, 0,
		19, 101, 1, 0, 0, 0, 21, 103, 1, 0, 0, 0, 23, 105, 1, 0, 0, 0, 25, 108,
		1, 0, 0, 0, 27, 112, 1, 0, 0, 0, 29, 124, 1, 0, 0, 0, 31, 142, 1, 0, 0,
		0, 33, 35, 7, 0, 0, 0, 34, 33, 1, 0, 0, 0, 35, 36, 1, 0, 0, 0, 36, 34,
		1, 0, 0, 0, 36, 37, 1, 0, 0, 0, 37, 38, 1, 0, 0, 0, 38, 39, 6, 0, 0, 0,
		39, 2, 1, 0, 0, 0, 40, 42, 7, 1, 0, 0, 41, 40, 1, 0, 0, 0, 42, 43, 1, 0,
		0, 0, 43, 41, 1, 0, 0, 0, 43, 44, 1, 0, 0, 0, 44, 4, 1, 0, 0, 0, 45, 46,
		5, 47, 0, 0, 46, 47, 5, 42, 0, 0, 47, 52, 1, 0, 0, 0, 48, 51, 3, 5, 2,
		0, 49, 51, 9, 0, 0, 0, 50, 48, 1, 0, 0, 0, 50, 49, 1, 0, 0, 0, 51, 54,
		1, 0, 0, 0, 52, 53, 1, 0, 0, 0, 52, 50, 1, 0, 0, 0, 53, 55, 1, 0, 0, 0,
		54, 52, 1, 0, 0, 0, 55, 56, 5, 42, 0, 0, 56, 57, 5, 47, 0, 0, 57, 58, 1,
		0, 0, 0, 58, 59, 6, 2, 0, 0, 59, 6, 1, 0, 0, 0, 60, 61, 5, 47, 0, 0, 61,
		62, 5, 47, 0, 0, 62, 66, 1, 0, 0, 0, 63, 65, 9, 0, 0, 0, 64, 63, 1, 0,
		0, 0, 65, 68, 1, 0, 0, 0, 66, 67, 1, 0, 0, 0, 66, 64, 1, 0, 0, 0, 67, 69,
		1, 0, 0, 0, 68, 66, 1, 0, 0, 0, 69, 70, 3, 3, 1, 0, 70, 71, 1, 0, 0, 0,
		71, 72, 6, 3, 0, 0, 72, 8, 1, 0, 0, 0, 73, 74, 5, 115, 0, 0, 74, 75, 5,
		111, 0, 0, 75, 76, 5, 117, 0, 0, 76, 77, 5, 114, 0, 0, 77, 78, 5, 99, 0,
		0, 78, 79, 5, 101, 0, 0, 79, 10, 1, 0, 0, 0, 80, 81, 5, 100, 0, 0, 81,
		82, 5, 101, 0, 0, 82, 83, 5, 115, 0, 0, 83, 84, 5, 116, 0, 0, 84, 85, 5,
		105, 0, 0, 85, 86, 5, 110, 0, 0, 86, 87, 5, 97, 0, 0, 87, 88, 5, 116, 0,
		0, 88, 89, 5, 105, 0, 0, 89, 90, 5, 111, 0, 0, 90, 91, 5, 110, 0, 0, 91,
		12, 1, 0, 0, 0, 92, 93, 5, 115, 0, 0, 93, 94, 5, 101, 0, 0, 94, 95, 5,
		110, 0, 0, 95, 96, 5, 100, 0, 0, 96, 14, 1, 0, 0, 0, 97, 98, 5, 40, 0,
		0, 98, 16, 1, 0, 0, 0, 99, 100, 5, 41, 0, 0, 100, 18, 1, 0, 0, 0, 101,
		102, 5, 91, 0, 0, 102, 20, 1, 0, 0, 0, 103, 104, 5, 93, 0, 0, 104, 22,
		1, 0, 0, 0, 105, 106, 5, 61, 0, 0, 106, 24, 1, 0, 0, 0, 107, 109, 7, 2,
		0, 0, 108, 107, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110, 108, 1, 0, 0, 0,
		110, 111, 1, 0, 0, 0, 111, 26, 1, 0, 0, 0, 112, 114, 5, 36, 0, 0, 113,
		115, 7, 3, 0, 0, 114, 113, 1, 0, 0, 0, 115, 116, 1, 0, 0, 0, 116, 114,
		1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117, 121, 1, 0, 0, 0, 118, 120, 7, 4,
		0, 0, 119, 118, 1, 0, 0, 0, 120, 123, 1, 0, 0, 0, 121, 119, 1, 0, 0, 0,
		121, 122, 1, 0, 0, 0, 122, 28, 1, 0, 0, 0, 123, 121, 1, 0, 0, 0, 124, 126,
		5, 64, 0, 0, 125, 127, 7, 5, 0, 0, 126, 125, 1, 0, 0, 0, 127, 128, 1, 0,
		0, 0, 128, 126, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129, 138, 1, 0, 0, 0,
		130, 132, 5, 58, 0, 0, 131, 133, 7, 5, 0, 0, 132, 131, 1, 0, 0, 0, 133,
		134, 1, 0, 0, 0, 134, 132, 1, 0, 0, 0, 134, 135, 1, 0, 0, 0, 135, 137,
		1, 0, 0, 0, 136, 130, 1, 0, 0, 0, 137, 140, 1, 0, 0, 0, 138, 136, 1, 0,
		0, 0, 138, 139, 1, 0, 0, 0, 139, 30, 1, 0, 0, 0, 140, 138, 1, 0, 0, 0,
		141, 143, 7, 6, 0, 0, 142, 141, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144,
		142, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145, 32, 1, 0, 0, 0, 13, 0, 36,
		43, 50, 52, 66, 110, 116, 121, 128, 134, 138, 144, 1, 6, 0, 0,
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
	NumscriptLexerEQ                = 12
	NumscriptLexerNUMBER            = 13
	NumscriptLexerVARIABLE_NAME     = 14
	NumscriptLexerACCOUNT           = 15
	NumscriptLexerASSET             = 16
)
