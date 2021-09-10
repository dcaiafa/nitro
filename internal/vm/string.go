package vm

import (
	"errors"
	"fmt"
	"io"
	"strings"
	"unicode"
)

type String struct {
	v string
}

var (
	_ Readable  = String{}
	_ Indexable = String{}
	_ Iterable  = String{}
)

func NewString(v string) String { return String{v} }

func (s String) String() string { return s.v }
func (s String) Type() string   { return "string" }
func (s String) Len() int       { return len(s.v) }

func (s String) Index(key Value) (Value, error) {
	switch key := key.(type) {
	case String:
		switch key.String() {
		case "find":
			return NativeFn(s.find), nil
		case "findlast":
			return NativeFn(s.findlast), nil
		case "match":
			return NativeFn(s.match), nil
		case "matchall":
			return NativeFn(s.matchall), nil
		case "replace":
			return NativeFn(s.replace), nil
		case "split":
			return NativeFn(s.split), nil
		case "trim":
			return NativeFn(s.trim), nil
		case "trimleft":
			return NativeFn(s.trimleft), nil
		case "trimright":
			return NativeFn(s.trimright), nil
		case "trimprefix":
			return NativeFn(s.trimprefix), nil
		case "trimsuffix":
			return NativeFn(s.trimsuffix), nil
		case "toupper":
			return NativeFn(s.toupper), nil
		case "tolower":
			return NativeFn(s.tolower), nil
		case "hasprefix":
			return NativeFn(s.hasprefix), nil
		case "hassuffix":
			return NativeFn(s.hassuffix), nil
		case "fields":
			return NativeFn(s.fields), nil
		case "repeat":
			return NativeFn(s.repeat), nil
		default:
			return nil, fmt.Errorf("string does not have method %q", key.String())
		}

	case Int:
		idx := int(key.Int64())
		if idx < 0 {
			idx = len(s.v) + idx
		}
		if idx < 0 || idx >= len(s.v) {
			return nil, nil
		}
		return NewInt(int64(s.v[idx])), nil

	default:
		return nil, fmt.Errorf(
			"cannot index string using key type %v",
			TypeName(key))
	}
}

func (s String) IndexRef(key Value) (ValueRef, error) {
	return NewValueRef(nil), fmt.Errorf("cannot modify string")
}

func (s String) Slice(b, e Value) (Value, error) {
	bi, ok := b.(Int)
	ei, ok2 := e.(Int)
	if !ok || !ok2 {
		return nil, fmt.Errorf(
			"slice indices must be Int; instead they are %q and %q",
			TypeName(b), TypeName(e))
	}

	begin := int(bi.Int64())
	end := int(ei.Int64())

	if begin < 0 {
		return nil, fmt.Errorf(
			"invalid slice begin index %v; begin index must be >= 0",
			begin)
	}

	if end < 0 {
		end = len(s.v) + end
	}
	if end > len(s.v) {
		end = len(s.v)
	}
	if end < begin {
		begin = 0
		end = 0
	}

	return NewString(s.v[begin:end]), nil
}

func (s String) EvalOp(op Op, operand Value) (Value, error) {
	operandStr, ok := operand.(String)
	if !ok {
		return nil, fmt.Errorf(
			"invalid operation between string and %v",
			TypeName(operand))
	}

	switch op {
	case OpAdd:
		return NewString(s.v + operandStr.String()), nil
	case OpLT:
		return NewBool(s.v < operandStr.String()), nil
	case OpLE:
		return NewBool(s.v <= operandStr.String()), nil
	case OpGT:
		return NewBool(s.v > operandStr.String()), nil
	case OpGE:
		return NewBool(s.v >= operandStr.String()), nil
	case OpEq:
		return NewBool(s.v == operandStr.String()), nil
	default:
		return nil, fmt.Errorf("string does not support this operation")
	}
}

func (s String) EvalUnaryMinus() (Value, error) {
	return nil, fmt.Errorf("string does not support this operation")
}

func (s String) MakeIterator() Iterator {
	i := &stringIter{
		str:  s.v,
		next: 0,
	}
	return NewIterator(i.Next, nil, 2)
}

type stringIter struct {
	str  string
	next int
}

func (i *stringIter) Next(m *VM, args []Value, nret int) ([]Value, error) {
	if i.next >= len(i.str) {
		return nil, nil
	}

	idx := i.next
	i.next++

	v := NewInt(int64(i.str[idx]))

	return []Value{v, NewInt(int64(idx))}, nil
}

func (s String) MakeReader() Reader {
	return &stringReader{
		BaseValue: BaseValue{"stringreader"},
		Reader:    strings.NewReader(s.v),
	}
}

type stringReader struct {
	BaseValue
	io.Reader
}

var errStringFindUsage = errors.New(
	`invalid usage. Expected <string>.find(string)`)

