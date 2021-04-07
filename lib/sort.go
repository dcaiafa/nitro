package lib

import (
	"sort"

	"github.com/dcaiafa/nitro"
)

type sorter struct {
	m    *nitro.Machine
	arr  *nitro.Array
	err  error
	less nitro.Value
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

	if s.less == nil {
		res, err := nitro.EvalBinOp(nitro.BinLT, a, b)
		if err != nil {
			s.err = err
			return false
		}
		return res.(nitro.Bool).Bool()
	}

	res, err := s.m.Call(s.less, []nitro.Value{a, b}, 1)
	if err != nil {
		s.err = err
		return false
	}

	return res[0].(nitro.Bool).Bool()
}

func (s *sorter) Swap(i, j int) {
	t := s.arr.Get(i)
	s.arr.Put(i, s.arr.Get(j))
	s.arr.Put(j, t)
}

func fnSort(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
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
		var err error
		s.less, err = getCallableArg(args, 1)
		if err != nil {
			return nil, err
		}
	}

	sort.Sort(s)

	if s.err != nil {
		return nil, s.err
	}

	return []nitro.Value{arr}, nil
}
