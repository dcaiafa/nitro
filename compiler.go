package nitro

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/dcaiafa/nitro/internal/ast"
	"github.com/dcaiafa/nitro/internal/errlogger"
	"github.com/dcaiafa/nitro/internal/parser2"
	"github.com/dcaiafa/nitro/internal/std"
	"github.com/dcaiafa/nitro/internal/token"
	"github.com/dcaiafa/nitro/internal/vm"
)

type FileLoader interface {
	LoadFile(name string) ([]byte, error)
}

type nativeFileLoader struct{}

func NewNativeFileLoader() FileLoader {
	return &nativeFileLoader{}
}

func (fs *nativeFileLoader) LoadFile(name string) ([]byte, error) {
	return ioutil.ReadFile(name)
}

type Compiler struct {
	fileLoader FileLoader
	moduleReg  *ast.ModuleRegistry
	diag       bool
	main       *ast.Root
}

func NewCompiler(fileLoader FileLoader) *Compiler {
	c := &Compiler{
		fileLoader: fileLoader,
		moduleReg:  ast.NewModuleRegistry(),
		main:       &ast.Root{},
	}
	std.Register(c)
	return c
}

func (c *Compiler) SetDiag(diag bool) {
	c.diag = diag
}

func (c *Compiler) AddNativeFn(name string, fn vm.NativeFn) {
	c.main.AddNativeFn(name, fn)
}

func (c *Compiler) RegisterNativeModuleLoader(
	name string,
	regFn func(r NativeModuleContext)) {
	c.moduleReg.RegisterNativeModuleLoader(name, regFn)
}

func (c *Compiler) Compile(
	filename string,
	errLogger errlogger.ErrLogger,
) (*vm.Program, error) {
	errLoggerWrapper := errlogger.NewErrLoggerBase(errLogger)

	unit, err := c.parseUnit(filename, errLoggerWrapper)
	if err != nil {
		return nil, err
	}

	pkg := &ast.Module{
		Units: []*ast.Unit{unit},
	}

	return c.compile(pkg, errLoggerWrapper)
}

func (c *Compiler) CompileInline(
	inline string,
	errLogger errlogger.ErrLogger,
) (*vm.Program, error) {
	errLoggerWrapper := errlogger.NewErrLoggerBase(errLogger)

	unit, err := parser2.ParseUnit("<inline>", inline, c.diag, errLoggerWrapper)
	if err != nil {
		return nil, err
	}

	pkg := &ast.Module{
		Units: []*ast.Unit{unit},
	}

	return c.compile(pkg, errLoggerWrapper)
}

func (c *Compiler) parseUnit(
	filename string,
	errLogger *errlogger.ErrLoggerWrapper,
) (*ast.Unit, error) {
	data, err := c.fileLoader.LoadFile(filename)
	if err != nil {
		errLogger.Failf(token.Pos{}, "Failed to load %q: %v", filename, err)
		return nil, errLogger.Error()
	}

	unit, err := parser2.ParseUnit(filename, string(data), c.diag, errLogger)
	if err != nil {
		return nil, err
	}

	return unit, nil
}

func (c *Compiler) compileModule(
	moduleName string,
	moduleDir string,
	errLogger *errlogger.ErrLoggerWrapper,
) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	unitFiles, err := filepath.Glob(filepath.Join(moduleDir, "*.n"))
	if err != nil {
		return err
	}

	moduleAST := new(ast.Module)

	for _, unitFile := range unitFiles {
		unitFile, err = filepath.Rel(wd, unitFile)
		if err != nil {
			return err
		}
		unitAST, err := c.parseUnit(unitFile, errLogger)
		if err != nil {
			return err
		}
		moduleAST.Units = append(moduleAST.Units, unitAST)
	}

	return nil

}

func (c *Compiler) compile(
	module *ast.Module,
	errLoggerWrapper *errlogger.ErrLoggerWrapper,
) (*vm.Program, error) {
	c.main.AddModule(module)

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

	program := ctx.Emitter().ToProgram()
	program.Metadata = c.main.Metadata()

	return program, nil
}
