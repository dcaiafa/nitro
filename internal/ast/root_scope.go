package ast

import (
	"github.com/dcaiafa/nitro/internal/errlogger"
	"github.com/dcaiafa/nitro/internal/symbol"
	"github.com/dcaiafa/nitro/internal/vm"
)

type RootScope struct {
	emitter        *vm.Emitter
	packageName    string
	funcRegistries []vm.ExportRegistry
	scope          symbol.Scope
}

func NewRootScope(emitter *vm.Emitter, packageName string, funcRegistries []vm.ExportRegistry) *RootScope {
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
		// Lazily create package symbol here.
		// TODO: do this only for simple scripts.
    // Notice that RootScope is also used as the scope for the Package. This
    // means this instance could be a package scope. Since native packages can
    // only be one level deep, and IsValidPackage starts from root, only create
    // package if this is root (i.e. packageName == "").
		if s.packageName == "" && reg.IsValidPackage(name) {
			pkgSym := &symbol.Package{
				Scope: NewRootScope(s.emitter, name, s.funcRegistries),
			}
			pkgSym.SetName(name)
			pkgSym.SetReadOnly(true)
			s.scope.PutSymbol(nil, pkgSym)
			return pkgSym
		}

		if v := reg.GetExport(s.packageName, name); v != nil {
      switch v := v.(type) {
      case *vm.NativeFn:
        fnSym := &symbol.FuncSymbol{
          External: true,
          IdxFunc:  s.emitter.AddExternalFunc(s.packageName, name, v),
        }
        fnSym.SetName(name)
        fnSym.SetReadOnly(true)
        s.scope.PutSymbol(nil, fnSym)
        return fnSym
      case vm.String, vm.Int, vm.Float:
        constSym := &symbol.ConstSymbol{
          LiteralNdx: s.emitter.AddLiteral(v),
        }
        constSym.SetName(name)
        constSym.SetReadOnly(true)
        s.scope.PutSymbol(nil, constSym)
      default:
        panic("invalid const type")
      }
		}
	}
	return nil
}

func (s *RootScope) PutSymbol(l errlogger.ErrLogger, sym symbol.Symbol) bool {
	return s.scope.PutSymbol(l, sym)
}
