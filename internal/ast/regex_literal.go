package ast

type RegexLiteral struct {
	astBase

	Regex string
}

func (l *RegexLiteral) isExpr() {}
