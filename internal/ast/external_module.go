package ast

import "github.com/dcaiafa/nitro/internal/vm"

type ExternalModule struct {
	astBase

	nativeFns ASTs
}

func (m *ExternalModule) AddNativeFn(name string, natFn vm.NativeFn) {
	m.nativeFns = append(
		m.nativeFns,
		&ExternFn{
			Name:     name,
			ExternFn: natFn,
		})
}

func (m *ExternalModule) RunPass(ctx *Context, pass Pass) {
	ctx.RunPassChild(m, m.nativeFns, pass)
}
