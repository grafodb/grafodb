// Code generated from internal/graph/schema/Schema.g4 by ANTLR 4.7.1. DO NOT EDIT.

package schema // Schema
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 22, 61, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 3, 2, 3, 2, 3, 2, 3, 3,
	3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 7, 5, 35, 10, 5, 12, 5,
	14, 5, 38, 11, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 6, 8,
	48, 10, 8, 13, 8, 14, 8, 49, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10,
	3, 11, 3, 11, 3, 11, 2, 2, 12, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 2, 3,
	3, 2, 7, 15, 2, 52, 2, 22, 3, 2, 2, 2, 4, 25, 3, 2, 2, 2, 6, 28, 3, 2,
	2, 2, 8, 32, 3, 2, 2, 2, 10, 39, 3, 2, 2, 2, 12, 43, 3, 2, 2, 2, 14, 45,
	3, 2, 2, 2, 16, 53, 3, 2, 2, 2, 18, 56, 3, 2, 2, 2, 20, 58, 3, 2, 2, 2,
	22, 23, 5, 4, 3, 2, 23, 24, 7, 2, 2, 3, 24, 3, 3, 2, 2, 2, 25, 26, 7, 3,
	2, 2, 26, 27, 5, 6, 4, 2, 27, 5, 3, 2, 2, 2, 28, 29, 7, 4, 2, 2, 29, 30,
	5, 8, 5, 2, 30, 31, 7, 5, 2, 2, 31, 7, 3, 2, 2, 2, 32, 36, 5, 10, 6, 2,
	33, 35, 5, 10, 6, 2, 34, 33, 3, 2, 2, 2, 35, 38, 3, 2, 2, 2, 36, 34, 3,
	2, 2, 2, 36, 37, 3, 2, 2, 2, 37, 9, 3, 2, 2, 2, 38, 36, 3, 2, 2, 2, 39,
	40, 7, 6, 2, 2, 40, 41, 5, 12, 7, 2, 41, 42, 5, 14, 8, 2, 42, 11, 3, 2,
	2, 2, 43, 44, 7, 15, 2, 2, 44, 13, 3, 2, 2, 2, 45, 47, 7, 4, 2, 2, 46,
	48, 5, 16, 9, 2, 47, 46, 3, 2, 2, 2, 48, 49, 3, 2, 2, 2, 49, 47, 3, 2,
	2, 2, 49, 50, 3, 2, 2, 2, 50, 51, 3, 2, 2, 2, 51, 52, 7, 5, 2, 2, 52, 15,
	3, 2, 2, 2, 53, 54, 5, 18, 10, 2, 54, 55, 5, 20, 11, 2, 55, 17, 3, 2, 2,
	2, 56, 57, 7, 16, 2, 2, 57, 19, 3, 2, 2, 2, 58, 59, 9, 2, 2, 2, 59, 21,
	3, 2, 2, 2, 4, 36, 49,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'schema'", "'{'", "'}'", "'type'", "'ID'", "'int'", "'float'", "'string'",
	"'bool'", "'datetime'", "'geopoint'", "", "", "", "'['", "']'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "TypeIdentArray", "TypeIdent",
	"FieldIdent", "LBRACK", "RBRACK", "WS", "NEWLINE", "COMMENT", "LINE_COMMENT",
}

