package lib

import (
	"fmt"
	"strings"

	"github.com/dcaiafa/nitro"
)

func replace(m *nitro.VM, caps []nitro.ValueRef, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) < 3 {
		return nil, errNotEnoughArgs
	}
	str, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}

	if arg2Str, ok := args[1].(nitro.String); ok {
		arg3Str, err := getStringArg(args, 2)
		if err != nil {
			return nil, err
		}

		n := -1
		if len(args) >= 4 {
			arg4, err := getIntArg(args, 3)
			if err != nil {
				return nil, err
			}
			n = int(arg4)
		}

		res := strings.Replace(str, arg2Str.String(), arg3Str, n)
		return []nitro.Value{nitro.NewString(res)}, nil
	}

	if arg2Regex, ok := args[1].(*nitro.Regex); ok {
		if arg3Str, ok := args[2].(nitro.String); ok {
			res := arg2Regex.Regexp.ReplaceAllString(str, arg3Str.String())
			return []nitro.Value{nitro.NewString(res)}, nil
		}
		if arg3Fn, ok := args[2].(nitro.Callable); ok {
			err = nil
			replFunc := func(s string) string {
				v, err2 := m.Call(arg3Fn, []nitro.Value{nitro.NewString(s)}, 1)
				if err2 != nil {
					if err == nil {
						err = err2
					}
					return ""
				}
				resStr, ok := v[0].(nitro.String)
				if !ok {
					if err == nil {
						err = fmt.Errorf(
							"user provided func must return string; returned %q instead",
							nitro.TypeName(v[0]))
					}
					return ""
				}
				return resStr.String()
			}

			res := arg2Regex.ReplaceAllStringFunc(str, replFunc)
			return []nitro.Value{nitro.NewString(res)}, nil
		}
	}

	return nil, fmt.Errorf(
		"allowed parameter combinations are: String, String, String; "+
			"String, String, String, Int; String, Regex, String; "+
			"or String, Regex, Func. Actual: %v, %v, %v"+
			nitro.TypeName(args[0]), nitro.TypeName(args[1]), nitro.TypeName(args[2]))
}
