package compiler

import (
	"errors"
	"fmt"
	"io/fs"
	"strings"

	"github.com/dcaiafa/nitro/internal/ast"
	"github.com/dcaiafa/nitro/internal/errlogger"
	"github.com/dcaiafa/nitro/internal/export"
	"github.com/dcaiafa/nitro/internal/mod"
	"github.com/dcaiafa/nitro/internal/parser2"
	"github.com/dcaiafa/nitro/internal/scope"
	"github.com/dcaiafa/nitro/internal/symbol"
	"github.com/dcaiafa/nitro/internal/token"
	"github.com/dcaiafa/nitro/internal/vm"
)

var ErrPackageCircularReference = errors.New("package circular reference")

type Compiler struct {
	diag            bool
	exports         export.Exports
	errLogger       *errlogger.ErrLoggerWrapper
	packageResolver *mod.PackageResolver
	packageMap      map[string]int
	packages        []*vm.CompiledPackage
}

func NewCompiler(exports export.Exports, errLogger errlogger.ErrLogger) *Compiler {
	return &Compiler{
		exports:         exports,
		errLogger:       errlogger.NewErrLoggerBase(errLogger),
		packageResolver: mod.NewPackageResolver(),
		packageMap:      make(map[string]int),
	}
}

func (c *Compiler) SetDiag(diag bool) {
	c.diag = diag
}

func (c *Compiler) RegisterModule(fs *mod.FS, moduleName string) {
	c.packageResolver.RegisterModule(fs, moduleName)
}

func (c *Compiler) CompileSimple(module string, filename string, program []byte) (*vm.Program, error) {
	pkgNdx := len(c.packages)
	c.packageMap[filename] = pkgNdx
	c.packages = append(c.packages, nil)

	unit, err := parser2.ParseUnit(filename, string(program), c.diag, c.errLogger)
	if err != nil {
		return nil, err
	}

	ast.SimpleScriptToPackage(unit)

	pkgAST := &ast.Package{
		Units: ast.ASTs{unit},
	}

	pkg, err := c.compile(module, pkgAST)
	if err != nil {
		return nil, err
	}

	c.packages[pkgNdx] = pkg

	prog := c.makeProgram(filename)

	return prog, nil
}

func (c *Compiler) makeProgram(mainPkg string) *vm.Program {
	prog := &vm.Program{}
	prog.Packages = c.packages
	return prog
}

func (c *Compiler) compilePackage(packageName string) error {
	pkgNdx := len(c.packages)
	c.packageMap[packageName] = pkgNdx
	c.packages = append(c.packages, nil)

	packageReader, err := c.packageResolver.GetPackage(packageName)
	if err != nil {
		return err
	}

	unitPaths, err := packageReader.ListUnits()
	if err != nil {
		c.errLogger.Failf(
			token.Pos{}, "Failed to load package %q: %v",
			packageName, err)
		return c.errLogger.Error()
	}

	root := &ast.Package{}

	for _, unitPath := range unitPaths {
		expandedUnitPath := packageReader.ExpandPath(unitPath)
		unitData, err := fs.ReadFile(packageReader, unitPath)
		if err != nil {
			c.errLogger.Failf(
				token.Pos{}, "Failed to load unit %q: %v", expandedUnitPath, err)
			return err
		}
		unit, err := parser2.ParseUnit(
			expandedUnitPath, string(unitData), c.diag, c.errLogger)
		if err != nil {
			return err
		}

		root.Units = append(root.Units, unit)
	}

	pkg, err := c.compile(packageReader.Module(), root)
	if err != nil {
		return err
	}

	c.packages[pkgNdx] = pkg

	return nil
}

func (c *Compiler) processImports(module string, root *ast.Package) ([]vm.DepPackage, error) {
	depSet := make(map[string]int)

	for _, unit := range root.Units {
		unit := unit.(*ast.Unit)

		for _, imp := range unit.Imports {
			imp := imp.(*ast.Import)

			if strings.HasPrefix(imp.Package, "//") {
				imp.ExpandedPackage = module + "/" + imp.Package[2:]
			} else {
				imp.ExpandedPackage = imp.Package
			}

			pkgIdx, ok := c.packageMap[imp.ExpandedPackage]

			if !ok {
				if !c.compileNativePackage(imp.ExpandedPackage) {
					err := c.compilePackage(imp.ExpandedPackage)
					if err != nil {
						return nil, err
					}
				}
				pkgIdx, ok = c.packageMap[imp.ExpandedPackage]
				if !ok {
					panic("not reached")
				}
			}

			if c.packages[pkgIdx] == nil {
				return nil, fmt.Errorf("%w: package %q", ErrPackageCircularReference, imp.ExpandedPackage)
			}

			depSet[imp.ExpandedPackage] = pkgIdx
		}
	}

	deps := make([]vm.DepPackage, 0, len(depSet))
	for pkg, ndx := range depSet {
		deps = append(deps, vm.DepPackage{
			Name:        pkg,
			ResolvedNdx: ndx,
		})
	}

	return deps, nil
}

func (c *Compiler) compileNativePackage(packageName string) bool {
	var syms map[string]int
	var lits []vm.Value

	c.exports.GetPackageExports(packageName, func(name string, v vm.Value) {
		if syms == nil {
			syms = make(map[string]int)
		}
		syms[name] = len(lits)
		lits = append(lits, v)
	})

	if syms == nil {
		return false
	}

	cpkg := &vm.CompiledPackage{
		Literals: lits,
		Symbols:  syms,
	}

	pkgNdx := len(c.packages)
	c.packageMap[packageName] = pkgNdx
	c.packages = append(c.packages, cpkg)

	return true
}

func (c *Compiler) compile(module string, root *ast.Package) (*vm.CompiledPackage, error) {
	deps, err := c.processImports(module, root)
	if err != nil {
		return nil, err
	}

	ctx := ast.NewContext(c.errLogger)

	root.RunPass(ctx, ast.Rewrite)
	if c.errLogger.Error() != nil {
		return nil, c.errLogger.Error()
	}
	root.RunPass(ctx, ast.CreateGlobals)
	if c.errLogger.Error() != nil {
		return nil, c.errLogger.Error()
	}
	root.RunPass(ctx, ast.Check)
	if c.errLogger.Error() != nil {
		return nil, c.errLogger.Error()
	}
	root.RunPass(ctx, ast.Emit)
	if c.errLogger.Error() != nil {
		return nil, c.errLogger.Error()
	}

	pkg := ctx.Emitter().ToCompiledPackage()
	pkg.Metadata = root.Metadata()

	mainFunc := root.Scope().GetSymbol("$main").(*symbol.FuncSymbol)
	pkg.MainFnNdx = mainFunc.IdxFunc
	pkg.Symbols = make(map[string]int)
	pkg.DepPackages = deps

	root.Scope().(*scope.SimpleScope).ForEachSymbol(func(sym symbol.Symbol) {
		switch sym := sym.(type) {
		case *symbol.FuncSymbol:
			pkg.Symbols[sym.Name()] = sym.IdxFunc
		}
	})

	return pkg, nil
}
