// Code generated from FOFA.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // FOFA

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by FOFAParser.
type FOFAVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by FOFAParser#start.
	VisitStart(ctx *StartContext) interface{}

	// Visit a parse tree produced by FOFAParser#compareExp.
	VisitCompareExp(ctx *CompareExpContext) interface{}

	// Visit a parse tree produced by FOFAParser#andLogicalExp.
	VisitAndLogicalExp(ctx *AndLogicalExpContext) interface{}

	// Visit a parse tree produced by FOFAParser#bracketExp.
	VisitBracketExp(ctx *BracketExpContext) interface{}

	// Visit a parse tree produced by FOFAParser#scompareExp.
	VisitScompareExp(ctx *ScompareExpContext) interface{}

	// Visit a parse tree produced by FOFAParser#orLogicalExp.
	VisitOrLogicalExp(ctx *OrLogicalExpContext) interface{}

	// Visit a parse tree produced by FOFAParser#attrPath.
	VisitAttrPath(ctx *AttrPathContext) interface{}

	// Visit a parse tree produced by FOFAParser#boolean.
	VisitBoolean(ctx *BooleanContext) interface{}

	// Visit a parse tree produced by FOFAParser#null.
	VisitNull(ctx *NullContext) interface{}

	// Visit a parse tree produced by FOFAParser#string.
	VisitString(ctx *StringContext) interface{}

	// Visit a parse tree produced by FOFAParser#double.
	VisitDouble(ctx *DoubleContext) interface{}

	// Visit a parse tree produced by FOFAParser#long.
	VisitLong(ctx *LongContext) interface{}
}