func (s String) find(m *VM, args []Value, nRet int) ([]Value, error) {
	if len(args) != 1 {
		return nil, errStringFindUsage
	}
	sub, ok := args[0].(String)
	if !ok {
		return nil, errStringFindUsage
	}
	idx := strings.Index(s.v, sub.String())
	if idx == -1 {
		return []Value{nil}, nil
	}
	return []Value{NewInt(int64(idx))}, nil
}

var errStringFindLastUsage = errors.New(
	`invalid usage. Expected <string>.findlast(string)`)

func (s String) findlast(m *VM, args []Value, nRet int) ([]Value, error) {
	if len(args) != 1 {
		return nil, errStringFindLastUsage
	}
	sub, ok := args[0].(String)
	if !ok {
		return nil, errStringFindLastUsage
	}
	idx := strings.LastIndex(s.v, sub.String())
	if idx == -1 {
		return []Value{nil}, nil
	}
	return []Value{NewInt(int64(idx))}, nil
}

var errStringMatchUsage = errors.New(
	`invalid usage. Expected <string>.match(regex)`)

func (s String) match(m *VM, args []Value, nRet int) ([]Value, error) {
	if len(args) != 1 {
		return nil, errStringMatchUsage
	}

	regex, ok := args[0].(*Regex)
	if !ok {
		return nil, errStringMatchUsage
	}

	matches := regex.FindStringSubmatch(s.v)
	if matches == nil {
		return []Value{nil}, nil
	}

	a := NewArrayWithSlice(make([]Value, 0, len(matches)))
	for _, match := range matches {
		a.Add(NewString(match))
	}

	return []Value{a}, nil
}

var errStringMatchAllUsage = NewInvalidUsageError("<string>.matchall(regex)")

func (s String) matchall(m *VM, args []Value, nRet int) ([]Value, error) {
	if len(args) != 1 {
		return nil, errStringMatchAllUsage
	}

	regex, ok := args[0].(*Regex)
	if !ok {
		return nil, errStringMatchAllUsage
	}

	matches := regex.FindAllStringSubmatch(s.v, -1)
	if matches == nil {
		return []Value{nil}, nil
	}

	a := NewArrayWithSlice(make([]Value, 0, len(matches)))
	for _, match := range matches {
		a2 := NewArrayWithSlice(make([]Value, 0, len(match)))
		for _, submatch := range match {
			a2.Add(NewString(submatch))
		}
		a.Add(a2)
	}

	return []Value{a}, nil
}

var errStringReplaceUsage = NewInvalidUsageError(
	`<string>.replace(string, string, num?) or ` +
		`<string>.replace(regex, string|func)`)

func (s String) replace(m *VM, args []Value, nRet int) ([]Value, error) {
	if len(args) < 1 {
		return nil, errStringReplaceUsage
	}

	if oldArg, ok := args[0].(String); ok {
		if len(args) != 2 && len(args) != 3 {
			return nil, errStringReplaceUsage
		}
		newArg, ok := args[1].(String)
		if !ok {
			return nil, errStringReplaceUsage
		}
		n := -1
		if len(args) == 3 {
			argN, ok := args[2].(Int)
			if !ok {
				return nil, errStringReplaceUsage
			}
			n = int(argN.Int64())
		}

		res := strings.Replace(s.v, oldArg.String(), newArg.String(), n)
		return []Value{NewString(res)}, nil
	}

	if oldArg, ok := args[0].(*Regex); ok {
		if len(args) != 2 {
			return nil, errStringReplaceUsage
		}

		switch newArg := args[1].(type) {
		case String:
			res := oldArg.ReplaceAllString(s.v, newArg.String())
			return []Value{NewString(res)}, nil

		case Callable:
			var err error
			replFunc := func(s string) string {
				v, err2 := m.Call(newArg, []Value{NewString(s)}, 1)
				if err2 != nil {
					if err == nil {
						err = err2
					}
					return ""
				}
				resStr, ok := v[0].(String)
				if !ok {
					if err == nil {
						err = fmt.Errorf(
							"user provided func must return string; returned %q instead",
							TypeName(v[0]))
					}
					return ""
				}
				return resStr.String()
			}

			res := oldArg.ReplaceAllStringFunc(s.v, replFunc)
			if err != nil {
				return nil, err
			}

			return []Value{NewString(res)}, nil
		}
	}

	return nil, errStringReplaceUsage
}

var errStringSplitUsage = NewInvalidUsageError(
	"<string>.split(string|regex, num?)")

func (s String) split(m *VM, args []Value, nRet int) ([]Value, error) {
	if len(args) != 1 && len(args) != 2 {
		return nil, errStringSplitUsage
	}

	n := -1
	if len(args) == 2 {
		intArg, ok := args[1].(Int)
		if !ok {
			return nil, errStringSplitUsage
		}
		n = int(intArg.Int64())
	}

	var parts []string
	switch sep := args[0].(type) {
	case String:
		parts = strings.SplitN(s.v, sep.String(), n)

	case *Regex:
		parts = sep.Split(s.v, n)

	default:
		return nil, errStringSplitUsage
	}

	a := NewArrayWithSlice(make([]Value, 0, len(parts)))
	for _, part := range parts {
		a.Add(NewString(part))
	}

	return []Value{a}, nil
}

