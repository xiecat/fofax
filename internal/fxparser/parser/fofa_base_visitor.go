// Code generated from FOFA.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // FOFA

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseFOFAVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseFOFAVisitor) VisitStart(ctx *StartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFOFAVisitor) VisitCompareExp(ctx *CompareExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFOFAVisitor) VisitAndLogicalExp(ctx *AndLogicalExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFOFAVisitor) VisitBracketExp(ctx *BracketExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFOFAVisitor) VisitScompareExp(ctx *ScompareExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFOFAVisitor) VisitOrLogicalExp(ctx *OrLogicalExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFOFAVisitor) VisitAttrPath(ctx *AttrPathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFOFAVisitor) VisitBoolean(ctx *BooleanContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFOFAVisitor) VisitNull(ctx *NullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFOFAVisitor) VisitString(ctx *StringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFOFAVisitor) VisitDouble(ctx *DoubleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFOFAVisitor) VisitLong(ctx *LongContext) interface{} {
	return v.VisitChildren(ctx)
}
