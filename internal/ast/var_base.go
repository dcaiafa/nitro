package ast

import (
	"github.com/dcaiafa/nitro/internal/runtime"
	"github.com/dcaiafa/nitro/internal/symbol"
	"github.com/dcaiafa/nitro/internal/token"
)

func AddVariable(ctx *Context, name string, pos token.Pos) symbol.Symbol {
	fn := ctx.CurrentFunc()
	if fn != nil {
		l := fn.NewLocal()
		l.SetName(name)
		l.SetPos(pos)
		if !ctx.CurrentScope().PutSymbol(ctx, l) {
			return nil
		}
		return l
	}

	g := ctx.Main().NewGlobal()
	g.SetName(name)
	g.SetPos(pos)
	if !ctx.CurrentScope().PutSymbol(ctx, g) {
		return nil
	}

	return g
}

func emitSymbolPush(pos token.Pos, emitter *runtime.Emitter, sym symbol.Symbol) {
	switch sym := sym.(type) {
	case *symbol.GlobalVarSymbol:
		emitter.Emit(pos, runtime.OpLoadGlobal, uint16(sym.GlobalNdx), 0)

	case *symbol.LocalVarSymbol:
		emitter.Emit(pos, runtime.OpLoadLocal, uint16(sym.LocalNdx), 0)

	case *symbol.CaptureSymbol:
		emitter.Emit(pos, runtime.OpLoadCapture, uint16(sym.CaptureNdx), 0)

	case *symbol.ParamSymbol:
		emitter.Emit(pos, runtime.OpLoadArg, uint16(sym.ParamNdx), 0)

	case *symbol.FuncSymbol:
		if sym.External {
			emitter.Emit(pos, runtime.OpLoadExternFn, uint16(sym.IdxFunc), 0)
		} else {
			emitter.Emit(pos, runtime.OpLoadFn, uint16(sym.IdxFunc), 0)
		}

	default:
		panic("not implemented")
	}
}

func emitSymbolRefPush(pos token.Pos, emitter *runtime.Emitter, sym symbol.Symbol) {
	switch sym := sym.(type) {
	case *symbol.GlobalVarSymbol:
		emitter.Emit(pos, runtime.OpLoadGlobalRef, uint16(sym.GlobalNdx), 0)

	case *symbol.LocalVarSymbol:
		emitter.Emit(pos, runtime.OpLoadLocalRef, uint16(sym.LocalNdx), 0)

	case *symbol.CaptureSymbol:
		emitter.Emit(pos, runtime.OpLoadCaptureRef, uint16(sym.CaptureNdx), 0)

	case *symbol.ParamSymbol:
		emitter.Emit(pos, runtime.OpLoadArgRef, uint16(sym.ParamNdx), 0)

	default:
		panic("not implemented")
	}
}
