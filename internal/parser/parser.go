package parser

//go:generate goyacc parser.y

func Parse(input string) (interface{}, error) {
	//yyDebug = 10
	yyErrorVerbose = true
	l := newLex(input)
	p := yyNewParser()
	p.Parse(l)
	return nil, l.err
}
