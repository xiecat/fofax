// Code generated from FOFA.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // FOFA

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 15, 57, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 3, 2, 3, 2, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 35, 10, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 43, 10, 3, 12, 3, 14, 3, 46, 11, 3, 3,
	4, 3, 4, 3, 4, 5, 4, 51, 10, 4, 3, 5, 3, 5, 5, 5, 55, 10, 5, 3, 5, 2, 3,
	4, 6, 2, 4, 6, 8, 2, 2, 2, 61, 2, 10, 3, 2, 2, 2, 4, 34, 3, 2, 2, 2, 6,
	47, 3, 2, 2, 2, 8, 54, 3, 2, 2, 2, 10, 11, 5, 4, 3, 2, 11, 3, 3, 2, 2,
	2, 12, 13, 8, 3, 1, 2, 13, 14, 7, 11, 2, 2, 14, 15, 5, 4, 3, 2, 15, 16,
	7, 12, 2, 2, 16, 35, 3, 2, 2, 2, 17, 18, 5, 6, 4, 2, 18, 19, 7, 8, 2, 2,
	19, 20, 5, 8, 5, 2, 20, 35, 3, 2, 2, 2, 21, 22, 5, 6, 4, 2, 22, 23, 7,
	9, 2, 2, 23, 24, 5, 8, 5, 2, 24, 35, 3, 2, 2, 2, 25, 26, 5, 6, 4, 2, 26,
	27, 7, 7, 2, 2, 27, 28, 5, 8, 5, 2, 28, 35, 3, 2, 2, 2, 29, 30, 5, 6, 4,
	2, 30, 31, 7, 10, 2, 2, 31, 32, 5, 8, 5, 2, 32, 35, 3, 2, 2, 2, 33, 35,
	5, 8, 5, 2, 34, 12, 3, 2, 2, 2, 34, 17, 3, 2, 2, 2, 34, 21, 3, 2, 2, 2,
	34, 25, 3, 2, 2, 2, 34, 29, 3, 2, 2, 2, 34, 33, 3, 2, 2, 2, 35, 44, 3,
	2, 2, 2, 36, 37, 12, 9, 2, 2, 37, 38, 7, 5, 2, 2, 38, 43, 5, 4, 3, 10,
	39, 40, 12, 8, 2, 2, 40, 41, 7, 6, 2, 2, 41, 43, 5, 4, 3, 9, 42, 36, 3,
	2, 2, 2, 42, 39, 3, 2, 2, 2, 43, 46, 3, 2, 2, 2, 44, 42, 3, 2, 2, 2, 44,
	45, 3, 2, 2, 2, 45, 5, 3, 2, 2, 2, 46, 44, 3, 2, 2, 2, 47, 50, 7, 13, 2,
	2, 48, 49, 7, 3, 2, 2, 49, 51, 7, 13, 2, 2, 50, 48, 3, 2, 2, 2, 50, 51,
	3, 2, 2, 2, 51, 7, 3, 2, 2, 2, 52, 55, 7, 4, 2, 2, 53, 55, 7, 14, 2, 2,
	54, 52, 3, 2, 2, 2, 54, 53, 3, 2, 2, 2, 55, 9, 3, 2, 2, 2, 7, 34, 42, 44,
	50, 54,
}
var literalNames = []string{
	"", "'.'", "", "'&&'", "'||'", "'!='", "'='", "'=='", "'=~'", "'('", "')'",
}
var symbolicNames = []string{
	"", "", "BOOLEAN", "AND", "OR", "NOT", "EQ", "SEQ", "CEQ", "BR_OPEN", "BR_CLOSE",
	"FOFA_KEY", "STRING", "WS",
}

var ruleNames = []string{
	"start", "query", "fofaKeyword", "fofaValue",
}

type FOFAParser struct {
	*antlr.BaseParser
}

// NewFOFAParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *FOFAParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewFOFAParser(input antlr.TokenStream) *FOFAParser {
	this := new(FOFAParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "FOFA.g4"

	return this
}

