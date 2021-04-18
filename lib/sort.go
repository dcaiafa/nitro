package lib

import (
	"fmt"
	gosort "sort"
	"strings"

	"github.com/dcaiafa/nitro"
)

type sortExpr struct {
	pathExpr *PathExpr
	desc     bool
}

type sorter struct {
	m     *nitro.Machine
	arr   *nitro.Array
	err   error
	less  nitro.Callable
	exprs []*sortExpr
}

func (s *sorter) Len() int {
	return s.arr.Len()
}

func (s *sorter) Less(i, j int) bool {
	a := s.arr.Get(i)
	b := s.arr.Get(j)

	if s.err != nil {
		return false
	}

	if len(s.exprs) > 0 {
		for i, expr := range s.exprs {
			a, _ := expr.pathExpr.Eval(a)
			if a == nil {
				return !expr.desc
			}
			b, _ := expr.pathExpr.Eval(b)
			if b == nil {
				return expr.desc
			}

			if i < len(s.exprs)-1 {
				areEq, err := evalCmpOp(nitro.BinEq, a, b)
				if err != nil {
					s.err = err
					return false
				}
				if areEq {
					continue
				}
			}

			op := nitro.BinLT
			if expr.desc {
				op = nitro.BinGT
			}
			res, err := evalCmpOp(op, a, b)
			if err != nil {
				s.err = err
				return false
			}
			return res
		}
	}

	if s.less != nil {
		res, err := s.m.Call(s.less, []nitro.Value{a, b}, 1)
		if err != nil {
			s.err = err
			return false
		}
		return res[0].(nitro.Bool).Bool()
	}

	if a == nil {
		return true
	}
	if b == nil {
		return false
	}

	res, err := nitro.EvalBinOp(nitro.BinLT, a, b)
	if err != nil {
		s.err = err
		return false
	}

	return res.(nitro.Bool).Bool()
}

func (s *sorter) Swap(i, j int) {
	t := s.arr.Get(i)
	s.arr.Put(i, s.arr.Get(j))
	s.arr.Put(j, t)
}

func sort(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	if len(args) < 1 {
		return nil, errNotEnoughArgs
	}

	arr, err := ToArray(m, args[0])
	if err != nil {
		return nil, err
	}

	s := &sorter{
		m:   m,
		arr: arr,
	}

	if len(args) >= 2 {
		switch arg2 := args[1].(type) {
		case nitro.Callable:
			s.less = arg2
			if len(args) != 2 {
				return nil, errTooManyArgs
			}
		case nitro.String:
			s.exprs = make([]*sortExpr, 0, len(args)-1)
			for _, arg := range args[1:] {
				argStr := arg.String()
				sortExpr := new(sortExpr)
				if strings.HasSuffix(argStr, ",a") {
					argStr = argStr[:len(argStr)-2]
				} else if strings.HasSuffix(argStr, ",d") {
					argStr = argStr[:len(argStr)-2]
					sortExpr.desc = true
				}
				sortExpr.pathExpr, err = ParsePathExpr(argStr)
				if err != nil {
					return nil, err
				}
				s.exprs = append(s.exprs, sortExpr)
			}
		}
	}

	gosort.Sort(s)

	if s.err != nil {
		return nil, s.err
	}

	return []nitro.Value{arr}, nil
}

func evalCmpOp(op nitro.BinOp, operand1, operand2 nitro.Value) (bool, error) {
	res, err := nitro.EvalBinOp(op, operand1, operand2)
	if err != nil {
		return false, err
	}
	boolRes, ok := res.(nitro.Bool)
	if !ok {
		return false, fmt.Errorf(
			"expected operation to return bool; returned %v instead",
			nitro.TypeName(res))
	}
	return boolRes.Bool(), nil
}
