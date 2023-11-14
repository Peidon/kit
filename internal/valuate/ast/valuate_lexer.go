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
		"", "'%'", "'!'", "'0'", "'true'", "'false'", "'nil'", "'('", "')'",
		"'{'", "'}'", "'['", "']'", "','", "'.'", "'||'", "'&&'", "'=='", "'!='",
		"'<'", "'<='", "'>'", "'>='", "'/'", "'+'", "'-'", "'*'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "TRUE", "FALSE", "NIL_LIT", "LP", "RP", "L_CURLY", "R_CURLY",
		"L_BRACKET", "R_BRACKET", "COMMA", "DOT", "LOGICAL_OR", "LOGICAL_AND",
		"EQUALS", "NOT_EQUALS", "LESS", "LESS_OR_EQUALS", "GREATER", "GREATER_OR_EQUALS",
		"DIV", "PLUS", "MINUS", "STAR", "SKIP_", "WHITESPACE", "IDENTIFIER",
		"VAR_ID", "STRING", "STRING_LITERAL", "BYTES_LITERAL", "DECIMAL_LIT",
		"FLOAT_NUMBER",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "TRUE", "FALSE", "NIL_LIT", "LP", "RP", "L_CURLY",
		"R_CURLY", "L_BRACKET", "R_BRACKET", "COMMA", "DOT", "LOGICAL_OR", "LOGICAL_AND",
		"EQUALS", "NOT_EQUALS", "LESS", "LESS_OR_EQUALS", "GREATER", "GREATER_OR_EQUALS",
		"DIV", "PLUS", "MINUS", "STAR", "SKIP_", "WHITESPACE", "IDENTIFIER",
		"VAR_ID", "STRING", "STRING_LITERAL", "BYTES_LITERAL", "DECIMAL_LIT",
		"FLOAT_NUMBER", "LETTER", "NON_ZERO_DIGIT", "DIGIT", "OCT_DIGIT", "HEX_DIGIT",
		"BIN_DIGIT", "POINT_FLOAT", "EXPONENT_FLOAT", "INT_PART", "FRACTION",
		"EXPONENT", "SHORT_BYTES", "LONG_BYTES", "LONG_BYTES_ITEM", "SHORT_BYTES_CHAR_NO_SINGLE_QUOTE",
		"SHORT_BYTES_CHAR_NO_DOUBLE_QUOTE", "LONG_BYTES_CHAR", "BYTES_ESCAPE_SEQ",
		"SPACES", "LINE_JOINING", "ESC", "UNICODE", "SAFECODEPOINT", "HEX",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 35, 402, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
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
		7, 57, 2, 58, 7, 58, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11,
		1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1,
		15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19,
		1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1,
		24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 3, 26, 189, 8, 26, 1, 26, 1, 26,
		1, 27, 4, 27, 194, 8, 27, 11, 27, 12, 27, 195, 1, 27, 1, 27, 1, 28, 1,
		28, 1, 28, 5, 28, 203, 8, 28, 10, 28, 12, 28, 206, 9, 28, 1, 29, 1, 29,
		1, 29, 1, 29, 1, 29, 5, 29, 213, 8, 29, 10, 29, 12, 29, 216, 9, 29, 1,
		29, 1, 29, 1, 29, 3, 29, 221, 8, 29, 1, 30, 1, 30, 1, 30, 3, 30, 226, 8,
		30, 1, 31, 1, 31, 1, 31, 5, 31, 231, 8, 31, 10, 31, 12, 31, 234, 9, 31,
		1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 3, 32, 243, 8, 32, 1,
		32, 1, 32, 3, 32, 247, 8, 32, 1, 33, 1, 33, 5, 33, 251, 8, 33, 10, 33,
		12, 33, 254, 9, 33, 1, 34, 1, 34, 3, 34, 258, 8, 34, 1, 35, 3, 35, 261,
		8, 35, 1, 36, 1, 36, 1, 37, 1, 37, 1, 38, 1, 38, 1, 39, 1, 39, 1, 40, 1,
		40, 1, 41, 3, 41, 274, 8, 41, 1, 41, 1, 41, 1, 41, 1, 41, 3, 41, 280, 8,
		41, 1, 42, 1, 42, 3, 42, 284, 8, 42, 1, 42, 1, 42, 1, 43, 4, 43, 289, 8,
		43, 11, 43, 12, 43, 290, 1, 44, 1, 44, 4, 44, 295, 8, 44, 11, 44, 12, 44,
		296, 1, 45, 1, 45, 3, 45, 301, 8, 45, 1, 45, 4, 45, 304, 8, 45, 11, 45,
		12, 45, 305, 1, 46, 1, 46, 1, 46, 5, 46, 311, 8, 46, 10, 46, 12, 46, 314,
		9, 46, 1, 46, 1, 46, 1, 46, 1, 46, 5, 46, 320, 8, 46, 10, 46, 12, 46, 323,
		9, 46, 1, 46, 3, 46, 326, 8, 46, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 5,
		47, 333, 8, 47, 10, 47, 12, 47, 336, 9, 47, 1, 47, 1, 47, 1, 47, 1, 47,
		1, 47, 1, 47, 1, 47, 1, 47, 5, 47, 346, 8, 47, 10, 47, 12, 47, 349, 9,
		47, 1, 47, 1, 47, 1, 47, 3, 47, 354, 8, 47, 1, 48, 1, 48, 3, 48, 358, 8,
		48, 1, 49, 3, 49, 361, 8, 49, 1, 50, 3, 50, 364, 8, 50, 1, 51, 3, 51, 367,
		8, 51, 1, 52, 1, 52, 1, 52, 1, 53, 4, 53, 373, 8, 53, 11, 53, 12, 53, 374,
		1, 54, 1, 54, 3, 54, 379, 8, 54, 1, 54, 3, 54, 382, 8, 54, 1, 54, 1, 54,
		3, 54, 386, 8, 54, 1, 55, 1, 55, 1, 55, 3, 55, 391, 8, 55, 1, 56, 1, 56,
		1, 56, 1, 56, 1, 56, 1, 56, 1, 57, 1, 57, 1, 58, 1, 58, 2, 334, 347, 0,
		59, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21,
		11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39,
		20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57,
		29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 0, 73, 0, 75, 0,
		77, 0, 79, 0, 81, 0, 83, 0, 85, 0, 87, 0, 89, 0, 91, 0, 93, 0, 95, 0, 97,
		0, 99, 0, 101, 0, 103, 0, 105, 0, 107, 0, 109, 0, 111, 0, 113, 0, 115,
		0, 117, 0, 1, 0, 18, 3, 0, 9, 10, 13, 13, 32, 32, 2, 0, 66, 66, 98, 98,
		2, 0, 82, 82, 114, 114, 3, 0, 65, 90, 95, 95, 97, 122, 1, 0, 49, 57, 1,
		0, 48, 57, 1, 0, 48, 55, 3, 0, 48, 57, 65, 70, 97, 102, 1, 0, 48, 49, 2,
		0, 69, 69, 101, 101, 2, 0, 43, 43, 45, 45, 5, 0, 0, 9, 11, 12, 14, 38,
		40, 91, 93, 127, 5, 0, 0, 9, 11, 12, 14, 33, 35, 91, 93, 127, 2, 0, 0,
		91, 93, 127, 1, 0, 0, 127, 2, 0, 9, 9, 32, 32, 8, 0, 34, 34, 47, 47, 92,
		92, 98, 98, 102, 102, 110, 110, 114, 114, 116, 116, 3, 0, 0, 31, 34, 34,
		92, 92, 413, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7,
		1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0,
		15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0,
		0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0,
		0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0,
		0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1,
		0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53,
		1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0,
		61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0,
		0, 69, 1, 0, 0, 0, 1, 119, 1, 0, 0, 0, 3, 121, 1, 0, 0, 0, 5, 123, 1, 0,
		0, 0, 7, 125, 1, 0, 0, 0, 9, 130, 1, 0, 0, 0, 11, 136, 1, 0, 0, 0, 13,
		140, 1, 0, 0, 0, 15, 142, 1, 0, 0, 0, 17, 144, 1, 0, 0, 0, 19, 146, 1,
		0, 0, 0, 21, 148, 1, 0, 0, 0, 23, 150, 1, 0, 0, 0, 25, 152, 1, 0, 0, 0,
		27, 154, 1, 0, 0, 0, 29, 156, 1, 0, 0, 0, 31, 159, 1, 0, 0, 0, 33, 162,
		1, 0, 0, 0, 35, 165, 1, 0, 0, 0, 37, 168, 1, 0, 0, 0, 39, 170, 1, 0, 0,
		0, 41, 173, 1, 0, 0, 0, 43, 175, 1, 0, 0, 0, 45, 178, 1, 0, 0, 0, 47, 180,
		1, 0, 0, 0, 49, 182, 1, 0, 0, 0, 51, 184, 1, 0, 0, 0, 53, 188, 1, 0, 0,
		0, 55, 193, 1, 0, 0, 0, 57, 199, 1, 0, 0, 0, 59, 220, 1, 0, 0, 0, 61, 225,
		1, 0, 0, 0, 63, 227, 1, 0, 0, 0, 65, 242, 1, 0, 0, 0, 67, 248, 1, 0, 0,
		0, 69, 257, 1, 0, 0, 0, 71, 260, 1, 0, 0, 0, 73, 262, 1, 0, 0, 0, 75, 264,
		1, 0, 0, 0, 77, 266, 1, 0, 0, 0, 79, 268, 1, 0, 0, 0, 81, 270, 1, 0, 0,
		0, 83, 279, 1, 0, 0, 0, 85, 283, 1, 0, 0, 0, 87, 288, 1, 0, 0, 0, 89, 292,
		1, 0, 0, 0, 91, 298, 1, 0, 0, 0, 93, 325, 1, 0, 0, 0, 95, 353, 1, 0, 0,
		0, 97, 357, 1, 0, 0, 0, 99, 360, 1, 0, 0, 0, 101, 363, 1, 0, 0, 0, 103,
		366, 1, 0, 0, 0, 105, 368, 1, 0, 0, 0, 107, 372, 1, 0, 0, 0, 109, 376,
		1, 0, 0, 0, 111, 387, 1, 0, 0, 0, 113, 392, 1, 0, 0, 0, 115, 398, 1, 0,
		0, 0, 117, 400, 1, 0, 0, 0, 119, 120, 5, 37, 0, 0, 120, 2, 1, 0, 0, 0,
		121, 122, 5, 33, 0, 0, 122, 4, 1, 0, 0, 0, 123, 124, 5, 48, 0, 0, 124,
		6, 1, 0, 0, 0, 125, 126, 5, 116, 0, 0, 126, 127, 5, 114, 0, 0, 127, 128,
		5, 117, 0, 0, 128, 129, 5, 101, 0, 0, 129, 8, 1, 0, 0, 0, 130, 131, 5,
		102, 0, 0, 131, 132, 5, 97, 0, 0, 132, 133, 5, 108, 0, 0, 133, 134, 5,
		115, 0, 0, 134, 135, 5, 101, 0, 0, 135, 10, 1, 0, 0, 0, 136, 137, 5, 110,
		0, 0, 137, 138, 5, 105, 0, 0, 138, 139, 5, 108, 0, 0, 139, 12, 1, 0, 0,
		0, 140, 141, 5, 40, 0, 0, 141, 14, 1, 0, 0, 0, 142, 143, 5, 41, 0, 0, 143,
		16, 1, 0, 0, 0, 144, 145, 5, 123, 0, 0, 145, 18, 1, 0, 0, 0, 146, 147,
		5, 125, 0, 0, 147, 20, 1, 0, 0, 0, 148, 149, 5, 91, 0, 0, 149, 22, 1, 0,
		0, 0, 150, 151, 5, 93, 0, 0, 151, 24, 1, 0, 0, 0, 152, 153, 5, 44, 0, 0,
		153, 26, 1, 0, 0, 0, 154, 155, 5, 46, 0, 0, 155, 28, 1, 0, 0, 0, 156, 157,
		5, 124, 0, 0, 157, 158, 5, 124, 0, 0, 158, 30, 1, 0, 0, 0, 159, 160, 5,
		38, 0, 0, 160, 161, 5, 38, 0, 0, 161, 32, 1, 0, 0, 0, 162, 163, 5, 61,
		0, 0, 163, 164, 5, 61, 0, 0, 164, 34, 1, 0, 0, 0, 165, 166, 5, 33, 0, 0,
		166, 167, 5, 61, 0, 0, 167, 36, 1, 0, 0, 0, 168, 169, 5, 60, 0, 0, 169,
		38, 1, 0, 0, 0, 170, 171, 5, 60, 0, 0, 171, 172, 5, 61, 0, 0, 172, 40,
		1, 0, 0, 0, 173, 174, 5, 62, 0, 0, 174, 42, 1, 0, 0, 0, 175, 176, 5, 62,
		0, 0, 176, 177, 5, 61, 0, 0, 177, 44, 1, 0, 0, 0, 178, 179, 5, 47, 0, 0,
		179, 46, 1, 0, 0, 0, 180, 181, 5, 43, 0, 0, 181, 48, 1, 0, 0, 0, 182, 183,
		5, 45, 0, 0, 183, 50, 1, 0, 0, 0, 184, 185, 5, 42, 0, 0, 185, 52, 1, 0,
		0, 0, 186, 189, 3, 107, 53, 0, 187, 189, 3, 109, 54, 0, 188, 186, 1, 0,
		0, 0, 188, 187, 1, 0, 0, 0, 189, 190, 1, 0, 0, 0, 190, 191, 6, 26, 0, 0,
		191, 54, 1, 0, 0, 0, 192, 194, 7, 0, 0, 0, 193, 192, 1, 0, 0, 0, 194, 195,
		1, 0, 0, 0, 195, 193, 1, 0, 0, 0, 195, 196, 1, 0, 0, 0, 196, 197, 1, 0,
		0, 0, 197, 198, 6, 27, 0, 0, 198, 56, 1, 0, 0, 0, 199, 204, 3, 71, 35,
		0, 200, 203, 3, 71, 35, 0, 201, 203, 3, 75, 37, 0, 202, 200, 1, 0, 0, 0,
		202, 201, 1, 0, 0, 0, 203, 206, 1, 0, 0, 0, 204, 202, 1, 0, 0, 0, 204,
		205, 1, 0, 0, 0, 205, 58, 1, 0, 0, 0, 206, 204, 1, 0, 0, 0, 207, 208, 3,
		17, 8, 0, 208, 214, 3, 57, 28, 0, 209, 210, 3, 27, 13, 0, 210, 211, 3,
		57, 28, 0, 211, 213, 1, 0, 0, 0, 212, 209, 1, 0, 0, 0, 213, 216, 1, 0,
		0, 0, 214, 212, 1, 0, 0, 0, 214, 215, 1, 0, 0, 0, 215, 217, 1, 0, 0, 0,
		216, 214, 1, 0, 0, 0, 217, 218, 3, 19, 9, 0, 218, 221, 1, 0, 0, 0, 219,
		221, 3, 57, 28, 0, 220, 207, 1, 0, 0, 0, 220, 219, 1, 0, 0, 0, 221, 60,
		1, 0, 0, 0, 222, 226, 3, 63, 31, 0, 223, 226, 3, 93, 46, 0, 224, 226, 3,
		95, 47, 0, 225, 222, 1, 0, 0, 0, 225, 223, 1, 0, 0, 0, 225, 224, 1, 0,
		0, 0, 226, 62, 1, 0, 0, 0, 227, 232, 5, 34, 0, 0, 228, 231, 3, 111, 55,
		0, 229, 231, 3, 115, 57, 0, 230, 228, 1, 0, 0, 0, 230, 229, 1, 0, 0, 0,
		231, 234, 1, 0, 0, 0, 232, 230, 1, 0, 0, 0, 232, 233, 1, 0, 0, 0, 233,
		235, 1, 0, 0, 0, 234, 232, 1, 0, 0, 0, 235, 236, 5, 34, 0, 0, 236, 64,
		1, 0, 0, 0, 237, 243, 7, 1, 0, 0, 238, 239, 7, 1, 0, 0, 239, 243, 7, 2,
		0, 0, 240, 241, 7, 2, 0, 0, 241, 243, 7, 1, 0, 0, 242, 237, 1, 0, 0, 0,
		242, 238, 1, 0, 0, 0, 242, 240, 1, 0, 0, 0, 243, 246, 1, 0, 0, 0, 244,
		247, 3, 93, 46, 0, 245, 247, 3, 95, 47, 0, 246, 244, 1, 0, 0, 0, 246, 245,
		1, 0, 0, 0, 247, 66, 1, 0, 0, 0, 248, 252, 3, 73, 36, 0, 249, 251, 3, 75,
		37, 0, 250, 249, 1, 0, 0, 0, 251, 254, 1, 0, 0, 0, 252, 250, 1, 0, 0, 0,
		252, 253, 1, 0, 0, 0, 253, 68, 1, 0, 0, 0, 254, 252, 1, 0, 0, 0, 255, 258,
		3, 83, 41, 0, 256, 258, 3, 85, 42, 0, 257, 255, 1, 0, 0, 0, 257, 256, 1,
		0, 0, 0, 258, 70, 1, 0, 0, 0, 259, 261, 7, 3, 0, 0, 260, 259, 1, 0, 0,
		0, 261, 72, 1, 0, 0, 0, 262, 263, 7, 4, 0, 0, 263, 74, 1, 0, 0, 0, 264,
		265, 7, 5, 0, 0, 265, 76, 1, 0, 0, 0, 266, 267, 7, 6, 0, 0, 267, 78, 1,
		0, 0, 0, 268, 269, 7, 7, 0, 0, 269, 80, 1, 0, 0, 0, 270, 271, 7, 8, 0,
		0, 271, 82, 1, 0, 0, 0, 272, 274, 3, 87, 43, 0, 273, 272, 1, 0, 0, 0, 273,
		274, 1, 0, 0, 0, 274, 275, 1, 0, 0, 0, 275, 280, 3, 89, 44, 0, 276, 277,
		3, 87, 43, 0, 277, 278, 5, 46, 0, 0, 278, 280, 1, 0, 0, 0, 279, 273, 1,
		0, 0, 0, 279, 276, 1, 0, 0, 0, 280, 84, 1, 0, 0, 0, 281, 284, 3, 87, 43,
		0, 282, 284, 3, 83, 41, 0, 283, 281, 1, 0, 0, 0, 283, 282, 1, 0, 0, 0,
		284, 285, 1, 0, 0, 0, 285, 286, 3, 91, 45, 0, 286, 86, 1, 0, 0, 0, 287,
		289, 3, 75, 37, 0, 288, 287, 1, 0, 0, 0, 289, 290, 1, 0, 0, 0, 290, 288,
		1, 0, 0, 0, 290, 291, 1, 0, 0, 0, 291, 88, 1, 0, 0, 0, 292, 294, 5, 46,
		0, 0, 293, 295, 3, 75, 37, 0, 294, 293, 1, 0, 0, 0, 295, 296, 1, 0, 0,
		0, 296, 294, 1, 0, 0, 0, 296, 297, 1, 0, 0, 0, 297, 90, 1, 0, 0, 0, 298,
		300, 7, 9, 0, 0, 299, 301, 7, 10, 0, 0, 300, 299, 1, 0, 0, 0, 300, 301,
		1, 0, 0, 0, 301, 303, 1, 0, 0, 0, 302, 304, 3, 75, 37, 0, 303, 302, 1,
		0, 0, 0, 304, 305, 1, 0, 0, 0, 305, 303, 1, 0, 0, 0, 305, 306, 1, 0, 0,
		0, 306, 92, 1, 0, 0, 0, 307, 312, 5, 39, 0, 0, 308, 311, 3, 99, 49, 0,
		309, 311, 3, 105, 52, 0, 310, 308, 1, 0, 0, 0, 310, 309, 1, 0, 0, 0, 311,
		314, 1, 0, 0, 0, 312, 310, 1, 0, 0, 0, 312, 313, 1, 0, 0, 0, 313, 315,
		1, 0, 0, 0, 314, 312, 1, 0, 0, 0, 315, 326, 5, 39, 0, 0, 316, 321, 5, 34,
		0, 0, 317, 320, 3, 101, 50, 0, 318, 320, 3, 105, 52, 0, 319, 317, 1, 0,
		0, 0, 319, 318, 1, 0, 0, 0, 320, 323, 1, 0, 0, 0, 321, 319, 1, 0, 0, 0,
		321, 322, 1, 0, 0, 0, 322, 324, 1, 0, 0, 0, 323, 321, 1, 0, 0, 0, 324,
		326, 5, 34, 0, 0, 325, 307, 1, 0, 0, 0, 325, 316, 1, 0, 0, 0, 326, 94,
		1, 0, 0, 0, 327, 328, 5, 39, 0, 0, 328, 329, 5, 39, 0, 0, 329, 330, 5,
		39, 0, 0, 330, 334, 1, 0, 0, 0, 331, 333, 3, 97, 48, 0, 332, 331, 1, 0,
		0, 0, 333, 336, 1, 0, 0, 0, 334, 335, 1, 0, 0, 0, 334, 332, 1, 0, 0, 0,
		335, 337, 1, 0, 0, 0, 336, 334, 1, 0, 0, 0, 337, 338, 5, 39, 0, 0, 338,
		339, 5, 39, 0, 0, 339, 354, 5, 39, 0, 0, 340, 341, 5, 34, 0, 0, 341, 342,
		5, 34, 0, 0, 342, 343, 5, 34, 0, 0, 343, 347, 1, 0, 0, 0, 344, 346, 3,
		97, 48, 0, 345, 344, 1, 0, 0, 0, 346, 349, 1, 0, 0, 0, 347, 348, 1, 0,
		0, 0, 347, 345, 1, 0, 0, 0, 348, 350, 1, 0, 0, 0, 349, 347, 1, 0, 0, 0,
		350, 351, 5, 34, 0, 0, 351, 352, 5, 34, 0, 0, 352, 354, 5, 34, 0, 0, 353,
		327, 1, 0, 0, 0, 353, 340, 1, 0, 0, 0, 354, 96, 1, 0, 0, 0, 355, 358, 3,
		103, 51, 0, 356, 358, 3, 105, 52, 0, 357, 355, 1, 0, 0, 0, 357, 356, 1,
		0, 0, 0, 358, 98, 1, 0, 0, 0, 359, 361, 7, 11, 0, 0, 360, 359, 1, 0, 0,
		0, 361, 100, 1, 0, 0, 0, 362, 364, 7, 12, 0, 0, 363, 362, 1, 0, 0, 0, 364,
		102, 1, 0, 0, 0, 365, 367, 7, 13, 0, 0, 366, 365, 1, 0, 0, 0, 367, 104,
		1, 0, 0, 0, 368, 369, 5, 92, 0, 0, 369, 370, 7, 14, 0, 0, 370, 106, 1,
		0, 0, 0, 371, 373, 7, 15, 0, 0, 372, 371, 1, 0, 0, 0, 373, 374, 1, 0, 0,
		0, 374, 372, 1, 0, 0, 0, 374, 375, 1, 0, 0, 0, 375, 108, 1, 0, 0, 0, 376,
		378, 5, 92, 0, 0, 377, 379, 3, 107, 53, 0, 378, 377, 1, 0, 0, 0, 378, 379,
		1, 0, 0, 0, 379, 385, 1, 0, 0, 0, 380, 382, 5, 13, 0, 0, 381, 380, 1, 0,
		0, 0, 381, 382, 1, 0, 0, 0, 382, 383, 1, 0, 0, 0, 383, 386, 5, 10, 0, 0,
		384, 386, 2, 12, 13, 0, 385, 381, 1, 0, 0, 0, 385, 384, 1, 0, 0, 0, 386,
		110, 1, 0, 0, 0, 387, 390, 5, 92, 0, 0, 388, 391, 7, 16, 0, 0, 389, 391,
		3, 113, 56, 0, 390, 388, 1, 0, 0, 0, 390, 389, 1, 0, 0, 0, 391, 112, 1,
		0, 0, 0, 392, 393, 5, 117, 0, 0, 393, 394, 3, 117, 58, 0, 394, 395, 3,
		117, 58, 0, 395, 396, 3, 117, 58, 0, 396, 397, 3, 117, 58, 0, 397, 114,
		1, 0, 0, 0, 398, 399, 8, 17, 0, 0, 399, 116, 1, 0, 0, 0, 400, 401, 7, 7,
		0, 0, 401, 118, 1, 0, 0, 0, 39, 0, 188, 195, 202, 204, 214, 220, 225, 230,
		232, 242, 246, 252, 257, 260, 273, 279, 283, 290, 296, 300, 305, 310, 312,
		319, 321, 325, 334, 347, 353, 357, 360, 363, 366, 374, 378, 381, 385, 390,
		1, 6, 0, 0,
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
	valuateLexerT__1              = 2
	valuateLexerT__2              = 3
	valuateLexerTRUE              = 4
	valuateLexerFALSE             = 5
	valuateLexerNIL_LIT           = 6
	valuateLexerLP                = 7
	valuateLexerRP                = 8
	valuateLexerL_CURLY           = 9
	valuateLexerR_CURLY           = 10
	valuateLexerL_BRACKET         = 11
	valuateLexerR_BRACKET         = 12
	valuateLexerCOMMA             = 13
	valuateLexerDOT               = 14
	valuateLexerLOGICAL_OR        = 15
	valuateLexerLOGICAL_AND       = 16
	valuateLexerEQUALS            = 17
	valuateLexerNOT_EQUALS        = 18
	valuateLexerLESS              = 19
	valuateLexerLESS_OR_EQUALS    = 20
	valuateLexerGREATER           = 21
	valuateLexerGREATER_OR_EQUALS = 22
	valuateLexerDIV               = 23
	valuateLexerPLUS              = 24
	valuateLexerMINUS             = 25
	valuateLexerSTAR              = 26
	valuateLexerSKIP_             = 27
	valuateLexerWHITESPACE        = 28
	valuateLexerIDENTIFIER        = 29
	valuateLexerVAR_ID            = 30
	valuateLexerSTRING            = 31
	valuateLexerSTRING_LITERAL    = 32
	valuateLexerBYTES_LITERAL     = 33
	valuateLexerDECIMAL_LIT       = 34
	valuateLexerFLOAT_NUMBER      = 35
)
