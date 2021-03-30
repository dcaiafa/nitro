package nitro

import (
	"context"

	"github.com/dcaiafa/nitro/internal/errlogger"
	"github.com/dcaiafa/nitro/internal/runtime"
	"github.com/dcaiafa/nitro/internal/token"
)

type (
	Value    = runtime.Value
	ValueRef = runtime.ValueRef
	String   = runtime.String
	Int      = runtime.Int
	Float    = runtime.Float
	Bool     = runtime.Bool
	Object   = runtime.Object
	Array    = runtime.Array
	ExternFn = runtime.ExternFn
	Program  = runtime.Program
	Frame    = runtime.FrameInfo
	Closure  = runtime.Closure

	RuntimeError = runtime.RuntimeError

	ErrLogger = errlogger.ErrLogger
	Pos       = token.Pos
)

var ErrCannotCallNil = runtime.ErrCannotCallNil

func NewClosure(extFn ExternFn, caps []ValueRef) *Closure {
	return runtime.NewClosure(extFn, caps)
}

func MakeEnumerator(ctx context.Context, v Value) (Value, error) {
	return runtime.MakeEnumerator(ctx, v)
}

func Next(ctx context.Context, e Value, n int) ([]Value, bool, error) {
	return runtime.Next(ctx, e, n)
}

func Call(ctx context.Context, callable Value, args []Value, retN int) ([]Value, error) {
	return runtime.Call(ctx, callable, args, retN)
}
