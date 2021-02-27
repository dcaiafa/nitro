package ast

import (
	"github.com/dcaiafa/nitro/internal/runtime"
	"github.com/dcaiafa/nitro/internal/symbol"
)

type Func struct {
	astBase

	Params ASTs
	Stmts  ASTs

	IsClosure bool
	FnNdx     int

	scope      *symbol.Scope
	paramCount int
	localCount int
	captures   []*symbol.CaptureSymbol
}

func (f *Func) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Check:
		f.scope = symbol.NewScope()
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
		emitter := ctx.Emitter()
		synthRet := len(f.Stmts) == 0
		if !synthRet {
			_, hasReturnStmt := f.Stmts[len(f.Stmts)-1].(*ReturnStmt)
			synthRet = !hasReturnStmt
		}
		if synthRet {
			emitter.Emit(runtime.OpRet, 0, 0)
		}
		emitter.PopFn()

		if f.IsClosure {
			// TODO: check if number of captures > 255 (can't fit in Operand2)
			for _, capture := range f.captures {
				emitSymbolRefPush(emitter, capture.Captured)
			}
			emitter.Emit(runtime.OpNewClosure, uint16(f.FnNdx), byte(len(f.captures)))
		}
	}
}

func (f *Func) NewLocal() *symbol.LocalVarSymbol {
	s := &symbol.LocalVarSymbol{}
	s.LocalNdx = f.localCount
	f.localCount++
	return s
}

func (f *Func) NewParam() *symbol.ParamSymbol {
	s := &symbol.ParamSymbol{}
	s.ParamNdx = f.paramCount
	f.paramCount++
	return s
}

func (f *Func) NewCapture(sym symbol.Symbol) *symbol.CaptureSymbol {
	c := &symbol.CaptureSymbol{
		Captured: sym,
	}
	c.SetName(sym.Name())
	c.SetPos(sym.Pos())
	f.captures = append(f.captures, c)
	if !f.scope.PutSymbol(nil, c) {
		// This cannot happen because the caller should have already ensured that
		// the symbol to be captured is external to this fn.
		panic("assert failed")
	}
	return c
}

func (f *Func) Scope() *symbol.Scope {
	return f.scope
}