// FOFAParser tokens.
const (
	FOFAParserEOF      = antlr.TokenEOF
	FOFAParserT__0     = 1
	FOFAParserBOOLEAN  = 2
	FOFAParserAND      = 3
	FOFAParserOR       = 4
	FOFAParserNOT      = 5
	FOFAParserEQ       = 6
	FOFAParserSEQ      = 7
	FOFAParserCEQ      = 8
	FOFAParserBR_OPEN  = 9
	FOFAParserBR_CLOSE = 10
	FOFAParserFOFA_KEY = 11
	FOFAParserSTRING   = 12
	FOFAParserWS       = 13
)

// FOFAParser rules.
const (
	FOFAParserRULE_start       = 0
	FOFAParserRULE_query       = 1
	FOFAParserRULE_fofaKeyword = 2
	FOFAParserRULE_fofaValue   = 3
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FOFAParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FOFAParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Query() IQueryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQueryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQueryContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FOFAListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FOFAListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *FOFAParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FOFAParserRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(8)
		p.query(0)
	}

	return localctx
}

// IQueryContext is an interface to support dynamic dispatch.
type IQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQueryContext differentiates from other interfaces.
	IsQueryContext()
}

type QueryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryContext() *QueryContext {
	var p = new(QueryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FOFAParserRULE_query
	return p
}

func (*QueryContext) IsQueryContext() {}

func NewQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryContext {
	var p = new(QueryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FOFAParserRULE_query

	return p
}

func (s *QueryContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryContext) CopyFrom(ctx *QueryContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *QueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type CompareExpContext struct {
	*QueryContext
	propertyName  IFofaKeywordContext
	op            antlr.Token
	propertyValue IFofaValueContext
}

func NewCompareExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompareExpContext {
	var p = new(CompareExpContext)

	p.QueryContext = NewEmptyQueryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*QueryContext))

	return p
}

func (s *CompareExpContext) GetOp() antlr.Token { return s.op }

func (s *CompareExpContext) SetOp(v antlr.Token) { s.op = v }

func (s *CompareExpContext) GetPropertyName() IFofaKeywordContext { return s.propertyName }

func (s *CompareExpContext) GetPropertyValue() IFofaValueContext { return s.propertyValue }

func (s *CompareExpContext) SetPropertyName(v IFofaKeywordContext) { s.propertyName = v }

func (s *CompareExpContext) SetPropertyValue(v IFofaValueContext) { s.propertyValue = v }

func (s *CompareExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareExpContext) FofaKeyword() IFofaKeywordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFofaKeywordContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFofaKeywordContext)
}

func (s *CompareExpContext) EQ() antlr.TerminalNode {
	return s.GetToken(FOFAParserEQ, 0)
}

func (s *CompareExpContext) FofaValue() IFofaValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFofaValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFofaValueContext)
}

func (s *CompareExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FOFAListener); ok {
		listenerT.EnterCompareExp(s)
	}
}

func (s *CompareExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FOFAListener); ok {
		listenerT.ExitCompareExp(s)
	}
}

type NoCompareExpContext struct {
	*QueryContext
	propertyName  IFofaKeywordContext
	op            antlr.Token
	propertyValue IFofaValueContext
}

func NewNoCompareExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NoCompareExpContext {
	var p = new(NoCompareExpContext)

	p.QueryContext = NewEmptyQueryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*QueryContext))

	return p
}

func (s *NoCompareExpContext) GetOp() antlr.Token { return s.op }

func (s *NoCompareExpContext) SetOp(v antlr.Token) { s.op = v }

func (s *NoCompareExpContext) GetPropertyName() IFofaKeywordContext { return s.propertyName }

func (s *NoCompareExpContext) GetPropertyValue() IFofaValueContext { return s.propertyValue }

func (s *NoCompareExpContext) SetPropertyName(v IFofaKeywordContext) { s.propertyName = v }

func (s *NoCompareExpContext) SetPropertyValue(v IFofaValueContext) { s.propertyValue = v }

func (s *NoCompareExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NoCompareExpContext) FofaKeyword() IFofaKeywordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFofaKeywordContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFofaKeywordContext)
}

func (s *NoCompareExpContext) NOT() antlr.TerminalNode {
	return s.GetToken(FOFAParserNOT, 0)
}

