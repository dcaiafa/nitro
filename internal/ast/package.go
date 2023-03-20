package ast

import (
	"github.com/dcaiafa/nitro/internal/meta"
	"github.com/dcaiafa/nitro/internal/scope"
	"github.com/dcaiafa/nitro/internal/symbol"
	"github.com/dcaiafa/nitro/internal/token"
)

type Package struct {
	PosImpl

	IsMain bool
	Units  ASTs

	scope    *scope.SimpleScope
	init     *FuncStmt
	main     *FuncStmt
	globals  int
	metadata *meta.Metadata
}

var _ Scope = (*Package)(nil)

func (p *Package) AddInitStmt(initStmt AST) {
	p.init.Block.Stmts = append(p.init.Block.Stmts, initStmt)
}

func (p *Package) Scope() scope.Scope {
	return p.scope
}

func (p *Package) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Rewrite:
		p.synthesizeInit(ctx)
		if p.IsMain {
			p.synthesizeMain()
		}

	case CreateGlobals:
		p.scope = scope.NewScope(scope.Package)

		p.metadata = new(meta.Metadata)

	case Emit:
		emitter := ctx.Emitter()
		emitter.SetGlobalCount(p.globals)
	}

	ctx.RunPassChild(p, p.Units, pass)
	ctx.RunPassChild(p, p.init, pass)
	ctx.RunPassChild(p, p.main, pass)
}

func (p *Package) synthesizeInit(ctx *Context) {
	initBlock := new(StmtBlock)
	p.init = &FuncStmt{
		Name: "$init",
		Func: Func{
			PosImpl: p.PosImpl,
			Block:   initBlock,
		},
	}

	if p.IsMain {
		for idx, imp := range ctx.Imports() {
			if idx == 0 {
				// Index 0 is "self".
				continue
			}
			literalIdx, ok := imp.Symbols["$init"]
			if !ok {
				continue
			}
			sym := new(symbol.LiteralSymbol)
			sym.SetName("$init")
			sym.SetReadOnly(true)
			sym.PackageIdx = idx
			sym.LiteralIdx = literalIdx

			initBlock.Stmts = append(initBlock.Stmts,
				&FuncCallExpr{
					PosImpl: p.PosImpl,
					Target: &PreResolvedReference{
						PosImpl: p.PosImpl,
						Symbol:  sym,
					},
				})
		}
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

func (m *Package) NewGlobal() *symbol.GlobalVarSymbol {
	g := &symbol.GlobalVarSymbol{}
	g.GlobalNdx = m.globals
	m.globals++
	return g
}

func (m *Package) AddGlobalParam(ctx *Context, name string, param *meta.Param, pos token.Pos) *symbol.GlobalVarSymbol {
	g := m.NewGlobal()
	g.SetName(name)
	g.SetPos(pos)

	if !ctx.GetScope(scope.Package).PutSymbol(ctx, g) {
		return nil
	}
	if !ctx.Emitter().AddGlobalParam(param.Name, g.GlobalNdx) {
		ctx.Failf(pos, "there is already a parameter named %q", name)
		return nil
	}
	m.metadata.Params = append(m.metadata.Params, param)
	return g
}

func (m *Package) Metadata() *meta.Metadata {
	return m.metadata
}
