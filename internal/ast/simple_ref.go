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
			return
		}

		s.sym = sym

	case Emit:
		emit := emitSymbolPush
		_, isLValue := ctx.Parent().(*LValue)
		if isLValue {
			emit = emitSymbolRefPush
		}
		emit(ctx.Emitter(), s.sym)
	}
}

func emitSymbolPush(emitter *runtime.Emitter, sym types.Symbol) {
	switch sym := sym.(type) {
	case *types.LocalVarSymbol:
		emitter.Emit(runtime.OpPushLocal, uint16(sym.Local), 0)

	//case *types.CaptureSymbol:
	case *types.ParamSymbol:
		emitter.Emit(runtime.OpPushArg, uint16(sym.Arg), 0)

	case *types.FuncSymbol:
		for _, capture := range sym.Captures {
			emitSymbolPush(emitter, capture)
		}

		capN := byte(len(sym.Captures))
		fn := uint16(sym.Fn)
		if sym.External {
			fn |= 1 << 15
		}
		emitter.Emit(runtime.OpMakeClosure, fn, capN)

	default:
		panic("not implemented")
	}
}

func emitSymbolRefPush(emitter *runtime.Emitter, sym types.Symbol) {
	switch sym := sym.(type) {
	case *types.LocalVarSymbol:
		emitter.Emit(runtime.OpPushLocalRef, uint16(sym.Local), 0)

	//case *types.CaptureSymbol:
	case *types.ParamSymbol:
		emitter.Emit(runtime.OpPushArgRef, uint16(sym.Arg), 0)

	default:
		panic("not implemented")
	}
}
