// Code generated from FOFA.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // FOFA

import "github.com/antlr/antlr4/runtime/Go/antlr"

// FOFAListener is a complete listener for a parse tree produced by FOFAParser.
type FOFAListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterCompareExp is called when entering the compareExp production.
	EnterCompareExp(c *CompareExpContext)

	// EnterAndLogicalExp is called when entering the andLogicalExp production.
	EnterAndLogicalExp(c *AndLogicalExpContext)

	// EnterBracketExp is called when entering the bracketExp production.
	EnterBracketExp(c *BracketExpContext)

	// EnterScompareExp is called when entering the scompareExp production.
	EnterScompareExp(c *ScompareExpContext)

	// EnterOrLogicalExp is called when entering the orLogicalExp production.
	EnterOrLogicalExp(c *OrLogicalExpContext)

	// EnterAttrPath is called when entering the attrPath production.
	EnterAttrPath(c *AttrPathContext)

	// EnterBoolean is called when entering the boolean production.
	EnterBoolean(c *BooleanContext)

	// EnterNull is called when entering the null production.
	EnterNull(c *NullContext)

	// EnterString is called when entering the string production.
	EnterString(c *StringContext)

	// EnterDouble is called when entering the double production.
	EnterDouble(c *DoubleContext)

	// EnterLong is called when entering the long production.
	EnterLong(c *LongContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitCompareExp is called when exiting the compareExp production.
	ExitCompareExp(c *CompareExpContext)

	// ExitAndLogicalExp is called when exiting the andLogicalExp production.
	ExitAndLogicalExp(c *AndLogicalExpContext)

	// ExitBracketExp is called when exiting the bracketExp production.
	ExitBracketExp(c *BracketExpContext)

	// ExitScompareExp is called when exiting the scompareExp production.
	ExitScompareExp(c *ScompareExpContext)

	// ExitOrLogicalExp is called when exiting the orLogicalExp production.
	ExitOrLogicalExp(c *OrLogicalExpContext)

	// ExitAttrPath is called when exiting the attrPath production.
	ExitAttrPath(c *AttrPathContext)

	// ExitBoolean is called when exiting the boolean production.
	ExitBoolean(c *BooleanContext)

	// ExitNull is called when exiting the null production.
	ExitNull(c *NullContext)

	// ExitString is called when exiting the string production.
	ExitString(c *StringContext)

	// ExitDouble is called when exiting the double production.
	ExitDouble(c *DoubleContext)

	// ExitLong is called when exiting the long production.
	ExitLong(c *LongContext)
}
