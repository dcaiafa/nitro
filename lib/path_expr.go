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
	ops []pathOp
}

func (e *PathExpr) Eval(v nitro.Value) (nitro.Value, bool) {
	return evalPathOp(v, e.ops)
}

func evalPathOp(v nitro.Value, ops []pathOp) (nitro.Value, bool) {
	if len(ops) == 0 {
		return v, true
	}

	indexable, ok := v.(nitro.Indexable)
	if !ok {
		return nil, false
	}

	var err error
	var child nitro.Value
	if ops[0].member.String() != "" {
		child, err = indexable.Index(ops[0].member)
		if err != nil {
			return nil, false
		}
	} else {
		child, err = indexable.Index(ops[0].index)
		if err != nil {
			return nil, false
		}
	}

	return evalPathOp(child, ops[1:])
}

func ParsePathExpr(expr string) (*PathExpr, error) {
	var ops []pathOp

	if len(expr) == 0 || expr[0] != '.' {
		return nil, fmt.Errorf("expression must start with '.'")
	}

	if len(expr) == 1 || expr[1] == '[' {
		// Special case. Ignore the first '.' when its the only thing in the
		// expression or when the first term is an array index.
		expr = expr[1:]
	}

	for len(expr) > 0 {
		if expr[0] == '[' {
			expr = expr[1:]
			end := strings.IndexByte(expr, ']')
			if end == 0 {
				return nil, fmt.Errorf("missing ]")
			}
			indexStr := expr[:end]
			index, err := strconv.ParseInt(indexStr, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid array index")
			}
			expr = expr[end+1:]
			ops = append(ops, pathOp{index: nitro.NewInt(index)})
		} else if expr[0] == '.' {
			expr = expr[1:]
			end := 0
			for ; end < len(expr); end++ {
				if expr[end] == '[' || expr[end] == '.' {
					break
				}
			}
			member := expr[:end]
			if len(member) == 0 {
				return nil, fmt.Errorf("member name cannot be empty")
			}
			expr = expr[end:]
			ops = append(ops, pathOp{member: nitro.NewString(member)})
		} else {
			return nil, fmt.Errorf("path expression term must start with . or [")
		}
	}

	return &PathExpr{ops: ops}, nil
}
