package parser2

import (
	"strings"

	"github.com/dcaiafa/nitro/internal/ast"
	"github.com/dcaiafa/nitro/internal/token"
)

type execArgMaker struct {
	partial ast.Expr
	args    ast.ASTs
}

func (t *execArgMaker) Reserve(n int) {
	t.args = make(ast.ASTs, 0, n)
}

func (t *execArgMaker) AddArg(arg ast.Expr) {
	_, isSep := arg.(*ast.ExecWS)
	if !isSep && arg != nil {
		if t.partial == nil {
			t.partial = arg
		} else {
			t.partial = &ast.BinaryExpr{
				Left:  t.partial,
				Right: arg,
        Op: ast.BinOpPlus,
			}
		}
	} else if t.partial != nil {
		elem := &ast.ArrayElement{
			Val: t.partial,
		}
		t.args = append(t.args, elem)
		t.partial = nil
	}
}

func (t *execArgMaker) ArrayLiteral() *ast.ArrayLiteral {
	return &ast.ArrayLiteral{
		Block: &ast.ArrayElementBlock{
			Elements: t.args,
		},
	}
}

func (t *execArgMaker) fix(arg ast.Expr) (isFinal bool) {
	if lit, ok := arg.(*ast.LiteralExpr); ok {
		if lit.Val.Type != token.String {
			panic("exec arg literal is not string")
		}
		fixed := strings.TrimSpace(lit.Val.Str)
		if len(fixed) != len(lit.Val.Str) {
			isFinal = true
			lit.Val.Str = fixed
		}
	}
	return isFinal
}
