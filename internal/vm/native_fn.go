package vm

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

type NativeFn struct {
	fn func(m *VM, args []Value, nRet int) ([]Value, error)
}

func NewNativeFn(fn func(m *VM, args []Value, nRet int) ([]Value, error)) *NativeFn {
	return &NativeFn{fn: fn}
}

func (f *NativeFn) String() string { return "<func>" }
func (f *NativeFn) Type() string   { return "Func" }

func (f *NativeFn) Name() string {
	fn := runtime.FuncForPC(reflect.ValueOf(f.fn).Pointer())
	fnName := fn.Name()
	lastSlash := strings.LastIndexByte(fnName, '/')
	if lastSlash != -1 {
		fnName = fnName[lastSlash+1:]
	}
	return fnName
}

func (f *NativeFn) EvalOp(op Op, operand Value) (Value, error) {
	if op == OpEq {
		return NewBool(f == operand), nil
	}
	return nil, fmt.Errorf("func does not support this operation")
}

func (f *NativeFn) Call(m *VM, args []Value, nRet int) ([]Value, error) {
	return f.fn(m, args, nRet)
}
