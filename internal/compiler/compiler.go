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
	mainMod    *ast.Module
	allMods    map[string]*ast.Module
}

func (c *Compiler) Compile(filename string) error {
	if c.mainMod != nil {
		panic("Compile cannot be called twice")
	}

	var err error
	c.mainMod, err = c.load(filename)
	if err != nil {
		return err
	}

	ctx := context.NewContext()
	for _, pass := range context.Passes {
		for _, mod := range c.allMods {
			mod.RunPass(ctx, pass)
			if ctx.HasErrors() {
				return ctx.Errors()[0]
			}
		}
	}

	return nil
}

func (c *Compiler) load(filename string) (*ast.Module, error) {
	absPath, err := c.fileSystem.AbsPath(filename)
	if err != nil {
		return nil, fmt.Errorf(
			"couldn't get absolute path for %q: %w", filename, err)
	}

	if c.allMods[absPath] != nil {
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

	c.allMods[absPath] = prog

	return prog, nil
}
