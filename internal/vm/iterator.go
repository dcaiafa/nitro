package vm

import "fmt"

type Iterator struct {
	fn         *Fn
	extFn      NativeFn
	captures   []ValueRef
	iterNRet   int
	tryCatches []tryCatch
	defers     []*Closure
	stack      []Value
	nlocals    int
	ip         int
	sp         int

	preAllocStack [stackSize]Value
}

func (e *Iterator) String() string { return "<Iterator>" }
func (e *Iterator) Type() string   { return "Iterator" }

func (e *Iterator) EvalBinOp(op BinOp, operand Value) (Value, error) {
	return nil, fmt.Errorf("iterator does not support this operation")
}

func (e *Iterator) EvalUnaryMinus() (Value, error) {
	return nil, fmt.Errorf("iterator does not support this operation")
}

func (e *Iterator) isCallable()   {}
func (e *Iterator) IterNRet() int { return e.iterNRet }

func NewIterator(extFn NativeFn, nret int) *Iterator {
	i := &Iterator{
		extFn:    extFn,
		iterNRet: nret,
	}
	return i
}
