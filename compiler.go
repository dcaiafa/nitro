package nitro

import (
	"context"
	"fmt"

	"github.com/dcaiafa/nitro/internal/ast"
	"github.com/dcaiafa/nitro/internal/parser"
	"github.com/dcaiafa/nitro/internal/runtime"
)

type Value = runtime.Value

type FileSystem interface {
	ReadFile(name string) ([]byte, error)
}

type Compiler struct {
	fileSystem  FileSystem
	externalFns map[string]runtime.ExternalFunc
}

func NewCompiler(fileSystem FileSystem) *Compiler {
	return &Compiler{
		fileSystem:  fileSystem,
		externalFns: make(map[string]runtime.ExternalFunc),
	}
}

func (c *Compiler) RegisterExternalFn(name string, fn runtime.ExternalFunc) {
	c.externalFns[name] = fn
}

func (c *Compiler) Compile(filename string) (*runtime.Program, error) {
	var err error

	data, err := c.fileSystem.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to load %q: %w", filename, err)
	}

	main, err := parser.Parse(filename, data)
	if err != nil {
		return nil, err
	}

	for name, extFn := range c.externalFns {
		main.AddExternalFn(name, extFn)
	}

	ctx := ast.NewContext()
	main.RunPass(ctx, ast.CreateAndResolveNames)
	if ctx.HasErrors() {
		return nil, ctx.Errors()[0]
	}
	main.RunPass(ctx, ast.Emit)
	if ctx.HasErrors() {
		return nil, ctx.Errors()[0]
	}

	return ctx.Emitter().ToProgram(), nil
}

func Run(ctx context.Context, p *runtime.Program) error {
	return runtime.NewMachine().Run(ctx, p)
}
