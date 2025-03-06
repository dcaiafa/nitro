package lib

import (
	"reflect"

	"github.com/dcaiafa/nitro"
)

func deepEqual(vm *nitro.VM, args []nitro.Value, nret int) ([]nitro.Value, error) {
	if len(args) < 2 {
		return nil, errNotEnoughArgs
	}

	ctx := &deepEqualContext{
		visiting: make(map[nitro.Value]bool),
	}

	equal := ctx.Equal(args[0], args[1])

	return []nitro.Value{nitro.NewBool(equal)}, nil
}

type deepEqualContext struct {
	visiting map[nitro.Value]bool
}

func (e *deepEqualContext) Equal(a, b nitro.Value) bool {
	if a == b {
		return true
	}

	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		return false
	}

	switch a := a.(type) {
	case *nitro.Array:
		return e.arrayEqual(a, b.(*nitro.Array))

	case *nitro.Object:
		return e.mapEqual(a, b.(*nitro.Object))

	default:
		res, err := nitro.EvalOp(nitro.OpEq, a, b)
		return err == nil && res == nitro.True
	}
}

func (e *deepEqualContext) arrayEqual(a, b *nitro.Array) bool {
	if a.Len() != b.Len() {
		return false
	}

	for i := 0; i < a.Len(); i++ {
		va := a.Get(i)
		vb := b.Get(i)

		if !e.Equal(va, vb) {
			return false
		}
	}

	return true
}

func (e *deepEqualContext) mapEqual(a, b *nitro.Object) bool {
	if a.Len() != b.Len() {
		return false
	}

	if e.visiting[a] {
		// Cycle detected.
		return false
	}

	e.visiting[a] = true
	defer delete(e.visiting, a)

	isEqual := true
	a.ForEach(func(k, va nitro.Value) bool {
		vb, ok := b.Get(k)
		if !ok {
			isEqual = false
			return false
		}
		if !e.Equal(va, vb) {
			isEqual = false
			return false
		}
		return true
	})

	return isEqual
}
