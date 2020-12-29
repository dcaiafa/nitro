package ast

import (
	"github.com/dcaiafa/nitro/internal/runtime"
	"github.com/dcaiafa/nitro/internal/token"
	"github.com/dcaiafa/nitro/internal/types"
)

type SimpleRef struct {
	astBase
	ID token.Token

	sym types.Symbol
}

func (s *SimpleRef) isExpr() {}

func (s *SimpleRef) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case CreateAndResolveNames:
		symName := s.ID.Str

		sym := ctx.FindSymbol(symName)
		if sym == nil {
			ctx.Failf(s.Pos(), "Symbol %q not found.", symName)
		}

		currentFunc := ctx.CurrentFunc()
		if sym.Owner() != nil && sym.Owner() != currentFunc {
			capture := &types.LocalVarSymbol{}
			capture.SetName(sym.Name())
			capture.SetOwner(currentFunc)
			capture.SetPos(sym.Pos())

			currentFunc.Captures = append(currentFunc.Captures, capture)
			sym = capture
		}
		s.sym = sym

	case Emit:
		emit := emitSymbolPush
		_, isRValue := ctx.Parent().(*RValue)
		if isRValue {
			emit = emitSymbolRefPush
		}
		emit(ctx.Emitter(), s.sym)
	}
}

func emitSymbolPush(emitter *runtime.Emitter, sym types.Symbol) {
	switch sym := sym.(type) {
	case *types.LocalVarSymbol:
		emitter.Emit(runtime.OpPushLocal, uint64(sym.StorageIndex()))

	//case *types.CaptureSymbol:
	//case *types.ParamSymbol:
	case *types.FuncSymbol:
		for _, capture := range sym.Captures {
			emitSymbolPush(emitter, capture)
		}

		captureN := len(sym.Captures)
		fn := uint32(sym.Fn)
		if sym.External {
			fn |= 1 << 31
		}
		emitter.Emit(
			runtime.OpMakeClosure,
			uint64(captureN)<<32|uint64(fn))
	default:
		panic("not implemented")
	}
}

func emitSymbolRefPush(emitter *runtime.Emitter, sym types.Symbol) {
	switch sym := sym.(type) {
	case *types.LocalVarSymbol:
		emitter.Emit(runtime.OpPushLocalRef, uint64(sym.StorageIndex()))

	//case *types.CaptureSymbol:
	//case *types.ParamSymbol:
	default:
		panic("not implemented")
	}
}
