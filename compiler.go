package nitro

import (
	"context"
	"fmt"
	"os"

	"github.com/dcaiafa/nitro/internal/ast"
	"github.com/dcaiafa/nitro/internal/errlogger"
	"github.com/dcaiafa/nitro/internal/parser"
	"github.com/dcaiafa/nitro/internal/runtime"
	"github.com/dcaiafa/nitro/internal/token"
)

type Value = runtime.Value
type ErrLogger = errlogger.ErrLogger
type Pos = token.Pos

type DefaultErrLogger struct{}

func (l *DefaultErrLogger) Failf(pos Pos, msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, pos.String()+": "+msg, args...)
}

func (l *DefaultErrLogger) Detailf(pos Pos, msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, pos.String()+": "+msg, args...)
}

type FileSystem interface {
	ReadFile(name string) ([]byte, error)
}

type Compiler struct {
	fileSystem  FileSystem
	externalFns map[string]runtime.ExternFn
}

func NewCompiler(fileSystem FileSystem) *Compiler {
	return &Compiler{
		fileSystem:  fileSystem,
		externalFns: make(map[string]runtime.ExternFn),
	}
}

func (c *Compiler) RegisterExternalFn(name string, fn runtime.ExternFn) {
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

	ctx := ast.NewContext(&DefaultErrLogger{})
	main.RunPass(ctx, ast.CreateAndResolveNames)
	if ctx.Error() != nil {
		return nil, ctx.Error()
	}
	main.RunPass(ctx, ast.Emit)
	if ctx.Error() != nil {
		return nil, ctx.Error()
	}

	return ctx.Emitter().ToProgram(), nil
}

func Run(ctx context.Context, p *runtime.Program) error {
	return runtime.NewMachine().Run(ctx, p)
}
