package grammar

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func NewParseErrorListener() *ParseErrorListener {
	return &ParseErrorListener{}
}

type ParseErrorListener struct {
	antlr.DefaultErrorListener
	errors []error
}

func (l *ParseErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.errors = append(l.errors, fmt.Errorf("[%d:%d] %s ", line, column, msg))
}

func (l ParseErrorListener) Err() error {
	if len(l.errors) > 0 {
		return l.errors[0]
	}
	return nil
}
