package compiler

import (
	"github.com/dcaiafa/nitro/internal/ast"
	"github.com/dcaiafa/nitro/internal/errlogger"
	"github.com/dcaiafa/nitro/internal/mod"
	"github.com/dcaiafa/nitro/internal/parser2"
	"github.com/dcaiafa/nitro/internal/symbol"
	"github.com/dcaiafa/nitro/internal/token"
	"github.com/dcaiafa/nitro/internal/vm"
)

type Compiler struct {
	diag           bool
	funcRegistries []vm.ExportRegistry
}

func NewCompiler() *Compiler {
	return &Compiler{}
}

func (c *Compiler) SetDiag(diag bool) {
	c.diag = diag
}

func (c *Compiler) AddFuncRegistry(reg vm.ExportRegistry) {
	c.funcRegistries = append(c.funcRegistries, reg)
}

func (c *Compiler) CompileSimple(
	filename string,
	scriptData []byte,
	errLogger errlogger.ErrLogger,
) (*vm.CompiledPackage, error) {
	errLoggerWrapper := errlogger.NewErrLoggerBase(errLogger)
	unit, err := parser2.ParseUnit(
		filename, string(scriptData), c.diag, errLoggerWrapper)
	if err != nil {
		return nil, err
	}
	ast.SimpleScriptToPackage(unit)
	root := &ast.Root{
		FuncRegistries: c.funcRegistries,
	}
	root.AddUnit(unit)
	return c.compilePackage(root, errLoggerWrapper)
}

func (c *Compiler) CompilePackage(
	packageReader mod.PackageReader,
	errLogger errlogger.ErrLogger,
) (*vm.CompiledPackage, error) {
	errLoggerWrapper := errlogger.NewErrLoggerBase(errLogger)
	unitPaths, err := packageReader.ListUnits()
	if err != nil {
		errLoggerWrapper.Failf(
			token.Pos{}, "Failed to load package at %q: %v",
			packageReader.Path(), err)
		return nil, errLoggerWrapper.Error()
	}
	root := &ast.Root{
		FuncRegistries: c.funcRegistries,
	}
	for _, unitPath := range unitPaths {
		unitData, err := packageReader.ReadUnit(unitPath)
		if err != nil {
			errLoggerWrapper.Failf(
				token.Pos{}, "Failed to load unit %q: %v",
				unitPath, err)
		}
		unit, err := parser2.ParseUnit(
			unitPath, string(unitData), c.diag, errLoggerWrapper)
		if err != nil {
			return nil, err
		}
		root.AddUnit(unit)
	}
	return c.compilePackage(root, errLoggerWrapper)
}

func (c *Compiler) compilePackage(
	root *ast.Root,
	errLoggerWrapper *errlogger.ErrLoggerWrapper,
) (*vm.CompiledPackage, error) {
	ctx := ast.NewContext(errLoggerWrapper)

	root.RunPass(ctx, ast.Rewrite)
	if errLoggerWrapper.Error() != nil {
		return nil, errLoggerWrapper.Error()
	}
	root.RunPass(ctx, ast.CreateGlobals)
	if errLoggerWrapper.Error() != nil {
		return nil, errLoggerWrapper.Error()
	}
	root.RunPass(ctx, ast.Check)
	if errLoggerWrapper.Error() != nil {
		return nil, errLoggerWrapper.Error()
	}
	root.RunPass(ctx, ast.Emit)
	if errLoggerWrapper.Error() != nil {
		return nil, errLoggerWrapper.Error()
	}

	pkg := ctx.Emitter().ToCompiledPackage()
	pkg.Metadata = root.Metadata()

	mainFunc := root.Package.Scope().GetSymbol("$main").(*symbol.FuncSymbol)
	pkg.MainFnNdx = mainFunc.IdxFunc
	pkg.Symbols = make(map[string]vm.PackageSymbol)

	root.Package.Scope().(*symbol.SimpleScope).ForEachSymbol(func(sym symbol.Symbol) {
		switch sym := sym.(type) {
		case *symbol.FuncSymbol:
			pkg.Symbols[sym.Name()] = vm.PackageSymbol{
				Type:  vm.PackageSymbolFunc,
				Index: uint32(sym.IdxFunc),
			}
		}
	})

	return pkg, nil
}
