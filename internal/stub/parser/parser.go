package parser

import "github.com/dcaiafa/nitro/internal/stub/ast"

//go:generate goyacc parser.y

func Parse(filename string, input []byte) (*ast.Unit, error) {
	//yyDebug = 10
	yyErrorVerbose = true
	l := newLex(input)
	p := yyNewParser()
	p.Parse(l)
	return l.Unit, l.err
}
