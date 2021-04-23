package runtime

import "fmt"

type Iterator struct {
	fn         *Fn
	extFn      NativeFn
	captures   []ValueRef
	iterNRet   int
	tryCatches []tryCatch
	defers     []*Closure
	ip         int
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

func NewIterator(extFn NativeFn, caps []ValueRef, nret int) *Iterator {
	return &Iterator{
		extFn:    extFn,
		captures: caps,
		iterNRet: nret,
	}
}
