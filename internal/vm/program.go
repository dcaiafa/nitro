package vm

import (
	"github.com/dcaiafa/nitro/internal/meta"
)

type PackageSymbolType uint8

type Program struct {
	Packages []*CompiledPackage
}

type CompiledPackage struct {
	// Index is the index of the CompiledPackage in the list of program packages.
	Index int

	Metadata  *meta.Metadata
	MainFnNdx int

	numGlobals int
	globals    []Value
	Literals   []Value
	params     map[string]*Param
	reqParamN  int

	Deps    map[string]*CompiledPackage // key is packageName
	Symbols map[string]int              // symbolName => LiteralsIndex
}
