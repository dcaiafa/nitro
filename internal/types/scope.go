package types

import (
	"github.com/dcaiafa/nitro/internal/errlogger"
)

type Scope struct {
	symbols map[string]Symbol
}

func NewScope() *Scope {
	return &Scope{
		symbols: make(map[string]Symbol),
	}
}

func (s *Scope) GetSymbol(name string) Symbol {
	return s.symbols[name]
}

func (s *Scope) PutSymbol(l errlogger.ErrLogger, sym Symbol) bool {
	if s.symbols[sym.Name()] != nil {
		l.Failf(
			sym.Pos(),
			"There is already something named %q in the current scope.",
			sym.Name())
		l.Detailf(sym.Pos(), "%q was previously declared here.", sym.Name())
		return false
	}
	s.symbols[sym.Name()] = sym
	return true
}

type ScopeStack struct {
	scopes []*Scope
}

func NewScopeStack() *ScopeStack {
	return &ScopeStack{}
}

func (s *ScopeStack) Push(scope *Scope) {
	s.scopes = append(s.scopes, scope)
}

func (s *ScopeStack) Pop() {
	if len(s.scopes) == 0 {
		panic("There are no scopes")
	}
	s.scopes = s.scopes[:len(s.scopes)-1]
}

func (s *ScopeStack) Current() *Scope {
	return s.scopes[0]
}

func (s *ScopeStack) FindSymbol(name string) (Symbol, *Scope) {
	for i := len(s.scopes) - 1; i >= 0; i-- {
		scope := s.scopes[i]
		sym := scope.GetSymbol(name)
		if sym != nil {
			return sym, scope
		}
	}
	return nil, nil
}
