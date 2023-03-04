package ast

import (
	"github.com/dcaiafa/nitro/internal/symbol"
	"github.com/dcaiafa/nitro/internal/token"
	"github.com/dcaiafa/nitro/internal/vm"
)

func emitVariableInit(ctx *Context, pos token.Pos, sym symbol.Symbol) {
	if sym.Lifted() {
		switch sym := sym.(type) {
		case *symbol.LocalVarSymbol:
			ctx.Emitter().Emit(pos, vm.OpInitLiftedLocal, uint32(sym.LocalNdx), 0)
		default:
			panic("unreachable")
		}
	} else if ctx.IsInRepeatableScope() {
		switch sym := sym.(type) {
		case *symbol.LocalVarSymbol:
			ctx.Emitter().Emit(pos, vm.OpInitLocal, uint32(sym.LocalNdx), 0)
		case *symbol.GlobalVarSymbol:
			ctx.Emitter().Emit(pos, vm.OpInitGlobal, uint32(sym.GlobalNdx), uint16(sym.PackageNdx))
		default:
			panic("unreachable")
		}
	}
}

func emitSymbolPush(pos token.Pos, emitter *vm.Emitter, sym symbol.Symbol) {
	switch sym := sym.(type) {
	case *symbol.GlobalVarSymbol:
		emitter.Emit(pos, vm.OpLoadGlobal, uint32(sym.GlobalNdx), uint16(sym.PackageNdx))

	case *symbol.LocalVarSymbol:
		if sym.Lifted() {
			emitter.Emit(pos, vm.OpLoadLocalDeref, uint32(sym.LocalNdx), 0)
		} else {
			emitter.Emit(pos, vm.OpLoadLocal, uint32(sym.LocalNdx), 0)
		}

	case *symbol.CaptureSymbol:
		emitter.Emit(pos, vm.OpLoadCapture, uint32(sym.CaptureNdx), 0)

	case *symbol.ParamSymbol:
		if sym.Lifted() {
			emitter.Emit(pos, vm.OpLoadArgDeref, uint32(sym.ParamNdx), 0)
		} else {
			emitter.Emit(pos, vm.OpLoadArg, uint32(sym.ParamNdx), 0)
		}

	default:
		panic("not implemented")
	}
}

func emitSymbolRefPush(pos token.Pos, emitter *vm.Emitter, sym symbol.Symbol) {
	switch sym := sym.(type) {
	case *symbol.GlobalVarSymbol:
		emitter.Emit(pos, vm.OpLoadGlobalRef, uint32(sym.GlobalNdx), uint16(sym.PackageNdx))

	case *symbol.LocalVarSymbol:
		if sym.Lifted() {
			emitter.Emit(pos, vm.OpLoadLocal, uint32(sym.LocalNdx), 0)
		} else {
			emitter.Emit(pos, vm.OpLoadLocalRef, uint32(sym.LocalNdx), 0)
		}

	case *symbol.CaptureSymbol:
		emitter.Emit(pos, vm.OpLoadCaptureRef, uint32(sym.CaptureNdx), 0)

	case *symbol.ParamSymbol:
		if sym.Lifted() {
			// TODO: fix name inconsistency: param vs arg
			emitter.Emit(pos, vm.OpLoadArg, uint32(sym.ParamNdx), 0)
		} else {
			emitter.Emit(pos, vm.OpLoadArgRef, uint32(sym.ParamNdx), 0)
		}

	default:
		panic("unreachable")
	}
}
