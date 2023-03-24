package scope

import (
	"github.com/dcaiafa/nitro/internal/errlogger"
	"github.com/dcaiafa/nitro/internal/symbol"
)

type Type int

const (
	Package Type = 1 << iota
	Unit
	Block
)

type Scope interface {
	Type() Type
	GetSymbol(name string) symbol.Symbol
	PutSymbol(l errlogger.ErrLogger, sym symbol.Symbol) bool
}

type SimpleScope struct {
	typ     Type
	symbols map[string]symbol.Symbol
	onMiss  func(name string) symbol.Symbol
}

// TODO: rename => New or NewSimple
func NewScope(typ Type) *SimpleScope {
	return &SimpleScope{
		typ:     typ,
		symbols: make(map[string]symbol.Symbol),
	}
}

func (s *SimpleScope) Type() Type {
	return s.typ
}

func (s *SimpleScope) GetSymbol(name string) symbol.Symbol {
	sym := s.symbols[name]
	if sym == nil && s.onMiss != nil {
		sym = s.onMiss(name)
	}
	return sym
}

func (s *SimpleScope) PutSymbol(l errlogger.ErrLogger, sym symbol.Symbol) bool {
	existing := s.symbols[sym.Name()]
	if existing != nil {
		l.Failf(
			sym.Pos(),
			"there is already something named %q in the current scope.",
			sym.Name())
		l.Detailf(existing.Pos(), "%q was previously declared here.", sym.Name())
		return false
	}
	s.symbols[sym.Name()] = sym
	return true
}

func (s *SimpleScope) SetMissHandler(onMiss func(name string) symbol.Symbol) {
	s.onMiss = onMiss
}

func (s *SimpleScope) ForEachSymbol(fn func(sym symbol.Symbol)) {
	for _, sym := range s.symbols {
		fn(sym)
	}
}