func (s *NoCompareExpContext) FofaValue() IFofaValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFofaValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFofaValueContext)
}

func (s *NoCompareExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FOFAListener); ok {
		listenerT.EnterNoCompareExp(s)
	}
}

func (s *NoCompareExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FOFAListener); ok {
		listenerT.ExitNoCompareExp(s)
	}
}

type AndLogicalExpContext struct {
	*QueryContext
	leftQuery  IQueryContext
	op         antlr.Token
	rightQuery IQueryContext
}

func NewAndLogicalExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndLogicalExpContext {
	var p = new(AndLogicalExpContext)

	p.QueryContext = NewEmptyQueryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*QueryContext))

	return p
}

func (s *AndLogicalExpContext) GetOp() antlr.Token { return s.op }

func (s *AndLogicalExpContext) SetOp(v antlr.Token) { s.op = v }

func (s *AndLogicalExpContext) GetLeftQuery() IQueryContext { return s.leftQuery }

func (s *AndLogicalExpContext) GetRightQuery() IQueryContext { return s.rightQuery }

func (s *AndLogicalExpContext) SetLeftQuery(v IQueryContext) { s.leftQuery = v }

func (s *AndLogicalExpContext) SetRightQuery(v IQueryContext) { s.rightQuery = v }

func (s *AndLogicalExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndLogicalExpContext) AllQuery() []IQueryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IQueryContext)(nil)).Elem())
	var tst = make([]IQueryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IQueryContext)
		}
	}

	return tst
}

func (s *AndLogicalExpContext) Query(i int) IQueryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQueryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IQueryContext)
}

func (s *AndLogicalExpContext) AND() antlr.TerminalNode {
	return s.GetToken(FOFAParserAND, 0)
}

func (s *AndLogicalExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FOFAListener); ok {
		listenerT.EnterAndLogicalExp(s)
	}
}

func (s *AndLogicalExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FOFAListener); ok {
		listenerT.ExitAndLogicalExp(s)
	}
}

type CcompareExpContext struct {
	*QueryContext
	propertyName  IFofaKeywordContext
	op            antlr.Token
	propertyValue IFofaValueContext
}

func NewCcompareExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CcompareExpContext {
	var p = new(CcompareExpContext)

	p.QueryContext = NewEmptyQueryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*QueryContext))

	return p
}

func (s *CcompareExpContext) GetOp() antlr.Token { return s.op }

func (s *CcompareExpContext) SetOp(v antlr.Token) { s.op = v }

func (s *CcompareExpContext) GetPropertyName() IFofaKeywordContext { return s.propertyName }

func (s *CcompareExpContext) GetPropertyValue() IFofaValueContext { return s.propertyValue }

func (s *CcompareExpContext) SetPropertyName(v IFofaKeywordContext) { s.propertyName = v }

func (s *CcompareExpContext) SetPropertyValue(v IFofaValueContext) { s.propertyValue = v }

func (s *CcompareExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CcompareExpContext) FofaKeyword() IFofaKeywordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFofaKeywordContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFofaKeywordContext)
}

func (s *CcompareExpContext) CEQ() antlr.TerminalNode {
	return s.GetToken(FOFAParserCEQ, 0)
}

func (s *CcompareExpContext) FofaValue() IFofaValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFofaValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFofaValueContext)
}

func (s *CcompareExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FOFAListener); ok {
		listenerT.EnterCcompareExp(s)
	}
}

func (s *CcompareExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FOFAListener); ok {
		listenerT.ExitCcompareExp(s)
	}
}

type BracketExpContext struct {
	*QueryContext
	leftBracket  antlr.Token
	rightBracket antlr.Token
}

func NewBracketExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BracketExpContext {
	var p = new(BracketExpContext)

	p.QueryContext = NewEmptyQueryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*QueryContext))

	return p
}

func (s *BracketExpContext) GetLeftBracket() antlr.Token { return s.leftBracket }

func (s *BracketExpContext) GetRightBracket() antlr.Token { return s.rightBracket }

func (s *BracketExpContext) SetLeftBracket(v antlr.Token) { s.leftBracket = v }

