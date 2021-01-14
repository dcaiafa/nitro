package parser

import "github.com/dcaiafa/nitro/internal/ast"

//go:generate goyacc parser.y

func Parse(filename string, input []byte) (*ast.Module, error) {
	//yyDebug = 10
	yyErrorVerbose = true
	l := newLex(input)
	p := yyNewParser()
	p.Parse(l)
	return l.Module, l.err
}
