package lib

import (
	"context"
	"fmt"
	"io"
	"os/exec"

	"github.com/dcaiafa/nitro"
)

type process struct {
	cmd        *exec.Cmd
	processOut io.ReadCloser
}

var _ io.Reader = (*process)(nil)

func newProcess(ctx context.Context, name string, args []string) *process {
	return &process{
		cmd: exec.CommandContext(ctx, name, args...),
	}
}

func (p *process) String() string { return "Process " + p.cmd.Path }
func (p *process) Type() string   { return "Process" }

func (p *process) SetStdin(stdin io.Reader) {
	p.cmd.Stdin = stdin
}

func (p *process) Run() error {
	var err error
	p.processOut, err = p.cmd.StdoutPipe()
	if err != nil {
		return err
	}
	err = p.cmd.Start()
	if err != nil {
		return err
	}
	return nil
}

func (p *process) Read(b []byte) (int, error) {
	n, err := p.processOut.Read(b)
	if err == nil {
		return n, nil
	}
	if err != io.EOF {
		return n, err
	}
	err = p.cmd.Wait()
	if err != nil {
		return n, err
	}
	return n, io.EOF
}

func fnExec(ctx context.Context, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	var stdin io.Reader
	var name string
	var pargs []string

	if len(args) < 1 {
		return nil, errNotEnoughArgs
	}
	if r, ok := args[0].(io.Reader); ok {
		stdin = r
		args = args[1:]
	}
	if len(args) < 1 {
		return nil, errNotEnoughArgs
	}

	if c, ok := args[0].(nitro.String); ok {
		name = c.String()
		args = args[1:]
		for i := 0; i < len(args); i++ {
			arg, ok := args[i].(nitro.String)
			if !ok {
				return nil, fmt.Errorf(
					"invalid argument of type %q",
					nitro.TypeName(args[i]))
			}
			pargs = append(pargs, arg.String())
		}
	} else {
		return nil, fmt.Errorf(
			"invalid argument of type %q",
			nitro.TypeName(args[0]))
	}

	p := newProcess(ctx, name, pargs)

	if stdin != nil {
		p.SetStdin(stdin)
	}

	err := p.Run()
	if err != nil {
		return nil, err
	}

	return []nitro.Value{p}, nil
}
