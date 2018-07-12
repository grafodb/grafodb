// Code generated from internal/graph/grafoql/GrafoQL.g4 by ANTLR 4.7.1. DO NOT EDIT.

package grafoql // GrafoQL
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 15, 43, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3, 2, 5,
	2, 15, 10, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 7, 4, 24, 10, 4,
	12, 4, 14, 4, 27, 11, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 7, 6,
	36, 10, 6, 12, 6, 14, 6, 39, 11, 6, 3, 6, 3, 6, 3, 6, 4, 25, 37, 2, 7,
	2, 4, 6, 8, 10, 2, 2, 2, 40, 2, 14, 3, 2, 2, 2, 4, 18, 3, 2, 2, 2, 6, 21,
	3, 2, 2, 2, 8, 30, 3, 2, 2, 2, 10, 33, 3, 2, 2, 2, 12, 15, 5, 4, 3, 2,
	13, 15, 5, 8, 5, 2, 14, 12, 3, 2, 2, 2, 14, 13, 3, 2, 2, 2, 15, 16, 3,
	2, 2, 2, 16, 17, 7, 2, 2, 3, 17, 3, 3, 2, 2, 2, 18, 19, 7, 3, 2, 2, 19,
	20, 5, 6, 4, 2, 20, 5, 3, 2, 2, 2, 21, 25, 7, 4, 2, 2, 22, 24, 11, 2, 2,
	2, 23, 22, 3, 2, 2, 2, 24, 27, 3, 2, 2, 2, 25, 26, 3, 2, 2, 2, 25, 23,
	3, 2, 2, 2, 26, 28, 3, 2, 2, 2, 27, 25, 3, 2, 2, 2, 28, 29, 7, 5, 2, 2,
	29, 7, 3, 2, 2, 2, 30, 31, 7, 6, 2, 2, 31, 32, 5, 6, 4, 2, 32, 9, 3, 2,
	2, 2, 33, 37, 7, 4, 2, 2, 34, 36, 11, 2, 2, 2, 35, 34, 3, 2, 2, 2, 36,
	39, 3, 2, 2, 2, 37, 38, 3, 2, 2, 2, 37, 35, 3, 2, 2, 2, 38, 40, 3, 2, 2,
	2, 39, 37, 3, 2, 2, 2, 40, 41, 7, 5, 2, 2, 41, 11, 3, 2, 2, 2, 5, 14, 25,
	37,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'query'", "'{'", "'}'", "'mutation'", "", "", "", "'['", "']'",
}
var symbolicNames = []string{
	"", "", "", "", "", "TypeIdentArray", "TypeIdent", "FieldIdent", "LBRACK",
	"RBRACK", "WS", "NEWLINE", "COMMENT", "LINE_COMMENT",
}

var ruleNames = []string{
	"document", "query", "schemaBody", "mutation", "mutationBody",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type GrafoQLParser struct {
	*antlr.BaseParser
}

func NewGrafoQLParser(input antlr.TokenStream) *GrafoQLParser {
	this := new(GrafoQLParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "GrafoQL.g4"

	return this
}

// GrafoQLParser tokens.
const (
	GrafoQLParserEOF            = antlr.TokenEOF
	GrafoQLParserT__0           = 1
	GrafoQLParserT__1           = 2
	GrafoQLParserT__2           = 3
	GrafoQLParserT__3           = 4
	GrafoQLParserTypeIdentArray = 5
	GrafoQLParserTypeIdent      = 6
	GrafoQLParserFieldIdent     = 7
	GrafoQLParserLBRACK         = 8
	GrafoQLParserRBRACK         = 9
	GrafoQLParserWS             = 10
	GrafoQLParserNEWLINE        = 11
	GrafoQLParserCOMMENT        = 12
	GrafoQLParserLINE_COMMENT   = 13
)

// GrafoQLParser rules.
const (
	GrafoQLParserRULE_document     = 0
	GrafoQLParserRULE_query        = 1
	GrafoQLParserRULE_schemaBody   = 2
	GrafoQLParserRULE_mutation     = 3
	GrafoQLParserRULE_mutationBody = 4
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
	p.RuleIndex = GrafoQLParserRULE_document
	return p
}

func (*DocumentContext) IsDocumentContext() {}

func NewDocumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DocumentContext {
	var p = new(DocumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrafoQLParserRULE_document

	return p
}

func (s *DocumentContext) GetParser() antlr.Parser { return s.parser }

func (s *DocumentContext) EOF() antlr.TerminalNode {
	return s.GetToken(GrafoQLParserEOF, 0)
}

func (s *DocumentContext) Query() IQueryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQueryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQueryContext)
}

func (s *DocumentContext) Mutation() IMutationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMutationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMutationContext)
}

func (s *DocumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DocumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DocumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrafoQLListener); ok {
		listenerT.EnterDocument(s)
	}
}

func (s *DocumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrafoQLListener); ok {
		listenerT.ExitDocument(s)
	}
}

