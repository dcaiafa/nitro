package compiler

import (
	"errors"
	"fmt"
	"io/fs"
	"strings"

	"github.com/dcaiafa/nitro/internal/ast"
	"github.com/dcaiafa/nitro/internal/errlogger"
	"github.com/dcaiafa/nitro/internal/mod"
	"github.com/dcaiafa/nitro/internal/parser2"
	"github.com/dcaiafa/nitro/internal/symbol"
	"github.com/dcaiafa/nitro/internal/token"
	"github.com/dcaiafa/nitro/internal/vm"
)

var ErrPackageCircularReference = errors.New("package circular reference")

type Compiler struct {
	diag            bool
	funcRegistries  []vm.ExportRegistry
	errLogger       *errlogger.ErrLoggerWrapper
	packageResolver *mod.PackageResolver
	packages        map[string]*vm.CompiledPackage
}

func NewCompiler(errLogger errlogger.ErrLogger) *Compiler {
	return &Compiler{
		errLogger:       errlogger.NewErrLoggerBase(errLogger),
		packageResolver: mod.NewPackageResolver(),
		packages:        make(map[string]*vm.CompiledPackage),
	}
}

func (c *Compiler) SetDiag(diag bool) {
	c.diag = diag
}

func (c *Compiler) AddFuncRegistry(reg vm.ExportRegistry) {
	c.funcRegistries = append(c.funcRegistries, reg)
}

func (c *Compiler) RegisterModule(fs *mod.FS, moduleName string) {
	c.packageResolver.RegisterModule(fs, moduleName)
}

/*
func (c *Compiler) CompileSimple(
	module string,
	filename string,
	scriptData []byte,
	errLogger errlogger.ErrLogger,
) (*vm.Program, error) {
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
	pkg, err := c.compile(root, errLoggerWrapper)
	if err != nil {
		return nil, err
	}
	prog := &vm.Program{
		Packages: []*vm.CompiledPackage{pkg},
	}
	return prog, nil
}
*/

func (c *Compiler) compilePackage(packageName string) error {
	// Mark this package as being compiled to detect circular references.
	// We will put the real value below.
	c.packages[packageName] = nil

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

	root := &ast.Root{
		FuncRegistries: c.funcRegistries,
	}

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

		root.AddUnit(unit)
	}

	err = c.processImports(packageReader.Module(), root)
	if err != nil {
		return err
	}

	pkg, err := c.compile(root)
	if err != nil {
		return err
	}

	c.packages[packageName] = pkg

	return nil
}

func (c *Compiler) processImports(module string, root *ast.Root) error {
	for _, unit := range root.Package.Units {
		unit := unit.(*ast.Unit)

		for _, imp := range unit.Imports {
			imp := imp.(*ast.Import)

			if strings.HasPrefix(imp.Package, "//") {
				imp.ExpandedPackage = module + "/" + imp.Package[2:]
			} else {
				imp.ExpandedPackage = imp.Package
			}

			pkg, ok := c.packages[imp.ExpandedPackage]
			if !ok {
				err := c.compilePackage(imp.ExpandedPackage)
				if err != nil {
					return err
				}
			}
			if pkg == nil {
				return fmt.Errorf("%w: package %q", ErrPackageCircularReference, imp.ExpandedPackage)
			}
		}
	}

	return nil
}

func (c *Compiler) compile(root *ast.Root) (*vm.CompiledPackage, error) {
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

/*
func (c *Compiler) parseImportsPackage(
	packageReader mod.PackageReader,
	errLogger errlogger.ErrLogger,
) (map[string]bool, error) {
	imports := make(map[string]bool)
	errLoggerWrapper := errlogger.NewErrLoggerBase(errLogger)
	unitPaths, err := packageReader.ListUnits()
	if err != nil {
		errLoggerWrapper.Failf(
			token.Pos{}, "Failed to load package at %q: %v",
			packageReader.Path(), err)
		return nil, errLoggerWrapper.Error()
	}
	for _, unitPath := range unitPaths {
		unitData, err := packageReader.ReadUnit(unitPath)
		if err != nil {
			errLoggerWrapper.Failf(
				token.Pos{}, "Failed to load unit %q: %v",
				unitPath, err)
			return nil, err
		}
		prologue, err := parser2.ParsePrologue(
			unitPath, string(unitData), c.diag, errLoggerWrapper)
		if err != nil {
			return nil, err
		}
		for _, imp := range prologue.Imports {
			imports[imp.(*ast.Import).Package] = true
		}
	}
	return imports, nil
}
*/
