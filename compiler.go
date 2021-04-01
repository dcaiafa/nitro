package nitro

import (
	"context"
	"io/ioutil"

	"github.com/dcaiafa/nitro/internal/ast"
	"github.com/dcaiafa/nitro/internal/errlogger"
	"github.com/dcaiafa/nitro/internal/parser2"
	"github.com/dcaiafa/nitro/internal/runtime"
	"github.com/dcaiafa/nitro/internal/std"
	"github.com/dcaiafa/nitro/internal/token"
)

type FileSystem interface {
	ReadFile(name string) ([]byte, error)
}

type nativeFileSystem struct{}

func NewNativeFileSystem() FileSystem {
	return &nativeFileSystem{}
}

func (fs *nativeFileSystem) ReadFile(name string) ([]byte, error) {
	return ioutil.ReadFile(name)
}

type Compiler struct {
	fileSystem  FileSystem
	diag        bool
	externalFns map[string]runtime.ExternFn
}

func NewCompiler(fileSystem FileSystem) *Compiler {
	c := &Compiler{
		fileSystem:  fileSystem,
		externalFns: make(map[string]runtime.ExternFn),
	}
	std.Register(c)
	return c
}

func (c *Compiler) SetDiag(diag bool) {
	c.diag = diag
}

func (c *Compiler) AddExternalFn(name string, fn runtime.ExternFn) {
	c.externalFns[name] = fn
}

func (c *Compiler) Compile(filename string, errLogger errlogger.ErrLogger) (*runtime.Program, error) {
	errLoggerWrapper := errlogger.NewErrLoggerBase(errLogger)

	data, err := c.fileSystem.ReadFile(filename)
	if err != nil {
		errLoggerWrapper.Failf(token.Pos{}, "Failed to load %q: %v", filename, err)
		return nil, errLoggerWrapper.Error()
	}

	module, err := parser2.Parse(filename, string(data), c.diag, errLoggerWrapper)
	if err != nil {
		return nil, err
	}

	main := &ast.Main{}
	main.AddModule(module)

	for name, extFn := range c.externalFns {
		main.AddExternalFn(name, extFn)
	}

	ctx := ast.NewContext(errLoggerWrapper)
	main.RunPass(ctx, ast.Check)

	if errLoggerWrapper.Error() != nil {
		return nil, errLoggerWrapper.Error()
	}

	main.RunPass(ctx, ast.Emit)

	if errLoggerWrapper.Error() != nil {
		return nil, errLoggerWrapper.Error()
	}

	return ctx.Emitter().ToProgram(), nil
}

func Run(ctx context.Context, prog *Program, params map[string]Value) error {
	return runtime.NewMachine(prog).Run(ctx, params)
}
