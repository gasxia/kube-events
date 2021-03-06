// Code generated from EventRule.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // EventRule

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// A complete Visitor for a parse tree produced by EventRuleParser.
type EventRuleVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by EventRuleParser#start.
	VisitStart(ctx *StartContext) interface{}

	// Visit a parse tree produced by EventRuleParser#Not.
	VisitNot(ctx *NotContext) interface{}

	// Visit a parse tree produced by EventRuleParser#Parenthesis.
	VisitParenthesis(ctx *ParenthesisContext) interface{}

	// Visit a parse tree produced by EventRuleParser#Variable.
	VisitVariable(ctx *VariableContext) interface{}

	// Visit a parse tree produced by EventRuleParser#CompareNumber.
	VisitCompareNumber(ctx *CompareNumberContext) interface{}

	// Visit a parse tree produced by EventRuleParser#StringEqualContains.
	VisitStringEqualContains(ctx *StringEqualContainsContext) interface{}

	// Visit a parse tree produced by EventRuleParser#AndOr.
	VisitAndOr(ctx *AndOrContext) interface{}

	// Visit a parse tree produced by EventRuleParser#StringIn.
	VisitStringIn(ctx *StringInContext) interface{}
}
