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
	Lifted() bool
	Lift()
	ReadOnly() bool
	SetReadOnly(ro bool)
}

type baseSymbol struct {
	name     string
	pos      token.Pos
	readOnly bool
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

func (b *baseSymbol) ReadOnly() bool {
	return b.readOnly
}

func (b *baseSymbol) SetReadOnly(ro bool) {
	b.readOnly = ro
}

type baseLiftable struct {
	lifted bool
}

func (b *baseLiftable) Liftable() bool {
	return true
}

func (b *baseLiftable) Lifted() bool {
	return b.lifted
}

func (b *baseLiftable) Lift() {
	b.lifted = true
}

type baseNonLiftable struct {
}

func (b *baseNonLiftable) Liftable() bool {
	return false
}

func (b *baseNonLiftable) Lifted() bool {
	return false
}

func (b *baseNonLiftable) Lift() {
	panic("not reachable")
}

type GlobalVarSymbol struct {
	baseSymbol
	baseNonLiftable

	PackageNdx int
	GlobalNdx  int
	Export     bool
}

type CaptureSymbol struct {
	baseSymbol
	baseLiftable

	Capture    Symbol
	CaptureNdx int
}

type ParamSymbol struct {
	baseSymbol
	baseLiftable

	ParamNdx int
}

type LocalVarSymbol struct {
	baseSymbol
	baseLiftable

	LocalNdx int
}
