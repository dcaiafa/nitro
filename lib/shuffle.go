package lib

import "github.com/dcaiafa/nitro"
import "math/rand"

func shuffle(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) < 1 {
		return nil, errNotEnoughArgs
	}

	arr, err := ToArray(m, args[0])
	if err != nil {
		return nil, err
	}

	rand.Shuffle(arr.Len(), func(i, j int) {
		t := arr.Get(i)
		arr.Put(i, arr.Get(j))
		arr.Put(j, t)
	})

	return []nitro.Value{arr}, nil
}
