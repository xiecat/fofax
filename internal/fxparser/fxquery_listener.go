package fxparser

import (
	"errors"
	"fmt"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/xiecat/fofax/internal/fx"
	"github.com/xiecat/fofax/internal/fxparser/parser"
	"github.com/xiecat/fofax/internal/fxparser/stack"
	"github.com/xiecat/fofax/internal/printer"
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
		printer.Successf("fx query: %s", fxvalue)
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

// ExitCcompareExp is called when production ccompareExp is exited.
func (ql *FxQueryListener) ExitCcompareExp(c *parser.CcompareExpContext) {
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

// ParseIssue describes one problem found while validating a query.
type ParseIssue struct {
	Line    int    `json:"line"`
	Column  int    `json:"column"`
	Message string `json:"message"`
}

// StrictErrorListener records every syntax error reported by ANTLR.
type StrictErrorListener struct {
	*antlr.DefaultErrorListener
	Issues []ParseIssue
}

// NewStrictErrorListener returns a listener that accumulates parse issues.
func NewStrictErrorListener() *StrictErrorListener {
	return &StrictErrorListener{}
}

// SyntaxError records one syntax error.
func (l *StrictErrorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {
	l.Issues = append(l.Issues, ParseIssue{Line: line, Column: column, Message: msg})
}

// ValidateQuery parses a query strictly: it reports syntax errors and also
// rejects input that the grammar only partially consumes. The bundled start
// rule has no EOF anchor, so trailing garbage would otherwise be ignored.
func ValidateQuery(query string) (string, []ParseIssue) {
	errListener := NewStrictErrorListener()
	stream := antlr.NewInputStream(query)
	lexer := parser.NewFOFALexer(stream)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errListener)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	fofaxParser := parser.NewFOFAParser(tokenStream)
	fofaxParser.RemoveErrorListeners()
	fofaxParser.AddErrorListener(errListener)
	tree := fofaxParser.Start()

	if remaining := tokenStream.LA(1); remaining != antlr.TokenEOF && len(errListener.Issues) == 0 {
		offending := tokenStream.LT(1)
		line, column := 0, 0
		if offending != nil {
			line = offending.GetLine()
			column = offending.GetColumn()
		}
		text := "<unknown>"
		if offending != nil {
			text = offending.GetText()
		}
		errListener.Issues = append(errListener.Issues, ParseIssue{
			Line:    line,
			Column:  column,
			Message: fmt.Sprintf("unexpected input near %q: the query is only partially parsed", text),
		})
	}

	if len(errListener.Issues) > 0 {
		return "", errListener.Issues
	}

	listener := NewFxQueryListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	normalized := listener.Stack.Pop()
	if normalized == nil {
		return "", []ParseIssue{{Message: "query produced no expression"}}
	}
	return normalized.(string), nil
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
	printer.Infof("Source: %s", query)
	printer.Infof("Parse: %s", tree.GetText())
	printer.Infof("Tree: %s", tree.ToStringTree([]string{""}, fofaxParser))
}

func Query(query string) string {
	normalized, err := NormalizeQuery(query)
	if err != nil {
		printer.Fatal(err.Error())
	}
	return normalized
}

// NormalizeQuery parses and rewrites a query without terminating the process,
// so callers such as the MCP server can report the error to a client.
func NormalizeQuery(query string) (string, error) {
	tree, err := parseTree(query)
	if err != nil {
		return "", err
	}
	listener := NewFxQueryListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	result := listener.Stack.Pop()
	if result == nil {
		return "", errors.New("query produced no expression")
	}
	return result.(string), nil
}
