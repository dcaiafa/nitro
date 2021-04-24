package nitro

import (
	"io/ioutil"

	"github.com/dcaiafa/nitro/internal/ast"
	"github.com/dcaiafa/nitro/internal/errlogger"
	"github.com/dcaiafa/nitro/internal/parser2"
	"github.com/dcaiafa/nitro/internal/vm"
	"github.com/dcaiafa/nitro/internal/std"
	"github.com/dcaiafa/nitro/internal/token"
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
	fileLoader  FileLoader
	diag        bool
	main        *ast.Main
	externalFns map[string]vm.NativeFn
}

func NewCompiler(fileLoader FileLoader) *Compiler {
	c := &Compiler{
		fileLoader:  fileLoader,
		main:        &ast.Main{},
		externalFns: make(map[string]vm.NativeFn),
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

	module, err := parser2.ParseModule(filename, string(data), false, c.diag, errLoggerWrapper)
	if err != nil {
		return nil, err
	}

	return c.compile(module, errLoggerWrapper)
}

func (c *Compiler) CompileShort(
	shortProg string,
	errLogger errlogger.ErrLogger,
) (*vm.Program, error) {
	errLoggerWrapper := errlogger.NewErrLoggerBase(errLogger)

	module, err := parser2.ParseModule("expr", shortProg, true, c.diag, errLoggerWrapper)
	if err != nil {
		return nil, err
	}

	return c.compile(module, errLoggerWrapper)
}

func (c *Compiler) compile(
	module *ast.Module,
	errLoggerWrapper *errlogger.ErrLoggerWrapper,
) (*vm.Program, error) {
	c.main.AddModule(module)

	ctx := ast.NewContext(errLoggerWrapper)

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
