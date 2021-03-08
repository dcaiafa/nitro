package parser2

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/dcaiafa/nitro/internal/ast"
	"github.com/dcaiafa/nitro/internal/errlogger"
	"github.com/dcaiafa/nitro/internal/parser2/parser"
)

func Parse(filename string, input string, diagMode bool, errLogger errlogger.ErrLogger) (*ast.Module, error) {
	inputStream := antlr.NewInputStream(string(input))
	lexer := newAugmentedLexer(inputStream)
	nitroParser := parser.NewNitroParser(
		antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel))
	nitroParser.RemoveErrorListeners()
	errListener := newErrorListener(filename, errLogger)
	nitroParser.AddErrorListener(errListener)
	if diagMode {
		nitroParser.GetInterpreter().SetPredictionMode(antlr.PredictionModeLLExactAmbigDetection)
		nitroParser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	}
	listener := newListener(filename)
	parseTree := nitroParser.Start()
	antlr.ParseTreeWalkerDefault.Walk(listener, parseTree)
	if errListener.Error() != nil {
		return nil, errListener.Error()
	}
	return listener.Module, nil
}
