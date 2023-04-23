package str

import (
	"fmt"
	"strings"

	"github.com/dcaiafa/nitro"
	"github.com/dcaiafa/nitro/internal/vm"
)

//go:generate stubgen str.stubgen

func find0(m *vm.VM, s string, substr string) (vm.Value, error) {
	idx := strings.Index(s, substr)
	if idx == -1 {
		return nil, nil
	}
	return vm.NewInt(int64(idx)), nil
}

func find_last0(m *vm.VM, s string, substr string) (vm.Value, error) {
	idx := strings.LastIndex(s, substr)
	if idx == -1 {
		return nil, nil
	}
	return vm.NewInt(int64(idx)), nil
}

func match0(m *vm.VM, s string, regex *vm.Regex) (vm.Value, error) {
	matches := regex.FindStringSubmatch(s)
	if matches == nil {
		return nil, nil
	}

	res := vm.NewArrayWithSlice(make([]vm.Value, 0, len(matches)))
	for _, match := range matches {
		res.Add(vm.NewString(match))
	}

	return res, nil
}

func match_all0(m *vm.VM, s string, regex *vm.Regex) (vm.Value, error) {
	matches := regex.FindAllStringSubmatch(s, -1)
	if matches == nil {
		return nil, nil
	}

	res := vm.NewArrayWithSlice(make([]vm.Value, 0, len(matches)))
	for _, match := range matches {
		subs := vm.NewArrayWithSlice(make([]vm.Value, 0, len(match)))
		for _, submatch := range match {
			subs.Add(vm.NewString(submatch))
		}
		res.Add(subs)
	}

	return res, nil
}

func replace0(m *vm.VM, s string, old string, rep string, n int64) (string, error) {
	return strings.Replace(s, old, rep, int(n)), nil
}

func replace1(m *vm.VM, s string, old *vm.Regex, rep string) (string, error) {
	return old.ReplaceAllString(s, rep), nil
}

func replace2(m *vm.VM, s string, old *vm.Regex, rep vm.Callable) (string, error) {
	var err error
	replFunc := func(s string) string {
		v, err2 := m.Call(rep, []vm.Value{vm.NewString(s)}, 1)
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
	res := old.ReplaceAllStringFunc(s, replFunc)
	if err != nil {
		return "", err
	}
	return res, nil
}

func split0(m *vm.VM, s string, sep string, n int64) (vm.Value, error) {
	parts := strings.SplitN(s, sep, int(n))
	partList := vm.NewArrayWithSlice(make([]vm.Value, 0, len(parts)))
	for _, part := range parts {
		partList.Add(vm.NewString(part))
	}
	return partList, nil
}

func split1(m *vm.VM, s string, sep *vm.Regex, n int64) (vm.Value, error) {
	parts := sep.Split(s, int(n))
	partList := vm.NewArrayWithSlice(make([]vm.Value, 0, len(parts)))
	for _, part := range parts {
		partList.Add(vm.NewString(part))
	}
	return partList, nil
}

func trim_space0(m *vm.VM, s string) (string, error) {
	return strings.TrimSpace(s), nil
}

func trim_prefix0(m *vm.VM, s string, prefix string) (string, error) {
	return strings.TrimPrefix(s, prefix), nil
}

func trim_suffix0(m *vm.VM, s string, prefix string) (string, error) {
	return strings.TrimSuffix(s, prefix), nil
}

func to_upper0(m *vm.VM, s string) (string, error) {
	return strings.ToUpper(s), nil
}

func to_lower0(m *vm.VM, s string) (string, error) {
	return strings.ToLower(s), nil
}

func has_prefix0(m *vm.VM, s string, prefix string) (bool, error) {
	return strings.HasPrefix(s, prefix), nil
}

func has_suffix(m *vm.VM, s string, suffix string) (bool, error) {
	return strings.HasSuffix(s, suffix), nil
}

func fields0(m *vm.VM, s string) (vm.Value, error) {
	fields := strings.Fields(s)
	fieldList := make([]vm.Value, len(fields))
	for i, field := range fields {
		fieldList[i] = vm.NewString(field)
	}
	return vm.NewArrayWithSlice(fieldList), nil
}

func repeat0(m *vm.VM, s string, n int64) (string, error) {
	return strings.Repeat(s, int(n)), nil
}
