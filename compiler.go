package nitro

import (
	"context"
	"fmt"

	"github.com/dcaiafa/nitro/internal/ast"
	"github.com/dcaiafa/nitro/internal/errlogger"
	"github.com/dcaiafa/nitro/internal/parser2"
	"github.com/dcaiafa/nitro/internal/runtime"
	"github.com/dcaiafa/nitro/internal/std"
)

type FileSystem interface {
	ReadFile(name string) ([]byte, error)
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

func (c *Compiler) Compile(filename string) (*runtime.Program, error) {
	var err error

	errLogger := &errlogger.ConsoleErrLogger{}

	data, err := c.fileSystem.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to load %q: %w", filename, err)
	}

	module, err := parser2.Parse(filename, string(data), c.diag, errLogger)
	if err != nil {
		return nil, err
	}

	main := &ast.Main{}
	main.AddModule(module)

	for name, extFn := range c.externalFns {
		main.AddExternalFn(name, extFn)
	}

	ctx := ast.NewContext(errLogger)
	main.RunPass(ctx, ast.Check)
	if ctx.Error() != nil {
		return nil, ctx.Error()
	}
	main.RunPass(ctx, ast.Emit)
	if ctx.Error() != nil {
		return nil, ctx.Error()
	}

	return ctx.Emitter().ToProgram(), nil
}

func Run(ctx context.Context, prog *Program, params map[string]Value) error {
	return runtime.NewMachine(prog).Run(ctx, params)
}
