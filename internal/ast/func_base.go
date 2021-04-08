package ast

import (
	"github.com/dcaiafa/nitro/internal/runtime"
	"github.com/dcaiafa/nitro/internal/symbol"
)

type Func struct {
	astBase

	Params       ASTs
	Block        *StmtBlock
	IsClosure    bool
	IsEnumerable bool

	idxFunc    int
	scope      *symbol.Scope
	paramCount int
	localCount int
	captures   []*symbol.CaptureSymbol
	enumerable bool
}

func (f *Func) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Check:
		f.scope = symbol.NewScope()
		f.idxFunc = ctx.Emitter().NewFn()

	case Emit:
		emitter := ctx.Emitter()
		emitter.PushFn(f.idxFunc)
		emitter.Emit(f.Pos(), runtime.OpInitCallFrame, uint32(f.localCount), 0)
	}

	ctx.Push(f)
	f.Params.RunPass(ctx, pass)
	f.Block.RunPass(ctx, pass)
	ctx.Pop()

	switch pass {
	case Rewrite:
		if f.enumerable {
			// func f(x, y) {         func g(x, y) {
			//    yield x+y      =>     return f() {
			// }                          yield x+y
			//                          }
			//                        }
			enumerator := new(LambdaExpr)
			enumerator.SetPos(f.Pos())
			enumerator.Block = f.Block

			f.Block = &StmtBlock{
				Stmts: ASTs{
					&ReturnStmt{
						Values: Exprs{enumerator},
					},
				},
			}
		}

	case Emit:
		emitter := ctx.Emitter()
		synthRet := len(f.Block.Stmts) == 0
		if !synthRet {
			_, hasReturnStmt := f.Block.Stmts[len(f.Block.Stmts)-1].(*ReturnStmt)
			synthRet = !hasReturnStmt
		}
		if synthRet {
			emitter.Emit(f.Pos(), runtime.OpRet, 0, 0)
		}
		emitter.PopFn()

		if f.IsClosure {
			// TODO: check if number of captures > 255 (can't fit in Operand2)
			for _, capture := range f.captures {
				emitSymbolRefPush(f.Pos(), emitter, capture.Captured)
			}
			emitter.Emit(f.Pos(), runtime.OpNewClosure, uint32(f.idxFunc), uint16(len(f.captures)))
		}
	}
}

func (f *Func) MarkEnumerable() {
	f.enumerable = true
}

func (f *Func) Enumerable() bool {
	return f.enumerable
}

func (f *Func) IdxFunc() int {
	return f.idxFunc
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