func (s *BracketExpContext) SetRightBracket(v antlr.Token) { s.rightBracket = v }

func (s *BracketExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BracketExpContext) Query() IQueryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQueryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQueryContext)
}

func (s *BracketExpContext) BR_OPEN() antlr.TerminalNode {
	return s.GetToken(FOFAParserBR_OPEN, 0)
}

func (s *BracketExpContext) BR_CLOSE() antlr.TerminalNode {
	return s.GetToken(FOFAParserBR_CLOSE, 0)
}

func (s *BracketExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FOFAListener); ok {
		listenerT.EnterBracketExp(s)
	}
}

func (s *BracketExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FOFAListener); ok {
		listenerT.ExitBracketExp(s)
	}
}

type SgExpContext struct {
	*QueryContext
	sgatom IFofaValueContext
}

func NewSgExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SgExpContext {
	var p = new(SgExpContext)

	p.QueryContext = NewEmptyQueryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*QueryContext))

	return p
}

func (s *SgExpContext) GetSgatom() IFofaValueContext { return s.sgatom }

func (s *SgExpContext) SetSgatom(v IFofaValueContext) { s.sgatom = v }

func (s *SgExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SgExpContext) FofaValue() IFofaValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFofaValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFofaValueContext)
}

func (s *SgExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FOFAListener); ok {
		listenerT.EnterSgExp(s)
	}
}

func (s *SgExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FOFAListener); ok {
		listenerT.ExitSgExp(s)
	}
}

type ScompareExpContext struct {
	*QueryContext
	propertyName  IFofaKeywordContext
	op            antlr.Token
	propertyValue IFofaValueContext
}

func NewScompareExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ScompareExpContext {
	var p = new(ScompareExpContext)

	p.QueryContext = NewEmptyQueryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*QueryContext))

	return p
}

func (s *ScompareExpContext) GetOp() antlr.Token { return s.op }

func (s *ScompareExpContext) SetOp(v antlr.Token) { s.op = v }

func (s *ScompareExpContext) GetPropertyName() IFofaKeywordContext { return s.propertyName }

func (s *ScompareExpContext) GetPropertyValue() IFofaValueContext { return s.propertyValue }

func (s *ScompareExpContext) SetPropertyName(v IFofaKeywordContext) { s.propertyName = v }

func (s *ScompareExpContext) SetPropertyValue(v IFofaValueContext) { s.propertyValue = v }

func (s *ScompareExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScompareExpContext) FofaKeyword() IFofaKeywordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFofaKeywordContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFofaKeywordContext)
}

func (s *ScompareExpContext) SEQ() antlr.TerminalNode {
	return s.GetToken(FOFAParserSEQ, 0)
}

func (s *ScompareExpContext) FofaValue() IFofaValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFofaValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFofaValueContext)
}

func (s *ScompareExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FOFAListener); ok {
		listenerT.EnterScompareExp(s)
	}
}

func (s *ScompareExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FOFAListener); ok {
		listenerT.ExitScompareExp(s)
	}
}

type OrLogicalExpContext struct {
	*QueryContext
	leftQuery  IQueryContext
	op         antlr.Token
	rightQuery IQueryContext
}

func NewOrLogicalExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrLogicalExpContext {
	var p = new(OrLogicalExpContext)

	p.QueryContext = NewEmptyQueryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*QueryContext))

	return p
}

func (s *OrLogicalExpContext) GetOp() antlr.Token { return s.op }

func (s *OrLogicalExpContext) SetOp(v antlr.Token) { s.op = v }

func (s *OrLogicalExpContext) GetLeftQuery() IQueryContext { return s.leftQuery }

func (s *OrLogicalExpContext) GetRightQuery() IQueryContext { return s.rightQuery }

func (s *OrLogicalExpContext) SetLeftQuery(v IQueryContext) { s.leftQuery = v }

func (s *OrLogicalExpContext) SetRightQuery(v IQueryContext) { s.rightQuery = v }

func (s *OrLogicalExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrLogicalExpContext) AllQuery() []IQueryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IQueryContext)(nil)).Elem())
	var tst = make([]IQueryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IQueryContext)
		}
	}

	return tst
}

