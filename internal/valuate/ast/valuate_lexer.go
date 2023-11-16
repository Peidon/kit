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
		"VAR_ID", "STRING", "STRING_LITERAL", "BYTES_LITERAL", "INT", "DECIMAL_LIT",
		"FLOAT_NUMBER",
	}
	staticData.ruleNames = []string{
		"T__0", "TRUE", "FALSE", "NIL_LIT", "LP", "RP", "L_CURLY", "R_CURLY",
		"L_BRACKET", "R_BRACKET", "COMMA", "DOT", "LOGICAL_OR", "LOGICAL_AND",
		"EQUALS", "NOT_EQUALS", "LESS", "LESS_OR_EQUALS", "GREATER", "GREATER_OR_EQUALS",
		"DIV", "PLUS", "MINUS", "STAR", "MODULUS", "SKIP_", "WHITESPACE", "IDENTIFIER",
		"VAR_ID", "STRING", "STRING_LITERAL", "BYTES_LITERAL", "INT", "DECIMAL_LIT",
		"FLOAT_NUMBER", "LETTER", "NON_ZERO_DIGIT", "DIGIT", "OCT_DIGIT", "HEX_DIGIT",
		"BIN_DIGIT", "POINT_FLOAT", "EXPONENT_FLOAT", "INT_PART", "FRACTION",
		"EXPONENT", "SHORT_BYTES", "LONG_BYTES", "LONG_BYTES_ITEM", "SHORT_BYTES_CHAR_NO_SINGLE_QUOTE",
		"SHORT_BYTES_CHAR_NO_DOUBLE_QUOTE", "LONG_BYTES_CHAR", "BYTES_ESCAPE_SEQ",
		"SPACES", "LINE_JOINING", "ESC", "UNICODE", "SAFECODEPOINT", "HEX",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 35, 404, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
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
		7, 57, 2, 58, 7, 58, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1,
		5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11,
		1, 11, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1,
		15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19,
		1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1,
		24, 1, 24, 1, 25, 1, 25, 3, 25, 187, 8, 25, 1, 25, 1, 25, 1, 26, 4, 26,
		192, 8, 26, 11, 26, 12, 26, 193, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 5,
		27, 201, 8, 27, 10, 27, 12, 27, 204, 9, 27, 1, 28, 1, 28, 1, 28, 1, 28,
		1, 28, 5, 28, 211, 8, 28, 10, 28, 12, 28, 214, 9, 28, 1, 28, 1, 28, 1,
		28, 3, 28, 219, 8, 28, 1, 29, 1, 29, 1, 29, 3, 29, 224, 8, 29, 1, 30, 1,
		30, 1, 30, 5, 30, 229, 8, 30, 10, 30, 12, 30, 232, 9, 30, 1, 30, 1, 30,
		1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 241, 8, 31, 1, 31, 1, 31, 3,
		31, 245, 8, 31, 1, 32, 1, 32, 3, 32, 249, 8, 32, 1, 33, 1, 33, 5, 33, 253,
		8, 33, 10, 33, 12, 33, 256, 9, 33, 1, 34, 1, 34, 3, 34, 260, 8, 34, 1,
		35, 3, 35, 263, 8, 35, 1, 36, 1, 36, 1, 37, 1, 37, 1, 38, 1, 38, 1, 39,
		1, 39, 1, 40, 1, 40, 1, 41, 3, 41, 276, 8, 41, 1, 41, 1, 41, 1, 41, 1,
		41, 3, 41, 282, 8, 41, 1, 42, 1, 42, 3, 42, 286, 8, 42, 1, 42, 1, 42, 1,
		43, 4, 43, 291, 8, 43, 11, 43, 12, 43, 292, 1, 44, 1, 44, 4, 44, 297, 8,
		44, 11, 44, 12, 44, 298, 1, 45, 1, 45, 3, 45, 303, 8, 45, 1, 45, 4, 45,
		306, 8, 45, 11, 45, 12, 45, 307, 1, 46, 1, 46, 1, 46, 5, 46, 313, 8, 46,
		10, 46, 12, 46, 316, 9, 46, 1, 46, 1, 46, 1, 46, 1, 46, 5, 46, 322, 8,
		46, 10, 46, 12, 46, 325, 9, 46, 1, 46, 3, 46, 328, 8, 46, 1, 47, 1, 47,
		1, 47, 1, 47, 1, 47, 5, 47, 335, 8, 47, 10, 47, 12, 47, 338, 9, 47, 1,
		47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 5, 47, 348, 8, 47,
		10, 47, 12, 47, 351, 9, 47, 1, 47, 1, 47, 1, 47, 3, 47, 356, 8, 47, 1,
		48, 1, 48, 3, 48, 360, 8, 48, 1, 49, 3, 49, 363, 8, 49, 1, 50, 3, 50, 366,
		8, 50, 1, 51, 3, 51, 369, 8, 51, 1, 52, 1, 52, 1, 52, 1, 53, 4, 53, 375,
		8, 53, 11, 53, 12, 53, 376, 1, 54, 1, 54, 3, 54, 381, 8, 54, 1, 54, 3,
		54, 384, 8, 54, 1, 54, 1, 54, 3, 54, 388, 8, 54, 1, 55, 1, 55, 1, 55, 3,
		55, 393, 8, 55, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 57, 1, 57,
		1, 58, 1, 58, 2, 336, 349, 0, 59, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6,
		13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31,
		16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49,
		25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67,
		34, 69, 35, 71, 0, 73, 0, 75, 0, 77, 0, 79, 0, 81, 0, 83, 0, 85, 0, 87,
		0, 89, 0, 91, 0, 93, 0, 95, 0, 97, 0, 99, 0, 101, 0, 103, 0, 105, 0, 107,
		0, 109, 0, 111, 0, 113, 0, 115, 0, 117, 0, 1, 0, 18, 3, 0, 9, 10, 13, 13,
		32, 32, 2, 0, 66, 66, 98, 98, 2, 0, 82, 82, 114, 114, 3, 0, 65, 90, 95,
		95, 97, 122, 1, 0, 49, 57, 1, 0, 48, 57, 1, 0, 48, 55, 3, 0, 48, 57, 65,
		70, 97, 102, 1, 0, 48, 49, 2, 0, 69, 69, 101, 101, 2, 0, 43, 43, 45, 45,
		5, 0, 0, 9, 11, 12, 14, 38, 40, 91, 93, 127, 5, 0, 0, 9, 11, 12, 14, 33,
		35, 91, 93, 127, 2, 0, 0, 91, 93, 127, 1, 0, 0, 127, 2, 0, 9, 9, 32, 32,
		8, 0, 34, 34, 47, 47, 92, 92, 98, 98, 102, 102, 110, 110, 114, 114, 116,
		116, 3, 0, 0, 31, 34, 34, 92, 92, 416, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0,
		0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0,
		0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0,
		0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1,
		0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35,
		1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0,
		43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0,
		0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0,
		0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0,
		0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 1, 119, 1, 0, 0, 0, 3, 121,
		1, 0, 0, 0, 5, 126, 1, 0, 0, 0, 7, 132, 1, 0, 0, 0, 9, 136, 1, 0, 0, 0,
		11, 138, 1, 0, 0, 0, 13, 140, 1, 0, 0, 0, 15, 142, 1, 0, 0, 0, 17, 144,
		1, 0, 0, 0, 19, 146, 1, 0, 0, 0, 21, 148, 1, 0, 0, 0, 23, 150, 1, 0, 0,
		0, 25, 152, 1, 0, 0, 0, 27, 155, 1, 0, 0, 0, 29, 158, 1, 0, 0, 0, 31, 161,
		1, 0, 0, 0, 33, 164, 1, 0, 0, 0, 35, 166, 1, 0, 0, 0, 37, 169, 1, 0, 0,
		0, 39, 171, 1, 0, 0, 0, 41, 174, 1, 0, 0, 0, 43, 176, 1, 0, 0, 0, 45, 178,
		1, 0, 0, 0, 47, 180, 1, 0, 0, 0, 49, 182, 1, 0, 0, 0, 51, 186, 1, 0, 0,
		0, 53, 191, 1, 0, 0, 0, 55, 197, 1, 0, 0, 0, 57, 218, 1, 0, 0, 0, 59, 223,
		1, 0, 0, 0, 61, 225, 1, 0, 0, 0, 63, 240, 1, 0, 0, 0, 65, 248, 1, 0, 0,
		0, 67, 250, 1, 0, 0, 0, 69, 259, 1, 0, 0, 0, 71, 262, 1, 0, 0, 0, 73, 264,
		1, 0, 0, 0, 75, 266, 1, 0, 0, 0, 77, 268, 1, 0, 0, 0, 79, 270, 1, 0, 0,
		0, 81, 272, 1, 0, 0, 0, 83, 281, 1, 0, 0, 0, 85, 285, 1, 0, 0, 0, 87, 290,
		1, 0, 0, 0, 89, 294, 1, 0, 0, 0, 91, 300, 1, 0, 0, 0, 93, 327, 1, 0, 0,
		0, 95, 355, 1, 0, 0, 0, 97, 359, 1, 0, 0, 0, 99, 362, 1, 0, 0, 0, 101,
		365, 1, 0, 0, 0, 103, 368, 1, 0, 0, 0, 105, 370, 1, 0, 0, 0, 107, 374,
		1, 0, 0, 0, 109, 378, 1, 0, 0, 0, 111, 389, 1, 0, 0, 0, 113, 394, 1, 0,
		0, 0, 115, 400, 1, 0, 0, 0, 117, 402, 1, 0, 0, 0, 119, 120, 5, 33, 0, 0,
		120, 2, 1, 0, 0, 0, 121, 122, 5, 116, 0, 0, 122, 123, 5, 114, 0, 0, 123,
		124, 5, 117, 0, 0, 124, 125, 5, 101, 0, 0, 125, 4, 1, 0, 0, 0, 126, 127,
		5, 102, 0, 0, 127, 128, 5, 97, 0, 0, 128, 129, 5, 108, 0, 0, 129, 130,
		5, 115, 0, 0, 130, 131, 5, 101, 0, 0, 131, 6, 1, 0, 0, 0, 132, 133, 5,
		110, 0, 0, 133, 134, 5, 105, 0, 0, 134, 135, 5, 108, 0, 0, 135, 8, 1, 0,
		0, 0, 136, 137, 5, 40, 0, 0, 137, 10, 1, 0, 0, 0, 138, 139, 5, 41, 0, 0,
		139, 12, 1, 0, 0, 0, 140, 141, 5, 123, 0, 0, 141, 14, 1, 0, 0, 0, 142,
		143, 5, 125, 0, 0, 143, 16, 1, 0, 0, 0, 144, 145, 5, 91, 0, 0, 145, 18,
		1, 0, 0, 0, 146, 147, 5, 93, 0, 0, 147, 20, 1, 0, 0, 0, 148, 149, 5, 44,
		0, 0, 149, 22, 1, 0, 0, 0, 150, 151, 5, 46, 0, 0, 151, 24, 1, 0, 0, 0,
		152, 153, 5, 124, 0, 0, 153, 154, 5, 124, 0, 0, 154, 26, 1, 0, 0, 0, 155,
		156, 5, 38, 0, 0, 156, 157, 5, 38, 0, 0, 157, 28, 1, 0, 0, 0, 158, 159,
		5, 61, 0, 0, 159, 160, 5, 61, 0, 0, 160, 30, 1, 0, 0, 0, 161, 162, 5, 33,
		0, 0, 162, 163, 5, 61, 0, 0, 163, 32, 1, 0, 0, 0, 164, 165, 5, 60, 0, 0,
		165, 34, 1, 0, 0, 0, 166, 167, 5, 60, 0, 0, 167, 168, 5, 61, 0, 0, 168,
		36, 1, 0, 0, 0, 169, 170, 5, 62, 0, 0, 170, 38, 1, 0, 0, 0, 171, 172, 5,
		62, 0, 0, 172, 173, 5, 61, 0, 0, 173, 40, 1, 0, 0, 0, 174, 175, 5, 47,
		0, 0, 175, 42, 1, 0, 0, 0, 176, 177, 5, 43, 0, 0, 177, 44, 1, 0, 0, 0,
		178, 179, 5, 45, 0, 0, 179, 46, 1, 0, 0, 0, 180, 181, 5, 42, 0, 0, 181,
		48, 1, 0, 0, 0, 182, 183, 5, 37, 0, 0, 183, 50, 1, 0, 0, 0, 184, 187, 3,
		107, 53, 0, 185, 187, 3, 109, 54, 0, 186, 184, 1, 0, 0, 0, 186, 185, 1,
		0, 0, 0, 187, 188, 1, 0, 0, 0, 188, 189, 6, 25, 0, 0, 189, 52, 1, 0, 0,
		0, 190, 192, 7, 0, 0, 0, 191, 190, 1, 0, 0, 0, 192, 193, 1, 0, 0, 0, 193,
		191, 1, 0, 0, 0, 193, 194, 1, 0, 0, 0, 194, 195, 1, 0, 0, 0, 195, 196,
		6, 26, 0, 0, 196, 54, 1, 0, 0, 0, 197, 202, 3, 71, 35, 0, 198, 201, 3,
		71, 35, 0, 199, 201, 3, 75, 37, 0, 200, 198, 1, 0, 0, 0, 200, 199, 1, 0,
		0, 0, 201, 204, 1, 0, 0, 0, 202, 200, 1, 0, 0, 0, 202, 203, 1, 0, 0, 0,
		203, 56, 1, 0, 0, 0, 204, 202, 1, 0, 0, 0, 205, 206, 3, 13, 6, 0, 206,
		212, 3, 55, 27, 0, 207, 208, 3, 23, 11, 0, 208, 209, 3, 55, 27, 0, 209,
		211, 1, 0, 0, 0, 210, 207, 1, 0, 0, 0, 211, 214, 1, 0, 0, 0, 212, 210,
		1, 0, 0, 0, 212, 213, 1, 0, 0, 0, 213, 215, 1, 0, 0, 0, 214, 212, 1, 0,
		0, 0, 215, 216, 3, 15, 7, 0, 216, 219, 1, 0, 0, 0, 217, 219, 3, 55, 27,
		0, 218, 205, 1, 0, 0, 0, 218, 217, 1, 0, 0, 0, 219, 58, 1, 0, 0, 0, 220,
		224, 3, 61, 30, 0, 221, 224, 3, 93, 46, 0, 222, 224, 3, 95, 47, 0, 223,
		220, 1, 0, 0, 0, 223, 221, 1, 0, 0, 0, 223, 222, 1, 0, 0, 0, 224, 60, 1,
		0, 0, 0, 225, 230, 5, 34, 0, 0, 226, 229, 3, 111, 55, 0, 227, 229, 3, 115,
		57, 0, 228, 226, 1, 0, 0, 0, 228, 227, 1, 0, 0, 0, 229, 232, 1, 0, 0, 0,
		230, 228, 1, 0, 0, 0, 230, 231, 1, 0, 0, 0, 231, 233, 1, 0, 0, 0, 232,
		230, 1, 0, 0, 0, 233, 234, 5, 34, 0, 0, 234, 62, 1, 0, 0, 0, 235, 241,
		7, 1, 0, 0, 236, 237, 7, 1, 0, 0, 237, 241, 7, 2, 0, 0, 238, 239, 7, 2,
		0, 0, 239, 241, 7, 1, 0, 0, 240, 235, 1, 0, 0, 0, 240, 236, 1, 0, 0, 0,
		240, 238, 1, 0, 0, 0, 241, 244, 1, 0, 0, 0, 242, 245, 3, 93, 46, 0, 243,
		245, 3, 95, 47, 0, 244, 242, 1, 0, 0, 0, 244, 243, 1, 0, 0, 0, 245, 64,
		1, 0, 0, 0, 246, 249, 3, 67, 33, 0, 247, 249, 5, 48, 0, 0, 248, 246, 1,
		0, 0, 0, 248, 247, 1, 0, 0, 0, 249, 66, 1, 0, 0, 0, 250, 254, 3, 73, 36,
		0, 251, 253, 3, 75, 37, 0, 252, 251, 1, 0, 0, 0, 253, 256, 1, 0, 0, 0,
		254, 252, 1, 0, 0, 0, 254, 255, 1, 0, 0, 0, 255, 68, 1, 0, 0, 0, 256, 254,
		1, 0, 0, 0, 257, 260, 3, 83, 41, 0, 258, 260, 3, 85, 42, 0, 259, 257, 1,
		0, 0, 0, 259, 258, 1, 0, 0, 0, 260, 70, 1, 0, 0, 0, 261, 263, 7, 3, 0,
		0, 262, 261, 1, 0, 0, 0, 263, 72, 1, 0, 0, 0, 264, 265, 7, 4, 0, 0, 265,
		74, 1, 0, 0, 0, 266, 267, 7, 5, 0, 0, 267, 76, 1, 0, 0, 0, 268, 269, 7,
		6, 0, 0, 269, 78, 1, 0, 0, 0, 270, 271, 7, 7, 0, 0, 271, 80, 1, 0, 0, 0,
		272, 273, 7, 8, 0, 0, 273, 82, 1, 0, 0, 0, 274, 276, 3, 87, 43, 0, 275,
		274, 1, 0, 0, 0, 275, 276, 1, 0, 0, 0, 276, 277, 1, 0, 0, 0, 277, 282,
		3, 89, 44, 0, 278, 279, 3, 87, 43, 0, 279, 280, 5, 46, 0, 0, 280, 282,
		1, 0, 0, 0, 281, 275, 1, 0, 0, 0, 281, 278, 1, 0, 0, 0, 282, 84, 1, 0,
		0, 0, 283, 286, 3, 87, 43, 0, 284, 286, 3, 83, 41, 0, 285, 283, 1, 0, 0,
		0, 285, 284, 1, 0, 0, 0, 286, 287, 1, 0, 0, 0, 287, 288, 3, 91, 45, 0,
		288, 86, 1, 0, 0, 0, 289, 291, 3, 75, 37, 0, 290, 289, 1, 0, 0, 0, 291,
		292, 1, 0, 0, 0, 292, 290, 1, 0, 0, 0, 292, 293, 1, 0, 0, 0, 293, 88, 1,
		0, 0, 0, 294, 296, 5, 46, 0, 0, 295, 297, 3, 75, 37, 0, 296, 295, 1, 0,
		0, 0, 297, 298, 1, 0, 0, 0, 298, 296, 1, 0, 0, 0, 298, 299, 1, 0, 0, 0,
		299, 90, 1, 0, 0, 0, 300, 302, 7, 9, 0, 0, 301, 303, 7, 10, 0, 0, 302,
		301, 1, 0, 0, 0, 302, 303, 1, 0, 0, 0, 303, 305, 1, 0, 0, 0, 304, 306,
		3, 75, 37, 0, 305, 304, 1, 0, 0, 0, 306, 307, 1, 0, 0, 0, 307, 305, 1,
		0, 0, 0, 307, 308, 1, 0, 0, 0, 308, 92, 1, 0, 0, 0, 309, 314, 5, 39, 0,
		0, 310, 313, 3, 99, 49, 0, 311, 313, 3, 105, 52, 0, 312, 310, 1, 0, 0,
		0, 312, 311, 1, 0, 0, 0, 313, 316, 1, 0, 0, 0, 314, 312, 1, 0, 0, 0, 314,
		315, 1, 0, 0, 0, 315, 317, 1, 0, 0, 0, 316, 314, 1, 0, 0, 0, 317, 328,
		5, 39, 0, 0, 318, 323, 5, 34, 0, 0, 319, 322, 3, 101, 50, 0, 320, 322,
		3, 105, 52, 0, 321, 319, 1, 0, 0, 0, 321, 320, 1, 0, 0, 0, 322, 325, 1,
		0, 0, 0, 323, 321, 1, 0, 0, 0, 323, 324, 1, 0, 0, 0, 324, 326, 1, 0, 0,
		0, 325, 323, 1, 0, 0, 0, 326, 328, 5, 34, 0, 0, 327, 309, 1, 0, 0, 0, 327,
		318, 1, 0, 0, 0, 328, 94, 1, 0, 0, 0, 329, 330, 5, 39, 0, 0, 330, 331,
		5, 39, 0, 0, 331, 332, 5, 39, 0, 0, 332, 336, 1, 0, 0, 0, 333, 335, 3,
		97, 48, 0, 334, 333, 1, 0, 0, 0, 335, 338, 1, 0, 0, 0, 336, 337, 1, 0,
		0, 0, 336, 334, 1, 0, 0, 0, 337, 339, 1, 0, 0, 0, 338, 336, 1, 0, 0, 0,
		339, 340, 5, 39, 0, 0, 340, 341, 5, 39, 0, 0, 341, 356, 5, 39, 0, 0, 342,
		343, 5, 34, 0, 0, 343, 344, 5, 34, 0, 0, 344, 345, 5, 34, 0, 0, 345, 349,
		1, 0, 0, 0, 346, 348, 3, 97, 48, 0, 347, 346, 1, 0, 0, 0, 348, 351, 1,
		0, 0, 0, 349, 350, 1, 0, 0, 0, 349, 347, 1, 0, 0, 0, 350, 352, 1, 0, 0,
		0, 351, 349, 1, 0, 0, 0, 352, 353, 5, 34, 0, 0, 353, 354, 5, 34, 0, 0,
		354, 356, 5, 34, 0, 0, 355, 329, 1, 0, 0, 0, 355, 342, 1, 0, 0, 0, 356,
		96, 1, 0, 0, 0, 357, 360, 3, 103, 51, 0, 358, 360, 3, 105, 52, 0, 359,
		357, 1, 0, 0, 0, 359, 358, 1, 0, 0, 0, 360, 98, 1, 0, 0, 0, 361, 363, 7,
		11, 0, 0, 362, 361, 1, 0, 0, 0, 363, 100, 1, 0, 0, 0, 364, 366, 7, 12,
		0, 0, 365, 364, 1, 0, 0, 0, 366, 102, 1, 0, 0, 0, 367, 369, 7, 13, 0, 0,
		368, 367, 1, 0, 0, 0, 369, 104, 1, 0, 0, 0, 370, 371, 5, 92, 0, 0, 371,
		372, 7, 14, 0, 0, 372, 106, 1, 0, 0, 0, 373, 375, 7, 15, 0, 0, 374, 373,
		1, 0, 0, 0, 375, 376, 1, 0, 0, 0, 376, 374, 1, 0, 0, 0, 376, 377, 1, 0,
		0, 0, 377, 108, 1, 0, 0, 0, 378, 380, 5, 92, 0, 0, 379, 381, 3, 107, 53,
		0, 380, 379, 1, 0, 0, 0, 380, 381, 1, 0, 0, 0, 381, 387, 1, 0, 0, 0, 382,
		384, 5, 13, 0, 0, 383, 382, 1, 0, 0, 0, 383, 384, 1, 0, 0, 0, 384, 385,
		1, 0, 0, 0, 385, 388, 5, 10, 0, 0, 386, 388, 2, 12, 13, 0, 387, 383, 1,
		0, 0, 0, 387, 386, 1, 0, 0, 0, 388, 110, 1, 0, 0, 0, 389, 392, 5, 92, 0,
		0, 390, 393, 7, 16, 0, 0, 391, 393, 3, 113, 56, 0, 392, 390, 1, 0, 0, 0,
		392, 391, 1, 0, 0, 0, 393, 112, 1, 0, 0, 0, 394, 395, 5, 117, 0, 0, 395,
		396, 3, 117, 58, 0, 396, 397, 3, 117, 58, 0, 397, 398, 3, 117, 58, 0, 398,
		399, 3, 117, 58, 0, 399, 114, 1, 0, 0, 0, 400, 401, 8, 17, 0, 0, 401, 116,
		1, 0, 0, 0, 402, 403, 7, 7, 0, 0, 403, 118, 1, 0, 0, 0, 40, 0, 186, 193,
		200, 202, 212, 218, 223, 228, 230, 240, 244, 248, 254, 259, 262, 275, 281,
		285, 292, 298, 302, 307, 312, 314, 321, 323, 327, 336, 349, 355, 359, 362,
		365, 368, 376, 380, 383, 387, 392, 1, 6, 0, 0,
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
	valuateLexerVAR_ID            = 29
	valuateLexerSTRING            = 30
	valuateLexerSTRING_LITERAL    = 31
	valuateLexerBYTES_LITERAL     = 32
	valuateLexerINT               = 33
	valuateLexerDECIMAL_LIT       = 34
	valuateLexerFLOAT_NUMBER      = 35
)
