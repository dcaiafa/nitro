package ast

import (
	"github.com/dcaiafa/nitro/internal/errlogger"
	"github.com/dcaiafa/nitro/internal/symbol"
	"github.com/dcaiafa/nitro/internal/vm"
)

type RootScope struct {
	emitter        *vm.Emitter
	packageName    string
	funcRegistries []vm.NativeFnRegistry
	scope          symbol.Scope
}

func NewRootScope(emitter *vm.Emitter, packageName string, funcRegistries []vm.NativeFnRegistry) *RootScope {
	return &RootScope{
		emitter:        emitter,
		packageName:    packageName,
		funcRegistries: funcRegistries,
		scope:          symbol.NewScope(),
	}
}

func (s *RootScope) GetSymbol(name string) symbol.Symbol {
	sym := s.scope.GetSymbol(name)
	if sym != nil {
		return sym
	}
	for _, reg := range s.funcRegistries {
		// Lazy create package symbol here.
		// TODO: do this only for simple scripts.
		// Notice that RootScope is also used as the scope for the Package. This
		// means this instance could be a package scope. Since native packages can
		// only be one level deep, and IsValidPackage starts from root, only lazy
		// create package if this is root (i.e. packageName == "").
		if s.packageName == "" && reg.IsValidPackage(name) {
			pkgSym := &symbol.Package{
				Scope: NewRootScope(s.emitter, name, s.funcRegistries),
			}
			pkgSym.SetName(name)
			pkgSym.SetReadOnly(true)
			s.scope.PutSymbol(nil, pkgSym)
			return pkgSym
		}

		if fn := reg.GetNativeFn(s.packageName, name); fn != nil {
			fnSym := &symbol.FuncSymbol{
				External: true,
				IdxFunc:  s.emitter.AddExternalFunc(s.packageName, name, fn),
			}
			fnSym.SetName(name)
			fnSym.SetReadOnly(true)
			s.scope.PutSymbol(nil, fnSym)
			return fnSym
		}
	}
	return nil
}

func (s *RootScope) PutSymbol(l errlogger.ErrLogger, sym symbol.Symbol) bool {
	return s.scope.PutSymbol(l, sym)
}
