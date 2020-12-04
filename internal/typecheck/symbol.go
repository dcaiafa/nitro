package typecheck

import "github.com/dcaiafa/nitro/internal/token"

type Symbol struct {
	Name string
	Pos  token.Pos
}

func NewSymbol(name string) *Symbol {
	return &Symbol{Name: name}
}
