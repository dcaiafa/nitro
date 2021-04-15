package lib

import (
	"strings"

	"github.com/dcaiafa/nitro"
)

func trim(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, expRetN int) ([]nitro.Value, error) {
	s, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}

	res := strings.TrimSpace(s)

	return []nitro.Value{nitro.NewString(res)}, nil
}

func hasprefix(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, expRetN int) ([]nitro.Value, error) {
	str, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}
	prefix, err := getStringArg(args, 1)
	if err != nil {
		return nil, err
	}

	res := strings.HasPrefix(str, prefix)

	return []nitro.Value{nitro.NewBool(res)}, nil
}
