// Code generated from internal/graph/schema/Schema.g4 by ANTLR 4.7.1. DO NOT EDIT.

package schema // Schema
import "github.com/antlr/antlr4/runtime/Go/antlr"

// SchemaListener is a complete listener for a parse tree produced by SchemaParser.
type SchemaListener interface {
	antlr.ParseTreeListener

	// EnterDocument is called when entering the document production.
	EnterDocument(c *DocumentContext)

	// EnterSchema is called when entering the schema production.
	EnterSchema(c *SchemaContext)

	// EnterSchemaBody is called when entering the schemaBody production.
	EnterSchemaBody(c *SchemaBodyContext)

	// EnterSchemaTypes is called when entering the schemaTypes production.
	EnterSchemaTypes(c *SchemaTypesContext)

	// EnterTypeDecl is called when entering the typeDecl production.
	EnterTypeDecl(c *TypeDeclContext)

	// EnterTypeName is called when entering the typeName production.
	EnterTypeName(c *TypeNameContext)

	// EnterTypeBody is called when entering the typeBody production.
	EnterTypeBody(c *TypeBodyContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterFieldName is called when entering the fieldName production.
	EnterFieldName(c *FieldNameContext)

	// EnterFieldType is called when entering the fieldType production.
	EnterFieldType(c *FieldTypeContext)

	// ExitDocument is called when exiting the document production.
	ExitDocument(c *DocumentContext)

	// ExitSchema is called when exiting the schema production.
	ExitSchema(c *SchemaContext)

	// ExitSchemaBody is called when exiting the schemaBody production.
	ExitSchemaBody(c *SchemaBodyContext)

	// ExitSchemaTypes is called when exiting the schemaTypes production.
	ExitSchemaTypes(c *SchemaTypesContext)

	// ExitTypeDecl is called when exiting the typeDecl production.
	ExitTypeDecl(c *TypeDeclContext)

	// ExitTypeName is called when exiting the typeName production.
	ExitTypeName(c *TypeNameContext)

	// ExitTypeBody is called when exiting the typeBody production.
	ExitTypeBody(c *TypeBodyContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitFieldName is called when exiting the fieldName production.
	ExitFieldName(c *FieldNameContext)

	// ExitFieldType is called when exiting the fieldType production.
	ExitFieldType(c *FieldTypeContext)
}
