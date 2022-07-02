package parser2

import "github.com/dcaiafa/nitro/internal/ast"

// Prologue contains a unit's Meta and Imports. It is not an AST, which is why
// it is declared in the parser package.
type Prologue struct {
	Meta    ast.ASTs
	Imports ast.ASTs
}
