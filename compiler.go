package nitro

import (
	"io/ioutil"

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
	main       *ast.Main
}

func NewCompiler(fileLoader FileLoader) *Compiler {
	c := &Compiler{
		fileLoader: fileLoader,
		moduleReg:  ast.NewModuleRegistry(),
		main:       &ast.Main{},
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

func (c *Compiler) Compile(
	filename string,
	errLogger errlogger.ErrLogger,
) (*vm.Program, error) {
	errLoggerWrapper := errlogger.NewErrLoggerBase(errLogger)

	data, err := c.fileLoader.LoadFile(filename)
	if err != nil {
		errLoggerWrapper.Failf(token.Pos{}, "Failed to load %q: %v", filename, err)
		return nil, errLoggerWrapper.Error()
	}

	module, err := parser2.ParseUnit(filename, string(data), c.diag, errLoggerWrapper)
	if err != nil {
		return nil, err
	}

	return c.compile(module, errLoggerWrapper)
}

func (c *Compiler) CompileInline(
	inline string,
	errLogger errlogger.ErrLogger,
) (*vm.Program, error) {
	errLoggerWrapper := errlogger.NewErrLoggerBase(errLogger)

	module, err := parser2.ParseUnit("<inline>", inline, c.diag, errLoggerWrapper)
	if err != nil {
		return nil, err
	}

	return c.compile(module, errLoggerWrapper)
}

func (c *Compiler) compile(
	module *ast.Unit,
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
