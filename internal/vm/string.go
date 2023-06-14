package vm

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

type String struct {
	v string
}

var (
	_ Readable  = String{}
	_ Indexable = String{}
)

func NewString(v string) String { return String{v} }

func (s String) String() string { return s.v }
func (s String) Type() string   { return "str" }
func (s String) Traits() Traits { return TraitEq }
func (s String) Len() int       { return len(s.v) }

func (s String) Index(key Value) (Value, error) {
	switch key := key.(type) {
	case Int:
		idx := int(key.Int64())
		if idx < 0 {
			idx = len(s.v) + idx
		}
		if idx < 0 || idx >= len(s.v) {
			return nil, nil
		}
		return NewInt(int64(s.v[idx])), nil

	default:
		return nil, fmt.Errorf(
			"cannot index str using key of type %v",
			TypeName(key))
	}
}

func (s String) IndexRef(key Value) (ValueRef, error) {
	return NewValueRef(nil), fmt.Errorf("cannot modify str")
}

func (s String) Slice(b, e Value) (Value, error) {
	bi, ok := b.(Int)
	ei, ok2 := e.(Int)
	if !ok || !ok2 {
		return nil, fmt.Errorf(
			"slice indices must be Int; instead they are %q and %q",
			TypeName(b), TypeName(e))
	}

	begin := int(bi.Int64())
	end := int(ei.Int64())

	if begin < 0 {
		return nil, fmt.Errorf(
			"invalid slice begin index %v; begin index must be >= 0",
			begin)
	}

	if end < 0 {
		end = len(s.v) + end
	}
	if end > len(s.v) {
		end = len(s.v)
	}
	if end < begin {
		begin = 0
		end = 0
	}

	return NewString(s.v[begin:end]), nil
}

func (s String) EvalOp(op Op, operand Value) (Value, error) {
	operandStr, ok := operand.(String)
	if !ok {
		if op == OpEq {
			return NewBool(false), nil
		}
		return nil, ErrOperationNotSupported
	}

	switch op {
	case OpAdd:
		return NewString(s.v + operandStr.String()), nil
	case OpLT:
		return NewBool(s.v < operandStr.String()), nil
	case OpLE:
		return NewBool(s.v <= operandStr.String()), nil
	case OpGT:
		return NewBool(s.v > operandStr.String()), nil
	case OpGE:
		return NewBool(s.v >= operandStr.String()), nil
	case OpEq:
		return NewBool(s.v == operandStr.String()), nil
	default:
		return nil, ErrOperationNotSupported
	}
}

func (s String) ToInt() (Int, error) {
	res, err := strconv.ParseInt(s.v, 0, 64)
	if err != nil {
		return NewInt(0), err
	}
	return NewInt(res), nil
}

func (s String) MakeReader() Reader {
	return &stringReader{
		Reader: strings.NewReader(s.v),
	}
}

type stringReader struct {
	io.Reader
}

func (r *stringReader) String() string { return "<reader>" }
func (r *stringReader) Type() string   { return "reader" }
func (r *stringReader) Traits() Traits { return TraitNone }
