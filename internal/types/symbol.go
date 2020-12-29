package types

import (
	"github.com/dcaiafa/nitro/internal/token"
)

type Symbol interface {
	Name() string
	SetName(name string)
	Owner() *FuncSymbol
	SetOwner(owner *FuncSymbol)
	StorageIndex() int
	SetStorageIndex(ndx int)
	Pos() token.Pos
	SetPos(pos token.Pos)
}

type baseSymbol struct {
	name         string
	owner        *FuncSymbol
	storageIndex int
	pos          token.Pos
}

func (b *baseSymbol) Name() string {
	return b.name
}

func (b *baseSymbol) SetName(name string) {
	b.name = name
}

func (b *baseSymbol) Owner() *FuncSymbol {
	return b.owner
}

func (b *baseSymbol) SetOwner(owner *FuncSymbol) {
	b.owner = owner
}

func (b *baseSymbol) StorageIndex() int {
	return b.storageIndex
}

func (b *baseSymbol) SetStorageIndex(ndx int) {
	b.storageIndex = ndx
}

func (b *baseSymbol) Pos() token.Pos {
	return b.pos
}

func (b *baseSymbol) SetPos(pos token.Pos) {
	b.pos = pos
}

type FuncSymbol struct {
	baseSymbol

	Captures []Symbol
	Params   []Symbol
	Locals   []Symbol
	External bool
	Fn       int

	Scope *Scope
}

type ExternalFuncSymbol struct {
	baseSymbol

	FuncIndex int
}

type LocalVarSymbol struct {
	baseSymbol
}

type CaptureSymbol struct {
	baseSymbol
	Captured Symbol
}

type ParamSymbol struct {
	baseSymbol
}