func (p *GrafoQLParser) Document() (localctx IDocumentContext) {
	localctx = NewDocumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GrafoQLParserRULE_document)

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
	p.SetState(12)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrafoQLParserT__0:
		{
			p.SetState(10)
			p.Query()
		}

	case GrafoQLParserT__3:
		{
			p.SetState(11)
			p.Mutation()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(14)
		p.Match(GrafoQLParserEOF)
	}

	return localctx
}

// IQueryContext is an interface to support dynamic dispatch.
type IQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQueryContext differentiates from other interfaces.
	IsQueryContext()
}

type QueryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryContext() *QueryContext {
	var p = new(QueryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrafoQLParserRULE_query
	return p
}

func (*QueryContext) IsQueryContext() {}

func NewQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryContext {
	var p = new(QueryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrafoQLParserRULE_query

	return p
}

func (s *QueryContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryContext) SchemaBody() ISchemaBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISchemaBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISchemaBodyContext)
}

func (s *QueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrafoQLListener); ok {
		listenerT.EnterQuery(s)
	}
}

func (s *QueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrafoQLListener); ok {
		listenerT.ExitQuery(s)
	}
}

func (p *GrafoQLParser) Query() (localctx IQueryContext) {
	localctx = NewQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GrafoQLParserRULE_query)

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
		p.SetState(16)
		p.Match(GrafoQLParserT__0)
	}
	{
		p.SetState(17)
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
	p.RuleIndex = GrafoQLParserRULE_schemaBody
	return p
}

func (*SchemaBodyContext) IsSchemaBodyContext() {}

func NewSchemaBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SchemaBodyContext {
	var p = new(SchemaBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrafoQLParserRULE_schemaBody

	return p
}

func (s *SchemaBodyContext) GetParser() antlr.Parser { return s.parser }
func (s *SchemaBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SchemaBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SchemaBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrafoQLListener); ok {
		listenerT.EnterSchemaBody(s)
	}
}

func (s *SchemaBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrafoQLListener); ok {
		listenerT.ExitSchemaBody(s)
	}
}

func (p *GrafoQLParser) SchemaBody() (localctx ISchemaBodyContext) {
	localctx = NewSchemaBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GrafoQLParserRULE_schemaBody)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(19)
		p.Match(GrafoQLParserT__1)
	}
	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			p.SetState(20)
			p.MatchWildcard()

		}
		p.SetState(25)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}
	{
		p.SetState(26)
		p.Match(GrafoQLParserT__2)
	}

	return localctx
}

// IMutationContext is an interface to support dynamic dispatch.
type IMutationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMutationContext differentiates from other interfaces.
	IsMutationContext()
}

type MutationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMutationContext() *MutationContext {
	var p = new(MutationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrafoQLParserRULE_mutation
	return p
}

func (*MutationContext) IsMutationContext() {}

func NewMutationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MutationContext {
	var p = new(MutationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrafoQLParserRULE_mutation

	return p
}

func (s *MutationContext) GetParser() antlr.Parser { return s.parser }

func (s *MutationContext) SchemaBody() ISchemaBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISchemaBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISchemaBodyContext)
}

func (s *MutationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MutationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MutationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrafoQLListener); ok {
		listenerT.EnterMutation(s)
	}
}

func (s *MutationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrafoQLListener); ok {
		listenerT.ExitMutation(s)
	}
}

func (p *GrafoQLParser) Mutation() (localctx IMutationContext) {
	localctx = NewMutationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GrafoQLParserRULE_mutation)

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
		p.SetState(28)
		p.Match(GrafoQLParserT__3)
	}
	{
		p.SetState(29)
		p.SchemaBody()
	}

	return localctx
}

// IMutationBodyContext is an interface to support dynamic dispatch.
type IMutationBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMutationBodyContext differentiates from other interfaces.
	IsMutationBodyContext()
}

type MutationBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMutationBodyContext() *MutationBodyContext {
	var p = new(MutationBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrafoQLParserRULE_mutationBody
	return p
}

func (*MutationBodyContext) IsMutationBodyContext() {}

func NewMutationBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MutationBodyContext {
	var p = new(MutationBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrafoQLParserRULE_mutationBody

	return p
}

func (s *MutationBodyContext) GetParser() antlr.Parser { return s.parser }
func (s *MutationBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MutationBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MutationBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrafoQLListener); ok {
		listenerT.EnterMutationBody(s)
	}
}

func (s *MutationBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrafoQLListener); ok {
		listenerT.ExitMutationBody(s)
	}
}

func (p *GrafoQLParser) MutationBody() (localctx IMutationBodyContext) {
	localctx = NewMutationBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GrafoQLParserRULE_mutationBody)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(31)
		p.Match(GrafoQLParserT__1)
	}
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			p.SetState(32)
			p.MatchWildcard()

		}
		p.SetState(37)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}
	{
		p.SetState(38)
		p.Match(GrafoQLParserT__2)
	}

	return localctx
}
