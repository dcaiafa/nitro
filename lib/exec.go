package lib

import (
	"context"
	"errors"
	"fmt"
	"io"

	osexec "os/exec"
	"sync"

	"github.com/dcaiafa/nitro"
	"github.com/dcaiafa/nitro/internal/ioqueue"
	"github.com/dcaiafa/nitro/lib/core"
)

var ErrAborted = errors.New("aborted")

const minOutBufferReady = 0
const minInBufferReady = 0

type process struct {
	vm         *nitro.VM
	crumb      nitro.FrameCrumb
	cmd        *osexec.Cmd
	closed     bool
	started    bool
	stdin      io.Reader
	stderr     io.Writer
	errSaver   *prefixSuffixSaver
	stdinPipe  io.WriteCloser
	stdoutPipe io.ReadCloser
	stderrPipe io.ReadCloser
	wg         sync.WaitGroup

	stdinQueue  *ioqueue.IOQueue
	stdoutQueue *ioqueue.IOQueue
	stderrQueue *ioqueue.IOQueue
	stdinData   *ioqueue.Data
	stdoutData  *ioqueue.Data

	errorQueue *ioqueue.IOQueue

	mu  sync.Mutex
	err error
}

var _ io.Reader = (*process)(nil)

func newProcess(m *nitro.VM, cmd *osexec.Cmd, stdin io.Reader, stderr io.Writer) *process {
	p := &process{
		vm:          m,
		cmd:         cmd,
		stdin:       stdin,
		stderr:      stderr,
		stdinQueue:  ioqueue.New(),
		stdoutQueue: ioqueue.New(),
		stderrQueue: ioqueue.New(),
		errorQueue:  ioqueue.New(),
	}
	return p
}

func (p *process) String() string { return "Process " + p.cmd.Path }
func (p *process) Type() string   { return "Process" }

func (p *process) EvalOp(op nitro.Op, operand nitro.Value) (nitro.Value, error) {
	return nil, fmt.Errorf("process does not support this operation")
}

func (p *process) Start() error {
	var err error

	if p.stdin != nil {
		if nativeReader, ok := p.stdin.(core.NativeReader); ok {
			p.cmd.Stdin = nativeReader.GetNativeReader()
			p.stdinQueue.Close()
		} else {
			p.stdinPipe, err = p.cmd.StdinPipe()
			if err != nil {
				return err
			}
			p.wg.Add(1)
			go p.inputLoop()
		}
	} else {
		p.stdinQueue.Close()
	}

	if p.stderr == nil {
		p.errSaver = &prefixSuffixSaver{N: 512}
		p.stderr = p.errSaver
	}

	if nativeWriter, ok := p.stderr.(core.NativeWriter); ok {
		p.cmd.Stderr = nativeWriter.GetNativeWriter()
		p.stderrQueue.Close()
	} else {
		p.stderrPipe, err = p.cmd.StderrPipe()
		if err != nil {
			return err
		}
		p.wg.Add(1)
		go p.errLoop()
	}

	p.stdoutPipe, err = p.cmd.StdoutPipe()
	if err != nil {
		return err
	}
	p.wg.Add(1)
	go p.outputLoop()

	err = p.cmd.Start()
	if err != nil {
		p.setError(err)
		p.Close()
		return err
	}

	return nil
}

// inputLoop transfers data from p.stdin (via p.stdinQueue) to the process's
// stdin (via p.stdinPipe). So it goes like this:
// 1. Read() reads from p.stdin, sends to p.stdinQueue
// 2. inputLoop() receives from p.stdinQueue, writes to p.stdinPipe.
func (p *process) inputLoop() {
	defer p.wg.Done()
	defer p.stdinPipe.Close()
	defer p.stdinQueue.SignalWantsClose()

	for {
		select {
		case data, ok := <-p.stdinQueue.C:
			if !ok {
				return
			}

			_, err := p.stdinPipe.Write(data.Data)
			if err != nil {
				// The process closed its input pipe. This will be the case for certain
				// programs like `head`. Leaving inputLoop() will trigger
				// SignalWantsClose(), which will tell Read() to close p.stdin.
				return
			}
			ioqueue.ReleaseData(data)

		case <-p.errorQueue.C:
			return
		}
	}
}

