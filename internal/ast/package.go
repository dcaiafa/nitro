package ast

import (
	"github.com/dcaiafa/nitro/internal/symbol"
	"github.com/dcaiafa/nitro/internal/token"
)

type Package struct {
	PosImpl

	Units ASTs

	scope symbol.Scope

	init *FuncStmt
	main *FuncStmt
}

var _ Scope = (*Package)(nil)

func (p *Package) AddInitStmt(initStmt AST) {
	p.init.Block.Stmts = append(p.init.Block.Stmts, initStmt)
}

func (p *Package) Scope() symbol.Scope {
	return p.scope
}

func (p *Package) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Rewrite:
		p.synthesizeInit()
		p.synthesizeMain()

	case CreateGlobals:
		p.scope = symbol.NewScope()
	}

	ctx.RunPassChild(p, p.Units, pass)
	ctx.RunPassChild(p, p.init, pass)
	ctx.RunPassChild(p, p.main, pass)
}

func (p *Package) synthesizeInit() {
	p.init = &FuncStmt{
		Name: "$init",
		Func: Func{
			PosImpl: p.PosImpl,
			Block:   &StmtBlock{},
		},
	}
}

func (p *Package) synthesizeMain() {
	callInit := &FuncCallExpr{
		PosImpl: p.PosImpl,
		Target: &SimpleRef{
			PosImpl: p.PosImpl,
			ID: token.Token{
				Pos:  p.Pos(),
				Type: token.String,
				Str:  "$init",
			},
		},
	}

	callMain := &FuncCallExpr{
		PosImpl: p.PosImpl,
		Target: &SimpleRef{
			PosImpl: p.PosImpl,
			ID: token.Token{
				Pos:  p.Pos(),
				Type: token.String,
				Str:  "main",
			},
		},
	}

	p.main = &FuncStmt{
		Name: "$main",
		Func: Func{
			PosImpl: p.PosImpl,
			Block: &StmtBlock{
				Stmts: ASTs{callInit, callMain},
			},
		},
	}
}

func (p *Package) GetImports() []string {
	var imports []string
	for _, unit := range p.Units {
		for _, importAST := range unit.(*Unit).Imports {
			imports = append(imports, importAST.(*Import).Package)
		}
	}
	return imports
}
