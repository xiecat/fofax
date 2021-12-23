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

	// EnterNoCompareExp is called when entering the noCompareExp production.
	EnterNoCompareExp(c *NoCompareExpContext)

	// EnterAndLogicalExp is called when entering the andLogicalExp production.
	EnterAndLogicalExp(c *AndLogicalExpContext)

	// EnterBracketExp is called when entering the bracketExp production.
	EnterBracketExp(c *BracketExpContext)

	// EnterScompareExp is called when entering the scompareExp production.
	EnterScompareExp(c *ScompareExpContext)

	// EnterOrLogicalExp is called when entering the orLogicalExp production.
	EnterOrLogicalExp(c *OrLogicalExpContext)

	// EnterFofaKeyword is called when entering the fofaKeyword production.
	EnterFofaKeyword(c *FofaKeywordContext)

	// EnterBoolean is called when entering the boolean production.
	EnterBoolean(c *BooleanContext)

	// EnterString is called when entering the string production.
	EnterString(c *StringContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitCompareExp is called when exiting the compareExp production.
	ExitCompareExp(c *CompareExpContext)

	// ExitNoCompareExp is called when exiting the noCompareExp production.
	ExitNoCompareExp(c *NoCompareExpContext)

	// ExitAndLogicalExp is called when exiting the andLogicalExp production.
	ExitAndLogicalExp(c *AndLogicalExpContext)

	// ExitBracketExp is called when exiting the bracketExp production.
	ExitBracketExp(c *BracketExpContext)

	// ExitScompareExp is called when exiting the scompareExp production.
	ExitScompareExp(c *ScompareExpContext)

	// ExitOrLogicalExp is called when exiting the orLogicalExp production.
	ExitOrLogicalExp(c *OrLogicalExpContext)

	// ExitFofaKeyword is called when exiting the fofaKeyword production.
	ExitFofaKeyword(c *FofaKeywordContext)

	// ExitBoolean is called when exiting the boolean production.
	ExitBoolean(c *BooleanContext)

	// ExitString is called when exiting the string production.
	ExitString(c *StringContext)
}
