package vm

import (
	"github.com/dcaiafa/nitro/internal/meta"
)

type PackageSymbolType uint8

type Program struct {
	Packages []*CompiledPackage
}

type DepPackage struct {
	Name        string
	ResolvedNdx int
}

type CompiledPackage struct {
	Metadata  *meta.Metadata
	MainFnNdx int

	numGlobals int
	globals    []Value
	Literals   []Value
	params     map[string]*Param
	reqParamN  int

	Symbols     map[string]int
	DepPackages []DepPackage
}