var ruleNames = []string{
	"document", "schema", "schemaBody", "schemaTypes", "typeDecl", "typeName",
	"typeBody", "field", "fieldName", "fieldType",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type SchemaParser struct {
	*antlr.BaseParser
}

func NewSchemaParser(input antlr.TokenStream) *SchemaParser {
	this := new(SchemaParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Schema.g4"

	return this
}

// SchemaParser tokens.
const (
	SchemaParserEOF            = antlr.TokenEOF
	SchemaParserT__0           = 1
	SchemaParserT__1           = 2
	SchemaParserT__2           = 3
	SchemaParserT__3           = 4
	SchemaParserT__4           = 5
	SchemaParserT__5           = 6
	SchemaParserT__6           = 7
	SchemaParserT__7           = 8
	SchemaParserT__8           = 9
	SchemaParserT__9           = 10
	SchemaParserT__10          = 11
	SchemaParserTypeIdentArray = 12
	SchemaParserTypeIdent      = 13
	SchemaParserFieldIdent     = 14
	SchemaParserLBRACK         = 15
	SchemaParserRBRACK         = 16
	SchemaParserWS             = 17
	SchemaParserNEWLINE        = 18
	SchemaParserCOMMENT        = 19
	SchemaParserLINE_COMMENT   = 20
)

// SchemaParser rules.
const (
	SchemaParserRULE_document    = 0
	SchemaParserRULE_schema      = 1
	SchemaParserRULE_schemaBody  = 2
	SchemaParserRULE_schemaTypes = 3
	SchemaParserRULE_typeDecl    = 4
	SchemaParserRULE_typeName    = 5
	SchemaParserRULE_typeBody    = 6
	SchemaParserRULE_field       = 7
	SchemaParserRULE_fieldName   = 8
	SchemaParserRULE_fieldType   = 9
)

// IDocumentContext is an interface to support dynamic dispatch.
type IDocumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDocumentContext differentiates from other interfaces.
	IsDocumentContext()
}

type DocumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDocumentContext() *DocumentContext {
	var p = new(DocumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SchemaParserRULE_document
	return p
}

func (*DocumentContext) IsDocumentContext() {}

func NewDocumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DocumentContext {
	var p = new(DocumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SchemaParserRULE_document

	return p
}

func (s *DocumentContext) GetParser() antlr.Parser { return s.parser }

func (s *DocumentContext) Schema() ISchemaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISchemaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISchemaContext)
}

func (s *DocumentContext) EOF() antlr.TerminalNode {
	return s.GetToken(SchemaParserEOF, 0)
}

func (s *DocumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DocumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DocumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SchemaListener); ok {
		listenerT.EnterDocument(s)
	}
}

func (s *DocumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SchemaListener); ok {
		listenerT.ExitDocument(s)
	}
}

func (p *SchemaParser) Document() (localctx IDocumentContext) {
	localctx = NewDocumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SchemaParserRULE_document)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(20)
		p.Schema()
	}
	{
		p.SetState(21)
		p.Match(SchemaParserEOF)
	}

	return localctx
}

// ISchemaContext is an interface to support dynamic dispatch.
type ISchemaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSchemaContext differentiates from other interfaces.
	IsSchemaContext()
}

type SchemaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySchemaContext() *SchemaContext {
	var p = new(SchemaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SchemaParserRULE_schema
	return p
}

func (*SchemaContext) IsSchemaContext() {}

func NewSchemaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SchemaContext {
	var p = new(SchemaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SchemaParserRULE_schema

	return p
}

func (s *SchemaContext) GetParser() antlr.Parser { return s.parser }

func (s *SchemaContext) SchemaBody() ISchemaBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISchemaBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISchemaBodyContext)
}

func (s *SchemaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SchemaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SchemaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SchemaListener); ok {
		listenerT.EnterSchema(s)
	}
}

func (s *SchemaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SchemaListener); ok {
		listenerT.ExitSchema(s)
	}
}

func (p *SchemaParser) Schema() (localctx ISchemaContext) {
	localctx = NewSchemaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SchemaParserRULE_schema)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(23)
		p.Match(SchemaParserT__0)
	}
	{
		p.SetState(24)
		p.SchemaBody()
	}

	return localctx
}

// ISchemaBodyContext is an interface to support dynamic dispatch.
type ISchemaBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSchemaBodyContext differentiates from other interfaces.
	IsSchemaBodyContext()
}

type SchemaBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySchemaBodyContext() *SchemaBodyContext {
	var p = new(SchemaBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SchemaParserRULE_schemaBody
	return p
}

func (*SchemaBodyContext) IsSchemaBodyContext() {}

func NewSchemaBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SchemaBodyContext {
	var p = new(SchemaBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SchemaParserRULE_schemaBody

	return p
}

func (s *SchemaBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *SchemaBodyContext) SchemaTypes() ISchemaTypesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISchemaTypesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISchemaTypesContext)
}

func (s *SchemaBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SchemaBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SchemaBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SchemaListener); ok {
		listenerT.EnterSchemaBody(s)
	}
}

func (s *SchemaBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SchemaListener); ok {
		listenerT.ExitSchemaBody(s)
	}
}

