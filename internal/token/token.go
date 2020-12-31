package token

import "fmt"

type Type int

const (
	Number Type = iota
	String
	Bool
	Keyword
	ID
)

type Pos struct {
	Filename string
	Line     int
	Col      int
}

func (p *Pos) String() string {
	return fmt.Sprintf("%v:%v:%v", p.Filename, p.Line, p.Col)
}

type Token struct {
	Pos  Pos
	Type Type
	Str  string
	Num  float64
	Bool bool
}
