package lib

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/dcaiafa/nitro"
)

func match(m *nitro.VM, caps []nitro.ValueRef, args []nitro.Value, nRet int) ([]nitro.Value, error) {
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

func split(m *nitro.VM, caps []nitro.ValueRef, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) < 2 {
		return nil, errNotEnoughArgs
	}

	str, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}

	n := -1
	if len(args) >= 3 {
		intArg, err := getIntArg(args, 2)
		if err != nil {
			return nil, err
		}
		n = int(intArg)
	}

	var parts []string
	switch sep := args[1].(type) {
	case nitro.String:
		parts = strings.SplitN(str, sep.String(), n)

	case *nitro.Regex:
		parts = sep.Split(str, n)

	default:
		return nil, fmt.Errorf(
			"Invalid arg#2 type %v. Expected String or Regex.",
			nitro.TypeName(sep))
	}

	a := nitro.NewArrayFromSlice(make([]nitro.Value, 0, len(parts)))
	for _, part := range parts {
		a.Push(nitro.NewString(part))
	}

	return []nitro.Value{a}, nil
}

func lines(m *nitro.VM, caps []nitro.ValueRef, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) < 1 {
		return nil, errNotEnoughArgs
	}
	input, err := ToReader(m, args[0])
	if err != nil {
		return nil, fmt.Errorf("invalid argument #1: %w", err)
	}

	l := &linesIter{
		input:   input,
		scanner: bufio.NewScanner(input),
	}

	outIter := nitro.NewIterator(l.Next, nil, 1)
	return []nitro.Value{outIter}, nil
}

type linesIter struct {
	input   io.Reader
	scanner *bufio.Scanner
	idxLine int
}

func (l *linesIter) Next(m *nitro.VM, caps []nitro.ValueRef, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if !l.scanner.Scan() {
		CloseReader(l.input)
		if l.scanner.Err() != nil {
			return nil, l.scanner.Err()
		}
		return iterDone(nRet)
	}
	l.idxLine++
	return []nitro.Value{
		nitro.NewBool(true),
		nitro.NewString(l.scanner.Text()),
		nitro.NewInt(int64(l.idxLine - 1)),
	}, nil
}

func trim(m *nitro.VM, caps []nitro.ValueRef, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	s, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}

	res := strings.TrimSpace(s)

	return []nitro.Value{nitro.NewString(res)}, nil
}

func hasprefix(m *nitro.VM, caps []nitro.ValueRef, args []nitro.Value, nRet int) ([]nitro.Value, error) {
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
