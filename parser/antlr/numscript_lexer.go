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
		"'send'", "'from'", "'up'", "'to'", "'remaining'", "'allowing'", "'unbounded'",
		"'overdraft'", "'('", "')'", "'['", "']'", "'{'", "'}'", "'='",
	}
	staticData.SymbolicNames = []string{
		"", "WS", "NEWLINE", "MULTILINE_COMMENT", "LINE_COMMENT", "VARS", "MAX",
		"SOURCE", "DESTINATION", "SEND", "FROM", "UP", "TO", "REMAINING", "ALLOWING",
		"UNBOUNDED", "OVERDRAFT", "LPARENS", "RPARENS", "LBRACKET", "RBRACKET",
		"LBRACE", "RBRACE", "EQ", "RATIO_PORTION_LITERAL", "PERCENTAGE_PORTION_LITERAL",
		"TYPE_IDENT", "NUMBER", "VARIABLE_NAME", "ACCOUNT", "ASSET",
	}
	staticData.RuleNames = []string{
		"WS", "NEWLINE", "MULTILINE_COMMENT", "LINE_COMMENT", "VARS", "MAX",
		"SOURCE", "DESTINATION", "SEND", "FROM", "UP", "TO", "REMAINING", "ALLOWING",
		"UNBOUNDED", "OVERDRAFT", "LPARENS", "RPARENS", "LBRACKET", "RBRACKET",
		"LBRACE", "RBRACE", "EQ", "RATIO_PORTION_LITERAL", "PERCENTAGE_PORTION_LITERAL",
		"TYPE_IDENT", "NUMBER", "VARIABLE_NAME", "ACCOUNT", "ASSET",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 30, 276, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 1, 0, 4, 0, 63,
		8, 0, 11, 0, 12, 0, 64, 1, 0, 1, 0, 1, 1, 4, 1, 70, 8, 1, 11, 1, 12, 1,
		71, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 79, 8, 2, 10, 2, 12, 2, 82, 9,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 93, 8, 3,
		10, 3, 12, 3, 96, 9, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1,
		20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 4, 23, 200, 8, 23, 11, 23, 12, 23,
		201, 1, 23, 3, 23, 205, 8, 23, 1, 23, 1, 23, 3, 23, 209, 8, 23, 1, 23,
		4, 23, 212, 8, 23, 11, 23, 12, 23, 213, 1, 24, 4, 24, 217, 8, 24, 11, 24,
		12, 24, 218, 1, 24, 1, 24, 4, 24, 223, 8, 24, 11, 24, 12, 24, 224, 3, 24,
		227, 8, 24, 1, 24, 1, 24, 1, 25, 4, 25, 232, 8, 25, 11, 25, 12, 25, 233,
		1, 26, 4, 26, 237, 8, 26, 11, 26, 12, 26, 238, 1, 27, 1, 27, 4, 27, 243,
		8, 27, 11, 27, 12, 27, 244, 1, 27, 5, 27, 248, 8, 27, 10, 27, 12, 27, 251,
		9, 27, 3, 27, 253, 8, 27, 1, 28, 1, 28, 4, 28, 257, 8, 28, 11, 28, 12,
		28, 258, 1, 28, 1, 28, 4, 28, 263, 8, 28, 11, 28, 12, 28, 264, 5, 28, 267,
		8, 28, 10, 28, 12, 28, 270, 9, 28, 1, 29, 4, 29, 273, 8, 29, 11, 29, 12,
		29, 274, 2, 80, 94, 0, 30, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7,
		15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33,
		17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51,
		26, 53, 27, 55, 28, 57, 29, 59, 30, 1, 0, 9, 3, 0, 9, 10, 13, 13, 32, 32,
		2, 0, 10, 10, 13, 13, 1, 0, 48, 57, 1, 0, 32, 32, 1, 0, 97, 122, 2, 0,
		95, 95, 97, 122, 3, 0, 48, 57, 95, 95, 97, 122, 5, 0, 45, 45, 48, 57, 65,
		90, 95, 95, 97, 122, 2, 0, 47, 57, 65, 90, 296, 0, 1, 1, 0, 0, 0, 0, 3,
		1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11,
		1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0,
		19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0,
		0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0,
		0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0,
		0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1,
		0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57,
		1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 1, 62, 1, 0, 0, 0, 3, 69, 1, 0, 0, 0, 5,
		73, 1, 0, 0, 0, 7, 88, 1, 0, 0, 0, 9, 101, 1, 0, 0, 0, 11, 106, 1, 0, 0,
		0, 13, 110, 1, 0, 0, 0, 15, 117, 1, 0, 0, 0, 17, 129, 1, 0, 0, 0, 19, 134,
		1, 0, 0, 0, 21, 139, 1, 0, 0, 0, 23, 142, 1, 0, 0, 0, 25, 145, 1, 0, 0,
		0, 27, 155, 1, 0, 0, 0, 29, 164, 1, 0, 0, 0, 31, 174, 1, 0, 0, 0, 33, 184,
		1, 0, 0, 0, 35, 186, 1, 0, 0, 0, 37, 188, 1, 0, 0, 0, 39, 190, 1, 0, 0,
		0, 41, 192, 1, 0, 0, 0, 43, 194, 1, 0, 0, 0, 45, 196, 1, 0, 0, 0, 47, 199,
		1, 0, 0, 0, 49, 216, 1, 0, 0, 0, 51, 231, 1, 0, 0, 0, 53, 236, 1, 0, 0,
		0, 55, 240, 1, 0, 0, 0, 57, 254, 1, 0, 0, 0, 59, 272, 1, 0, 0, 0, 61, 63,
		7, 0, 0, 0, 62, 61, 1, 0, 0, 0, 63, 64, 1, 0, 0, 0, 64, 62, 1, 0, 0, 0,
		64, 65, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 67, 6, 0, 0, 0, 67, 2, 1, 0,
		0, 0, 68, 70, 7, 1, 0, 0, 69, 68, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 69,
		1, 0, 0, 0, 71, 72, 1, 0, 0, 0, 72, 4, 1, 0, 0, 0, 73, 74, 5, 47, 0, 0,
		74, 75, 5, 42, 0, 0, 75, 80, 1, 0, 0, 0, 76, 79, 3, 5, 2, 0, 77, 79, 9,
		0, 0, 0, 78, 76, 1, 0, 0, 0, 78, 77, 1, 0, 0, 0, 79, 82, 1, 0, 0, 0, 80,
		81, 1, 0, 0, 0, 80, 78, 1, 0, 0, 0, 81, 83, 1, 0, 0, 0, 82, 80, 1, 0, 0,
		0, 83, 84, 5, 42, 0, 0, 84, 85, 5, 47, 0, 0, 85, 86, 1, 0, 0, 0, 86, 87,
		6, 2, 0, 0, 87, 6, 1, 0, 0, 0, 88, 89, 5, 47, 0, 0, 89, 90, 5, 47, 0, 0,
		90, 94, 1, 0, 0, 0, 91, 93, 9, 0, 0, 0, 92, 91, 1, 0, 0, 0, 93, 96, 1,
		0, 0, 0, 94, 95, 1, 0, 0, 0, 94, 92, 1, 0, 0, 0, 95, 97, 1, 0, 0, 0, 96,
		94, 1, 0, 0, 0, 97, 98, 3, 3, 1, 0, 98, 99, 1, 0, 0, 0, 99, 100, 6, 3,
		0, 0, 100, 8, 1, 0, 0, 0, 101, 102, 5, 118, 0, 0, 102, 103, 5, 97, 0, 0,
		103, 104, 5, 114, 0, 0, 104, 105, 5, 115, 0, 0, 105, 10, 1, 0, 0, 0, 106,
		107, 5, 109, 0, 0, 107, 108, 5, 97, 0, 0, 108, 109, 5, 120, 0, 0, 109,
		12, 1, 0, 0, 0, 110, 111, 5, 115, 0, 0, 111, 112, 5, 111, 0, 0, 112, 113,
		5, 117, 0, 0, 113, 114, 5, 114, 0, 0, 114, 115, 5, 99, 0, 0, 115, 116,
		5, 101, 0, 0, 116, 14, 1, 0, 0, 0, 117, 118, 5, 100, 0, 0, 118, 119, 5,
		101, 0, 0, 119, 120, 5, 115, 0, 0, 120, 121, 5, 116, 0, 0, 121, 122, 5,
		105, 0, 0, 122, 123, 5, 110, 0, 0, 123, 124, 5, 97, 0, 0, 124, 125, 5,
		116, 0, 0, 125, 126, 5, 105, 0, 0, 126, 127, 5, 111, 0, 0, 127, 128, 5,
		110, 0, 0, 128, 16, 1, 0, 0, 0, 129, 130, 5, 115, 0, 0, 130, 131, 5, 101,
		0, 0, 131, 132, 5, 110, 0, 0, 132, 133, 5, 100, 0, 0, 133, 18, 1, 0, 0,
		0, 134, 135, 5, 102, 0, 0, 135, 136, 5, 114, 0, 0, 136, 137, 5, 111, 0,
		0, 137, 138, 5, 109, 0, 0, 138, 20, 1, 0, 0, 0, 139, 140, 5, 117, 0, 0,
		140, 141, 5, 112, 0, 0, 141, 22, 1, 0, 0, 0, 142, 143, 5, 116, 0, 0, 143,
		144, 5, 111, 0, 0, 144, 24, 1, 0, 0, 0, 145, 146, 5, 114, 0, 0, 146, 147,
		5, 101, 0, 0, 147, 148, 5, 109, 0, 0, 148, 149, 5, 97, 0, 0, 149, 150,
		5, 105, 0, 0, 150, 151, 5, 110, 0, 0, 151, 152, 5, 105, 0, 0, 152, 153,
		5, 110, 0, 0, 153, 154, 5, 103, 0, 0, 154, 26, 1, 0, 0, 0, 155, 156, 5,
		97, 0, 0, 156, 157, 5, 108, 0, 0, 157, 158, 5, 108, 0, 0, 158, 159, 5,
		111, 0, 0, 159, 160, 5, 119, 0, 0, 160, 161, 5, 105, 0, 0, 161, 162, 5,
		110, 0, 0, 162, 163, 5, 103, 0, 0, 163, 28, 1, 0, 0, 0, 164, 165, 5, 117,
		0, 0, 165, 166, 5, 110, 0, 0, 166, 167, 5, 98, 0, 0, 167, 168, 5, 111,
		0, 0, 168, 169, 5, 117, 0, 0, 169, 170, 5, 110, 0, 0, 170, 171, 5, 100,
		0, 0, 171, 172, 5, 101, 0, 0, 172, 173, 5, 100, 0, 0, 173, 30, 1, 0, 0,
		0, 174, 175, 5, 111, 0, 0, 175, 176, 5, 118, 0, 0, 176, 177, 5, 101, 0,
		0, 177, 178, 5, 114, 0, 0, 178, 179, 5, 100, 0, 0, 179, 180, 5, 114, 0,
		0, 180, 181, 5, 97, 0, 0, 181, 182, 5, 102, 0, 0, 182, 183, 5, 116, 0,
		0, 183, 32, 1, 0, 0, 0, 184, 185, 5, 40, 0, 0, 185, 34, 1, 0, 0, 0, 186,
		187, 5, 41, 0, 0, 187, 36, 1, 0, 0, 0, 188, 189, 5, 91, 0, 0, 189, 38,
		1, 0, 0, 0, 190, 191, 5, 93, 0, 0, 191, 40, 1, 0, 0, 0, 192, 193, 5, 123,
		0, 0, 193, 42, 1, 0, 0, 0, 194, 195, 5, 125, 0, 0, 195, 44, 1, 0, 0, 0,
		196, 197, 5, 61, 0, 0, 197, 46, 1, 0, 0, 0, 198, 200, 7, 2, 0, 0, 199,
		198, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0, 201, 199, 1, 0, 0, 0, 201, 202,
		1, 0, 0, 0, 202, 204, 1, 0, 0, 0, 203, 205, 7, 3, 0, 0, 204, 203, 1, 0,
		0, 0, 204, 205, 1, 0, 0, 0, 205, 206, 1, 0, 0, 0, 206, 208, 5, 47, 0, 0,
		207, 209, 7, 3, 0, 0, 208, 207, 1, 0, 0, 0, 208, 209, 1, 0, 0, 0, 209,
		211, 1, 0, 0, 0, 210, 212, 7, 2, 0, 0, 211, 210, 1, 0, 0, 0, 212, 213,
		1, 0, 0, 0, 213, 211, 1, 0, 0, 0, 213, 214, 1, 0, 0, 0, 214, 48, 1, 0,
		0, 0, 215, 217, 7, 2, 0, 0, 216, 215, 1, 0, 0, 0, 217, 218, 1, 0, 0, 0,
		218, 216, 1, 0, 0, 0, 218, 219, 1, 0, 0, 0, 219, 226, 1, 0, 0, 0, 220,
		222, 5, 46, 0, 0, 221, 223, 7, 2, 0, 0, 222, 221, 1, 0, 0, 0, 223, 224,
		1, 0, 0, 0, 224, 222, 1, 0, 0, 0, 224, 225, 1, 0, 0, 0, 225, 227, 1, 0,
		0, 0, 226, 220, 1, 0, 0, 0, 226, 227, 1, 0, 0, 0, 227, 228, 1, 0, 0, 0,
		228, 229, 5, 37, 0, 0, 229, 50, 1, 0, 0, 0, 230, 232, 7, 4, 0, 0, 231,
		230, 1, 0, 0, 0, 232, 233, 1, 0, 0, 0, 233, 231, 1, 0, 0, 0, 233, 234,
		1, 0, 0, 0, 234, 52, 1, 0, 0, 0, 235, 237, 7, 2, 0, 0, 236, 235, 1, 0,
		0, 0, 237, 238, 1, 0, 0, 0, 238, 236, 1, 0, 0, 0, 238, 239, 1, 0, 0, 0,
		239, 54, 1, 0, 0, 0, 240, 252, 5, 36, 0, 0, 241, 243, 7, 5, 0, 0, 242,
		241, 1, 0, 0, 0, 243, 244, 1, 0, 0, 0, 244, 242, 1, 0, 0, 0, 244, 245,
		1, 0, 0, 0, 245, 249, 1, 0, 0, 0, 246, 248, 7, 6, 0, 0, 247, 246, 1, 0,
		0, 0, 248, 251, 1, 0, 0, 0, 249, 247, 1, 0, 0, 0, 249, 250, 1, 0, 0, 0,
		250, 253, 1, 0, 0, 0, 251, 249, 1, 0, 0, 0, 252, 242, 1, 0, 0, 0, 252,
		253, 1, 0, 0, 0, 253, 56, 1, 0, 0, 0, 254, 256, 5, 64, 0, 0, 255, 257,
		7, 7, 0, 0, 256, 255, 1, 0, 0, 0, 257, 258, 1, 0, 0, 0, 258, 256, 1, 0,
		0, 0, 258, 259, 1, 0, 0, 0, 259, 268, 1, 0, 0, 0, 260, 262, 5, 58, 0, 0,
		261, 263, 7, 7, 0, 0, 262, 261, 1, 0, 0, 0, 263, 264, 1, 0, 0, 0, 264,
		262, 1, 0, 0, 0, 264, 265, 1, 0, 0, 0, 265, 267, 1, 0, 0, 0, 266, 260,
		1, 0, 0, 0, 267, 270, 1, 0, 0, 0, 268, 266, 1, 0, 0, 0, 268, 269, 1, 0,
		0, 0, 269, 58, 1, 0, 0, 0, 270, 268, 1, 0, 0, 0, 271, 273, 7, 8, 0, 0,
		272, 271, 1, 0, 0, 0, 273, 274, 1, 0, 0, 0, 274, 272, 1, 0, 0, 0, 274,
		275, 1, 0, 0, 0, 275, 60, 1, 0, 0, 0, 22, 0, 64, 71, 78, 80, 94, 201, 204,
		208, 213, 218, 224, 226, 233, 238, 244, 249, 252, 258, 264, 268, 274, 1,
		6, 0, 0,
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
	NumscriptLexerUP                         = 11
	NumscriptLexerTO                         = 12
	NumscriptLexerREMAINING                  = 13
	NumscriptLexerALLOWING                   = 14
	NumscriptLexerUNBOUNDED                  = 15
	NumscriptLexerOVERDRAFT                  = 16
	NumscriptLexerLPARENS                    = 17
	NumscriptLexerRPARENS                    = 18
	NumscriptLexerLBRACKET                   = 19
	NumscriptLexerRBRACKET                   = 20
	NumscriptLexerLBRACE                     = 21
	NumscriptLexerRBRACE                     = 22
	NumscriptLexerEQ                         = 23
	NumscriptLexerRATIO_PORTION_LITERAL      = 24
	NumscriptLexerPERCENTAGE_PORTION_LITERAL = 25
	NumscriptLexerTYPE_IDENT                 = 26
	NumscriptLexerNUMBER                     = 27
	NumscriptLexerVARIABLE_NAME              = 28
	NumscriptLexerACCOUNT                    = 29
	NumscriptLexerASSET                      = 30
)
