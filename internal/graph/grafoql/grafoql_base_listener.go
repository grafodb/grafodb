// Code generated from internal/graph/grafoql/GrafoQL.g4 by ANTLR 4.7.1. DO NOT EDIT.

package grafoql // GrafoQL
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseGrafoQLListener is a complete listener for a parse tree produced by GrafoQLParser.
type BaseGrafoQLListener struct{}

var _ GrafoQLListener = &BaseGrafoQLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGrafoQLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGrafoQLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGrafoQLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGrafoQLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterDocument is called when production document is entered.
func (s *BaseGrafoQLListener) EnterDocument(ctx *DocumentContext) {}

// ExitDocument is called when production document is exited.
func (s *BaseGrafoQLListener) ExitDocument(ctx *DocumentContext) {}

// EnterQuery is called when production query is entered.
func (s *BaseGrafoQLListener) EnterQuery(ctx *QueryContext) {}

// ExitQuery is called when production query is exited.
func (s *BaseGrafoQLListener) ExitQuery(ctx *QueryContext) {}

// EnterSchemaBody is called when production schemaBody is entered.
func (s *BaseGrafoQLListener) EnterSchemaBody(ctx *SchemaBodyContext) {}

// ExitSchemaBody is called when production schemaBody is exited.
func (s *BaseGrafoQLListener) ExitSchemaBody(ctx *SchemaBodyContext) {}

// EnterMutation is called when production mutation is entered.
func (s *BaseGrafoQLListener) EnterMutation(ctx *MutationContext) {}

// ExitMutation is called when production mutation is exited.
func (s *BaseGrafoQLListener) ExitMutation(ctx *MutationContext) {}

// EnterMutationBody is called when production mutationBody is entered.
func (s *BaseGrafoQLListener) EnterMutationBody(ctx *MutationBodyContext) {}

// ExitMutationBody is called when production mutationBody is exited.
func (s *BaseGrafoQLListener) ExitMutationBody(ctx *MutationBodyContext) {}
