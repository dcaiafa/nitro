package vm

import (
	"github.com/dcaiafa/nitro/internal/meta"
)

type PackageSymbolType uint8

const (
	PackageSymbolInvalid PackageSymbolType = iota
	PackageSymbolFunc
	PackageSymbolConst
)

type Program struct {
	Packages []*CompiledPackage
}

type DepPackage struct {
	Name        string
	ResolvedNdx int
}

type PackageSymbol struct {
	Type  PackageSymbolType
	Index uint32
}

type CompiledPackage struct {
	Metadata  *meta.Metadata
	MainFnNdx int

	numGlobals int
	globals    []Value
	literals   []Value
	params     map[string]*Param
	reqParamN  int

	Symbols     map[string]PackageSymbol
	DepPackages []DepPackage
}

type ExportRegistry interface {
	IsValidPackage(pkg string) bool
	GetExport(pkg, name string) Value
}
