package lib

import (
	"fmt"
	"io"
	osexec "os/exec"

	"github.com/dcaiafa/nitro"
)

type process struct {
	cmd          *osexec.Cmd
	inputReader  io.Reader
	input        chan interface{}
	output       chan interface{}
	outQueue     ByteQueue
	stdin        io.WriteCloser
	stdout       io.ReadCloser
	outputClosed bool
}

var _ io.Reader = (*process)(nil)

func newProcess(
	m *nitro.Machine,
	name string,
	args []string,
	input io.Reader,
) *process {
	return &process{
		cmd:         osexec.CommandContext(m.Context(), name, args...),
		inputReader: input,
		input:       make(chan interface{}, 1),
		output:      make(chan interface{}, 1),
	}
}

func (p *process) String() string { return "Process " + p.cmd.Path }
func (p *process) Type() string   { return "Process" }

func (p *process) Run() error {
	var err error
	p.stdin, err = p.cmd.StdinPipe()
	if err != nil {
		return err
	}
	p.stdout, err = p.cmd.StdoutPipe()
	if err != nil {
		return err
	}
	err = p.cmd.Start()
	if err != nil {
		return err
	}

	if p.inputReader != nil {
		go func() {
			for {
				buf := make([]byte, 512)
				n, err := p.inputReader.Read(buf)
				if err == io.EOF {
					close(p.input)
					return
				} else if err != nil {
					p.input <- err
					close(p.input)
					return
				}
				p.input <- buf[:n]
			}
		}()
	} else {
		close(p.input)
	}

	go func() {
		for {
			buf := make([]byte, 512)
			n, err := p.stdout.Read(buf)
			if err == io.EOF {
				close(p.output)
				return
			} else if err != nil {
				p.output <- err
				close(p.output)
				return
			}
			p.output <- buf[:n]
		}
	}()

	return nil
}

func (p *process) Read(b []byte) (int, error) {
	if len(p.outQueue.Peek()) == 0 {
		err := p.feedProcessUntilOutputAvailable()
		if err != nil {
			return 0, err
		}
	}
	if len(p.outQueue.Peek()) == 0 {
		return 0, io.EOF
	}
	n := len(p.outQueue.Peek())
	if n > len(b) {
		n = len(b)
	}
	copy(b, p.outQueue.Peek()[:n])
	p.outQueue.Pop(n)
	return n, nil
}

func (p *process) feedProcessUntilOutputAvailable() error {
	input := p.input

	for {
		if len(p.outQueue.Peek()) > 0 {
			return nil
		}

		select {
		case bufOrErr, ok := <-input:
			if !ok {
				input = nil
				p.stdin.Close()
				continue
			}
			switch bufOrErr := bufOrErr.(type) {
			case []byte:
				_, err := p.stdin.Write(bufOrErr)
				if err != nil {
					return err
				}
			case error:
				return bufOrErr
			}

		case bufOrErr, ok := <-p.output:
			if !ok {
				// Output was closed, wait for process to finish.
				return p.cmd.Wait()
			}
			switch bufOrErr := bufOrErr.(type) {
			case []byte:
				p.outQueue.Write(bufOrErr)
			case error:
				return bufOrErr
			}
		}
	}
}

func exec(
	m *nitro.Machine,
	caps []nitro.ValueRef,
	args []nitro.Value,
	retN int,
) ([]nitro.Value, error) {
	var err error
	var stdin io.Reader
	var name string
	var pargs []string

	if len(args) < 1 {
		return nil, errNotEnoughArgs
	}

	if _, ok := args[0].(nitro.String); !ok {
		stdin, err = ToReader(m, args[0])
		if err != nil {
			return nil, err
		}
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

	p := newProcess(m, name, pargs, stdin)

	err = p.Run()
	if err != nil {
		return nil, err
	}

	return []nitro.Value{p}, nil
}
