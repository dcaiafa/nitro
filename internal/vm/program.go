package vm

import (
	"github.com/dcaiafa/nitro/internal/meta"
)

type PackageSymbolType uint8

type Program struct {
	Packages []*CompiledPackage
}

func (p *Program) MainPackage() *CompiledPackage {
	return p.Packages[len(p.Packages)-1]
}

func (p *Program) Metadata() *meta.Metadata {
	return p.MainPackage().Metadata
}

type CompiledPackage struct {
	Name string

	// Index is the index of the CompiledPackage in the list of program packages.
	Index int

	Metadata  *meta.Metadata
	MainFnNdx int

	numGlobals int
	globals    []Value
	Literals   []Value
	params     map[string]*Param
	reqParamN  int

	Deps    []*CompiledPackage
	Symbols map[string]int // symbolName => LiteralsIndex
}
