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

		sym := ctx.FindSymbol(symName)
		if sym == nil {
			ctx.Failf(r.Pos(), "Symbol %q not found.", symName)
			return
		}

		r.sym = sym

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
		emitter.Emit(runtime.OpPushGlobal, uint16(sym.GlobalNdx), 0)

	case *types.LocalVarSymbol:
		emitter.Emit(runtime.OpPushLocal, uint16(sym.LocalNdx), 0)

	//case *types.CaptureSymbol:
	case *types.ParamSymbol:
		emitter.Emit(runtime.OpPushArg, uint16(sym.ParamNdx), 0)

	case *types.FuncSymbol:
		if sym.External {
			emitter.Emit(runtime.OpPushExternFn, uint16(sym.FnNdx), 0)
		} else {
			emitter.Emit(runtime.OpPushFn, uint16(sym.FnNdx), 0)
		}

	default:
		panic("not implemented")
	}
}

func emitSymbolRefPush(emitter *runtime.Emitter, sym types.Symbol) {
	switch sym := sym.(type) {
	case *types.GlobalVarSymbol:
		emitter.Emit(runtime.OpPushGlobalRef, uint16(sym.GlobalNdx), 0)

	case *types.LocalVarSymbol:
		emitter.Emit(runtime.OpPushLocalRef, uint16(sym.LocalNdx), 0)

	//case *types.CaptureSymbol:
	case *types.ParamSymbol:
		emitter.Emit(runtime.OpPushArgRef, uint16(sym.ParamNdx), 0)

	default:
		panic("not implemented")
	}
}
