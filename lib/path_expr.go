package lib

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dcaiafa/nitro"
)

type pathOp struct {
	index  nitro.Int
	member nitro.String
}

type PathExpr struct {
	fn  nitro.Callable
	ops []pathOp
}

func (e *PathExpr) Eval(m *nitro.Machine, v nitro.Value) (nitro.Value, error) {
	if e.fn != nil {
		res, err := m.Call(e.fn, []nitro.Value{v}, 1)
		if err != nil {
			return nil, err
		}
		return res[0], nil
	}
	return evalPathOp(v, e.ops)
}

func evalPathOp(v nitro.Value, ops []pathOp) (nitro.Value, error) {
	if len(ops) == 0 {
		return v, nil
	}

	indexable, ok := v.(nitro.Indexable)
	if !ok {
		return nil, nil
	}

	var err error
	var child nitro.Value
	if ops[0].member.String() != "" {
		child, err = indexable.Index(ops[0].member)
		if err != nil {
			return nil, nil
		}
	} else {
		child, err = indexable.Index(ops[0].index)
		if err != nil {
			return nil, nil
		}
	}

	return evalPathOp(child, ops[1:])
}

func ParsePathExpr(v nitro.Value) (*PathExpr, string, error) {
	callable, ok := v.(nitro.Callable)
	if ok {
		return &PathExpr{fn: callable}, "", nil
	}

	exprValue, ok := v.(nitro.String)
	if !ok {
		return nil, "", fmt.Errorf(
			"path expression must be a string or a function")
	}

	expr := exprValue.String()

	var ops []pathOp
	var extra string

	if len(expr) == 0 || expr[0] != '.' {
		return nil, "", fmt.Errorf("expression must start with '.'")
	}

	for len(expr) > 0 {
		if expr[0] == '[' {
			expr = expr[1:]
			var indexStr string
			var err error
			indexStr, expr, err = parsePathExprComponent(expr)
			if err != nil {
				return nil, "", err
			}
			if len(expr) == 0 || expr[0] != ']' {
				return nil, "", fmt.Errorf("missing ]")
			}
			expr = expr[1:]
			index, err := strconv.ParseInt(indexStr, 10, 64)
			if err != nil {
				return nil, "", fmt.Errorf("invalid array index")
			}
			ops = append(ops, pathOp{index: nitro.NewInt(index)})
		} else if expr[0] == '.' {
			expr = expr[1:]
			var member string
			var err error
			member, expr, err = parsePathExprComponent(expr)
			if err != nil {
				return nil, "", err
			}
			if len(member) == 0 {
				if len(ops) == 0 {
					continue
				} else {
					return nil, "", fmt.Errorf("member cannot be empty")
				}
			}
			ops = append(ops, pathOp{member: nitro.NewString(member)})
		} else if expr[0] == ',' {
			extra = expr[1:]
			break
		} else {
			return nil, "", fmt.Errorf("path expression term must start with '.', ',' or '['")
		}
	}

	return &PathExpr{ops: ops}, extra, nil
}

func parsePathExprComponent(expr string) (string, string, error) {
	buf := strings.Builder{}

Loop:
	for len(expr) > 0 {
		if expr[0] == '\\' {
			if len(expr) < 2 {
				return "", "", fmt.Errorf("invalid escape sequence")
			}
			switch expr[1] {
			case '\\', '.', ',', '[', ']':
				buf.WriteByte(expr[1])
			default:
				return "", "", fmt.Errorf("invalid escape sequence")
			}
			expr = expr[2:]
			continue
		}

		switch expr[0] {
		case '[', ']', '.', ',':
			break Loop
		}

		buf.WriteByte(expr[0])
		expr = expr[1:]
	}

	return buf.String(), expr, nil
}
