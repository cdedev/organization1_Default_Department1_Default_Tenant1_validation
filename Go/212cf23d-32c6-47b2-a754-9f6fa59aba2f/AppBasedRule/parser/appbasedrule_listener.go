// Code generated from /usr/local/lib/Go/212cf23d-32c6-47b2-a754-9f6fa59aba2f/1709267325888/AppBasedRule.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // AppBasedRule

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AppBasedRuleListener is a complete listener for a parse tree produced by AppBasedRuleParser.
type AppBasedRuleListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterGlucose is called when entering the glucose production.
	EnterGlucose(c *GlucoseContext)

	// EnterGlucose_acosh is called when entering the glucose_acosh production.
	EnterGlucose_acosh(c *Glucose_acoshContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitGlucose is called when exiting the glucose production.
	ExitGlucose(c *GlucoseContext)

	// ExitGlucose_acosh is called when exiting the glucose_acosh production.
	ExitGlucose_acosh(c *Glucose_acoshContext)
}
