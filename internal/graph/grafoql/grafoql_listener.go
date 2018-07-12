// Code generated from internal/graph/grafoql/GrafoQL.g4 by ANTLR 4.7.1. DO NOT EDIT.

package grafoql // GrafoQL
import "github.com/antlr/antlr4/runtime/Go/antlr"

// GrafoQLListener is a complete listener for a parse tree produced by GrafoQLParser.
type GrafoQLListener interface {
	antlr.ParseTreeListener

	// EnterDocument is called when entering the document production.
	EnterDocument(c *DocumentContext)

	// EnterQuery is called when entering the query production.
	EnterQuery(c *QueryContext)

	// EnterSchemaBody is called when entering the schemaBody production.
	EnterSchemaBody(c *SchemaBodyContext)

	// EnterMutation is called when entering the mutation production.
	EnterMutation(c *MutationContext)

	// EnterMutationBody is called when entering the mutationBody production.
	EnterMutationBody(c *MutationBodyContext)

	// ExitDocument is called when exiting the document production.
	ExitDocument(c *DocumentContext)

	// ExitQuery is called when exiting the query production.
	ExitQuery(c *QueryContext)

	// ExitSchemaBody is called when exiting the schemaBody production.
	ExitSchemaBody(c *SchemaBodyContext)

	// ExitMutation is called when exiting the mutation production.
	ExitMutation(c *MutationContext)

	// ExitMutationBody is called when exiting the mutationBody production.
	ExitMutationBody(c *MutationBodyContext)
}