func (p *SchemaParser) SchemaBody() (localctx ISchemaBodyContext) {
	localctx = NewSchemaBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SchemaParserRULE_schemaBody)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(26)
		p.Match(SchemaParserT__1)
	}
	{
		p.SetState(27)
		p.SchemaTypes()
	}
	{
		p.SetState(28)
		p.Match(SchemaParserT__2)
	}

	return localctx
}

// ISchemaTypesContext is an interface to support dynamic dispatch.
type ISchemaTypesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSchemaTypesContext differentiates from other interfaces.
	IsSchemaTypesContext()
}

type SchemaTypesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySchemaTypesContext() *SchemaTypesContext {
	var p = new(SchemaTypesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SchemaParserRULE_schemaTypes
	return p
}

func (*SchemaTypesContext) IsSchemaTypesContext() {}

func NewSchemaTypesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SchemaTypesContext {
	var p = new(SchemaTypesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SchemaParserRULE_schemaTypes

	return p
}

func (s *SchemaTypesContext) GetParser() antlr.Parser { return s.parser }

func (s *SchemaTypesContext) AllTypeDecl() []ITypeDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeDeclContext)(nil)).Elem())
	var tst = make([]ITypeDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeDeclContext)
		}
	}

	return tst
}

func (s *SchemaTypesContext) TypeDecl(i int) ITypeDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeDeclContext)
}

func (s *SchemaTypesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SchemaTypesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SchemaTypesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SchemaListener); ok {
		listenerT.EnterSchemaTypes(s)
	}
}

func (s *SchemaTypesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SchemaListener); ok {
		listenerT.ExitSchemaTypes(s)
	}
}

func (p *SchemaParser) SchemaTypes() (localctx ISchemaTypesContext) {
	localctx = NewSchemaTypesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SchemaParserRULE_schemaTypes)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(30)
		p.TypeDecl()
	}
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SchemaParserT__3 {
		{
			p.SetState(31)
			p.TypeDecl()
		}

		p.SetState(36)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITypeDeclContext is an interface to support dynamic dispatch.
type ITypeDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeDeclContext differentiates from other interfaces.
	IsTypeDeclContext()
}

type TypeDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDeclContext() *TypeDeclContext {
	var p = new(TypeDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SchemaParserRULE_typeDecl
	return p
}

func (*TypeDeclContext) IsTypeDeclContext() {}

func NewTypeDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDeclContext {
	var p = new(TypeDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SchemaParserRULE_typeDecl

	return p
}

func (s *TypeDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDeclContext) TypeName() ITypeNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *TypeDeclContext) TypeBody() ITypeBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeBodyContext)
}

func (s *TypeDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SchemaListener); ok {
		listenerT.EnterTypeDecl(s)
	}
}

func (s *TypeDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SchemaListener); ok {
		listenerT.ExitTypeDecl(s)
	}
}

func (p *SchemaParser) TypeDecl() (localctx ITypeDeclContext) {
	localctx = NewTypeDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SchemaParserRULE_typeDecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(37)
		p.Match(SchemaParserT__3)
	}
	{
		p.SetState(38)
		p.TypeName()
	}
	{
		p.SetState(39)
		p.TypeBody()
	}

	return localctx
}

// ITypeNameContext is an interface to support dynamic dispatch.
type ITypeNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeNameContext differentiates from other interfaces.
	IsTypeNameContext()
}

type TypeNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeNameContext() *TypeNameContext {
	var p = new(TypeNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SchemaParserRULE_typeName
	return p
}

func (*TypeNameContext) IsTypeNameContext() {}

func NewTypeNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeNameContext {
	var p = new(TypeNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SchemaParserRULE_typeName

	return p
}

func (s *TypeNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeNameContext) TypeIdent() antlr.TerminalNode {
	return s.GetToken(SchemaParserTypeIdent, 0)
}

func (s *TypeNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SchemaListener); ok {
		listenerT.EnterTypeName(s)
	}
}

func (s *TypeNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SchemaListener); ok {
		listenerT.ExitTypeName(s)
	}
}

func (p *SchemaParser) TypeName() (localctx ITypeNameContext) {
	localctx = NewTypeNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SchemaParserRULE_typeName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(41)
		p.Match(SchemaParserTypeIdent)
	}

	return localctx
}

// ITypeBodyContext is an interface to support dynamic dispatch.
type ITypeBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeBodyContext differentiates from other interfaces.
	IsTypeBodyContext()
}

type TypeBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeBodyContext() *TypeBodyContext {
	var p = new(TypeBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SchemaParserRULE_typeBody
	return p
}

func (*TypeBodyContext) IsTypeBodyContext() {}

func NewTypeBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeBodyContext {
	var p = new(TypeBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SchemaParserRULE_typeBody

	return p
}

func (s *TypeBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeBodyContext) AllField() []IFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldContext)(nil)).Elem())
	var tst = make([]IFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldContext)
		}
	}

	return tst
}

func (s *TypeBodyContext) Field(i int) IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *TypeBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SchemaListener); ok {
		listenerT.EnterTypeBody(s)
	}
}

func (s *TypeBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SchemaListener); ok {
		listenerT.ExitTypeBody(s)
	}
}

func (p *SchemaParser) TypeBody() (localctx ITypeBodyContext) {
	localctx = NewTypeBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SchemaParserRULE_typeBody)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(43)
		p.Match(SchemaParserT__1)
	}
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SchemaParserFieldIdent {
		{
			p.SetState(44)
			p.Field()
		}

		p.SetState(47)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(49)
		p.Match(SchemaParserT__2)
	}

	return localctx
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SchemaParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SchemaParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) FieldName() IFieldNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *FieldContext) FieldType() IFieldTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldTypeContext)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SchemaListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SchemaListener); ok {
		listenerT.ExitField(s)
	}
}

func (p *SchemaParser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SchemaParserRULE_field)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(51)
		p.FieldName()
	}
	{
		p.SetState(52)
		p.FieldType()
	}

	return localctx
}

// IFieldNameContext is an interface to support dynamic dispatch.
type IFieldNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldNameContext differentiates from other interfaces.
	IsFieldNameContext()
}

type FieldNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldNameContext() *FieldNameContext {
	var p = new(FieldNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SchemaParserRULE_fieldName
	return p
}

func (*FieldNameContext) IsFieldNameContext() {}

func NewFieldNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldNameContext {
	var p = new(FieldNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SchemaParserRULE_fieldName

	return p
}

func (s *FieldNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldNameContext) FieldIdent() antlr.TerminalNode {
	return s.GetToken(SchemaParserFieldIdent, 0)
}

func (s *FieldNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SchemaListener); ok {
		listenerT.EnterFieldName(s)
	}
}

func (s *FieldNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SchemaListener); ok {
		listenerT.ExitFieldName(s)
	}
}

func (p *SchemaParser) FieldName() (localctx IFieldNameContext) {
	localctx = NewFieldNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SchemaParserRULE_fieldName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(54)
		p.Match(SchemaParserFieldIdent)
	}

	return localctx
}

// IFieldTypeContext is an interface to support dynamic dispatch.
type IFieldTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldTypeContext differentiates from other interfaces.
	IsFieldTypeContext()
}

type FieldTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldTypeContext() *FieldTypeContext {
	var p = new(FieldTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SchemaParserRULE_fieldType
	return p
}

func (*FieldTypeContext) IsFieldTypeContext() {}

func NewFieldTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldTypeContext {
	var p = new(FieldTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SchemaParserRULE_fieldType

	return p
}

func (s *FieldTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldTypeContext) TypeIdent() antlr.TerminalNode {
	return s.GetToken(SchemaParserTypeIdent, 0)
}

func (s *FieldTypeContext) TypeIdentArray() antlr.TerminalNode {
	return s.GetToken(SchemaParserTypeIdentArray, 0)
}

func (s *FieldTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SchemaListener); ok {
		listenerT.EnterFieldType(s)
	}
}

func (s *FieldTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SchemaListener); ok {
		listenerT.ExitFieldType(s)
	}
}

func (p *SchemaParser) FieldType() (localctx IFieldTypeContext) {
	localctx = NewFieldTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SchemaParserRULE_fieldType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SchemaParserT__4)|(1<<SchemaParserT__5)|(1<<SchemaParserT__6)|(1<<SchemaParserT__7)|(1<<SchemaParserT__8)|(1<<SchemaParserT__9)|(1<<SchemaParserT__10)|(1<<SchemaParserTypeIdentArray)|(1<<SchemaParserTypeIdent))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
