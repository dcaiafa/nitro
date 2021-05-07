package nitro

import (
	"context"
	"regexp"

	"github.com/dcaiafa/nitro/internal/errlogger"
	"github.com/dcaiafa/nitro/internal/meta"
	"github.com/dcaiafa/nitro/internal/token"
	"github.com/dcaiafa/nitro/internal/vm"
)

type (
	VM = vm.VM

	Array     = vm.Array
	Bool      = vm.Bool
	Closure   = vm.Closure
	Iterator  = vm.Iterator
	Iterable  = vm.Iterable
	ExternFn  = vm.NativeFn
	Float     = vm.Float
	Frame     = vm.FrameInfo
	Func      = vm.Fn
	Int       = vm.Int
	Object    = vm.Object
	Program   = vm.Program
	Regex     = vm.Regex
	String    = vm.String
	Value     = vm.Value
	ValueRef  = vm.ValueRef
	Callable  = vm.Callable
	Indexable = vm.Indexable
	NativeFn  = vm.NativeFn
	Metadata  = meta.Metadata

	RuntimeError = vm.RuntimeError

	ErrLogger = errlogger.ErrLogger
	Pos       = token.Pos
)

type Op = vm.Op

const (
	OpUMinus = vm.OpUMinus
	OpAdd    = vm.OpAdd
	OpSub    = vm.OpSub
	OpMult   = vm.OpMult
	OpDiv    = vm.OpDiv
	OpMod    = vm.OpMod
	OpLT     = vm.OpLT
	OpLE     = vm.OpLE
	OpGT     = vm.OpGT
	OpGE     = vm.OpGE
	OpEq     = vm.OpEq
	OpNE     = vm.OpNE
)

var ErrCannotCallNil = vm.ErrCannotCallNil

func NewVM(ctx context.Context, p *Program) *VM {
	return vm.NewVM(ctx, p)
}

func MakeIterator(m *VM, v Value) (*Iterator, error) {
	return vm.MakeIterator(m, v)
}

func Next(m *VM, e Value, n int) ([]Value, bool, error) {
	return vm.Next(m, e, n)
}

func NewConsoleErrLogger() ErrLogger {
	return &errlogger.ConsoleErrLogger{}
}

func EvalOp(op Op, operand1, operand2 Value) (Value, error) {
	return vm.EvalOp(op, operand1, operand2)
}

func TypeName(v Value) string            { return vm.TypeName(v) }
func NewString(v string) String          { return vm.NewString(v) }
func NewInt(v int64) Int                 { return vm.NewInt(v) }
func NewFloat(v float64) Float           { return vm.NewFloat(v) }
func NewBool(v bool) Bool                { return vm.NewBool(v) }
func NewValueRef(ref *Value) ValueRef    { return vm.NewValueRef(ref) }
func NewArray() *Array                   { return vm.NewArray() }
func NewArrayFromSlice(s []Value) *Array { return vm.NewArrayWithSlice(s) }
func NewObject() *Object                 { return vm.NewObject() }
func NewRegex(r *regexp.Regexp) *Regex   { return vm.NewRegex(r) }
func CoerceToBool(v Value) bool          { return vm.CoerceToBool(v) }

func NewIterator(f ExternFn, nret int) *Iterator {
	return vm.NewIterator(f, nret)
}
