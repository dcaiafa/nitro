package compiler

import (
	"errors"
	"fmt"

	"github.com/dcaiafa/nitro/internal/ast"
	"github.com/dcaiafa/nitro/internal/context"
	"github.com/dcaiafa/nitro/internal/parser"
)

var ErrAlreadyLoaded = errors.New("already loaded")

type FileSystem interface {
	AbsPath(name string) (string, error)
	ReadFile(name string) ([]byte, error)
}

type Compiler struct {
	fileSystem FileSystem
	mainProg   *ast.Program
	allProgs   map[string]*ast.Program
}

func (c *Compiler) Compile(filename string) error {
	if c.mainProg != nil {
		panic("Compile cannot be called twice")
	}

	var err error
	c.mainProg, err = c.load(filename)
	if err != nil {
		return err
	}

	ctx := context.NewContext()
	for _, pass := range context.Passes {
		for _, prog := range c.allProgs {
			prog.RunPass(ctx, pass)
			if ctx.HasErrors() {
				return ctx.Errors()[0]
			}
		}
	}

	return nil
}

func (c *Compiler) load(filename string) (*ast.Program, error) {
	absPath, err := c.fileSystem.AbsPath(filename)
	if err != nil {
		return nil, fmt.Errorf(
			"couldn't get absolute path for %q: %w", filename, err)
	}

	if c.allProgs[absPath] != nil {
		return nil, ErrAlreadyLoaded
	}

	data, err := c.fileSystem.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to load %q: %w", filename, err)
	}

	prog, err := parser.Parse(filename, data)
	if err != nil {
		return nil, err
	}

	c.allProgs[absPath] = prog

	return prog, nil
}
