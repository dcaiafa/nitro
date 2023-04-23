package main

import (
	"bytes"
	"io"
	"os/exec"

	"github.com/dcaiafa/nitro/internal/errlogger"
	"github.com/dcaiafa/nitro/internal/stub/ast"
)

type Generator struct {
	unit   *ast.Unit
	output io.Writer
}

func NewGenerator(unit *ast.Unit, output io.Writer) *Generator {
	return &Generator{
		unit:   unit,
		output: output,
	}
}

func (g *Generator) Run() error {
	errLogger := errlogger.NewErrLoggerBase(
		&errlogger.ConsoleErrLogger{})

	ctx := ast.NewContext(errLogger)

	passes := []ast.Pass{
		ast.Check,
		ast.Emit,
	}

	for _, pass := range passes {
		g.unit.RunPass(ctx, pass)
		if errLogger.Error() != nil {
			return errLogger.Error()
		}
	}

	outBuf := &bytes.Buffer{}
	ctx.Analysis.Emit(outBuf)

	err := goFmt(outBuf, g.output)
	if err != nil {
		return err
	}

	return nil
}

func goFmt(r io.Reader, w io.Writer) error {
	cmd := exec.Command("gofmt")
	cmd.Stdin = r
	cmd.Stdout = w
	return cmd.Run()
}
