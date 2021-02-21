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

func (r *SimpleRef) isExpr() {}

func (r *SimpleRef) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Check:
		symName := r.ID.Str

		r.sym = ctx.FindSymbol(symName)
		if r.sym == nil {
			ctx.Failf(r.Pos(), "Symbol %q not found.", symName)
			return
		}

	case Emit:
		emit := emitSymbolPush
		_, isLValue := ctx.Parent().(*LValue)
		if isLValue {
			emit = emitSymbolRefPush
		}
		emit(ctx.Emitter(), r.sym)
	}
}

func emitSymbolPush(emitter *runtime.Emitter, sym types.Symbol) {
	switch sym := sym.(type) {
	case *types.GlobalVarSymbol:
		emitter.Emit(runtime.OpLoadGlobal, uint16(sym.GlobalNdx), 0)

	case *types.LocalVarSymbol:
		emitter.Emit(runtime.OpLoadLocal, uint16(sym.LocalNdx), 0)

	case *types.CaptureSymbol:
		emitter.Emit(runtime.OpLoadCapture, uint16(sym.CaptureNdx), 0)

	case *types.ParamSymbol:
		emitter.Emit(runtime.OpLoadArg, uint16(sym.ParamNdx), 0)

	case *types.FuncSymbol:
		if sym.External {
			emitter.Emit(runtime.OpLoadExternFn, uint16(sym.FnNdx), 0)
		} else {
			emitter.Emit(runtime.OpLoadFn, uint16(sym.FnNdx), 0)
		}

	default:
		panic("not implemented")
	}
}

func emitSymbolRefPush(emitter *runtime.Emitter, sym types.Symbol) {
	switch sym := sym.(type) {
	case *types.GlobalVarSymbol:
		emitter.Emit(runtime.OpLoadGlobalRef, uint16(sym.GlobalNdx), 0)

	case *types.LocalVarSymbol:
		emitter.Emit(runtime.OpLoadLocalRef, uint16(sym.LocalNdx), 0)

	case *types.CaptureSymbol:
		emitter.Emit(runtime.OpLoadCaptureRef, uint16(sym.CaptureNdx), 0)

	case *types.ParamSymbol:
		emitter.Emit(runtime.OpLoadArgRef, uint16(sym.ParamNdx), 0)

	default:
		panic("not implemented")
	}
}