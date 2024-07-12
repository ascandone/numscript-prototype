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
		"", "", "", "", "", "'vars'", "'max'", "'source'", "'destination'",
		"'send'", "'from'", "'to'", "'('", "')'", "'['", "']'", "'{'", "'}'",
		"'='",
	}
	staticData.SymbolicNames = []string{
		"", "WS", "NEWLINE", "MULTILINE_COMMENT", "LINE_COMMENT", "VARS", "MAX",
		"SOURCE", "DESTINATION", "SEND", "FROM", "TO", "LPARENS", "RPARENS",
		"LBRACKET", "RBRACKET", "LBRACE", "RBRACE", "EQ", "RATIO_PORTION_LITERAL",
		"PERCENTAGE_PORTION_LITERAL", "TYPE_IDENT", "NUMBER", "VARIABLE_NAME",
		"ACCOUNT", "ASSET",
	}
	staticData.RuleNames = []string{
		"WS", "NEWLINE", "MULTILINE_COMMENT", "LINE_COMMENT", "VARS", "MAX",
		"SOURCE", "DESTINATION", "SEND", "FROM", "TO", "LPARENS", "RPARENS",
		"LBRACKET", "RBRACKET", "LBRACE", "RBRACE", "EQ", "RATIO_PORTION_LITERAL",
		"PERCENTAGE_PORTION_LITERAL", "TYPE_IDENT", "NUMBER", "VARIABLE_NAME",
		"ACCOUNT", "ASSET",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 25, 222, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 1, 0, 4, 0,
		53, 8, 0, 11, 0, 12, 0, 54, 1, 0, 1, 0, 1, 1, 4, 1, 60, 8, 1, 11, 1, 12,
		1, 61, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 69, 8, 2, 10, 2, 12, 2, 72,
		9, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 83, 8,
		3, 10, 3, 12, 3, 86, 9, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1,
		10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15,
		1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 4, 18, 148, 8, 18, 11, 18, 12,
		18, 149, 1, 18, 3, 18, 153, 8, 18, 1, 18, 1, 18, 3, 18, 157, 8, 18, 1,
		18, 4, 18, 160, 8, 18, 11, 18, 12, 18, 161, 1, 19, 4, 19, 165, 8, 19, 11,
		19, 12, 19, 166, 1, 19, 1, 19, 4, 19, 171, 8, 19, 11, 19, 12, 19, 172,
		3, 19, 175, 8, 19, 1, 19, 1, 19, 1, 20, 4, 20, 180, 8, 20, 11, 20, 12,
		20, 181, 1, 21, 4, 21, 185, 8, 21, 11, 21, 12, 21, 186, 1, 22, 1, 22, 4,
		22, 191, 8, 22, 11, 22, 12, 22, 192, 1, 22, 5, 22, 196, 8, 22, 10, 22,
		12, 22, 199, 9, 22, 1, 23, 1, 23, 4, 23, 203, 8, 23, 11, 23, 12, 23, 204,
		1, 23, 1, 23, 4, 23, 209, 8, 23, 11, 23, 12, 23, 210, 5, 23, 213, 8, 23,
		10, 23, 12, 23, 216, 9, 23, 1, 24, 4, 24, 219, 8, 24, 11, 24, 12, 24, 220,
		2, 70, 84, 0, 25, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17,
		9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35,
		18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 1, 0, 9, 3,
		0, 9, 10, 13, 13, 32, 32, 2, 0, 10, 10, 13, 13, 1, 0, 48, 57, 1, 0, 32,
		32, 1, 0, 97, 122, 2, 0, 95, 95, 97, 122, 3, 0, 48, 57, 95, 95, 97, 122,
		5, 0, 45, 45, 48, 57, 65, 90, 95, 95, 97, 122, 2, 0, 47, 57, 65, 90, 241,
		0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0,
		0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0,
		0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0,
		0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1,
		0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39,
		1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0,
		47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 1, 52, 1, 0, 0, 0, 3, 59, 1, 0, 0, 0,
		5, 63, 1, 0, 0, 0, 7, 78, 1, 0, 0, 0, 9, 91, 1, 0, 0, 0, 11, 96, 1, 0,
		0, 0, 13, 100, 1, 0, 0, 0, 15, 107, 1, 0, 0, 0, 17, 119, 1, 0, 0, 0, 19,
		124, 1, 0, 0, 0, 21, 129, 1, 0, 0, 0, 23, 132, 1, 0, 0, 0, 25, 134, 1,
		0, 0, 0, 27, 136, 1, 0, 0, 0, 29, 138, 1, 0, 0, 0, 31, 140, 1, 0, 0, 0,
		33, 142, 1, 0, 0, 0, 35, 144, 1, 0, 0, 0, 37, 147, 1, 0, 0, 0, 39, 164,
		1, 0, 0, 0, 41, 179, 1, 0, 0, 0, 43, 184, 1, 0, 0, 0, 45, 188, 1, 0, 0,
		0, 47, 200, 1, 0, 0, 0, 49, 218, 1, 0, 0, 0, 51, 53, 7, 0, 0, 0, 52, 51,
		1, 0, 0, 0, 53, 54, 1, 0, 0, 0, 54, 52, 1, 0, 0, 0, 54, 55, 1, 0, 0, 0,
		55, 56, 1, 0, 0, 0, 56, 57, 6, 0, 0, 0, 57, 2, 1, 0, 0, 0, 58, 60, 7, 1,
		0, 0, 59, 58, 1, 0, 0, 0, 60, 61, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 61, 62,
		1, 0, 0, 0, 62, 4, 1, 0, 0, 0, 63, 64, 5, 47, 0, 0, 64, 65, 5, 42, 0, 0,
		65, 70, 1, 0, 0, 0, 66, 69, 3, 5, 2, 0, 67, 69, 9, 0, 0, 0, 68, 66, 1,
		0, 0, 0, 68, 67, 1, 0, 0, 0, 69, 72, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 70,
		68, 1, 0, 0, 0, 71, 73, 1, 0, 0, 0, 72, 70, 1, 0, 0, 0, 73, 74, 5, 42,
		0, 0, 74, 75, 5, 47, 0, 0, 75, 76, 1, 0, 0, 0, 76, 77, 6, 2, 0, 0, 77,
		6, 1, 0, 0, 0, 78, 79, 5, 47, 0, 0, 79, 80, 5, 47, 0, 0, 80, 84, 1, 0,
		0, 0, 81, 83, 9, 0, 0, 0, 82, 81, 1, 0, 0, 0, 83, 86, 1, 0, 0, 0, 84, 85,
		1, 0, 0, 0, 84, 82, 1, 0, 0, 0, 85, 87, 1, 0, 0, 0, 86, 84, 1, 0, 0, 0,
		87, 88, 3, 3, 1, 0, 88, 89, 1, 0, 0, 0, 89, 90, 6, 3, 0, 0, 90, 8, 1, 0,
		0, 0, 91, 92, 5, 118, 0, 0, 92, 93, 5, 97, 0, 0, 93, 94, 5, 114, 0, 0,
		94, 95, 5, 115, 0, 0, 95, 10, 1, 0, 0, 0, 96, 97, 5, 109, 0, 0, 97, 98,
		5, 97, 0, 0, 98, 99, 5, 120, 0, 0, 99, 12, 1, 0, 0, 0, 100, 101, 5, 115,
		0, 0, 101, 102, 5, 111, 0, 0, 102, 103, 5, 117, 0, 0, 103, 104, 5, 114,
		0, 0, 104, 105, 5, 99, 0, 0, 105, 106, 5, 101, 0, 0, 106, 14, 1, 0, 0,
		0, 107, 108, 5, 100, 0, 0, 108, 109, 5, 101, 0, 0, 109, 110, 5, 115, 0,
		0, 110, 111, 5, 116, 0, 0, 111, 112, 5, 105, 0, 0, 112, 113, 5, 110, 0,
		0, 113, 114, 5, 97, 0, 0, 114, 115, 5, 116, 0, 0, 115, 116, 5, 105, 0,
		0, 116, 117, 5, 111, 0, 0, 117, 118, 5, 110, 0, 0, 118, 16, 1, 0, 0, 0,
		119, 120, 5, 115, 0, 0, 120, 121, 5, 101, 0, 0, 121, 122, 5, 110, 0, 0,
		122, 123, 5, 100, 0, 0, 123, 18, 1, 0, 0, 0, 124, 125, 5, 102, 0, 0, 125,
		126, 5, 114, 0, 0, 126, 127, 5, 111, 0, 0, 127, 128, 5, 109, 0, 0, 128,
		20, 1, 0, 0, 0, 129, 130, 5, 116, 0, 0, 130, 131, 5, 111, 0, 0, 131, 22,
		1, 0, 0, 0, 132, 133, 5, 40, 0, 0, 133, 24, 1, 0, 0, 0, 134, 135, 5, 41,
		0, 0, 135, 26, 1, 0, 0, 0, 136, 137, 5, 91, 0, 0, 137, 28, 1, 0, 0, 0,
		138, 139, 5, 93, 0, 0, 139, 30, 1, 0, 0, 0, 140, 141, 5, 123, 0, 0, 141,
		32, 1, 0, 0, 0, 142, 143, 5, 125, 0, 0, 143, 34, 1, 0, 0, 0, 144, 145,
		5, 61, 0, 0, 145, 36, 1, 0, 0, 0, 146, 148, 7, 2, 0, 0, 147, 146, 1, 0,
		0, 0, 148, 149, 1, 0, 0, 0, 149, 147, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0,
		150, 152, 1, 0, 0, 0, 151, 153, 7, 3, 0, 0, 152, 151, 1, 0, 0, 0, 152,
		153, 1, 0, 0, 0, 153, 154, 1, 0, 0, 0, 154, 156, 5, 47, 0, 0, 155, 157,
		7, 3, 0, 0, 156, 155, 1, 0, 0, 0, 156, 157, 1, 0, 0, 0, 157, 159, 1, 0,
		0, 0, 158, 160, 7, 2, 0, 0, 159, 158, 1, 0, 0, 0, 160, 161, 1, 0, 0, 0,
		161, 159, 1, 0, 0, 0, 161, 162, 1, 0, 0, 0, 162, 38, 1, 0, 0, 0, 163, 165,
		7, 2, 0, 0, 164, 163, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 164, 1, 0,
		0, 0, 166, 167, 1, 0, 0, 0, 167, 174, 1, 0, 0, 0, 168, 170, 5, 46, 0, 0,
		169, 171, 7, 2, 0, 0, 170, 169, 1, 0, 0, 0, 171, 172, 1, 0, 0, 0, 172,
		170, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0, 173, 175, 1, 0, 0, 0, 174, 168,
		1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 175, 176, 1, 0, 0, 0, 176, 177, 5, 37,
		0, 0, 177, 40, 1, 0, 0, 0, 178, 180, 7, 4, 0, 0, 179, 178, 1, 0, 0, 0,
		180, 181, 1, 0, 0, 0, 181, 179, 1, 0, 0, 0, 181, 182, 1, 0, 0, 0, 182,
		42, 1, 0, 0, 0, 183, 185, 7, 2, 0, 0, 184, 183, 1, 0, 0, 0, 185, 186, 1,
		0, 0, 0, 186, 184, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 44, 1, 0, 0,
		0, 188, 190, 5, 36, 0, 0, 189, 191, 7, 5, 0, 0, 190, 189, 1, 0, 0, 0, 191,
		192, 1, 0, 0, 0, 192, 190, 1, 0, 0, 0, 192, 193, 1, 0, 0, 0, 193, 197,
		1, 0, 0, 0, 194, 196, 7, 6, 0, 0, 195, 194, 1, 0, 0, 0, 196, 199, 1, 0,
		0, 0, 197, 195, 1, 0, 0, 0, 197, 198, 1, 0, 0, 0, 198, 46, 1, 0, 0, 0,
		199, 197, 1, 0, 0, 0, 200, 202, 5, 64, 0, 0, 201, 203, 7, 7, 0, 0, 202,
		201, 1, 0, 0, 0, 203, 204, 1, 0, 0, 0, 204, 202, 1, 0, 0, 0, 204, 205,
		1, 0, 0, 0, 205, 214, 1, 0, 0, 0, 206, 208, 5, 58, 0, 0, 207, 209, 7, 7,
		0, 0, 208, 207, 1, 0, 0, 0, 209, 210, 1, 0, 0, 0, 210, 208, 1, 0, 0, 0,
		210, 211, 1, 0, 0, 0, 211, 213, 1, 0, 0, 0, 212, 206, 1, 0, 0, 0, 213,
		216, 1, 0, 0, 0, 214, 212, 1, 0, 0, 0, 214, 215, 1, 0, 0, 0, 215, 48, 1,
		0, 0, 0, 216, 214, 1, 0, 0, 0, 217, 219, 7, 8, 0, 0, 218, 217, 1, 0, 0,
		0, 219, 220, 1, 0, 0, 0, 220, 218, 1, 0, 0, 0, 220, 221, 1, 0, 0, 0, 221,
		50, 1, 0, 0, 0, 21, 0, 54, 61, 68, 70, 84, 149, 152, 156, 161, 166, 172,
		174, 181, 186, 192, 197, 204, 210, 214, 220, 1, 6, 0, 0,
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
	NumscriptLexerWS                         = 1
	NumscriptLexerNEWLINE                    = 2
	NumscriptLexerMULTILINE_COMMENT          = 3
	NumscriptLexerLINE_COMMENT               = 4
	NumscriptLexerVARS                       = 5
	NumscriptLexerMAX                        = 6
	NumscriptLexerSOURCE                     = 7
	NumscriptLexerDESTINATION                = 8
	NumscriptLexerSEND                       = 9
	NumscriptLexerFROM                       = 10
	NumscriptLexerTO                         = 11
	NumscriptLexerLPARENS                    = 12
	NumscriptLexerRPARENS                    = 13
	NumscriptLexerLBRACKET                   = 14
	NumscriptLexerRBRACKET                   = 15
	NumscriptLexerLBRACE                     = 16
	NumscriptLexerRBRACE                     = 17
	NumscriptLexerEQ                         = 18
	NumscriptLexerRATIO_PORTION_LITERAL      = 19
	NumscriptLexerPERCENTAGE_PORTION_LITERAL = 20
	NumscriptLexerTYPE_IDENT                 = 21
	NumscriptLexerNUMBER                     = 22
	NumscriptLexerVARIABLE_NAME              = 23
	NumscriptLexerACCOUNT                    = 24
	NumscriptLexerASSET                      = 25
)
