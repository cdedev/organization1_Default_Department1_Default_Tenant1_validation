// Code generated from /usr/local/lib/Go/212cf23d-32c6-47b2-a754-9f6fa59aba2f/1710560357049/Rule4.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Rule4

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRule4Listener is a complete listener for a parse tree produced by Rule4Parser.
type BaseRule4Listener struct{}

var _ Rule4Listener = &BaseRule4Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRule4Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRule4Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRule4Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRule4Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseRule4Listener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseRule4Listener) ExitExpression(ctx *ExpressionContext) {}

// EnterField1 is called when production field1 is entered.
func (s *BaseRule4Listener) EnterField1(ctx *Field1Context) {}

// ExitField1 is called when production field1 is exited.
func (s *BaseRule4Listener) ExitField1(ctx *Field1Context) {}

// EnterField2 is called when production field2 is entered.
func (s *BaseRule4Listener) EnterField2(ctx *Field2Context) {}

// ExitField2 is called when production field2 is exited.
func (s *BaseRule4Listener) ExitField2(ctx *Field2Context) {}