func (p *process) outputLoop() {
	defer p.wg.Done()

	for {
		data := ioqueue.NewData()
		n, err := p.stdoutPipe.Read(data.Data)
		if err != nil {
			if err != io.EOF {
				p.setError(err)
			}
			p.stdoutQueue.Close()
			return
		}
		data.Data = data.Data[:n]
		select {
		case p.stdoutQueue.C <- data:
		case <-p.errorQueue.C:
			return
		}
	}
}

func (p *process) errLoop() {
	defer p.wg.Done()

	for {
		data := ioqueue.NewData()
		n, err := p.stderrPipe.Read(data.Data)
		if err != nil {
			if err != io.EOF {
				p.setError(err)
			}
			p.stderrQueue.Close()
			return
		}

		data.Data = data.Data[:n]
		select {
		case p.stderrQueue.C <- data:
		case <-p.errorQueue.C:
			return
		}
	}
}

func (p *process) getError() error {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.err
}

func (p *process) setError(err error) {
	p.mu.Lock()
	if p.err == nil {
		p.err = err
	}
	p.mu.Unlock()
	p.errorQueue.Close()
}

func (p *process) Close() error {
	if p.closed {
		return nil
	}

	p.closed = true
	if p.vm.ShuttingDown() {
		fmt.Fprintf(
			Stderr(p.vm),
			"WARNING: open process %v created at %v. "+
				"Did you forget to pipe the output?\n",
			p.cmd.Path,
			p.vm.GetFrameInfo(p.crumb),
		)
	}
	p.vm.UnregisterCloser(p)

	p.setError(ErrAborted)

	if p.cmd != nil && p.cmd.Process != nil {
		if err := p.getError(); err != io.EOF {
			p.cmd.Process.Kill()
		}

		wait := func(ctx context.Context) {
			err := p.cmd.Wait()
			if err != nil {
				if p.errSaver != nil {
					err = fmt.Errorf("%w\n%v", err, string(p.errSaver.Bytes()))
				}
				p.mu.Lock()
				p.err = err
				p.mu.Unlock()
			}
		}

		if p.vm.ShuttingDown() {
			wait(context.Background())
		} else {
			p.vm.Block(wait)
		}
	}

	p.wg.Wait()
	core.CloseReader(p.stdin)

	return nil
}

