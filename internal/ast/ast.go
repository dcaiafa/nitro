package ast

import "github.com/dcaiafa/go-expr/expr/internal/context"

type AST interface {
	context.PassRunner
	context.Printer
}
