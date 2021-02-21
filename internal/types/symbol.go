package types

import (
	"github.com/dcaiafa/nitro/internal/token"
)

type Symbol interface {
	Name() string
	SetName(name string)
	Pos() token.Pos
	SetPos(pos token.Pos)
}

type baseSymbol struct {
	name string
	pos  token.Pos
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

type FuncSymbol struct {
	baseSymbol
	External bool
	FnNdx    int
}

type GlobalVarSymbol struct {
	baseSymbol
	GlobalNdx int
}

type Capturable interface {
	isCapturable()
}

type CaptureSymbol struct {
	baseSymbol

	Captured   Symbol
	CaptureNdx int
}

func (s *CaptureSymbol) isCapturable() {}

type ParamSymbol struct {
	baseSymbol
	ParamNdx int
}

func (s *ParamSymbol) isCapturable() {}

type LocalVarSymbol struct {
	baseSymbol
	LocalNdx int
}

func (s *LocalVarSymbol) isCapturable() {}