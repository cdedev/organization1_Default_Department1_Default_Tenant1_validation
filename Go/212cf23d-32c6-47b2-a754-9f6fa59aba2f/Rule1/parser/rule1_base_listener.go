// Code generated from /usr/local/lib/Go/212cf23d-32c6-47b2-a754-9f6fa59aba2f/1709266634448/Rule1.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Rule1

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRule1Listener is a complete listener for a parse tree produced by Rule1Parser.
type BaseRule1Listener struct{}

var _ Rule1Listener = &BaseRule1Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRule1Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRule1Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRule1Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRule1Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseRule1Listener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseRule1Listener) ExitExpression(ctx *ExpressionContext) {}

// EnterGlucose is called when production glucose is entered.
func (s *BaseRule1Listener) EnterGlucose(ctx *GlucoseContext) {}

// ExitGlucose is called when production glucose is exited.
func (s *BaseRule1Listener) ExitGlucose(ctx *GlucoseContext) {}
