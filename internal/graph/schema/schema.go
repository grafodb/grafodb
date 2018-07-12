package schema

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/grafodb/grafodb/internal/util/grammar"
)

type Schema struct {
	Types map[string]Type
}

type Type struct {
	Name   string
	Fields []Field
}

type Field struct {
	Name string
}

func Parse(raw []byte) (*Schema, error) {
	lexer := NewSchemaLexer(antlr.NewInputStream(string(raw)))
	parser := NewSchemaParser(antlr.NewCommonTokenStream(lexer, 0))

	errorListener := grammar.NewParseErrorListener()
	parser.AddErrorListener(errorListener)
	parser.BuildParseTrees = true

	listener := &schemaListener{}
	antlr.NewParseTreeWalker().Walk(listener, parser.Document())

	if errorListener.Err() != nil {
		return nil, errorListener.Err()
	}
	return listener.Schema(), nil
}

type schemaListener struct {
	BaseSchemaListener

	schema   *Schema
	tmpTypes []Type
	tmpType  *Type
}

func (b schemaListener) Schema() *Schema {
	return b.schema
}

// func (b *schemaListener) VisitTerminal(node antlr.TerminalNode) {
//
// }
//
// func (b *schemaListener) VisitErrorNode(node antlr.ErrorNode) {
//
// }
//
// func (b *schemaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
//
// }
//
// func (b *schemaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
//
// }
