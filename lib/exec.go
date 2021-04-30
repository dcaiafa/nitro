package lib

import (
	"errors"
	"fmt"
	"io"
	osexec "os/exec"
	"strings"
	"sync/atomic"

	"github.com/dcaiafa/nitro"
)

type process struct {
	cmd          *osexec.Cmd
	stdin        io.Reader
	input        chan interface{}
	output       chan interface{}
	outQueue     ByteQueue
	inPipe       io.WriteCloser
	outPipe      io.ReadCloser
	errPipe      io.ReadCloser
	outPipeCount int32
	stdout       io.Writer
	stderr       io.Writer
	errSaver     *prefixSuffixSaver
}

var _ io.Reader = (*process)(nil)

func newProcess(
	m *nitro.VM,
	cmd *osexec.Cmd,
	input io.Reader,
) *process {
	return &process{
		cmd:    cmd,
		stdin:  input,
		input:  make(chan interface{}, 1),
		output: make(chan interface{}, 1),
	}
}

func (p *process) String() string { return "Process " + p.cmd.Path }
func (p *process) Type() string   { return "Process" }

func (p *process) EvalBinOp(op nitro.BinOp, operand nitro.Value) (nitro.Value, error) {
	return nil, fmt.Errorf("process does not support this operation")
}

func (p *process) EvalUnaryMinus() (nitro.Value, error) {
	return nil, fmt.Errorf("process does not support this operation")
}

func (p *process) SetStdout(w io.Writer) {
	p.stdout = w
}

func (p *process) SetStderr(w io.Writer) {
	p.stderr = w
}

func (p *process) Run() error {
	var err error
	p.inPipe, err = p.cmd.StdinPipe()
	if err != nil {
		return err
	}

	if p.stdout == nil {
		p.outPipe, err = p.cmd.StdoutPipe()
		if err != nil {
			return err
		}
	} else {
		p.cmd.Stdout = p.stdout
	}

	p.errSaver = &prefixSuffixSaver{N: 1024}
	p.cmd.Stderr = p.errSaver

	err = p.cmd.Start()
	if err != nil {
		return err
	}

	if p.stdin != nil {
		go func() {
			for {
				buf := make([]byte, 512)
				n, err := p.stdin.Read(buf)
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

	if p.outPipe != nil {
		atomic.AddInt32(&p.outPipeCount, 1)
		go func() {
			defer func() {
				if atomic.AddInt32(&p.outPipeCount, -1) == 0 {
					close(p.output)
				}
			}()

			for {
				buf := make([]byte, 512)
				n, err := p.outPipe.Read(buf)
				if err == io.EOF {
					return
				} else if err != nil {
					p.output <- err
					return
				}
				p.output <- buf[:n]
			}
		}()
	} else {
		close(p.output)
		err := p.wait()
		if err != nil {
			return err
		}
	}

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
				p.inPipe.Close()
				continue
			}
			switch bufOrErr := bufOrErr.(type) {
			case []byte:
				_, err := p.inPipe.Write(bufOrErr)
				if err != nil {
					return err
				}
			case error:
				return bufOrErr
			}

		case bufOrErr, ok := <-p.output:
			if !ok {
				// Output was closed, wait for process to finish.
				return p.wait()
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

func (p *process) wait() error {
	err := p.cmd.Wait()
	if err != nil {
		if p.errSaver != nil {
			err = fmt.Errorf("%w\n%v", err, string(p.errSaver.Bytes()))
		}
		return err
	}
	return nil
}

type execOptions struct {
	Cmd           []string    `nitro:"cmd"`
	Env           []string    `nitro:"env"`
	Dir           string      `nitro:"string"`
	Stdout        nitro.Value `nitro:"stdout"`
	Stderr        nitro.Value `nitro:"stderr"`
	CombineOutput bool        `nitro:"combineoutput"`
}

var execOptionsConv Value2Structer

var errExecUsage = errors.New(
	`invalid usage. Expected exec((string|reader|iter)?, object) or exec((string|reader|iter)?, string, string...)`)

func exec(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	var err error
	var stdin io.Reader
	var opt execOptions
	var cmd *osexec.Cmd

	if len(args) < 1 {
		return nil, errExecUsage
	}

	switch arg := args[0].(type) {
	case io.Reader:
		stdin = arg
		args = args[1:]

	case nitro.String:
		// Only use this string argument as "stdin" if it was provided via pipeline.
		// Otherwise, this is probably the "command" argument in the concise form.
		if m.IsPipeline() {
			stdin = strings.NewReader(arg.String())
			args = args[1:]
		}

	case *nitro.Iterator:
		stdin = &iterReader{
			m: m,
			e: arg,
		}
		args = args[1:]
	}

	if len(args) < 1 {
		return nil, errExecUsage
	}

	optv, ok := args[0].(*nitro.Object)
	if ok {
		err = execOptionsConv.Convert(optv, &opt)
		if err != nil {
			return nil, err
		}
	} else {
		for _, arg := range args {
			argStr, ok := arg.(nitro.String)
			if !ok {
				return nil, errExecUsage
			}
			opt.Cmd = append(opt.Cmd, argStr.String())
		}
	}

	if len(opt.Cmd) == 0 {
		return nil, fmt.Errorf("option \"cmd\" cannot be empty")
	}

	cmd = osexec.Command(opt.Cmd[0], opt.Cmd[1:]...)

	p := newProcess(m, cmd, stdin)

	err = p.Run()
	if err != nil {
		return nil, err
	}

	return []nitro.Value{p}, nil
}