func (s *OrLogicalExpContext) Query(i int) IQueryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQueryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IQueryContext)
}

func (s *OrLogicalExpContext) OR() antlr.TerminalNode {
	return s.GetToken(FOFAParserOR, 0)
}

func (s *OrLogicalExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FOFAListener); ok {
		listenerT.EnterOrLogicalExp(s)
	}
}

func (s *OrLogicalExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FOFAListener); ok {
		listenerT.ExitOrLogicalExp(s)
	}
}

func (p *FOFAParser) Query() (localctx IQueryContext) {
	return p.query(0)
}

func (p *FOFAParser) query(_p int) (localctx IQueryContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewQueryContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IQueryContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, FOFAParserRULE_query, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(32)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		localctx = NewBracketExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(11)

			var _m = p.Match(FOFAParserBR_OPEN)

			localctx.(*BracketExpContext).leftBracket = _m
		}
		{
			p.SetState(12)
			p.query(0)
		}
		{
			p.SetState(13)

			var _m = p.Match(FOFAParserBR_CLOSE)

			localctx.(*BracketExpContext).rightBracket = _m
		}

	case 2:
		localctx = NewCompareExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(15)

			var _x = p.FofaKeyword()

			localctx.(*CompareExpContext).propertyName = _x
		}
		{
			p.SetState(16)

			var _m = p.Match(FOFAParserEQ)

			localctx.(*CompareExpContext).op = _m
		}
		{
			p.SetState(17)

			var _x = p.FofaValue()

			localctx.(*CompareExpContext).propertyValue = _x
		}

	case 3:
		localctx = NewScompareExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(19)

			var _x = p.FofaKeyword()

			localctx.(*ScompareExpContext).propertyName = _x
		}
		{
			p.SetState(20)

			var _m = p.Match(FOFAParserSEQ)

			localctx.(*ScompareExpContext).op = _m
		}
		{
			p.SetState(21)

			var _x = p.FofaValue()

			localctx.(*ScompareExpContext).propertyValue = _x
		}

	case 4:
		localctx = NewNoCompareExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(23)

			var _x = p.FofaKeyword()

			localctx.(*NoCompareExpContext).propertyName = _x
		}
		{
			p.SetState(24)

			var _m = p.Match(FOFAParserNOT)

			localctx.(*NoCompareExpContext).op = _m
		}
		{
			p.SetState(25)

			var _x = p.FofaValue()

			localctx.(*NoCompareExpContext).propertyValue = _x
		}

	case 5:
		localctx = NewCcompareExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(27)

			var _x = p.FofaKeyword()

			localctx.(*CcompareExpContext).propertyName = _x
		}
		{
			p.SetState(28)

			var _m = p.Match(FOFAParserCEQ)

			localctx.(*CcompareExpContext).op = _m
		}
		{
			p.SetState(29)

			var _x = p.FofaValue()

			localctx.(*CcompareExpContext).propertyValue = _x
		}

	case 6:
		localctx = NewSgExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(31)

			var _x = p.FofaValue()

			localctx.(*SgExpContext).sgatom = _x
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(40)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAndLogicalExpContext(p, NewQueryContext(p, _parentctx, _parentState))
				localctx.(*AndLogicalExpContext).leftQuery = _prevctx

				p.PushNewRecursionContext(localctx, _startState, FOFAParserRULE_query)
				p.SetState(34)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(35)

					var _m = p.Match(FOFAParserAND)

					localctx.(*AndLogicalExpContext).op = _m
				}
				{
					p.SetState(36)

					var _x = p.query(8)

					localctx.(*AndLogicalExpContext).rightQuery = _x
				}

			case 2:
				localctx = NewOrLogicalExpContext(p, NewQueryContext(p, _parentctx, _parentState))
				localctx.(*OrLogicalExpContext).leftQuery = _prevctx

				p.PushNewRecursionContext(localctx, _startState, FOFAParserRULE_query)
				p.SetState(37)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(38)

					var _m = p.Match(FOFAParserOR)

					localctx.(*OrLogicalExpContext).op = _m
				}
				{
					p.SetState(39)

					var _x = p.query(7)

					localctx.(*OrLogicalExpContext).rightQuery = _x
				}

			}

		}
		p.SetState(44)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

