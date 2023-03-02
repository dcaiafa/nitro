package symbol

import (
	"github.com/dcaiafa/nitro/internal/vm"
)

type Import struct {
	baseSymbol

	Package *vm.CompiledPackage

	index   int
	symbols map[string]*GlobalVarSymbol
}

func NewImport(pkg *vm.CompiledPackage, pkgIndex int) *Import {
	i := &Import{
		Package: pkg,
		index:   pkgIndex,
		symbols: make(map[string]*GlobalVarSymbol, len(pkg.Symbols)),
	}
	i.SetReadOnly(true)
	return i
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

	sym = new(GlobalVarSymbol)
	sym.SetName(name)
	sym.SetReadOnly(true)
	sym.PackageNdx = i.index
	sym.GlobalNdx = ndx

	i.symbols[name] = sym

	return sym
}
