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
	Func     = runtime.Fn

	RuntimeError = runtime.RuntimeError

	ErrLogger = errlogger.ErrLogger
	Pos       = token.Pos
)

var ErrCannotCallNil = runtime.ErrCannotCallNil

func MakeEnumerator(ctx context.Context, v Value) (Value, error) {
	return runtime.MakeEnumerator(ctx, v)
}

func Next(ctx context.Context, e Value, n int) ([]Value, bool, error) {
	return runtime.Next(ctx, e, n)
}

func Call(ctx context.Context, callable Value, args []Value, retN int) ([]Value, error) {
	return runtime.Call(ctx, callable, args, retN)
}

func NewConsoleErrLogger() ErrLogger {
	return &errlogger.ConsoleErrLogger{}
}

func TypeName(v Value) string                      { return runtime.TypeName(v) }
func NewString(v string) String                    { return runtime.NewString(v) }
func NewInt(v int64) Int                           { return runtime.NewInt(v) }
func NewFloat(v float64) Float                     { return runtime.NewFloat(v) }
func NewBool(v bool) Bool                          { return runtime.NewBool(v) }
func NewValueRef(ref *Value) ValueRef              { return runtime.NewValueRef(ref) }
func NewClosure(f ExternFn, c []ValueRef) *Closure { return runtime.NewClosure(f, c) }
func NewArray() *Array                             { return runtime.NewArray() }
func NewObject() *Object                           { return runtime.NewObject() }
