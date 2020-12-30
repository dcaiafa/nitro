package types

import (
	"github.com/dcaiafa/nitro/internal/token"
)

type Symbol interface {
	Name() string
	SetName(name string)
	Owner() *FuncSymbol
	SetOwner(owner *FuncSymbol)
	Pos() token.Pos
	SetPos(pos token.Pos)
}

type baseSymbol struct {
	name  string
	owner *FuncSymbol
	pos   token.Pos
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
	Locals   []*LocalVarSymbol
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

	Local int
}

type CaptureSymbol struct {
	baseSymbol
	Captured Symbol
}

type ParamSymbol struct {
	baseSymbol
}
