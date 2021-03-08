package parser2

import (
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/dcaiafa/nitro/internal/parser2/parser"
)

func TestHelloWorld(t *testing.T) {
	prog := `
	var a = {
	  hello: "world"
		a: 123
	if x > 123 then
	  if: true, end: false,
		elif: 123
		else: xxx
	end
	  d: 123
	}
		`
	s := antlr.NewInputStream(prog)
	l := newAugmentedLexer(s)
	p := parser.NewNitroParser(antlr.NewCommonTokenStream(l, antlr.TokenDefaultChannel))
	p.GetInterpreter().SetPredictionMode(antlr.PredictionModeLLExactAmbigDetection)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	antlr.ParseTreeWalkerDefault.Walk(newListener(), p.Start())
}
