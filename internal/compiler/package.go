package compiler

import (
	"path/filepath"
	"strings"

	"github.com/dcaiafa/nitro/internal/ast"
	"github.com/dcaiafa/nitro/internal/errlogger"
	"github.com/dcaiafa/nitro/internal/fs"
	"github.com/dcaiafa/nitro/internal/parser2"
	"github.com/dcaiafa/nitro/internal/scope"
	"github.com/dcaiafa/nitro/internal/symbol"
	"github.com/dcaiafa/nitro/internal/vm"
)

type packageGetter interface {
	getPackage(pkgName string) (*vm.CompiledPackage, int, error)
}

type packageCompiler struct {
	ModuleRoot    string
	ModuleName    string
	PackageName   string
	PackageGetter packageGetter
	PackagePath   string
	ErrLogger     *errlogger.ErrLoggerWrapper
	FS            fs.FS

	packageAST *ast.Package
	deps       []*vm.CompiledPackage
}

func (c *packageCompiler) parse() error {
	c.packageAST = new(ast.Package)

	unitFIs, err := c.FS.List(c.PackagePath)
	if err != nil {
		return err
	}

	for _, unitFI := range unitFIs {
		ext := filepath.Ext(unitFI.Name)
		if unitFI.IsDir || (ext != ".n" && ext != ".bg") {
			continue
		}
		unitPath := filepath.Join(c.PackagePath, unitFI.Name)
		unitData, err := c.FS.Read(unitPath)
		if err != nil {
			return err
		}
		unit, err := parser2.ParseUnit(
			unitPath, string(unitData), false, c.ErrLogger)
		if err != nil {
			return err
		}
		c.packageAST.Units = append(c.packageAST.Units, unit)
	}

	return nil
}

func (c *packageCompiler) processImports() error {
	deps := make(map[string]*vm.CompiledPackage)
	for _, unit := range c.packageAST.Units {
		unit := unit.(*ast.Unit)
		for _, imp := range unit.Imports {
			imp := imp.(*ast.Import)
			if strings.HasPrefix(imp.Package, "//") {
				imp.ExpandedPackage = c.ModuleName + "/" + imp.Package[2:]
			} else {
				imp.ExpandedPackage = imp.Package
			}
			if _, ok := deps[imp.ExpandedPackage]; !ok {
				pkg, _, err := c.PackageGetter.getPackage(imp.ExpandedPackage)
				if err != nil {
					return err
				}
				deps[imp.ExpandedPackage] = pkg
			}
		}
	}

	c.deps = make([]*vm.CompiledPackage, 0, len(deps))
	for _, pkg := range deps {
		c.deps = append(c.deps, pkg)
	}

	return nil
}

var passes = []ast.Pass{
	ast.Rewrite,
	ast.CreateGlobals,
	ast.Check,
	ast.Emit,
}

func (c *packageCompiler) Compile() (*vm.CompiledPackage, error) {
	err := c.parse()
	if err != nil {
		return nil, err
	}

	err = c.processImports()
	if err != nil {
		return nil, err
	}

	ctx := ast.NewContext(c.ErrLogger, c.deps)

	for _, pass := range passes {
		c.packageAST.RunPass(ctx, pass)
		// TODO: errlogger should just indicate that an error happened.
		// Then we return a generic "errors happened".
		if c.ErrLogger.Error() != nil {
			return nil, c.ErrLogger.Error()
		}
	}

	pkg := ctx.Emitter().ToCompiledPackage()
	pkg.Name = c.PackageName
	pkg.Metadata = c.packageAST.Metadata()

	mainFunc := c.packageAST.Scope().GetSymbol("$main").(*symbol.GlobalVarSymbol)
	pkg.MainFnNdx = mainFunc.GlobalNdx
	pkg.Deps = c.deps

	// TODO: ugly
	pkg.Symbols = make(map[string]int)
	c.packageAST.Scope().(*scope.SimpleScope).ForEachSymbol(func(sym symbol.Symbol) {
		global, ok := sym.(*symbol.GlobalVarSymbol)
		if !ok || !global.Export {
			return
		}
		pkg.Symbols[global.Name()] = global.GlobalNdx
	})

	return pkg, nil
}