// IFofaKeywordContext is an interface to support dynamic dispatch.
type IFofaKeywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFofaKeywordContext differentiates from other interfaces.
	IsFofaKeywordContext()
}

type FofaKeywordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFofaKeywordContext() *FofaKeywordContext {
	var p = new(FofaKeywordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FOFAParserRULE_fofaKeyword
	return p
}

func (*FofaKeywordContext) IsFofaKeywordContext() {}

func NewFofaKeywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FofaKeywordContext {
	var p = new(FofaKeywordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FOFAParserRULE_fofaKeyword

	return p
}

func (s *FofaKeywordContext) GetParser() antlr.Parser { return s.parser }

func (s *FofaKeywordContext) AllFOFA_KEY() []antlr.TerminalNode {
	return s.GetTokens(FOFAParserFOFA_KEY)
}

func (s *FofaKeywordContext) FOFA_KEY(i int) antlr.TerminalNode {
	return s.GetToken(FOFAParserFOFA_KEY, i)
}

func (s *FofaKeywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FofaKeywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FofaKeywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FOFAListener); ok {
		listenerT.EnterFofaKeyword(s)
	}
}

func (s *FofaKeywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FOFAListener); ok {
		listenerT.ExitFofaKeyword(s)
	}
}

func (p *FOFAParser) FofaKeyword() (localctx IFofaKeywordContext) {
	localctx = NewFofaKeywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, FOFAParserRULE_fofaKeyword)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(45)
		p.Match(FOFAParserFOFA_KEY)
	}
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FOFAParserT__0 {
		{
			p.SetState(46)
			p.Match(FOFAParserT__0)
		}
		{
			p.SetState(47)
			p.Match(FOFAParserFOFA_KEY)
		}

	}

	return localctx
}

// IFofaValueContext is an interface to support dynamic dispatch.
type IFofaValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFofaValueContext differentiates from other interfaces.
	IsFofaValueContext()
}

type FofaValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFofaValueContext() *FofaValueContext {
	var p = new(FofaValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FOFAParserRULE_fofaValue
	return p
}

func (*FofaValueContext) IsFofaValueContext() {}

func NewFofaValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FofaValueContext {
	var p = new(FofaValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FOFAParserRULE_fofaValue

	return p
}

func (s *FofaValueContext) GetParser() antlr.Parser { return s.parser }

func (s *FofaValueContext) CopyFrom(ctx *FofaValueContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *FofaValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FofaValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BooleanContext struct {
	*FofaValueContext
}

func NewBooleanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanContext {
	var p = new(BooleanContext)

	p.FofaValueContext = NewEmptyFofaValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FofaValueContext))

	return p
}

func (s *BooleanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(FOFAParserBOOLEAN, 0)
}

func (s *BooleanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FOFAListener); ok {
		listenerT.EnterBoolean(s)
	}
}

func (s *BooleanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FOFAListener); ok {
		listenerT.ExitBoolean(s)
	}
}

type StringContext struct {
	*FofaValueContext
}

func NewStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringContext {
	var p = new(StringContext)

	p.FofaValueContext = NewEmptyFofaValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FofaValueContext))

	return p
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) STRING() antlr.TerminalNode {
	return s.GetToken(FOFAParserSTRING, 0)
}

func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FOFAListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FOFAListener); ok {
		listenerT.ExitString(s)
	}
}

func (p *FOFAParser) FofaValue() (localctx IFofaValueContext) {
	localctx = NewFofaValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, FOFAParserRULE_fofaValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(52)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FOFAParserBOOLEAN:
		localctx = NewBooleanContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(50)
			p.Match(FOFAParserBOOLEAN)
		}

	case FOFAParserSTRING:
		localctx = NewStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(51)
			p.Match(FOFAParserSTRING)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *FOFAParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *QueryContext = nil
		if localctx != nil {
			t = localctx.(*QueryContext)
		}
		return p.Query_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *FOFAParser) Query_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 6)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
