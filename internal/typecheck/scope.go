package typecheck

import "log"

type Scope struct {
	symbols map[string]*Symbol
}

func NewScope() *Scope {
	return &Scope{}
}

func (s *Scope) GetSymbol(name string) *Symbol {
	return s.symbols[name]
}

func (s *Scope) PutSymbol(sym *Symbol) {
	if s.symbols[sym.Name] != nil {
		log.Panicf("Symbol with name %v already exists", sym.Name)
	}
	s.symbols[sym.Name] = sym
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

func (s *ScopeStack) FindSymbol(name string) (*Symbol, *Scope) {
	for i := len(s.scopes) - 1; i >= 0; i-- {
		scope := s.scopes[i]
		sym := scope.GetSymbol(name)
		if sym != nil {
			return sym, scope
		}
	}
	return nil, nil
}
