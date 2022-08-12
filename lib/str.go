package lib

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/dcaiafa/nitro"
	"github.com/dcaiafa/nitro/internal/vm"
)

func strFind(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) > 2 {
		return nil, errTooManyArgs
	}

	s, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}

	sub, err := getStringArg(args, 1)
	if err != nil {
		return nil, err
	}

	idx := strings.Index(s, sub)
	if idx == -1 {
		return []nitro.Value{nil}, nil
	}

	return []nitro.Value{nitro.NewInt(int64(idx))}, nil
}

func strFindLast(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) > 2 {
		return nil, errTooManyArgs
	}

	s, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}

	sub, err := getStringArg(args, 1)
	if err != nil {
		return nil, err
	}

	idx := strings.LastIndex(s, sub)
	if idx == -1 {
		return []nitro.Value{nil}, nil
	}

	return []nitro.Value{nitro.NewInt(int64(idx))}, nil
}

func strMatch(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) > 2 {
		return nil, errTooManyArgs
	}

	s, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}

	regex, err := getRegexArg(args, 1)
	if err != nil {
		return nil, err
	}

	matches := regex.FindStringSubmatch(s)
	if matches == nil {
		return []nitro.Value{nil}, nil
	}

	a := vm.NewArrayWithSlice(make([]nitro.Value, 0, len(matches)))
	for _, match := range matches {
		a.Add(nitro.NewString(match))
	}

	return []nitro.Value{a}, nil
}

func strMatchAll(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) > 2 {
		return nil, errTooManyArgs
	}

	s, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}

	regex, err := getRegexArg(args, 1)
	if err != nil {
		return nil, err
	}

	matches := regex.FindAllStringSubmatch(s, -1)
	if matches == nil {
		return []nitro.Value{nil}, nil
	}

	a := vm.NewArrayWithSlice(make([]nitro.Value, 0, len(matches)))
	for _, match := range matches {
		a2 := vm.NewArrayWithSlice(make([]nitro.Value, 0, len(match)))
		for _, submatch := range match {
			a2.Add(nitro.NewString(submatch))
		}
		a.Add(a2)
	}

	return []nitro.Value{a}, nil
}

func strReplace(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	s, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}

	if oldArg, err := getStringArg(args, 1); err == nil {
		if len(args) > 4 {
			return nil, errTooManyArgs
		}
		newArg, err := getStringArg(args, 2)
		if err != nil {
			return nil, err
		}
		var n int64 = -1
		if len(args) == 4 {
			n, err = getIntArg(args, 3)
			if err != nil {
				return nil, err
			}
		}
		res := strings.Replace(s, oldArg, newArg, int(n))
		return []nitro.Value{nitro.NewString(res)}, nil
	} else if oldArg, err := getRegexArg(args, 1); err == nil {
		if len(args) < 3 {
			return nil, errNotEnoughArgs
		} else if len(args) > 3 {
			return nil, errTooManyArgs
		}

		if newArg, err := getStringArg(args, 2); err == nil {
			res := oldArg.ReplaceAllString(s, newArg)
			return []nitro.Value{nitro.NewString(res)}, nil
		} else if replFn, err := getCallableArg(args, 2); err == nil {
			var err error
			replFunc := func(s string) string {
				v, err2 := m.Call(replFn, []nitro.Value{nitro.NewString(s)}, 1)
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
							"replace function must return string; returned %q instead",
							nitro.TypeName(v[0]))
					}
					return ""
				}
				return resStr.String()
			}

			res := oldArg.ReplaceAllStringFunc(s, replFunc)
			if err != nil {
				return nil, err
			}
			return []nitro.Value{nitro.NewString(res)}, nil
		} else {
			return nil, fmt.Errorf(
				"expected argument 3 to be str or callable, but it was %v",
				nitro.TypeName(args[2]))
		}
	} else {
		return nil, fmt.Errorf(
			"expected argument 2 to be str or regex, but it was %v",
			nitro.TypeName(args[1]))
	}
}

func strSplit(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	s, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}

	var n int64 = -1
	if len(args) == 3 {
		n, err = getIntArg(args, 2)
		if err != nil {
			return nil, err
		}
	}

	var parts []string
	switch sep := args[1].(type) {
	case nitro.String:
		parts = strings.SplitN(s, sep.String(), int(n))
	case *nitro.Regex:
		parts = sep.Split(s, int(n))
	default:
		return nil, errExpectedArg2(1, "str", "regex", args[1])
	}

	a := vm.NewArrayWithSlice(make([]nitro.Value, 0, len(parts)))
	for _, part := range parts {
		a.Add(nitro.NewString(part))
	}

	return []nitro.Value{a}, nil
}

func strTrim(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) > 1 {
		return nil, errTooManyArgs
	}
	s, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}
	res := strings.TrimSpace(s)
	return []nitro.Value{nitro.NewString(res)}, nil
}

func strTrimLeft(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) > 1 {
		return nil, errTooManyArgs
	}
	s, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}
	res := strings.TrimLeftFunc(s, unicode.IsSpace)
	return []nitro.Value{nitro.NewString(res)}, nil
}

func strTrimRight(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) > 1 {
		return nil, errTooManyArgs
	}
	s, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}
	res := strings.TrimRightFunc(s, unicode.IsSpace)
	return []nitro.Value{nitro.NewString(res)}, nil
}

func strTrimPrefix(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) > 2 {
		return nil, errTooManyArgs
	}
	s, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}
	prefix, err := getStringArg(args, 1)
	if err != nil {
		return nil, err
	}
	res := strings.TrimPrefix(s, prefix)
	return []nitro.Value{nitro.NewString(res)}, nil
}

func strTrimSuffix(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) > 2 {
		return nil, errTooManyArgs
	}
	s, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}
	suffix, err := getStringArg(args, 1)
	if err != nil {
		return nil, err
	}
	res := strings.TrimSuffix(s, suffix)
	return []nitro.Value{nitro.NewString(res)}, nil
}

func strToUpper(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) > 1 {
		return nil, errTooManyArgs
	}
	s, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}
	res := strings.ToUpper(s)
	return []nitro.Value{nitro.NewString(res)}, nil
}

func strToLower(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) > 1 {
		return nil, errTooManyArgs
	}
	s, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}
	res := strings.ToLower(s)
	return []nitro.Value{nitro.NewString(res)}, nil
}

func strHasPrefix(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) > 2 {
		return nil, errTooManyArgs
	}
	s, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}
	prefix, err := getStringArg(args, 1)
	if err != nil {
		return nil, err
	}
	res := strings.HasPrefix(s, prefix)
	return []nitro.Value{nitro.NewBool(res)}, nil
}

func strHasSuffix(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) > 2 {
		return nil, errTooManyArgs
	}
	s, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}
	suffix, err := getStringArg(args, 1)
	if err != nil {
		return nil, err
	}
	res := strings.HasSuffix(s, suffix)
	return []nitro.Value{nitro.NewBool(res)}, nil
}

func strFields(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) > 1 {
		return nil, errTooManyArgs
	}
	s, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}
	res := strings.Fields(s)
	resList := make([]nitro.Value, len(res))
	for i, v := range res {
		resList[i] = nitro.NewString(v)
	}
	return []nitro.Value{vm.NewArrayWithSlice(resList)}, nil
}

func strRepeat(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) > 2 {
		return nil, errTooManyArgs
	}
	s, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}
	n, err := getIntArg(args, 1)
	if err != nil {
		return nil, err
	}
	res := strings.Repeat(s, int(n))
	return []nitro.Value{nitro.NewString(res)}, nil
}
