package symbol

import (
	"github.com/dcaiafa/nitro/internal/token"
)

type Symbol interface {
	Name() string
	SetName(name string)
	Pos() token.Pos
	SetPos(pos token.Pos)
	Liftable() bool
	SetLiftable(l bool)
	Lifted() bool
	Lift()
}

type baseSymbol struct {
	name     string
	pos      token.Pos
	liftable bool
	lifted   bool
}

func (b *baseSymbol) Name() string {
	return b.name
}

func (b *baseSymbol) SetName(name string) {
	b.name = name
}

func (b *baseSymbol) Pos() token.Pos {
	return b.pos
}

func (b *baseSymbol) SetPos(pos token.Pos) {
	b.pos = pos
}

func (b *baseSymbol) Liftable() bool {
	return b.liftable
}

func (b *baseSymbol) SetLiftable(l bool) {
	b.liftable = l
}

func (b *baseSymbol) Lifted() bool {
	return b.lifted
}

func (b *baseSymbol) Lift() {
	b.lifted = true
}

type FuncSymbol struct {
	baseSymbol
	External bool
	IdxFunc  int
}

type GlobalVarSymbol struct {
	baseSymbol
	GlobalNdx int
}

type CaptureSymbol struct {
	baseSymbol

	Capture    Symbol
	CaptureNdx int
}

type ParamSymbol struct {
	baseSymbol
	ParamNdx int
}

type LocalVarSymbol struct {
	baseSymbol
	LocalNdx int
}

type ModuleRef struct {
	baseSymbol
	Module *Module
}

type Module struct {
	Name  string
	Scope *Scope
}

func NewModule(name string) *Module {
	return &Module{
		Name:  name,
		Scope: NewScope(),
	}
}
