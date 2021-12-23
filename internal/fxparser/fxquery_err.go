package fxparser

import (
	"fofax/internal/printer"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// https://github.com/antlr/antlr4/issues/2158

func NewFxErrorListener() *FxErrorListener {
	return &FxErrorListener{
		errors: 0,
	}
}

type FxErrorListener struct {
	errors int
}

func (l *FxErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	printer.Info(msg)
	l.errors += 1
}

func (l *FxErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	l.errors += 1
}
func (l *FxErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	l.errors += 1
}
func (l *FxErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	l.errors += 1
}
