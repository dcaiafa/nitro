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
	extFns    []*NativeFn
	literals  []Value
	params    map[string]*Param
	reqParamN int

	Symbols map[string]PackageSymbol
}

type NativeFnRegistry interface {
	IsValidPackage(pkg string) bool
	GetNativeFn(pkg, name string) *NativeFn
}
