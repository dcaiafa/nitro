package lib

import (
	"github.com/dcaiafa/nitro"
)

func match(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	str, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}

	regex, err := getRegexArg(args, 1)
	if err != nil {
		return nil, err
	}

	matches := regex.FindStringSubmatch(str)
	if matches == nil {
		return []nitro.Value{nitro.NewBool(false), nil}, nil
	}
	a := nitro.NewArrayFromSlice(make([]nitro.Value, 0, len(matches)))
	for _, match := range matches {
		a.Push(nitro.NewString(match))
	}
	return []nitro.Value{a}, nil
}
