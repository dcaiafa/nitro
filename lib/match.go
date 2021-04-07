package lib

import (
	"github.com/dcaiafa/nitro"
)

func fnMatch(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	str, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}

	regex, err := getRegexArg(args, 1)
	if err != nil {
		return nil, err
	}

	if retN == 1 {
		ok := regex.MatchString(str)
		return []nitro.Value{nitro.NewBool(ok)}, nil
	} else {
		matches := regex.FindStringSubmatch(str)
		if matches == nil {
			return []nitro.Value{nitro.NewBool(false), nil}, nil
		}
		a := nitro.NewArrayFromSlice(make([]nitro.Value, 0, len(matches)))
		for _, match := range matches {
			a.Push(nitro.NewString(match))
		}
		return []nitro.Value{nitro.NewBool(true), a}, nil
	}
}
