package symbol

import (
	"github.com/dcaiafa/nitro/internal/vm"
)

type Package interface {
	Symbol

	GetSymbol(name string) Symbol
}

type PackageAlias struct {
	baseSymbol
	baseNonLiftable

	Package Package
}

var _ Package = (*PackageAlias)(nil)

func (a *PackageAlias) GetSymbol(name string) Symbol {
	return a.Package.GetSymbol(name)
}

type RegularPackage struct {
	baseSymbol
	baseNonLiftable

	compiledPackage *vm.CompiledPackage
	symbols         map[string]Symbol
}

var _ Package = (*RegularPackage)(nil)

func NewRegularPackage(name string, compiledPackage *vm.CompiledPackage) *RegularPackage {
	p := &RegularPackage{
		compiledPackage: compiledPackage,
		symbols:         make(map[string]Symbol, len(compiledPackage.Symbols)),
	}
	p.SetName(name)
	p.SetReadOnly(true)
	return p
}

func (p *RegularPackage) GetSymbol(name string) Symbol {
	sym := p.symbols[name]
	if sym == nil {
		psym, ok := p.compiledPackage.Symbols[name]
		if ok {
			switch psym.Type {
			case vm.PackageSymbolFunc:
				fsym := new(FuncSymbol)
				fsym.IdxFunc = int(psym.Index)
				fsym.SetName(name)
				fsym.SetReadOnly(true)
				sym = fsym

			case vm.PackageSymbolConst:
				csym := new(ConstSymbol)
				csym.LiteralNdx = int(psym.Index)
				sym = csym

			default:
				panic("not reached")
			}
			// TODO: do we need to set position for these symbols?
			p.symbols[name] = sym
		}
	}

	return sym
}
