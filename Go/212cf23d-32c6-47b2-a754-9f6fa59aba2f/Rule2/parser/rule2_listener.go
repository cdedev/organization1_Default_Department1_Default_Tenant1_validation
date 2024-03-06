// Code generated from /usr/local/lib/Go/212cf23d-32c6-47b2-a754-9f6fa59aba2f/1709693485583/Rule2.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Rule2

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Rule2Listener is a complete listener for a parse tree produced by Rule2Parser.
type Rule2Listener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterEducation is called when entering the education production.
	EnterEducation(c *EducationContext)

	// EnterFor is called when entering the for production.
	EnterFor(c *ForContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitEducation is called when exiting the education production.
	ExitEducation(c *EducationContext)

	// ExitFor is called when exiting the for production.
	ExitFor(c *ForContext)
}
