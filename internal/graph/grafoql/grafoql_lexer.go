// Code generated from internal/graph/grafoql/GrafoQL.g4 by ANTLR 4.7.1. DO NOT EDIT.

package grafoql

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 15, 131,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7,
	3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11,
	3, 11, 7, 11, 72, 10, 11, 12, 11, 14, 11, 75, 11, 11, 3, 12, 3, 12, 3,
	12, 7, 12, 80, 10, 12, 12, 12, 14, 12, 83, 11, 12, 3, 13, 3, 13, 3, 14,
	3, 14, 3, 15, 6, 15, 90, 10, 15, 13, 15, 14, 15, 91, 3, 15, 3, 15, 3, 16,
	5, 16, 97, 10, 16, 3, 16, 3, 16, 6, 16, 101, 10, 16, 13, 16, 14, 16, 102,
	3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 7, 17, 111, 10, 17, 12, 17, 14,
	17, 114, 11, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18,
	3, 18, 7, 18, 125, 10, 18, 12, 18, 14, 18, 128, 11, 18, 3, 18, 3, 18, 3,
	112, 2, 19, 3, 3, 5, 4, 7, 5, 9, 6, 11, 2, 13, 2, 15, 2, 17, 2, 19, 7,
	21, 8, 23, 9, 25, 10, 27, 11, 29, 12, 31, 13, 33, 14, 35, 15, 3, 2, 8,
	3, 2, 67, 92, 4, 2, 97, 97, 99, 124, 5, 2, 67, 92, 97, 97, 99, 124, 3,
	2, 50, 59, 5, 2, 11, 11, 14, 14, 34, 34, 4, 2, 12, 12, 15, 15, 2, 136,
	2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2,
	2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2,
	2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2,
	2, 2, 2, 35, 3, 2, 2, 2, 3, 37, 3, 2, 2, 2, 5, 43, 3, 2, 2, 2, 7, 45, 3,
	2, 2, 2, 9, 47, 3, 2, 2, 2, 11, 56, 3, 2, 2, 2, 13, 58, 3, 2, 2, 2, 15,
	60, 3, 2, 2, 2, 17, 62, 3, 2, 2, 2, 19, 64, 3, 2, 2, 2, 21, 68, 3, 2, 2,
	2, 23, 76, 3, 2, 2, 2, 25, 84, 3, 2, 2, 2, 27, 86, 3, 2, 2, 2, 29, 89,
	3, 2, 2, 2, 31, 100, 3, 2, 2, 2, 33, 106, 3, 2, 2, 2, 35, 120, 3, 2, 2,
	2, 37, 38, 7, 115, 2, 2, 38, 39, 7, 119, 2, 2, 39, 40, 7, 103, 2, 2, 40,
	41, 7, 116, 2, 2, 41, 42, 7, 123, 2, 2, 42, 4, 3, 2, 2, 2, 43, 44, 7, 125,
	2, 2, 44, 6, 3, 2, 2, 2, 45, 46, 7, 127, 2, 2, 46, 8, 3, 2, 2, 2, 47, 48,
	7, 111, 2, 2, 48, 49, 7, 119, 2, 2, 49, 50, 7, 118, 2, 2, 50, 51, 7, 99,
	2, 2, 51, 52, 7, 118, 2, 2, 52, 53, 7, 107, 2, 2, 53, 54, 7, 113, 2, 2,
	54, 55, 7, 112, 2, 2, 55, 10, 3, 2, 2, 2, 56, 57, 9, 2, 2, 2, 57, 12, 3,
	2, 2, 2, 58, 59, 9, 3, 2, 2, 59, 14, 3, 2, 2, 2, 60, 61, 9, 4, 2, 2, 61,
	16, 3, 2, 2, 2, 62, 63, 9, 5, 2, 2, 63, 18, 3, 2, 2, 2, 64, 65, 5, 25,
	13, 2, 65, 66, 5, 21, 11, 2, 66, 67, 5, 27, 14, 2, 67, 20, 3, 2, 2, 2,
	68, 73, 5, 11, 6, 2, 69, 72, 5, 15, 8, 2, 70, 72, 5, 17, 9, 2, 71, 69,
	3, 2, 2, 2, 71, 70, 3, 2, 2, 2, 72, 75, 3, 2, 2, 2, 73, 71, 3, 2, 2, 2,
	73, 74, 3, 2, 2, 2, 74, 22, 3, 2, 2, 2, 75, 73, 3, 2, 2, 2, 76, 81, 5,
	13, 7, 2, 77, 80, 5, 15, 8, 2, 78, 80, 5, 17, 9, 2, 79, 77, 3, 2, 2, 2,
	79, 78, 3, 2, 2, 2, 80, 83, 3, 2, 2, 2, 81, 79, 3, 2, 2, 2, 81, 82, 3,
	2, 2, 2, 82, 24, 3, 2, 2, 2, 83, 81, 3, 2, 2, 2, 84, 85, 7, 93, 2, 2, 85,
	26, 3, 2, 2, 2, 86, 87, 7, 95, 2, 2, 87, 28, 3, 2, 2, 2, 88, 90, 9, 6,
	2, 2, 89, 88, 3, 2, 2, 2, 90, 91, 3, 2, 2, 2, 91, 89, 3, 2, 2, 2, 91, 92,
	3, 2, 2, 2, 92, 93, 3, 2, 2, 2, 93, 94, 8, 15, 2, 2, 94, 30, 3, 2, 2, 2,
	95, 97, 7, 15, 2, 2, 96, 95, 3, 2, 2, 2, 96, 97, 3, 2, 2, 2, 97, 98, 3,
	2, 2, 2, 98, 101, 7, 12, 2, 2, 99, 101, 7, 15, 2, 2, 100, 96, 3, 2, 2,
	2, 100, 99, 3, 2, 2, 2, 101, 102, 3, 2, 2, 2, 102, 100, 3, 2, 2, 2, 102,
	103, 3, 2, 2, 2, 103, 104, 3, 2, 2, 2, 104, 105, 8, 16, 2, 2, 105, 32,
	3, 2, 2, 2, 106, 107, 7, 49, 2, 2, 107, 108, 7, 44, 2, 2, 108, 112, 3,
	2, 2, 2, 109, 111, 11, 2, 2, 2, 110, 109, 3, 2, 2, 2, 111, 114, 3, 2, 2,
	2, 112, 113, 3, 2, 2, 2, 112, 110, 3, 2, 2, 2, 113, 115, 3, 2, 2, 2, 114,
	112, 3, 2, 2, 2, 115, 116, 7, 44, 2, 2, 116, 117, 7, 49, 2, 2, 117, 118,
	3, 2, 2, 2, 118, 119, 8, 17, 2, 2, 119, 34, 3, 2, 2, 2, 120, 121, 7, 49,
	2, 2, 121, 122, 7, 49, 2, 2, 122, 126, 3, 2, 2, 2, 123, 125, 10, 7, 2,
	2, 124, 123, 3, 2, 2, 2, 125, 128, 3, 2, 2, 2, 126, 124, 3, 2, 2, 2, 126,
	127, 3, 2, 2, 2, 127, 129, 3, 2, 2, 2, 128, 126, 3, 2, 2, 2, 129, 130,
	8, 18, 2, 2, 130, 36, 3, 2, 2, 2, 13, 2, 71, 73, 79, 81, 91, 96, 100, 102,
	112, 126, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'query'", "'{'", "'}'", "'mutation'", "", "", "", "'['", "']'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "TypeIdentArray", "TypeIdent", "FieldIdent", "LBRACK",
	"RBRACK", "WS", "NEWLINE", "COMMENT", "LINE_COMMENT",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "ULetter", "LLetter", "Letter", "Digit",
	"TypeIdentArray", "TypeIdent", "FieldIdent", "LBRACK", "RBRACK", "WS",
	"NEWLINE", "COMMENT", "LINE_COMMENT",
}

type GrafoQLLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewGrafoQLLexer(input antlr.CharStream) *GrafoQLLexer {

	l := new(GrafoQLLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "GrafoQL.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// GrafoQLLexer tokens.
const (
	GrafoQLLexerT__0           = 1
	GrafoQLLexerT__1           = 2
	GrafoQLLexerT__2           = 3
	GrafoQLLexerT__3           = 4
	GrafoQLLexerTypeIdentArray = 5
	GrafoQLLexerTypeIdent      = 6
	GrafoQLLexerFieldIdent     = 7
	GrafoQLLexerLBRACK         = 8
	GrafoQLLexerRBRACK         = 9
	GrafoQLLexerWS             = 10
	GrafoQLLexerNEWLINE        = 11
	GrafoQLLexerCOMMENT        = 12
	GrafoQLLexerLINE_COMMENT   = 13
)
