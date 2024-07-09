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
		"", "", "", "", "", "'source'", "'destination'", "'send'", "'from'",
		"'('", "')'", "'['", "']'", "'{'", "'}'", "'='",
	}
	staticData.SymbolicNames = []string{
		"", "WS", "NEWLINE", "MULTILINE_COMMENT", "LINE_COMMENT", "SOURCE",
		"DESTINATION", "SEND", "FROM", "LPARENS", "RPARENS", "LBRACKET", "RBRACKET",
		"LBRACE", "RBRACE", "EQ", "RATIO_PORTION_LITERAL", "PERCENTAGE_PORTION_LITERAL",
		"NUMBER", "VARIABLE_NAME", "ACCOUNT", "ASSET",
	}
	staticData.RuleNames = []string{
		"WS", "NEWLINE", "MULTILINE_COMMENT", "LINE_COMMENT", "SOURCE", "DESTINATION",
		"SEND", "FROM", "LPARENS", "RPARENS", "LBRACKET", "RBRACKET", "LBRACE",
		"RBRACE", "EQ", "RATIO_PORTION_LITERAL", "PERCENTAGE_PORTION_LITERAL",
		"NUMBER", "VARIABLE_NAME", "ACCOUNT", "ASSET",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 21, 197, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 1, 0, 4, 0, 45, 8, 0, 11, 0, 12, 0, 46, 1, 0, 1, 0, 1, 1, 4, 1, 52,
		8, 1, 11, 1, 12, 1, 53, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 61, 8, 2, 10,
		2, 12, 2, 64, 9, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1,
		3, 5, 3, 75, 8, 3, 10, 3, 12, 3, 78, 9, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1,
		11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 4, 15, 128, 8, 15,
		11, 15, 12, 15, 129, 1, 15, 3, 15, 133, 8, 15, 1, 15, 1, 15, 3, 15, 137,
		8, 15, 1, 15, 4, 15, 140, 8, 15, 11, 15, 12, 15, 141, 1, 16, 4, 16, 145,
		8, 16, 11, 16, 12, 16, 146, 1, 16, 1, 16, 4, 16, 151, 8, 16, 11, 16, 12,
		16, 152, 3, 16, 155, 8, 16, 1, 16, 1, 16, 1, 17, 4, 17, 160, 8, 17, 11,
		17, 12, 17, 161, 1, 18, 1, 18, 4, 18, 166, 8, 18, 11, 18, 12, 18, 167,
		1, 18, 5, 18, 171, 8, 18, 10, 18, 12, 18, 174, 9, 18, 1, 19, 1, 19, 4,
		19, 178, 8, 19, 11, 19, 12, 19, 179, 1, 19, 1, 19, 4, 19, 184, 8, 19, 11,
		19, 12, 19, 185, 5, 19, 188, 8, 19, 10, 19, 12, 19, 191, 9, 19, 1, 20,
		4, 20, 194, 8, 20, 11, 20, 12, 20, 195, 2, 62, 76, 0, 21, 1, 1, 3, 2, 5,
		3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25,
		13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 1,
		0, 8, 3, 0, 9, 10, 13, 13, 32, 32, 2, 0, 10, 10, 13, 13, 1, 0, 48, 57,
		1, 0, 32, 32, 2, 0, 95, 95, 97, 122, 3, 0, 48, 57, 95, 95, 97, 122, 5,
		0, 45, 45, 48, 57, 65, 90, 95, 95, 97, 122, 2, 0, 47, 57, 65, 90, 215,
		0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0,
		0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0,
		0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0,
		0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1,
		0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39,
		1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 1, 44, 1, 0, 0, 0, 3, 51, 1, 0, 0, 0, 5,
		55, 1, 0, 0, 0, 7, 70, 1, 0, 0, 0, 9, 83, 1, 0, 0, 0, 11, 90, 1, 0, 0,
		0, 13, 102, 1, 0, 0, 0, 15, 107, 1, 0, 0, 0, 17, 112, 1, 0, 0, 0, 19, 114,
		1, 0, 0, 0, 21, 116, 1, 0, 0, 0, 23, 118, 1, 0, 0, 0, 25, 120, 1, 0, 0,
		0, 27, 122, 1, 0, 0, 0, 29, 124, 1, 0, 0, 0, 31, 127, 1, 0, 0, 0, 33, 144,
		1, 0, 0, 0, 35, 159, 1, 0, 0, 0, 37, 163, 1, 0, 0, 0, 39, 175, 1, 0, 0,
		0, 41, 193, 1, 0, 0, 0, 43, 45, 7, 0, 0, 0, 44, 43, 1, 0, 0, 0, 45, 46,
		1, 0, 0, 0, 46, 44, 1, 0, 0, 0, 46, 47, 1, 0, 0, 0, 47, 48, 1, 0, 0, 0,
		48, 49, 6, 0, 0, 0, 49, 2, 1, 0, 0, 0, 50, 52, 7, 1, 0, 0, 51, 50, 1, 0,
		0, 0, 52, 53, 1, 0, 0, 0, 53, 51, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0, 54, 4,
		1, 0, 0, 0, 55, 56, 5, 47, 0, 0, 56, 57, 5, 42, 0, 0, 57, 62, 1, 0, 0,
		0, 58, 61, 3, 5, 2, 0, 59, 61, 9, 0, 0, 0, 60, 58, 1, 0, 0, 0, 60, 59,
		1, 0, 0, 0, 61, 64, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 62, 60, 1, 0, 0, 0,
		63, 65, 1, 0, 0, 0, 64, 62, 1, 0, 0, 0, 65, 66, 5, 42, 0, 0, 66, 67, 5,
		47, 0, 0, 67, 68, 1, 0, 0, 0, 68, 69, 6, 2, 0, 0, 69, 6, 1, 0, 0, 0, 70,
		71, 5, 47, 0, 0, 71, 72, 5, 47, 0, 0, 72, 76, 1, 0, 0, 0, 73, 75, 9, 0,
		0, 0, 74, 73, 1, 0, 0, 0, 75, 78, 1, 0, 0, 0, 76, 77, 1, 0, 0, 0, 76, 74,
		1, 0, 0, 0, 77, 79, 1, 0, 0, 0, 78, 76, 1, 0, 0, 0, 79, 80, 3, 3, 1, 0,
		80, 81, 1, 0, 0, 0, 81, 82, 6, 3, 0, 0, 82, 8, 1, 0, 0, 0, 83, 84, 5, 115,
		0, 0, 84, 85, 5, 111, 0, 0, 85, 86, 5, 117, 0, 0, 86, 87, 5, 114, 0, 0,
		87, 88, 5, 99, 0, 0, 88, 89, 5, 101, 0, 0, 89, 10, 1, 0, 0, 0, 90, 91,
		5, 100, 0, 0, 91, 92, 5, 101, 0, 0, 92, 93, 5, 115, 0, 0, 93, 94, 5, 116,
		0, 0, 94, 95, 5, 105, 0, 0, 95, 96, 5, 110, 0, 0, 96, 97, 5, 97, 0, 0,
		97, 98, 5, 116, 0, 0, 98, 99, 5, 105, 0, 0, 99, 100, 5, 111, 0, 0, 100,
		101, 5, 110, 0, 0, 101, 12, 1, 0, 0, 0, 102, 103, 5, 115, 0, 0, 103, 104,
		5, 101, 0, 0, 104, 105, 5, 110, 0, 0, 105, 106, 5, 100, 0, 0, 106, 14,
		1, 0, 0, 0, 107, 108, 5, 102, 0, 0, 108, 109, 5, 114, 0, 0, 109, 110, 5,
		111, 0, 0, 110, 111, 5, 109, 0, 0, 111, 16, 1, 0, 0, 0, 112, 113, 5, 40,
		0, 0, 113, 18, 1, 0, 0, 0, 114, 115, 5, 41, 0, 0, 115, 20, 1, 0, 0, 0,
		116, 117, 5, 91, 0, 0, 117, 22, 1, 0, 0, 0, 118, 119, 5, 93, 0, 0, 119,
		24, 1, 0, 0, 0, 120, 121, 5, 123, 0, 0, 121, 26, 1, 0, 0, 0, 122, 123,
		5, 125, 0, 0, 123, 28, 1, 0, 0, 0, 124, 125, 5, 61, 0, 0, 125, 30, 1, 0,
		0, 0, 126, 128, 7, 2, 0, 0, 127, 126, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0,
		129, 127, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130, 132, 1, 0, 0, 0, 131,
		133, 7, 3, 0, 0, 132, 131, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 134,
		1, 0, 0, 0, 134, 136, 5, 47, 0, 0, 135, 137, 7, 3, 0, 0, 136, 135, 1, 0,
		0, 0, 136, 137, 1, 0, 0, 0, 137, 139, 1, 0, 0, 0, 138, 140, 7, 2, 0, 0,
		139, 138, 1, 0, 0, 0, 140, 141, 1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 141,
		142, 1, 0, 0, 0, 142, 32, 1, 0, 0, 0, 143, 145, 7, 2, 0, 0, 144, 143, 1,
		0, 0, 0, 145, 146, 1, 0, 0, 0, 146, 144, 1, 0, 0, 0, 146, 147, 1, 0, 0,
		0, 147, 154, 1, 0, 0, 0, 148, 150, 5, 46, 0, 0, 149, 151, 7, 2, 0, 0, 150,
		149, 1, 0, 0, 0, 151, 152, 1, 0, 0, 0, 152, 150, 1, 0, 0, 0, 152, 153,
		1, 0, 0, 0, 153, 155, 1, 0, 0, 0, 154, 148, 1, 0, 0, 0, 154, 155, 1, 0,
		0, 0, 155, 156, 1, 0, 0, 0, 156, 157, 5, 37, 0, 0, 157, 34, 1, 0, 0, 0,
		158, 160, 7, 2, 0, 0, 159, 158, 1, 0, 0, 0, 160, 161, 1, 0, 0, 0, 161,
		159, 1, 0, 0, 0, 161, 162, 1, 0, 0, 0, 162, 36, 1, 0, 0, 0, 163, 165, 5,
		36, 0, 0, 164, 166, 7, 4, 0, 0, 165, 164, 1, 0, 0, 0, 166, 167, 1, 0, 0,
		0, 167, 165, 1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168, 172, 1, 0, 0, 0, 169,
		171, 7, 5, 0, 0, 170, 169, 1, 0, 0, 0, 171, 174, 1, 0, 0, 0, 172, 170,
		1, 0, 0, 0, 172, 173, 1, 0, 0, 0, 173, 38, 1, 0, 0, 0, 174, 172, 1, 0,
		0, 0, 175, 177, 5, 64, 0, 0, 176, 178, 7, 6, 0, 0, 177, 176, 1, 0, 0, 0,
		178, 179, 1, 0, 0, 0, 179, 177, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180,
		189, 1, 0, 0, 0, 181, 183, 5, 58, 0, 0, 182, 184, 7, 6, 0, 0, 183, 182,
		1, 0, 0, 0, 184, 185, 1, 0, 0, 0, 185, 183, 1, 0, 0, 0, 185, 186, 1, 0,
		0, 0, 186, 188, 1, 0, 0, 0, 187, 181, 1, 0, 0, 0, 188, 191, 1, 0, 0, 0,
		189, 187, 1, 0, 0, 0, 189, 190, 1, 0, 0, 0, 190, 40, 1, 0, 0, 0, 191, 189,
		1, 0, 0, 0, 192, 194, 7, 7, 0, 0, 193, 192, 1, 0, 0, 0, 194, 195, 1, 0,
		0, 0, 195, 193, 1, 0, 0, 0, 195, 196, 1, 0, 0, 0, 196, 42, 1, 0, 0, 0,
		20, 0, 46, 53, 60, 62, 76, 129, 132, 136, 141, 146, 152, 154, 161, 167,
		172, 179, 185, 189, 195, 1, 6, 0, 0,
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
	NumscriptLexerSOURCE                     = 5
	NumscriptLexerDESTINATION                = 6
	NumscriptLexerSEND                       = 7
	NumscriptLexerFROM                       = 8
	NumscriptLexerLPARENS                    = 9
	NumscriptLexerRPARENS                    = 10
	NumscriptLexerLBRACKET                   = 11
	NumscriptLexerRBRACKET                   = 12
	NumscriptLexerLBRACE                     = 13
	NumscriptLexerRBRACE                     = 14
	NumscriptLexerEQ                         = 15
	NumscriptLexerRATIO_PORTION_LITERAL      = 16
	NumscriptLexerPERCENTAGE_PORTION_LITERAL = 17
	NumscriptLexerNUMBER                     = 18
	NumscriptLexerVARIABLE_NAME              = 19
	NumscriptLexerACCOUNT                    = 20
	NumscriptLexerASSET                      = 21
)
