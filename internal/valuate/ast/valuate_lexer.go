// Code generated from valuate.g4 by ANTLR 4.10.1. DO NOT EDIT.

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

type valuateLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var valuatelexerLexerStaticData struct {
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

func valuatelexerLexerInit() {
	staticData := &valuatelexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'!'", "'true'", "'false'", "'nil'", "'('", "')'", "'{'", "'}'",
		"'['", "']'", "','", "'.'", "'||'", "'&&'", "'=='", "'!='", "'<'", "'<='",
		"'>'", "'>='", "'/'", "'+'", "'-'", "'*'", "'%'",
	}
	staticData.symbolicNames = []string{
		"", "", "TRUE", "FALSE", "NIL_LIT", "LP", "RP", "L_CURLY", "R_CURLY",
		"L_BRACKET", "R_BRACKET", "COMMA", "DOT", "LOGICAL_OR", "LOGICAL_AND",
		"EQUALS", "NOT_EQUALS", "LESS", "LESS_OR_EQUALS", "GREATER", "GREATER_OR_EQUALS",
		"DIV", "PLUS", "MINUS", "STAR", "MODULUS", "SKIP_", "WHITESPACE", "IDENTIFIER",
		"STRING", "STRING_LITERAL", "BYTES_LITERAL", "INT", "DECIMAL_LIT", "FLOAT_NUMBER",
	}
	staticData.ruleNames = []string{
		"T__0", "TRUE", "FALSE", "NIL_LIT", "LP", "RP", "L_CURLY", "R_CURLY",
		"L_BRACKET", "R_BRACKET", "COMMA", "DOT", "LOGICAL_OR", "LOGICAL_AND",
		"EQUALS", "NOT_EQUALS", "LESS", "LESS_OR_EQUALS", "GREATER", "GREATER_OR_EQUALS",
		"DIV", "PLUS", "MINUS", "STAR", "MODULUS", "SKIP_", "WHITESPACE", "IDENTIFIER",
		"STRING", "STRING_LITERAL", "BYTES_LITERAL", "INT", "DECIMAL_LIT", "FLOAT_NUMBER",
		"LETTER", "NON_ZERO_DIGIT", "DIGIT", "OCT_DIGIT", "HEX_DIGIT", "BIN_DIGIT",
		"POINT_FLOAT", "EXPONENT_FLOAT", "INT_PART", "FRACTION", "EXPONENT",
		"SHORT_BYTES", "LONG_BYTES", "LONG_BYTES_ITEM", "SHORT_BYTES_CHAR_NO_SINGLE_QUOTE",
		"SHORT_BYTES_CHAR_NO_DOUBLE_QUOTE", "LONG_BYTES_CHAR", "BYTES_ESCAPE_SEQ",
		"SPACES", "LINE_JOINING", "ESC", "UNICODE", "SAFECODEPOINT", "HEX",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 34, 387, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6,
		1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12,
		1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1,
		15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19,
		1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1,
		25, 1, 25, 3, 25, 185, 8, 25, 1, 25, 1, 25, 1, 26, 4, 26, 190, 8, 26, 11,
		26, 12, 26, 191, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 5, 27, 199, 8, 27,
		10, 27, 12, 27, 202, 9, 27, 1, 28, 1, 28, 1, 28, 3, 28, 207, 8, 28, 1,
		29, 1, 29, 1, 29, 5, 29, 212, 8, 29, 10, 29, 12, 29, 215, 9, 29, 1, 29,
		1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 3, 30, 224, 8, 30, 1, 30, 1,
		30, 3, 30, 228, 8, 30, 1, 31, 1, 31, 3, 31, 232, 8, 31, 1, 32, 1, 32, 5,
		32, 236, 8, 32, 10, 32, 12, 32, 239, 9, 32, 1, 33, 1, 33, 3, 33, 243, 8,
		33, 1, 34, 3, 34, 246, 8, 34, 1, 35, 1, 35, 1, 36, 1, 36, 1, 37, 1, 37,
		1, 38, 1, 38, 1, 39, 1, 39, 1, 40, 3, 40, 259, 8, 40, 1, 40, 1, 40, 1,
		40, 1, 40, 3, 40, 265, 8, 40, 1, 41, 1, 41, 3, 41, 269, 8, 41, 1, 41, 1,
		41, 1, 42, 4, 42, 274, 8, 42, 11, 42, 12, 42, 275, 1, 43, 1, 43, 4, 43,
		280, 8, 43, 11, 43, 12, 43, 281, 1, 44, 1, 44, 3, 44, 286, 8, 44, 1, 44,
		4, 44, 289, 8, 44, 11, 44, 12, 44, 290, 1, 45, 1, 45, 1, 45, 5, 45, 296,
		8, 45, 10, 45, 12, 45, 299, 9, 45, 1, 45, 1, 45, 1, 45, 1, 45, 5, 45, 305,
		8, 45, 10, 45, 12, 45, 308, 9, 45, 1, 45, 3, 45, 311, 8, 45, 1, 46, 1,
		46, 1, 46, 1, 46, 1, 46, 5, 46, 318, 8, 46, 10, 46, 12, 46, 321, 9, 46,
		1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 5, 46, 331, 8,
		46, 10, 46, 12, 46, 334, 9, 46, 1, 46, 1, 46, 1, 46, 3, 46, 339, 8, 46,
		1, 47, 1, 47, 3, 47, 343, 8, 47, 1, 48, 3, 48, 346, 8, 48, 1, 49, 3, 49,
		349, 8, 49, 1, 50, 3, 50, 352, 8, 50, 1, 51, 1, 51, 1, 51, 1, 52, 4, 52,
		358, 8, 52, 11, 52, 12, 52, 359, 1, 53, 1, 53, 3, 53, 364, 8, 53, 1, 53,
		3, 53, 367, 8, 53, 1, 53, 1, 53, 3, 53, 371, 8, 53, 1, 54, 1, 54, 1, 54,
		3, 54, 376, 8, 54, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 56, 1,
		56, 1, 57, 1, 57, 2, 319, 332, 0, 58, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11,
		6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15,
		31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24,
		49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33,
		67, 34, 69, 0, 71, 0, 73, 0, 75, 0, 77, 0, 79, 0, 81, 0, 83, 0, 85, 0,
		87, 0, 89, 0, 91, 0, 93, 0, 95, 0, 97, 0, 99, 0, 101, 0, 103, 0, 105, 0,
		107, 0, 109, 0, 111, 0, 113, 0, 115, 0, 1, 0, 18, 3, 0, 9, 10, 13, 13,
		32, 32, 2, 0, 66, 66, 98, 98, 2, 0, 82, 82, 114, 114, 3, 0, 65, 90, 95,
		95, 97, 122, 1, 0, 49, 57, 1, 0, 48, 57, 1, 0, 48, 55, 3, 0, 48, 57, 65,
		70, 97, 102, 1, 0, 48, 49, 2, 0, 69, 69, 101, 101, 2, 0, 43, 43, 45, 45,
		5, 0, 0, 9, 11, 12, 14, 38, 40, 91, 93, 127, 5, 0, 0, 9, 11, 12, 14, 33,
		35, 91, 93, 127, 2, 0, 0, 91, 93, 127, 1, 0, 0, 127, 2, 0, 9, 9, 32, 32,
		8, 0, 34, 34, 47, 47, 92, 92, 98, 98, 102, 102, 110, 110, 114, 114, 116,
		116, 3, 0, 0, 31, 34, 34, 92, 92, 397, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0,
		0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0,
		0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0,
		0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1,
		0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35,
		1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0,
		43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0,
		0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0,
		0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0,
		0, 0, 0, 67, 1, 0, 0, 0, 1, 117, 1, 0, 0, 0, 3, 119, 1, 0, 0, 0, 5, 124,
		1, 0, 0, 0, 7, 130, 1, 0, 0, 0, 9, 134, 1, 0, 0, 0, 11, 136, 1, 0, 0, 0,
		13, 138, 1, 0, 0, 0, 15, 140, 1, 0, 0, 0, 17, 142, 1, 0, 0, 0, 19, 144,
		1, 0, 0, 0, 21, 146, 1, 0, 0, 0, 23, 148, 1, 0, 0, 0, 25, 150, 1, 0, 0,
		0, 27, 153, 1, 0, 0, 0, 29, 156, 1, 0, 0, 0, 31, 159, 1, 0, 0, 0, 33, 162,
		1, 0, 0, 0, 35, 164, 1, 0, 0, 0, 37, 167, 1, 0, 0, 0, 39, 169, 1, 0, 0,
		0, 41, 172, 1, 0, 0, 0, 43, 174, 1, 0, 0, 0, 45, 176, 1, 0, 0, 0, 47, 178,
		1, 0, 0, 0, 49, 180, 1, 0, 0, 0, 51, 184, 1, 0, 0, 0, 53, 189, 1, 0, 0,
		0, 55, 195, 1, 0, 0, 0, 57, 206, 1, 0, 0, 0, 59, 208, 1, 0, 0, 0, 61, 223,
		1, 0, 0, 0, 63, 231, 1, 0, 0, 0, 65, 233, 1, 0, 0, 0, 67, 242, 1, 0, 0,
		0, 69, 245, 1, 0, 0, 0, 71, 247, 1, 0, 0, 0, 73, 249, 1, 0, 0, 0, 75, 251,
		1, 0, 0, 0, 77, 253, 1, 0, 0, 0, 79, 255, 1, 0, 0, 0, 81, 264, 1, 0, 0,
		0, 83, 268, 1, 0, 0, 0, 85, 273, 1, 0, 0, 0, 87, 277, 1, 0, 0, 0, 89, 283,
		1, 0, 0, 0, 91, 310, 1, 0, 0, 0, 93, 338, 1, 0, 0, 0, 95, 342, 1, 0, 0,
		0, 97, 345, 1, 0, 0, 0, 99, 348, 1, 0, 0, 0, 101, 351, 1, 0, 0, 0, 103,
		353, 1, 0, 0, 0, 105, 357, 1, 0, 0, 0, 107, 361, 1, 0, 0, 0, 109, 372,
		1, 0, 0, 0, 111, 377, 1, 0, 0, 0, 113, 383, 1, 0, 0, 0, 115, 385, 1, 0,
		0, 0, 117, 118, 5, 33, 0, 0, 118, 2, 1, 0, 0, 0, 119, 120, 5, 116, 0, 0,
		120, 121, 5, 114, 0, 0, 121, 122, 5, 117, 0, 0, 122, 123, 5, 101, 0, 0,
		123, 4, 1, 0, 0, 0, 124, 125, 5, 102, 0, 0, 125, 126, 5, 97, 0, 0, 126,
		127, 5, 108, 0, 0, 127, 128, 5, 115, 0, 0, 128, 129, 5, 101, 0, 0, 129,
		6, 1, 0, 0, 0, 130, 131, 5, 110, 0, 0, 131, 132, 5, 105, 0, 0, 132, 133,
		5, 108, 0, 0, 133, 8, 1, 0, 0, 0, 134, 135, 5, 40, 0, 0, 135, 10, 1, 0,
		0, 0, 136, 137, 5, 41, 0, 0, 137, 12, 1, 0, 0, 0, 138, 139, 5, 123, 0,
		0, 139, 14, 1, 0, 0, 0, 140, 141, 5, 125, 0, 0, 141, 16, 1, 0, 0, 0, 142,
		143, 5, 91, 0, 0, 143, 18, 1, 0, 0, 0, 144, 145, 5, 93, 0, 0, 145, 20,
		1, 0, 0, 0, 146, 147, 5, 44, 0, 0, 147, 22, 1, 0, 0, 0, 148, 149, 5, 46,
		0, 0, 149, 24, 1, 0, 0, 0, 150, 151, 5, 124, 0, 0, 151, 152, 5, 124, 0,
		0, 152, 26, 1, 0, 0, 0, 153, 154, 5, 38, 0, 0, 154, 155, 5, 38, 0, 0, 155,
		28, 1, 0, 0, 0, 156, 157, 5, 61, 0, 0, 157, 158, 5, 61, 0, 0, 158, 30,
		1, 0, 0, 0, 159, 160, 5, 33, 0, 0, 160, 161, 5, 61, 0, 0, 161, 32, 1, 0,
		0, 0, 162, 163, 5, 60, 0, 0, 163, 34, 1, 0, 0, 0, 164, 165, 5, 60, 0, 0,
		165, 166, 5, 61, 0, 0, 166, 36, 1, 0, 0, 0, 167, 168, 5, 62, 0, 0, 168,
		38, 1, 0, 0, 0, 169, 170, 5, 62, 0, 0, 170, 171, 5, 61, 0, 0, 171, 40,
		1, 0, 0, 0, 172, 173, 5, 47, 0, 0, 173, 42, 1, 0, 0, 0, 174, 175, 5, 43,
		0, 0, 175, 44, 1, 0, 0, 0, 176, 177, 5, 45, 0, 0, 177, 46, 1, 0, 0, 0,
		178, 179, 5, 42, 0, 0, 179, 48, 1, 0, 0, 0, 180, 181, 5, 37, 0, 0, 181,
		50, 1, 0, 0, 0, 182, 185, 3, 105, 52, 0, 183, 185, 3, 107, 53, 0, 184,
		182, 1, 0, 0, 0, 184, 183, 1, 0, 0, 0, 185, 186, 1, 0, 0, 0, 186, 187,
		6, 25, 0, 0, 187, 52, 1, 0, 0, 0, 188, 190, 7, 0, 0, 0, 189, 188, 1, 0,
		0, 0, 190, 191, 1, 0, 0, 0, 191, 189, 1, 0, 0, 0, 191, 192, 1, 0, 0, 0,
		192, 193, 1, 0, 0, 0, 193, 194, 6, 26, 0, 0, 194, 54, 1, 0, 0, 0, 195,
		200, 3, 69, 34, 0, 196, 199, 3, 69, 34, 0, 197, 199, 3, 73, 36, 0, 198,
		196, 1, 0, 0, 0, 198, 197, 1, 0, 0, 0, 199, 202, 1, 0, 0, 0, 200, 198,
		1, 0, 0, 0, 200, 201, 1, 0, 0, 0, 201, 56, 1, 0, 0, 0, 202, 200, 1, 0,
		0, 0, 203, 207, 3, 59, 29, 0, 204, 207, 3, 91, 45, 0, 205, 207, 3, 93,
		46, 0, 206, 203, 1, 0, 0, 0, 206, 204, 1, 0, 0, 0, 206, 205, 1, 0, 0, 0,
		207, 58, 1, 0, 0, 0, 208, 213, 5, 34, 0, 0, 209, 212, 3, 109, 54, 0, 210,
		212, 3, 113, 56, 0, 211, 209, 1, 0, 0, 0, 211, 210, 1, 0, 0, 0, 212, 215,
		1, 0, 0, 0, 213, 211, 1, 0, 0, 0, 213, 214, 1, 0, 0, 0, 214, 216, 1, 0,
		0, 0, 215, 213, 1, 0, 0, 0, 216, 217, 5, 34, 0, 0, 217, 60, 1, 0, 0, 0,
		218, 224, 7, 1, 0, 0, 219, 220, 7, 1, 0, 0, 220, 224, 7, 2, 0, 0, 221,
		222, 7, 2, 0, 0, 222, 224, 7, 1, 0, 0, 223, 218, 1, 0, 0, 0, 223, 219,
		1, 0, 0, 0, 223, 221, 1, 0, 0, 0, 224, 227, 1, 0, 0, 0, 225, 228, 3, 91,
		45, 0, 226, 228, 3, 93, 46, 0, 227, 225, 1, 0, 0, 0, 227, 226, 1, 0, 0,
		0, 228, 62, 1, 0, 0, 0, 229, 232, 3, 65, 32, 0, 230, 232, 5, 48, 0, 0,
		231, 229, 1, 0, 0, 0, 231, 230, 1, 0, 0, 0, 232, 64, 1, 0, 0, 0, 233, 237,
		3, 71, 35, 0, 234, 236, 3, 73, 36, 0, 235, 234, 1, 0, 0, 0, 236, 239, 1,
		0, 0, 0, 237, 235, 1, 0, 0, 0, 237, 238, 1, 0, 0, 0, 238, 66, 1, 0, 0,
		0, 239, 237, 1, 0, 0, 0, 240, 243, 3, 81, 40, 0, 241, 243, 3, 83, 41, 0,
		242, 240, 1, 0, 0, 0, 242, 241, 1, 0, 0, 0, 243, 68, 1, 0, 0, 0, 244, 246,
		7, 3, 0, 0, 245, 244, 1, 0, 0, 0, 246, 70, 1, 0, 0, 0, 247, 248, 7, 4,
		0, 0, 248, 72, 1, 0, 0, 0, 249, 250, 7, 5, 0, 0, 250, 74, 1, 0, 0, 0, 251,
		252, 7, 6, 0, 0, 252, 76, 1, 0, 0, 0, 253, 254, 7, 7, 0, 0, 254, 78, 1,
		0, 0, 0, 255, 256, 7, 8, 0, 0, 256, 80, 1, 0, 0, 0, 257, 259, 3, 85, 42,
		0, 258, 257, 1, 0, 0, 0, 258, 259, 1, 0, 0, 0, 259, 260, 1, 0, 0, 0, 260,
		265, 3, 87, 43, 0, 261, 262, 3, 85, 42, 0, 262, 263, 5, 46, 0, 0, 263,
		265, 1, 0, 0, 0, 264, 258, 1, 0, 0, 0, 264, 261, 1, 0, 0, 0, 265, 82, 1,
		0, 0, 0, 266, 269, 3, 85, 42, 0, 267, 269, 3, 81, 40, 0, 268, 266, 1, 0,
		0, 0, 268, 267, 1, 0, 0, 0, 269, 270, 1, 0, 0, 0, 270, 271, 3, 89, 44,
		0, 271, 84, 1, 0, 0, 0, 272, 274, 3, 73, 36, 0, 273, 272, 1, 0, 0, 0, 274,
		275, 1, 0, 0, 0, 275, 273, 1, 0, 0, 0, 275, 276, 1, 0, 0, 0, 276, 86, 1,
		0, 0, 0, 277, 279, 5, 46, 0, 0, 278, 280, 3, 73, 36, 0, 279, 278, 1, 0,
		0, 0, 280, 281, 1, 0, 0, 0, 281, 279, 1, 0, 0, 0, 281, 282, 1, 0, 0, 0,
		282, 88, 1, 0, 0, 0, 283, 285, 7, 9, 0, 0, 284, 286, 7, 10, 0, 0, 285,
		284, 1, 0, 0, 0, 285, 286, 1, 0, 0, 0, 286, 288, 1, 0, 0, 0, 287, 289,
		3, 73, 36, 0, 288, 287, 1, 0, 0, 0, 289, 290, 1, 0, 0, 0, 290, 288, 1,
		0, 0, 0, 290, 291, 1, 0, 0, 0, 291, 90, 1, 0, 0, 0, 292, 297, 5, 39, 0,
		0, 293, 296, 3, 97, 48, 0, 294, 296, 3, 103, 51, 0, 295, 293, 1, 0, 0,
		0, 295, 294, 1, 0, 0, 0, 296, 299, 1, 0, 0, 0, 297, 295, 1, 0, 0, 0, 297,
		298, 1, 0, 0, 0, 298, 300, 1, 0, 0, 0, 299, 297, 1, 0, 0, 0, 300, 311,
		5, 39, 0, 0, 301, 306, 5, 34, 0, 0, 302, 305, 3, 99, 49, 0, 303, 305, 3,
		103, 51, 0, 304, 302, 1, 0, 0, 0, 304, 303, 1, 0, 0, 0, 305, 308, 1, 0,
		0, 0, 306, 304, 1, 0, 0, 0, 306, 307, 1, 0, 0, 0, 307, 309, 1, 0, 0, 0,
		308, 306, 1, 0, 0, 0, 309, 311, 5, 34, 0, 0, 310, 292, 1, 0, 0, 0, 310,
		301, 1, 0, 0, 0, 311, 92, 1, 0, 0, 0, 312, 313, 5, 39, 0, 0, 313, 314,
		5, 39, 0, 0, 314, 315, 5, 39, 0, 0, 315, 319, 1, 0, 0, 0, 316, 318, 3,
		95, 47, 0, 317, 316, 1, 0, 0, 0, 318, 321, 1, 0, 0, 0, 319, 320, 1, 0,
		0, 0, 319, 317, 1, 0, 0, 0, 320, 322, 1, 0, 0, 0, 321, 319, 1, 0, 0, 0,
		322, 323, 5, 39, 0, 0, 323, 324, 5, 39, 0, 0, 324, 339, 5, 39, 0, 0, 325,
		326, 5, 34, 0, 0, 326, 327, 5, 34, 0, 0, 327, 328, 5, 34, 0, 0, 328, 332,
		1, 0, 0, 0, 329, 331, 3, 95, 47, 0, 330, 329, 1, 0, 0, 0, 331, 334, 1,
		0, 0, 0, 332, 333, 1, 0, 0, 0, 332, 330, 1, 0, 0, 0, 333, 335, 1, 0, 0,
		0, 334, 332, 1, 0, 0, 0, 335, 336, 5, 34, 0, 0, 336, 337, 5, 34, 0, 0,
		337, 339, 5, 34, 0, 0, 338, 312, 1, 0, 0, 0, 338, 325, 1, 0, 0, 0, 339,
		94, 1, 0, 0, 0, 340, 343, 3, 101, 50, 0, 341, 343, 3, 103, 51, 0, 342,
		340, 1, 0, 0, 0, 342, 341, 1, 0, 0, 0, 343, 96, 1, 0, 0, 0, 344, 346, 7,
		11, 0, 0, 345, 344, 1, 0, 0, 0, 346, 98, 1, 0, 0, 0, 347, 349, 7, 12, 0,
		0, 348, 347, 1, 0, 0, 0, 349, 100, 1, 0, 0, 0, 350, 352, 7, 13, 0, 0, 351,
		350, 1, 0, 0, 0, 352, 102, 1, 0, 0, 0, 353, 354, 5, 92, 0, 0, 354, 355,
		7, 14, 0, 0, 355, 104, 1, 0, 0, 0, 356, 358, 7, 15, 0, 0, 357, 356, 1,
		0, 0, 0, 358, 359, 1, 0, 0, 0, 359, 357, 1, 0, 0, 0, 359, 360, 1, 0, 0,
		0, 360, 106, 1, 0, 0, 0, 361, 363, 5, 92, 0, 0, 362, 364, 3, 105, 52, 0,
		363, 362, 1, 0, 0, 0, 363, 364, 1, 0, 0, 0, 364, 370, 1, 0, 0, 0, 365,
		367, 5, 13, 0, 0, 366, 365, 1, 0, 0, 0, 366, 367, 1, 0, 0, 0, 367, 368,
		1, 0, 0, 0, 368, 371, 5, 10, 0, 0, 369, 371, 2, 12, 13, 0, 370, 366, 1,
		0, 0, 0, 370, 369, 1, 0, 0, 0, 371, 108, 1, 0, 0, 0, 372, 375, 5, 92, 0,
		0, 373, 376, 7, 16, 0, 0, 374, 376, 3, 111, 55, 0, 375, 373, 1, 0, 0, 0,
		375, 374, 1, 0, 0, 0, 376, 110, 1, 0, 0, 0, 377, 378, 5, 117, 0, 0, 378,
		379, 3, 115, 57, 0, 379, 380, 3, 115, 57, 0, 380, 381, 3, 115, 57, 0, 381,
		382, 3, 115, 57, 0, 382, 112, 1, 0, 0, 0, 383, 384, 8, 17, 0, 0, 384, 114,
		1, 0, 0, 0, 385, 386, 7, 7, 0, 0, 386, 116, 1, 0, 0, 0, 38, 0, 184, 191,
		198, 200, 206, 211, 213, 223, 227, 231, 237, 242, 245, 258, 264, 268, 275,
		281, 285, 290, 295, 297, 304, 306, 310, 319, 332, 338, 342, 345, 348, 351,
		359, 363, 366, 370, 375, 1, 6, 0, 0,
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

// valuateLexerInit initializes any static state used to implement valuateLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewvaluateLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func ValuateLexerInit() {
	staticData := &valuatelexerLexerStaticData
	staticData.once.Do(valuatelexerLexerInit)
}

// NewvaluateLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewvaluateLexer(input antlr.CharStream) *valuateLexer {
	ValuateLexerInit()
	l := new(valuateLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &valuatelexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "valuate.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// valuateLexer tokens.
const (
	valuateLexerT__0              = 1
	valuateLexerTRUE              = 2
	valuateLexerFALSE             = 3
	valuateLexerNIL_LIT           = 4
	valuateLexerLP                = 5
	valuateLexerRP                = 6
	valuateLexerL_CURLY           = 7
	valuateLexerR_CURLY           = 8
	valuateLexerL_BRACKET         = 9
	valuateLexerR_BRACKET         = 10
	valuateLexerCOMMA             = 11
	valuateLexerDOT               = 12
	valuateLexerLOGICAL_OR        = 13
	valuateLexerLOGICAL_AND       = 14
	valuateLexerEQUALS            = 15
	valuateLexerNOT_EQUALS        = 16
	valuateLexerLESS              = 17
	valuateLexerLESS_OR_EQUALS    = 18
	valuateLexerGREATER           = 19
	valuateLexerGREATER_OR_EQUALS = 20
	valuateLexerDIV               = 21
	valuateLexerPLUS              = 22
	valuateLexerMINUS             = 23
	valuateLexerSTAR              = 24
	valuateLexerMODULUS           = 25
	valuateLexerSKIP_             = 26
	valuateLexerWHITESPACE        = 27
	valuateLexerIDENTIFIER        = 28
	valuateLexerSTRING            = 29
	valuateLexerSTRING_LITERAL    = 30
	valuateLexerBYTES_LITERAL     = 31
	valuateLexerINT               = 32
	valuateLexerDECIMAL_LIT       = 33
	valuateLexerFLOAT_NUMBER      = 34
)
