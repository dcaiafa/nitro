package ast

import (
	"github.com/dcaiafa/nitro/internal/runtime"
	"github.com/dcaiafa/nitro/internal/symbol"
)

type Func struct {
	astBase

	Params    ASTs
	Block     *StmtBlock
	IsClosure bool

	idxFunc    int
	scope      *symbol.Scope
	paramCount int
	localCount int
	captures   []*symbol.CaptureSymbol
	isIterator bool
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
		if f.isIterator {
			// func f(x, y) {         func f(x, y) {
			//    yield x+y      =>     return g() {
			// }                          yield x+y
			//                          }
			//                        }
			g := new(LambdaExpr)
			g.SetPos(f.Pos())
			g.Block = f.Block
			g.isIterator = true

			f.isIterator = false
			f.Block = &StmtBlock{
				Stmts: ASTs{
					&ReturnStmt{
						Values: Exprs{g},
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
			if f.isIterator {
				emitter.Emit(f.Pos(), runtime.OpIterRet, 0, 0)
			} else {
				emitter.Emit(f.Pos(), runtime.OpRet, 0, 0)
			}
		}
		emitter.PopFn()

		if f.IsClosure {
			for _, capture := range f.captures {
				emitSymbolRefPush(f.Pos(), emitter, capture.Captured)
			}
			op := runtime.OpNewClosure
			if f.isIterator {
				op = runtime.OpNewIter
			}
			emitter.Emit(f.Pos(), op, uint32(f.idxFunc), uint16(len(f.captures)))
		}
	}
}

func (f *Func) MarkIterator() {
	f.isIterator = true
}

func (f *Func) IsIterator() bool {
	return f.isIterator
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
		Captured:   sym,
		CaptureNdx: len(f.captures),
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
