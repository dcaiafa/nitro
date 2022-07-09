package symbol

import (
	"github.com/dcaiafa/nitro/internal/errlogger"
)

type Scope interface {
	GetSymbol(name string) Symbol
	PutSymbol(l errlogger.ErrLogger, sym Symbol) bool
}

type SimpleScope struct {
	symbols map[string]Symbol
}

func NewScope() *SimpleScope {
	return &SimpleScope{
		symbols: make(map[string]Symbol),
	}
}

func (s *SimpleScope) GetSymbol(name string) Symbol {
	return s.symbols[name]
}

func (s *SimpleScope) PutSymbol(l errlogger.ErrLogger, sym Symbol) bool {
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

func (s *SimpleScope) ForEachSymbol(fn func(sym Symbol)) {
	for _, sym := range s.symbols {
		fn(sym)
	}
}

type ScopeStack struct {
	scopes []Scope
}

func NewScopeStack() *ScopeStack {
	return &ScopeStack{}
}

func (s *ScopeStack) Push(scope Scope) {
	s.scopes = append(s.scopes, scope)
}

func (s *ScopeStack) Pop() {
	if len(s.scopes) == 0 {
		panic("There are no scopes")
	}
	s.scopes = s.scopes[:len(s.scopes)-1]
}

func (s *ScopeStack) Current() Scope {
	return s.scopes[0]
}

func (s *ScopeStack) FindSymbol(name string) (Symbol, Scope) {
	for i := len(s.scopes) - 1; i >= 0; i-- {
		scope := s.scopes[i]
		sym := scope.GetSymbol(name)
		if sym != nil {
			return sym, scope
		}
	}
	return nil, nil
}
