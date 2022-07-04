package nitro

import (
	"os"
	"path/filepath"

	"github.com/dcaiafa/nitro/internal/ast"
	"github.com/dcaiafa/nitro/internal/errlogger"
	"github.com/dcaiafa/nitro/internal/parser2"
	"github.com/dcaiafa/nitro/internal/std"
	"github.com/dcaiafa/nitro/internal/symbol"
	"github.com/dcaiafa/nitro/internal/token"
	"github.com/dcaiafa/nitro/internal/vm"
)

type PackageReader interface {
	Path() string
	ListUnits() ([]string, error)
	ReadUnit(unit string) ([]byte, error)
}

type NativePackageReader struct {
	path string
}

func NewNativePackageReader(path string) *NativePackageReader {
	return &NativePackageReader{
		path: path,
	}
}

func (r *NativePackageReader) Path() string {
	return r.path
}

func (r *NativePackageReader) ListUnits() ([]string, error) {
	allFiles, err := os.ReadDir(r.path)
	if err != nil {
		return nil, err
	}
	units := make([]string, 0, len(allFiles))
	for _, file := range allFiles {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".n" {
			units = append(units, filepath.Join(r.path, file.Name()))
		}
	}
	return units, nil
}

func (r *NativePackageReader) ReadUnit(unit string) ([]byte, error) {
	unitData, err := os.ReadFile(filepath.Join(r.path, unit))
	if err != nil {
		return nil, err
	}
	return unitData, nil
}

type Compiler struct {
	moduleReg *ast.ModuleRegistry
	diag      bool
	main      *ast.Root
}

func NewCompiler() *Compiler {
	c := &Compiler{
		moduleReg: ast.NewModuleRegistry(),
		main:      &ast.Root{},
	}
	std.Register(c)
	return c
}

func (c *Compiler) SetDiag(diag bool) {
	c.diag = diag
}

func (c *Compiler) AddNativeFn(name string, fn func(m *VM, args []Value, nRet int) ([]Value, error)) {
	c.main.AddNativeFn(name, vm.NewNativeFn(fn))
}

func (c *Compiler) RegisterNativeModuleLoader(
	name string,
	regFn func(r NativeModuleContext)) {
	c.moduleReg.RegisterNativeModuleLoader(name, regFn)
}

func (c *Compiler) CompileSimple(
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
	c.main.AddUnit(unit)
	return c.compilePackage(errLoggerWrapper)
}

func (c *Compiler) CompilePackage(
	packageReader PackageReader,
	errLogger errlogger.ErrLogger,
) (*vm.Program, error) {
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
		}
		unit, err := parser2.ParseUnit(
			unitPath, string(unitData), c.diag, errLoggerWrapper)
		if err != nil {
			return nil, err
		}
		c.main.AddUnit(unit)
	}
	return c.compilePackage(errLoggerWrapper)
}

func (c *Compiler) compilePackage(
	errLoggerWrapper *errlogger.ErrLoggerWrapper,
) (*vm.Program, error) {
	ctx := ast.NewContext(c.moduleReg, errLoggerWrapper)

	c.main.RunPass(ctx, ast.Rewrite)
	if errLoggerWrapper.Error() != nil {
		return nil, errLoggerWrapper.Error()
	}

	c.main.RunPass(ctx, ast.Check)
	if errLoggerWrapper.Error() != nil {
		return nil, errLoggerWrapper.Error()
	}

	c.main.RunPass(ctx, ast.Emit)
	if errLoggerWrapper.Error() != nil {
		return nil, errLoggerWrapper.Error()
	}

	mainFunc := c.main.Package.Scope().GetSymbol("$main").(*symbol.FuncSymbol)

	program := ctx.Emitter().ToProgram()
	program.Metadata = c.main.Metadata()
	program.MainFnNdx = mainFunc.IdxFunc

	return program, nil
}
