// Code generated from /usr/local/lib/Go/212cf23d-32c6-47b2-a754-9f6fa59aba2f/1709693485583/Rule2.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Rule2

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRule2Listener is a complete listener for a parse tree produced by Rule2Parser.
type BaseRule2Listener struct{}

var _ Rule2Listener = &BaseRule2Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRule2Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRule2Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRule2Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRule2Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseRule2Listener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseRule2Listener) ExitExpression(ctx *ExpressionContext) {}

// EnterEducation is called when production education is entered.
func (s *BaseRule2Listener) EnterEducation(ctx *EducationContext) {}

// ExitEducation is called when production education is exited.
func (s *BaseRule2Listener) ExitEducation(ctx *EducationContext) {}

// EnterFor is called when production for is entered.
func (s *BaseRule2Listener) EnterFor(ctx *ForContext) {}

// ExitFor is called when production for is exited.
func (s *BaseRule2Listener) ExitFor(ctx *ForContext) {}
