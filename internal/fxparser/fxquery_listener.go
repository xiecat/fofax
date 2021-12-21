package fxparser

import (
	"fmt"
	"fofax/internal/fx"
	"fofax/internal/fxparser/parser"
	"fofax/internal/fxparser/stack"
	"fofax/internal/printer"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"strings"
)

type FxQueryListener struct {
	*parser.BaseFOFAListener
	Stack *stack.Stack
}

func NewFxQueryListener() *FxQueryListener {
	return &FxQueryListener{
		Stack: stack.New(),
	}
}

func (ql *FxQueryListener) ExitCompareExp(c *parser.CompareExpContext) {
	if c.GetPropertyName().GetText() == "fx" {
		fxvalue := strings.TrimSpace(c.GetPropertyValue().GetText())
		fxvalue = fxvalue[1 : len(fxvalue)-1]
		fxinfo, err := fx.Info.SearchSingle(fxvalue)
		if err != nil {
			printer.Fatalf("%s Err:%s", c.GetPropertyValue().GetText(), err.Error())
		}
		ql.Stack.Push(fxinfo.QueryString())
	} else {
		ql.Stack.Push(c.GetText())
	}
}

// ExitScompareExp is called when production scompareExp is exited.
func (ql *FxQueryListener) ExitScompareExp(c *parser.ScompareExpContext) {
	ql.Stack.Push(c.GetText())
}

// ExitBracketExp is called when production bracketExp is exited.
func (ql *FxQueryListener) ExitBracketExp(c *parser.BracketExpContext) {
	ql.Stack.Push(fmt.Sprintf("(%s)", ql.Stack.Pop()))
}

func (ql *FxQueryListener) ExitAndLogicalExp(c *parser.AndLogicalExpContext) {
	right := ql.Stack.Pop()
	left := ql.Stack.Pop()
	ql.Stack.Push(fmt.Sprintf("%s&&%s", left, right))
}

func (ql *FxQueryListener) ExitOrLogicalExp(c *parser.OrLogicalExpContext) {
	right := ql.Stack.Pop()
	left := ql.Stack.Pop()
	ql.Stack.Push(fmt.Sprintf("%s||%s", left, right))
}

func PrintParserTree(query string) {
	stream := antlr.NewInputStream(query)
	lexer := parser.NewFOFALexer(stream)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	fofaxParser := parser.NewFOFAParser(tokenStream)
	tree := fofaxParser.Start()
	fmt.Println(tree.GetText())
	fmt.Println(tree.ToStringTree([]string{""}, fofaxParser))

}

func Query(query string) string {
	stream := antlr.NewInputStream(query)
	lexer := parser.NewFOFALexer(stream)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	dslParser := parser.NewFOFAParser(tokenStream)
	tree := dslParser.Start()
	listener := NewFxQueryListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	return listener.Stack.Pop().(string)
}
