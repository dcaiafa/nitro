package ast

type Module struct {
	astBase

	Units []*Unit
}

func (m *Module) RunPass(ctx *Context, pass Pass) {
	for _, unit := range m.Units {
		ctx.RunPassChild(m, unit, pass)
	}
}
