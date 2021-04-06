package nitro

import (
	"context"
	"regexp"

	"github.com/dcaiafa/nitro/internal/errlogger"
	"github.com/dcaiafa/nitro/internal/runtime"
	"github.com/dcaiafa/nitro/internal/token"
)

type (
	Machine = runtime.Machine

	Array    = runtime.Array
	Bool     = runtime.Bool
	Closure  = runtime.Closure
	ExternFn = runtime.ExternFn
	Float    = runtime.Float
	Frame    = runtime.FrameInfo
	Func     = runtime.Fn
	Int      = runtime.Int
	Object   = runtime.Object
	Program  = runtime.Program
	Regex    = runtime.Regex
	String   = runtime.String
	Value    = runtime.Value
	ValueRef = runtime.ValueRef

	RuntimeError = runtime.RuntimeError

	ErrLogger = errlogger.ErrLogger
	Pos       = token.Pos
)

type BinOp = runtime.BinOp

const (
	BinAdd  = runtime.BinAdd
	BinSub  = runtime.BinSub
	BinMult = runtime.BinMult
	BinDiv  = runtime.BinDiv
	BinMod  = runtime.BinMod
	BinLT   = runtime.BinLT
	BinLE   = runtime.BinLE
	BinGT   = runtime.BinGT
	BinGE   = runtime.BinGE
	BinEq   = runtime.BinEq
	BinNE   = runtime.BinNE
)

var ErrCannotCallNil = runtime.ErrCannotCallNil

func NewMachine(p *Program) *Machine {
	return runtime.NewMachine(p)
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

func NewConsoleErrLogger() ErrLogger {
	return &errlogger.ConsoleErrLogger{}
}

func SetUserData(ctx context.Context, k, v interface{}) {
	runtime.MachineFromContext(ctx).SetUserData(k, v)
}

func GetUserData(ctx context.Context, k interface{}) interface{} {
	return runtime.MachineFromContext(ctx).GetUserData(k)
}

func EvalBinOp(op BinOp, operand1, operand2 Value) (Value, error) {
	return runtime.EvalBinOp(op, operand1, operand2)
}

func TypeName(v Value) string                      { return runtime.TypeName(v) }
func NewString(v string) String                    { return runtime.NewString(v) }
func NewInt(v int64) Int                           { return runtime.NewInt(v) }
func NewFloat(v float64) Float                     { return runtime.NewFloat(v) }
func NewBool(v bool) Bool                          { return runtime.NewBool(v) }
func NewValueRef(ref *Value) ValueRef              { return runtime.NewValueRef(ref) }
func NewClosure(f ExternFn, c []ValueRef) *Closure { return runtime.NewClosure(f, c) }
func NewArray() *Array                             { return runtime.NewArray() }
func NewArrayFromSlice(s []Value) *Array           { return runtime.NewArrayWithSlice(s) }
func NewObject() *Object                           { return runtime.NewObject() }
func NewRegex(r *regexp.Regexp) *Regex             { return runtime.NewRegex(r) }
func CoerceToBool(v Value) bool                    { return runtime.CoerceToBool(v) }
