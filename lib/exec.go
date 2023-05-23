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
	"github.com/dcaiafa/nitro/internal/vm"
	"github.com/dcaiafa/nitro/lib/core"
	libio "github.com/dcaiafa/nitro/lib/io"
)

var ErrAborted = errors.New("aborted")
var ErrProcessAlreadyStarted = errors.New("process already started")

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

func newProcess(m *nitro.VM, cmd *osexec.Cmd, stdin io.Reader) *process {
	p := &process{
		vm:          m,
		cmd:         cmd,
		stdin:       stdin,
		stdinQueue:  ioqueue.New(),
		stdoutQueue: ioqueue.New(),
		stderrQueue: ioqueue.New(),
		errorQueue:  ioqueue.New(),
	}
	return p
}

func (p *process) String() string    { return "Process " + p.cmd.Path }
func (p *process) Type() string      { return "Process" }
func (p *process) Traits() vm.Traits { return vm.TraitNone }

func (p *process) SetStderr(w io.Writer) error {
	if p.started {
		return ErrProcessAlreadyStarted
	}
	p.stderr = w
	return nil
}

func (p *process) start() error {
	var err error

	if p.started {
		panic("already started")
	}
	p.started = true

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
			libio.Stderr(p.vm),
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

	if !p.started {
		err = p.start()
		if err != nil {
			return 0, err
		}
	}

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

func execExec(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	var err error
	var stdin io.Reader
	var cmd *osexec.Cmd

	if len(args) > 2 {
		return nil, errTooManyArgs
	}

	cmdArgsIndex := 0
	if len(args) == 2 {
		stdin, err = getReaderArg(m, args, 0)
		if err != nil {
			return nil, err
		}
		cmdArgsIndex = 1
	}

	cmdArgsList, err := getListArg(args, cmdArgsIndex)
	if err != nil {
		return nil, err
	}

	cmdArgs := make([]string, 0, cmdArgsList.Len())

	for i := 0; i < cmdArgsList.Len(); i++ {
		v := cmdArgsList.Get(i)
		if v == nil {
			continue
		}
		switch v := v.(type) {
		case vm.Int, vm.Float, vm.String:
			cmdArgs = append(cmdArgs, v.String())
		default:
			return nil, fmt.Errorf(
				"command argument allowed types are string, int and float, "+
					"but argument #%v is %v",
				i+1, vm.TypeName(v))
		}
	}

	cmd = osexec.Command(cmdArgs[0], cmdArgs[1:]...)

	p := newProcess(m, cmd, stdin)
	p.crumb = m.GetFrameCrumb(1)
	m.RegisterCloser(p)

	return []nitro.Value{p}, nil
}

func execWithStderr(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) > 2 {
		return nil, errTooManyArgs
	}

	p, err := getProcessArg(args, 0)
	if err != nil {
		return nil, err
	}

	w, err := getWriterArg(args, 1)
	if err != nil {
		return nil, err
	}

	err = p.SetStderr(w)
	if err != nil {
		return nil, err
	}

	return []nitro.Value{p}, nil
}
