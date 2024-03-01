// Code generated from /usr/local/lib/Go/212cf23d-32c6-47b2-a754-9f6fa59aba2f/1709267325888/AppBasedRule.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // AppBasedRule

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAppBasedRuleListener is a complete listener for a parse tree produced by AppBasedRuleParser.
type BaseAppBasedRuleListener struct{}

var _ AppBasedRuleListener = &BaseAppBasedRuleListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAppBasedRuleListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAppBasedRuleListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAppBasedRuleListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAppBasedRuleListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseAppBasedRuleListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseAppBasedRuleListener) ExitExpression(ctx *ExpressionContext) {}

// EnterGlucose is called when production glucose is entered.
func (s *BaseAppBasedRuleListener) EnterGlucose(ctx *GlucoseContext) {}

// ExitGlucose is called when production glucose is exited.
func (s *BaseAppBasedRuleListener) ExitGlucose(ctx *GlucoseContext) {}

// EnterGlucose_acosh is called when production glucose_acosh is entered.
func (s *BaseAppBasedRuleListener) EnterGlucose_acosh(ctx *Glucose_acoshContext) {}

// ExitGlucose_acosh is called when production glucose_acosh is exited.
func (s *BaseAppBasedRuleListener) ExitGlucose_acosh(ctx *Glucose_acoshContext) {}
