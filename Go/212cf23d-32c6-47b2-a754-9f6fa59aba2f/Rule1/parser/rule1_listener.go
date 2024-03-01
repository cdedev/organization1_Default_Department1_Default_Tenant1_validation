// Code generated from /usr/local/lib/Go/212cf23d-32c6-47b2-a754-9f6fa59aba2f/1709266634448/Rule1.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Rule1

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Rule1Listener is a complete listener for a parse tree produced by Rule1Parser.
type Rule1Listener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterGlucose is called when entering the glucose production.
	EnterGlucose(c *GlucoseContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitGlucose is called when exiting the glucose production.
	ExitGlucose(c *GlucoseContext)
}
