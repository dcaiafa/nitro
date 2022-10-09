package parser2

import (
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/dcaiafa/nitro/internal/ast"
	"github.com/dcaiafa/nitro/internal/errlogger"
	"github.com/dcaiafa/nitro/internal/parser2/parser"
)

var parserPool = sync.Pool{
	New: func() interface{} {
		return parser.NewNitroParser(nil)
	},
}

var lexerPool = sync.Pool{
	New: func() interface{} {
		return newAugmentedLexer(nil)
	},
}

func ParseUnit(filename string, input string, diagMode bool, errLogger errlogger.ErrLogger) (*ast.Unit, error) {
	lexer := lexerPool.Get().(*augmentedLexer)
	lexer.SetInputStream(antlr.NewInputStream(string(input)))

	nitroParser := parserPool.Get().(*parser.NitroParser)
	nitroParser.SetInputStream(antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel))

	nitroParser.RemoveErrorListeners()
	errListener := newErrorListener(filename, errLogger)
	nitroParser.AddErrorListener(errListener)
	if diagMode {
		nitroParser.GetInterpreter().SetPredictionMode(antlr.PredictionModeLLExactAmbigDetection)
		nitroParser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	}
	listener := newListener(filename, errListener)
	parseTree := nitroParser.Start()

	parserPool.Put(nitroParser)
	lexerPool.Put(lexer)
	antlr.ParseTreeWalkerDefault.Walk(listener, parseTree)
	if errListener.Error() != nil {
		return nil, errListener.Error()
	}
	return listener.Unit, nil
}

func ParsePrologue(filename string, input string, diagMode bool, errLogger errlogger.ErrLogger) (*Prologue, error) {
	lexer := lexerPool.Get().(*augmentedLexer)
	lexer.SetInputStream(antlr.NewInputStream(string(input)))

	nitroParser := parserPool.Get().(*parser.NitroParser)
	nitroParser.SetInputStream(antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel))

	nitroParser.RemoveErrorListeners()
	errListener := newErrorListener(filename, errLogger)
	nitroParser.AddErrorListener(errListener)
	if diagMode {
		nitroParser.GetInterpreter().SetPredictionMode(antlr.PredictionModeLLExactAmbigDetection)
		nitroParser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	}
	listener := newListener(filename, errListener)
	parseTree := nitroParser.Prologue()

	parserPool.Put(nitroParser)
	lexerPool.Put(lexer)
	antlr.ParseTreeWalkerDefault.Walk(listener, parseTree)
	if errListener.Error() != nil {
		return nil, errListener.Error()
	}
	return listener.Prologue, nil
}
