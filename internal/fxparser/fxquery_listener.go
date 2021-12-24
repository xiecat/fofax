package fxparser

import (
	"errors"
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
		fxvalue := strings.Trim(c.GetPropertyValue().GetText(), `\"`)
		printer.Successf("fx query id:%s", fxvalue)
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

// ExitNoCompareExp is called when production scompareExp is exited.
func (ql *FxQueryListener) ExitNoCompareExp(c *parser.NoCompareExpContext) {
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

func (ql *FxQueryListener) ExitSgExp(c *parser.SgExpContext) {
	ql.Stack.Push(c.GetText())
}

func (ql *FxQueryListener) ExitOrLogicalExp(c *parser.OrLogicalExpContext) {
	right := ql.Stack.Pop()
	left := ql.Stack.Pop()
	ql.Stack.Push(fmt.Sprintf("%s||%s", left, right))
}

func parseTree(query string) (parser.IStartContext, error) {
	errListener := NewFxErrorListener()
	stream := antlr.NewInputStream(query)
	lexer := parser.NewFOFALexer(stream)
	lexer.AddErrorListener(errListener)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	fofaxParser := parser.NewFOFAParser(tokenStream)
	fofaxParser.AddErrorListener(errListener)
	tree := fofaxParser.Start()
	if errListener.errors > 0 {
		return tree, errors.New("found syntax errors in input")
	}
	return tree, nil
}

func PrintParserTree(query string) {
	errListener := NewFxErrorListener()
	stream := antlr.NewInputStream(query)
	lexer := parser.NewFOFALexer(stream)
	lexer.AddErrorListener(errListener)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	fofaxParser := parser.NewFOFAParser(tokenStream)
	fofaxParser.AddErrorListener(errListener)
	tree := fofaxParser.Start()
	if errListener.errors > 0 {
		printer.Fatal("found syntax errors in input")
	}
	fmt.Println("Source: ", query)
	fmt.Println("Parse: ")
	printer.Infof("Source: %s", query)
	printer.Infof("Parse: %s", tree.GetText())
	printer.Infof("Tree: %s", tree.ToStringTree([]string{""}, fofaxParser))
}

func Query(query string) string {
	tree, err := parseTree(query)
	if err != nil {
		printer.Fatal(err.Error())
	}
	listener := NewFxQueryListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	return listener.Stack.Pop().(string)
}
