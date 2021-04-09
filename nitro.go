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
	Iterator = runtime.Iterator
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

func NewMachine(ctx context.Context, p *Program) *Machine {
	return runtime.NewMachine(ctx, p)
}

func MakeIterator(m *Machine, v Value) (Value, error) {
	return runtime.MakeIterator(m, v)
}

func Next(m *Machine, e Value, n int) ([]Value, bool, error) {
	return runtime.Next(m, e, n)
}

func NewConsoleErrLogger() ErrLogger {
	return &errlogger.ConsoleErrLogger{}
}

func EvalBinOp(op BinOp, operand1, operand2 Value) (Value, error) {
	return runtime.EvalBinOp(op, operand1, operand2)
}

func TypeName(v Value) string                          { return runtime.TypeName(v) }
func NewString(v string) String                        { return runtime.NewString(v) }
func NewInt(v int64) Int                               { return runtime.NewInt(v) }
func NewFloat(v float64) Float                         { return runtime.NewFloat(v) }
func NewBool(v bool) Bool                              { return runtime.NewBool(v) }
func NewValueRef(ref *Value) ValueRef                  { return runtime.NewValueRef(ref) }
func NewClosure(f ExternFn, c []ValueRef) *Closure     { return runtime.NewClosure(f, c) }
func NewIterator(f ExternFn, c []ValueRef) *Iterator { return runtime.NewIterator(f, c) }
func NewArray() *Array                                 { return runtime.NewArray() }
func NewArrayFromSlice(s []Value) *Array               { return runtime.NewArrayWithSlice(s) }
func NewObject() *Object                               { return runtime.NewObject() }
func NewRegex(r *regexp.Regexp) *Regex                 { return runtime.NewRegex(r) }
func CoerceToBool(v Value) bool                        { return runtime.CoerceToBool(v) }
