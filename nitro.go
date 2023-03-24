package nitro

import (
	"regexp"

	"github.com/dcaiafa/nitro/internal/compiler"
	"github.com/dcaiafa/nitro/internal/errlogger"
	"github.com/dcaiafa/nitro/internal/meta"
	"github.com/dcaiafa/nitro/internal/token"
	"github.com/dcaiafa/nitro/internal/vm"
)

type (
	VM = vm.VM

	Array        = vm.Array
	Bool         = vm.Bool
	Callable     = vm.Callable
	CloseFn      = vm.CloseFn
	Closure      = vm.Closure
	Float        = vm.Float
	Frame        = vm.FrameInfo
	FrameCrumb   = vm.FrameCrumb
	Func         = vm.Fn
	Indexable    = vm.Indexable
	Int          = vm.Int
	Iterable     = vm.Iterable
	Iterator     = vm.Iterator
	NativeFn     = vm.NativeFn
	Object       = vm.Object
	Program      = vm.Program
	Readable     = vm.Readable
	Reader       = vm.Reader
	Regex        = vm.Regex
	String       = vm.String
	Value        = vm.Value
	ValueRef     = vm.ValueRef
	RuntimeError = vm.RuntimeError

	Metadata = meta.Metadata

	Compiler = compiler.Compiler

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

var (
	True  = vm.True
	False = vm.False
)

var ErrCannotCallNil = vm.ErrCannotCallNil

func NewVM(p *Program) *VM {
	return vm.NewVM(p)
}

func IsIterable(v Value) bool {
	return vm.IsIterable(v)
}

func MakeIterator(m *VM, v Value) (Iterator, error) {
	return vm.MakeIterator(m, v)
}

func IsReadable(v Value) bool {
	return vm.IsReadable(v)
}

func MakeReader(m *VM, v Value) (Reader, error) {
	return vm.MakeReader(m, v)
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

func NewIterator(f func(m *VM, args []Value, nRet int) ([]Value, error), c CloseFn, nret int) Iterator {
	return vm.NewIterator(f, c, nret)
}

func NewNativeFn(f func(m *VM, args []Value, nRet int) ([]Value, error)) *NativeFn {
	return vm.NewNativeFn(f)
}
