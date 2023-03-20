package symbol

import (
	"github.com/dcaiafa/nitro/internal/vm"
)

// TODO: rename file: package.go => import.go

type Import struct {
	baseSymbol
	baseNonLiftable

	Package *vm.CompiledPackage

	index   int
	symbols map[string]*LiteralSymbol
}

var _ Symbol = (*Import)(nil)

func NewImport(pkg *vm.CompiledPackage, pkgIndex int) *Import {
	return &Import{
		Package: pkg,
		index:   pkgIndex,
		symbols: make(map[string]*LiteralSymbol, len(pkg.Symbols)),
	}
}

func (i *Import) GetSymbol(name string) Symbol {
	sym := i.symbols[name]
	if sym != nil {
		return sym
	}

	ndx, ok := i.Package.Symbols[name]
	if !ok {
		return nil
	}

	sym = new(LiteralSymbol)
	sym.SetName(name)
	sym.SetReadOnly(true)
	sym.PackageIdx = i.index
	sym.LiteralIdx = ndx

	i.symbols[name] = sym

	return sym
}
