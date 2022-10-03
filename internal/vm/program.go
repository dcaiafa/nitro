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

type PackageSymbol struct {
	Type  PackageSymbolType
	Index uint32
}

type CompiledPackage struct {
	Metadata  *meta.Metadata
	MainFnNdx int

	globals   int
	fns       []Fn
	literals  []Value
	params    map[string]*Param
	reqParamN int

	Symbols map[string]PackageSymbol
}

type ExportRegistry interface {
	IsValidPackage(pkg string) bool
	GetExport(pkg, name string) Value
}
