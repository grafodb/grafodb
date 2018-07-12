// Code generated from internal/graph/schema/Schema.g4 by ANTLR 4.7.1. DO NOT EDIT.

package schema // Schema
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSchemaListener is a complete listener for a parse tree produced by SchemaParser.
type BaseSchemaListener struct{}

var _ SchemaListener = &BaseSchemaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSchemaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSchemaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSchemaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSchemaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterDocument is called when production document is entered.
func (s *BaseSchemaListener) EnterDocument(ctx *DocumentContext) {}

// ExitDocument is called when production document is exited.
func (s *BaseSchemaListener) ExitDocument(ctx *DocumentContext) {}

// EnterSchema is called when production schema is entered.
func (s *BaseSchemaListener) EnterSchema(ctx *SchemaContext) {}

// ExitSchema is called when production schema is exited.
func (s *BaseSchemaListener) ExitSchema(ctx *SchemaContext) {}

// EnterSchemaBody is called when production schemaBody is entered.
func (s *BaseSchemaListener) EnterSchemaBody(ctx *SchemaBodyContext) {}

// ExitSchemaBody is called when production schemaBody is exited.
func (s *BaseSchemaListener) ExitSchemaBody(ctx *SchemaBodyContext) {}

// EnterSchemaTypes is called when production schemaTypes is entered.
func (s *BaseSchemaListener) EnterSchemaTypes(ctx *SchemaTypesContext) {}

// ExitSchemaTypes is called when production schemaTypes is exited.
func (s *BaseSchemaListener) ExitSchemaTypes(ctx *SchemaTypesContext) {}

// EnterTypeDecl is called when production typeDecl is entered.
func (s *BaseSchemaListener) EnterTypeDecl(ctx *TypeDeclContext) {}

// ExitTypeDecl is called when production typeDecl is exited.
func (s *BaseSchemaListener) ExitTypeDecl(ctx *TypeDeclContext) {}

// EnterTypeName is called when production typeName is entered.
func (s *BaseSchemaListener) EnterTypeName(ctx *TypeNameContext) {}

// ExitTypeName is called when production typeName is exited.
func (s *BaseSchemaListener) ExitTypeName(ctx *TypeNameContext) {}

// EnterTypeBody is called when production typeBody is entered.
func (s *BaseSchemaListener) EnterTypeBody(ctx *TypeBodyContext) {}

// ExitTypeBody is called when production typeBody is exited.
func (s *BaseSchemaListener) ExitTypeBody(ctx *TypeBodyContext) {}

// EnterField is called when production field is entered.
func (s *BaseSchemaListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseSchemaListener) ExitField(ctx *FieldContext) {}

// EnterFieldName is called when production fieldName is entered.
func (s *BaseSchemaListener) EnterFieldName(ctx *FieldNameContext) {}

// ExitFieldName is called when production fieldName is exited.
func (s *BaseSchemaListener) ExitFieldName(ctx *FieldNameContext) {}

// EnterFieldType is called when production fieldType is entered.
func (s *BaseSchemaListener) EnterFieldType(ctx *FieldTypeContext) {}

// ExitFieldType is called when production fieldType is exited.
func (s *BaseSchemaListener) ExitFieldType(ctx *FieldTypeContext) {}
