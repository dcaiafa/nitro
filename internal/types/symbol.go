package types

import "github.com/dcaiafa/nitro/internal/token"

type Type int

const (
	NoType Type = iota
	Module
	Function
	Dynamic
)

type Symbol struct {
	Name  string
	Type  Type
	Pos   token.Pos
	Scope *Scope
}

func NewSymbol(name string) *Symbol {
	return &Symbol{Name: name}
}
