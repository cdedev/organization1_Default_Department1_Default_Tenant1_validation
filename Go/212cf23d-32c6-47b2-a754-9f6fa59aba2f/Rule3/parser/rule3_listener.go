// Code generated from /usr/local/lib/Go/212cf23d-32c6-47b2-a754-9f6fa59aba2f/1709694127334/Rule3.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Rule3

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Rule3Listener is a complete listener for a parse tree produced by Rule3Parser.
type Rule3Listener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterEducation is called when entering the education production.
	EnterEducation(c *EducationContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitEducation is called when exiting the education production.
	ExitEducation(c *EducationContext)
}
