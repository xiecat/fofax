// Code generated from FOFA.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // FOFA

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseFOFAListener is a complete listener for a parse tree produced by FOFAParser.
type BaseFOFAListener struct{}

var _ FOFAListener = &BaseFOFAListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFOFAListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFOFAListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFOFAListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFOFAListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseFOFAListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseFOFAListener) ExitStart(ctx *StartContext) {}

// EnterCompareExp is called when production compareExp is entered.
func (s *BaseFOFAListener) EnterCompareExp(ctx *CompareExpContext) {}

// ExitCompareExp is called when production compareExp is exited.
func (s *BaseFOFAListener) ExitCompareExp(ctx *CompareExpContext) {}

// EnterAndLogicalExp is called when production andLogicalExp is entered.
func (s *BaseFOFAListener) EnterAndLogicalExp(ctx *AndLogicalExpContext) {}

// ExitAndLogicalExp is called when production andLogicalExp is exited.
func (s *BaseFOFAListener) ExitAndLogicalExp(ctx *AndLogicalExpContext) {}

// EnterBracketExp is called when production bracketExp is entered.
func (s *BaseFOFAListener) EnterBracketExp(ctx *BracketExpContext) {}

// ExitBracketExp is called when production bracketExp is exited.
func (s *BaseFOFAListener) ExitBracketExp(ctx *BracketExpContext) {}

// EnterScompareExp is called when production scompareExp is entered.
func (s *BaseFOFAListener) EnterScompareExp(ctx *ScompareExpContext) {}

// ExitScompareExp is called when production scompareExp is exited.
func (s *BaseFOFAListener) ExitScompareExp(ctx *ScompareExpContext) {}

// EnterOrLogicalExp is called when production orLogicalExp is entered.
func (s *BaseFOFAListener) EnterOrLogicalExp(ctx *OrLogicalExpContext) {}

// ExitOrLogicalExp is called when production orLogicalExp is exited.
func (s *BaseFOFAListener) ExitOrLogicalExp(ctx *OrLogicalExpContext) {}

// EnterAttrPath is called when production attrPath is entered.
func (s *BaseFOFAListener) EnterAttrPath(ctx *AttrPathContext) {}

// ExitAttrPath is called when production attrPath is exited.
func (s *BaseFOFAListener) ExitAttrPath(ctx *AttrPathContext) {}

// EnterBoolean is called when production boolean is entered.
func (s *BaseFOFAListener) EnterBoolean(ctx *BooleanContext) {}

// ExitBoolean is called when production boolean is exited.
func (s *BaseFOFAListener) ExitBoolean(ctx *BooleanContext) {}

// EnterNull is called when production null is entered.
func (s *BaseFOFAListener) EnterNull(ctx *NullContext) {}

// ExitNull is called when production null is exited.
func (s *BaseFOFAListener) ExitNull(ctx *NullContext) {}

// EnterString is called when production string is entered.
func (s *BaseFOFAListener) EnterString(ctx *StringContext) {}

// ExitString is called when production string is exited.
func (s *BaseFOFAListener) ExitString(ctx *StringContext) {}

// EnterDouble is called when production double is entered.
func (s *BaseFOFAListener) EnterDouble(ctx *DoubleContext) {}

// ExitDouble is called when production double is exited.
func (s *BaseFOFAListener) ExitDouble(ctx *DoubleContext) {}

// EnterLong is called when production long is entered.
func (s *BaseFOFAListener) EnterLong(ctx *LongContext) {}

// ExitLong is called when production long is exited.
func (s *BaseFOFAListener) ExitLong(ctx *LongContext) {}