var errStringTrimUsage = NewInvalidUsageError("<string>.trim()")

func (s String) trim(m *VM, args []Value, nRet int) ([]Value, error) {
	if len(args) != 0 {
		return nil, errStringTrimUsage
	}
	res := strings.TrimSpace(s.v)
	return []Value{NewString(res)}, nil
}

var errStringTrimLeftUsage = NewInvalidUsageError("<string>.trimleft()")

func (s String) trimleft(m *VM, args []Value, nRet int) ([]Value, error) {
	if len(args) != 0 {
		return nil, errStringTrimLeftUsage
	}
	res := strings.TrimLeftFunc(s.v, unicode.IsSpace)
	return []Value{NewString(res)}, nil
}

var errStringTrimRightUsage = NewInvalidUsageError("<string>.trimright()")

func (s String) trimright(m *VM, args []Value, nRet int) ([]Value, error) {
	if len(args) != 0 {
		return nil, errStringTrimRightUsage
	}
	res := strings.TrimRightFunc(s.v, unicode.IsSpace)
	return []Value{NewString(res)}, nil
}

var errStringTrimPrefixUsage = NewInvalidUsageError("<string>.trimprefix(prefix)")

func (s String) trimprefix(m *VM, args []Value, nRet int) ([]Value, error) {
	if len(args) != 1 {
		return nil, errStringTrimPrefixUsage
	}
	prefix, ok := args[0].(String)
	if !ok {
		return nil, errStringTrimPrefixUsage
	}
	res := strings.TrimPrefix(s.v, prefix.String())
	return []Value{NewString(res)}, nil
}

var errStringTrimSuffixUsage = NewInvalidUsageError("<string>.trimsuffix(suffix)")

func (s String) trimsuffix(m *VM, args []Value, nRet int) ([]Value, error) {
	if len(args) != 1 {
		return nil, errStringTrimSuffixUsage
	}
	suffix, ok := args[0].(String)
	if !ok {
		return nil, errStringTrimSuffixUsage
	}
	res := strings.TrimSuffix(s.v, suffix.String())
	return []Value{NewString(res)}, nil
}

var errStringToUpperUsage = NewInvalidUsageError("<string>.toupper()")

func (s String) toupper(m *VM, args []Value, nRet int) ([]Value, error) {
	if len(args) != 0 {
		return nil, errStringToUpperUsage
	}
	res := strings.ToUpper(s.v)
	return []Value{NewString(res)}, nil
}

var errStringToLowerUsage = NewInvalidUsageError("<string>.tolower()")

func (s String) tolower(m *VM, args []Value, nRet int) ([]Value, error) {
	if len(args) != 0 {
		return nil, errStringToLowerUsage
	}
	res := strings.ToLower(s.v)
	return []Value{NewString(res)}, nil
}

var errStringHasPrefixUsage = NewInvalidUsageError("<string>.hasprefix(string)")

func (s String) hasprefix(m *VM, args []Value, nRet int) ([]Value, error) {
	if len(args) != 1 {
		return nil, errStringHasPrefixUsage
	}
	prefix, ok := args[0].(String)
	if !ok {
		return nil, errStringHasPrefixUsage
	}
	res := strings.HasPrefix(s.v, prefix.String())
	return []Value{NewBool(res)}, nil
}

var errStringHasSuffixUsage = NewInvalidUsageError("<string>.hassuffix(string)")

func (s String) hassuffix(m *VM, args []Value, nRet int) ([]Value, error) {
	if len(args) != 1 {
		return nil, errStringHasSuffixUsage
	}
	prefix, ok := args[0].(String)
	if !ok {
		return nil, errStringHasSuffixUsage
	}
	res := strings.HasSuffix(s.v, prefix.String())
	return []Value{NewBool(res)}, nil
}

var errStringFieldsUsage = NewInvalidUsageError("<string>.fields()")

func (s String) fields(m *VM, args []Value, nRet int) ([]Value, error) {
	if len(args) != 0 {
		return nil, errStringFieldsUsage
	}
	res := strings.Fields(s.v)
	resList := make([]Value, len(res))
	for i, v := range res {
		resList[i] = NewString(v)
	}
	return []Value{NewArrayWithSlice(resList)}, nil
}

var errStringRepeatUsage = NewInvalidUsageError("<string>.repeat(int)")

func (s String) repeat(m *VM, args []Value, nRet int) ([]Value, error) {
	if len(args) != 1 {
		return nil, errStringRepeatUsage
	}
	n, ok := args[0].(Int)
	if !ok {
		return nil, errStringRepeatUsage
	}
	res := strings.Repeat(s.v, int(n.Int64()))
	return []Value{NewString(res)}, nil
}
