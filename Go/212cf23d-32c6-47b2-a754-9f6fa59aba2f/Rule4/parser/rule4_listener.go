// Code generated from /usr/local/lib/Go/212cf23d-32c6-47b2-a754-9f6fa59aba2f/1710560357049/Rule4.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Rule4

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Rule4Listener is a complete listener for a parse tree produced by Rule4Parser.
type Rule4Listener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterField1 is called when entering the field1 production.
	EnterField1(c *Field1Context)

	// EnterField2 is called when entering the field2 production.
	EnterField2(c *Field2Context)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitField1 is called when exiting the field1 production.
	ExitField1(c *Field1Context)

	// ExitField2 is called when exiting the field2 production.
	ExitField2(c *Field2Context)
}
