package ast

import (
	"github.com/dcaiafa/nitro/internal/runtime"
	"github.com/dcaiafa/nitro/internal/types"
)

type Func struct {
	astBase

	Params ASTs
	Stmts  ASTs

	FnNdx int

	scope      *types.Scope
	paramCount int
	localCount int
}

func (f *Func) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Check:
		f.scope = types.NewScope()
		f.FnNdx = ctx.Emitter().NewFn()

	case Emit:
		emitter := ctx.Emitter()
		emitter.PushFn(f.FnNdx)
		emitter.Emit(runtime.OpInitCallFrame, uint16(f.localCount), 0)
	}

	ctx.Push(f)
	f.Params.RunPass(ctx, pass)
	f.Stmts.RunPass(ctx, pass)
	ctx.Pop()

	switch pass {
	case Emit:
		synthRet := len(f.Stmts) == 0
		if !synthRet {
			_, hasReturnStmt := f.Stmts[len(f.Stmts)-1].(*ReturnStmt)
			synthRet = !hasReturnStmt
		}
		if synthRet {
			ctx.Emitter().Emit(runtime.OpRet, 0, 0)
		}

		ctx.Emitter().PopFn()
	}
}

func (f *Func) NewLocal() *types.LocalVarSymbol {
	s := &types.LocalVarSymbol{}
	s.LocalNdx = f.localCount
	f.localCount++
	return s
}

func (f *Func) NewParam() *types.ParamSymbol {
	s := &types.ParamSymbol{}
	s.ParamNdx = f.paramCount
	f.paramCount++
	return s
}

func (f *Func) Scope() *types.Scope {
	return f.scope
}