func (p *process) Read(b []byte) (written int, err error) {
	defer func() {
		if err != nil {
			p.setError(err)
			p.Close()
		}
		err = p.getError()
	}()

	// There could be stdout left overs to be processed - a block of data received
	// from outputLoop that was not fully consumed by the previous Read call.
	if p.stdoutData != nil {
		n := copy(b, p.stdoutData.Data)
		b = b[n:]
		written += n
		if len(p.stdoutData.Data) == n {
			ioqueue.ReleaseData(p.stdoutData)
			p.stdoutData = nil
		} else {
			p.stdoutData.Data = p.stdoutData.Data[n:]
		}
	}

	// These are the channels we will be selecting from down below. Store them in
	// variables so we can set each to nil individually as we determine that the
	// channel is closed. Sending/Receiving to a nil channel is ignored by the
	// select statement.
	var (
		stdinC  = p.stdinQueue.C
		stderrC = p.stderrQueue.C
		stdoutC = p.stdoutQueue.C
	)

	for written == 0 {
		if err := p.getError(); err != nil {
			return 0, err
		}

		// The process closed both stdout and stderr.
		// There is no more data to read.
		if stdoutC == nil && stderrC == nil {
			break
		}

		// If there is a stdin, ensure we always have enough data read from it to
		// send to the inputLoop.
		stdinQueueClosed, stdinQueueWantsClose := p.stdinQueue.State()
		if p.stdinData == nil && !stdinQueueClosed {
			if stdinQueueWantsClose {
				p.stdinQueue.Close()
				core.CloseReader(p.stdin)
			} else {
				p.stdinData = ioqueue.NewData()
				n, err := p.stdin.Read(p.stdinData.Data)
				if err == nil {
					p.stdinData.Data = p.stdinData.Data[:n]
				} else {
					ioqueue.ReleaseData(p.stdinData)
					p.stdinData = nil
					p.stdinQueue.Close()
					core.CloseReader(p.stdin)
					if err != io.EOF {
						p.setError(err)
					}
				}
			}
		}

		if p.stdinData == nil {
			stdinC = nil
		}

		var ok bool
		var data *ioqueue.Data
		var stderrData *ioqueue.Data

		p.vm.Block(func(ctx context.Context) {
			select {
			case stdinC <- p.stdinData:
				p.stdinData = nil

			case data, ok = <-stdoutC:
				if ok {
					n := copy(b, data.Data)
					written += n
					b = b[n:]
					if len(data.Data) == n {
						ioqueue.ReleaseData(data)
					} else {
						data.Data = data.Data[n:]
						p.stdoutData = data
					}
				} else {
					stdoutC = nil
				}

			case stderrData, ok = <-stderrC:
				if !ok {
					stderrC = nil
				}
			case <-ctx.Done():
				p.setError(ctx.Err())
			}
		})

		// Processing the stderr data must happen outside of Block() because it
		// could be reentrant.
		if stderrData != nil {
			_, err := p.stderr.Write(stderrData.Data)
			if err == nil {
				ioqueue.ReleaseData(stderrData)
			} else {
				p.setError(err)
			}
		}
	}

	if written == 0 {
		return 0, io.EOF
	}

	return written, nil
}

type execOptions struct {
	Cmd    []string    `nitro:"cmd"`
	Env    []string    `nitro:"env"`
	Dir    string      `nitro:"string"`
	Stderr nitro.Value `nitro:"stderr"`
}

var execOptionsConv core.Value2Structer

var errExecUsage = errors.New(
	`invalid usage. Expected exec((string|reader|iter)?, object) or ` +
		`exec((string|reader|iter)?, string, string...)`)

func exec(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	var err error
	var stdin io.Reader
	var opt execOptions
	var cmd *osexec.Cmd

	if len(args) < 1 {
		return nil, errExecUsage
	}

	switch arg := args[0].(type) {
	case nitro.String, *nitro.Object:
		// Only use the first string or map argument as "stdin" if it was
		// provided via pipeline. Otherwise, this is probably the command argument
		// in concise form, or the options map.
		if m.IsPipeline() {
			stdin, err = nitro.MakeReader(m, arg)
			if err != nil {
				return nil, err
			}
			args = args[1:]
		}

	default:
		if nitro.IsReadable(args[0]) {
			stdin, err = nitro.MakeReader(m, args[0])
			if err != nil {
				return nil, err
			}
			args = args[1:]
		}
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
		return nil, fmt.Errorf(`option "cmd" cannot be empty`)
	}

	var stderr io.Writer
	if opt.Stderr != nil {
		stderr, ok = opt.Stderr.(io.Writer)
		if !ok {
			return nil, fmt.Errorf(
				"option stderr %q is not a writer", nitro.TypeName(opt.Stderr))
		}
	}

	cmd = osexec.Command(opt.Cmd[0], opt.Cmd[1:]...)

	if opt.Env != nil {
		cmd.Env = opt.Env
	}

	p := newProcess(m, cmd, stdin, stderr)

	err = p.Start()
	if err != nil {
		return nil, err
	}

	p.crumb = m.GetFrameCrumb(1)
	m.RegisterCloser(p)

	return []nitro.Value{p}, nil
}
