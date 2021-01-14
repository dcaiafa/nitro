package ast

import (
	"github.com/dcaiafa/nitro/internal/runtime"
	"github.com/dcaiafa/nitro/internal/types"
)

type Func struct {
	astBase

	Name   string
	Params ASTs
	Stmts  ASTs

	FnNdx int

	scope      *types.Scope
	paramCount int
	localCount int
}

func (a *Func) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case CreateAndResolveNames:
		a.scope = types.NewScope()
		a.FnNdx = ctx.Emitter().NewFn()

	case Emit:
		emitter := ctx.Emitter()
		emitter.PushFn(a.FnNdx)
		emitter.Emit(runtime.OpInitCallFrame, uint16(a.localCount), 0)
	}

	ctx.Push(a)
	a.Params.RunPass(ctx, pass)
	a.Stmts.RunPass(ctx, pass)
	ctx.Pop()

	switch pass {
	case Emit:
		synthRet := len(a.Stmts) == 0
		if !synthRet {
			_, hasReturnStmt := a.Stmts[len(a.Stmts)-1].(*ReturnStmt)
			synthRet = !hasReturnStmt
		}
		if synthRet {
			ctx.Emitter().Emit(runtime.OpRet, 0, 0)
		}

		ctx.Emitter().PopFn()
	}
}

func (a *Func) NewLocal() *types.LocalVarSymbol {
	s := &types.LocalVarSymbol{}
	s.Local = a.localCount
	a.localCount++
	return s
}

func (a *Func) NewParam() *types.ParamSymbol {
	s := &types.ParamSymbol{}
	s.ParamNdx = a.paramCount
	a.paramCount++
	return s
}

func (a *Func) Scope() *types.Scope {
	return a.scope
}
