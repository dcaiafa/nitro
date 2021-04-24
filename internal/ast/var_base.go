package ast

import (
	"github.com/dcaiafa/nitro/internal/vm"
	"github.com/dcaiafa/nitro/internal/symbol"
	"github.com/dcaiafa/nitro/internal/token"
)

func AddVariable(ctx *Context, name string, pos token.Pos) symbol.Symbol {
	return AddVariableToScope(ctx, ctx.CurrentScope(), name, pos)
}

func AddVariableToScope(ctx *Context, scope *symbol.Scope, name string, pos token.Pos) symbol.Symbol {
	fn := ctx.CurrentFunc()
	if fn != nil {
		l := fn.NewLocal()
		l.SetName(name)
		l.SetPos(pos)
		if !scope.PutSymbol(ctx, l) {
			return nil
		}
		return l
	}

	g := ctx.Main().NewGlobal()
	g.SetName(name)
	g.SetPos(pos)
	if !scope.PutSymbol(ctx, g) {
		return nil
	}

	return g
}

func emitSymbolPush(pos token.Pos, emitter *vm.Emitter, sym symbol.Symbol) {
	switch sym := sym.(type) {
	case *symbol.GlobalVarSymbol:
		emitter.Emit(pos, vm.OpLoadGlobal, uint32(sym.GlobalNdx), 0)

	case *symbol.LocalVarSymbol:
		emitter.Emit(pos, vm.OpLoadLocal, uint32(sym.LocalNdx), 0)

	case *symbol.CaptureSymbol:
		emitter.Emit(pos, vm.OpLoadCapture, uint32(sym.CaptureNdx), 0)

	case *symbol.ParamSymbol:
		emitter.Emit(pos, vm.OpLoadArg, uint32(sym.ParamNdx), 0)

	case *symbol.FuncSymbol:
		if sym.External {
			emitter.Emit(pos, vm.OpLoadNativeFn, uint32(sym.IdxFunc), 0)
		} else {
			emitter.Emit(pos, vm.OpLoadFn, uint32(sym.IdxFunc), 0)
		}

	default:
		panic("not implemented")
	}
}

func emitSymbolRefPush(pos token.Pos, emitter *vm.Emitter, sym symbol.Symbol) {
	switch sym := sym.(type) {
	case *symbol.GlobalVarSymbol:
		emitter.Emit(pos, vm.OpLoadGlobalRef, uint32(sym.GlobalNdx), 0)

	case *symbol.LocalVarSymbol:
		emitter.Emit(pos, vm.OpLoadLocalRef, uint32(sym.LocalNdx), 0)

	case *symbol.CaptureSymbol:
		emitter.Emit(pos, vm.OpLoadCaptureRef, uint32(sym.CaptureNdx), 0)

	case *symbol.ParamSymbol:
		emitter.Emit(pos, vm.OpLoadArgRef, uint32(sym.ParamNdx), 0)

	default:
		panic("not implemented")
	}
}

func emitSymbolCapture(pos token.Pos, emitter *vm.Emitter, sym symbol.Symbol) {
	switch sym := sym.(type) {
	case *symbol.LocalVarSymbol:
		emitter.Emit(pos, vm.OpCaptureLocal, uint32(sym.LocalNdx), 0)

	case *symbol.ParamSymbol:
		emitter.Emit(pos, vm.OpCaptureArg, uint32(sym.ParamNdx), 0)

	default:
		emitSymbolRefPush(pos, emitter, sym)
	}
}
