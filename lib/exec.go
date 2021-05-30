package lib

import (
	"container/list"
	"errors"
	"fmt"
	"io"
	//golog "log"
	osexec "os/exec"
	"runtime"
	"strings"
	"sync"

	"github.com/dcaiafa/nitro"
)

var ErrAborted = errors.New("aborted")

const minOutBufferReady = 0
const minInBufferReady = 0

type processBuffer struct {
	data []byte
	arr  [512]byte
}

var processBufferPool = sync.Pool{
	New: func() interface{} {
		return &processBuffer{}
	},
}

func newProcessBuffer() *processBuffer {
	b := processBufferPool.Get().(*processBuffer)
	b.data = b.arr[:]
	return b
}

func releaseProcessBuffer(b *processBuffer) {
	processBufferPool.Put(b)
}

type process struct {
	cmd      *osexec.Cmd
	stdin    io.Reader
	stdout   io.Writer
	stderr   io.Writer
	errSaver *prefixSuffixSaver
	inPipe   io.WriteCloser
	outPipe  io.ReadCloser
	errPipe  io.ReadCloser
	wg       sync.WaitGroup

	mu        sync.Mutex
	inBufs    list.List
	cvInLoop  *sync.Cond
	inClosed  bool
	outBufs   list.List
	cvOutLoop *sync.Cond
	outClosed bool
	errBufs   list.List
	cvErrLoop *sync.Cond
	errClosed bool
	cvRead    *sync.Cond
	err       error
}

var _ io.Reader = (*process)(nil)

func newProcess(m *nitro.VM, cmd *osexec.Cmd, stdin io.Reader) *process {
	p := &process{
		cmd:   cmd,
		stdin: stdin,
	}
	p.inBufs.Init()
	p.cvInLoop = sync.NewCond(&p.mu)
	p.outBufs.Init()
	p.cvOutLoop = sync.NewCond(&p.mu)
	p.errBufs.Init()
	p.cvErrLoop = sync.NewCond(&p.mu)
	p.cvRead = sync.NewCond(&p.mu)
	return p
}

func (p *process) String() string { return "Process " + p.cmd.Path }
func (p *process) Type() string   { return "Process" }

func (p *process) EvalOp(op nitro.Op, operand nitro.Value) (nitro.Value, error) {
	return nil, fmt.Errorf("process does not support this operation")
}

func (p *process) SetStdout(w io.Writer) {
	p.stdout = w
}

func (p *process) SetStderr(w io.Writer) {
	p.stderr = w
}

func (p *process) Start() error {
	var err error
	if p.stdin != nil {
		p.inPipe, err = p.cmd.StdinPipe()
		if err != nil {
			return err
		}
		p.wg.Add(1)
		go p.inputLoop()
	} else {
		p.inClosed = true
	}

	if p.stderr == nil {
		p.errSaver = &prefixSuffixSaver{N: 512}
		p.stderr = p.errSaver
	}

	p.errPipe, err = p.cmd.StderrPipe()
	if err != nil {
		return err
	}
	p.wg.Add(1)
	go p.errLoop()

	p.outPipe, err = p.cmd.StdoutPipe()
	if err != nil {
		return err
	}
	p.wg.Add(1)
	go p.outputLoop()

	err = p.cmd.Start()
	if err != nil {
		p.setError(err)
		p.wg.Wait()
		return err
	}

	return nil
}

func (p *process) inputLoop() {
	defer p.wg.Done()
	defer p.inPipe.Close()

	for p.inLoopWait() {
		buf, wakeUp := p.dequeueInput()
		if wakeUp {
			p.cvRead.Signal()
		}
		if buf == nil {
			break
		}

		_, err := p.inPipe.Write(buf.data)
		if err != nil {
			p.setError(err)
			break
		}
		releaseProcessBuffer(buf)
	}
}

func (p *process) inLoopWait() bool {
	p.mu.Lock()
	defer p.mu.Unlock()

	for {
		if p.err != nil {
			return false
		}
		if p.inBufs.Len() > 0 {
			return true
		}
		if p.inClosed {
			return false
		}

		p.cvInLoop.Wait()
	}
}

func (p *process) outputLoop() {
	defer p.wg.Done()

	for p.outLoopWait() {
		buf := newProcessBuffer()
		n, err := p.outPipe.Read(buf.data)
		if err == io.EOF {
			p.closeOutput()
			break
		} else if err != nil {
			p.setError(err)
			break
		}
		buf.data = buf.data[:n]
		p.enqueueOutput(buf)
	}
}

func (p *process) outLoopWait() bool {
	p.mu.Lock()
	defer p.mu.Unlock()

	for {
		if p.err != nil {
			return false
		}

		if p.outBufs.Len() == 0 {
			return true
		}

		p.cvOutLoop.Wait()
	}
}

func (p *process) closeOutput() {
	p.mu.Lock()
	p.outClosed = true
	p.mu.Unlock()
	p.cvRead.Signal()
}

func (p *process) enqueueOutput(buf *processBuffer) {
	p.mu.Lock()
	p.outBufs.PushBack(buf)
	p.mu.Unlock()
	p.cvRead.Signal()
}

func (p *process) errLoop() {
	defer p.wg.Done()
	for {
		err := p.errLoopWait()
		if err != nil {
			break
		}
		buf := newProcessBuffer()
		n, err := p.errPipe.Read(buf.data)
		if err == io.EOF {
			p.closeErr()
			break
		} else if err != nil {
			p.setError(err)
			break
		}
		buf.data = buf.data[:n]
		p.enqueueErr(buf)
	}
}

func (p *process) errLoopWait() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	for {
		if p.err != nil {
			return p.err
		}

		ready := p.errBufs.Len() == 0

		if ready {
			return nil
		}

		p.cvErrLoop.Wait()
	}
}

func (p *process) enqueueErr(buf *processBuffer) {
	p.mu.Lock()
	p.errBufs.PushBack(buf)
	p.mu.Unlock()
	p.cvRead.Signal()
}

func (p *process) dequeueErr(b []byte) int {
	p.mu.Lock()

	total := 0
	for len(b) > 0 && p.errBufs.Len() > 0 {
		buf := p.errBufs.Front().Value.(*processBuffer)
		n := copy(b, buf.data)
		if n < len(buf.data) {
			buf.data = buf.data[n:]
		} else {
			p.errBufs.Remove(p.errBufs.Front())
			releaseProcessBuffer(buf)
		}
		total += n
		b = b[n:]
	}

	shouldSignal := p.errBufs.Len() == 0

	p.mu.Unlock()

	if shouldSignal {
		p.cvErrLoop.Signal()
	}

	return total
}

func (p *process) hasErr() bool {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.errBufs.Len() > 0
}

func (p *process) closeErr() {
	p.mu.Lock()
	p.errClosed = true
	p.mu.Unlock()
	p.cvRead.Signal()
}

func (p *process) isErrClosed() bool {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.errBufs.Len() == 0 && p.errClosed
}

func (p *process) dequeueInput() (buf *processBuffer, wakeUp bool) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.inBufs.Len() == 0 {
		if !p.inClosed {
			panic("dequeueInput called but no data")
		}
		return nil, false
	}

	buf = p.inBufs.Front().Value.(*processBuffer)
	p.inBufs.Remove(p.inBufs.Front())
	wakeUp = p.inBufs.Len() <= minInBufferReady
	return buf, wakeUp
}

func (p *process) error() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	return p.err
}

func (p *process) setError(err error) {
	stack := make([]byte, 10240)
	n := runtime.Stack(stack, false)
	stack = stack[:n]

	p.mu.Lock()
	if p.err == nil {
		p.err = err
	}
	p.mu.Unlock()
	p.cvRead.Signal()
	p.cvInLoop.Signal()
	p.cvOutLoop.Signal()
	p.cvErrLoop.Signal()
}

func (p *process) Close() error {
	p.setError(ErrAborted)
	p.cmd.Process.Kill()
	CloseReader(p.stdin)
	p.cmd.Wait()
	p.wg.Wait()
	return nil
}

func (p *process) Read(b []byte) (int, error) {
	for {
		if err := p.error(); err != nil {
			if err != io.EOF {
				p.cmd.Process.Kill()
			}
			CloseReader(p.stdin)
			p.cmd.Wait()
			p.wg.Wait()
			p.Close()
			return 0, err
		}

		if p.isReadyForInput() {
			buf := newProcessBuffer()
			n, err := p.stdin.Read(buf.data)
			if err == io.EOF {
				p.stdin = nil
				p.closeInput()
				continue
			} else if err != nil {
				p.setError(err)
				continue
			}
			buf.data = buf.data[:n]
			p.enqueueInput(buf)
		}

		if p.hasErr() {
			buf := newProcessBuffer()
			n := p.dequeueErr(buf.data)
			_, err := p.stderr.Write(buf.data[:n])
			if err != nil {
				p.setError(err)
				continue
			}
		}

		if p.hasOutput() {
			n := p.dequeueOutput(b)
			return n, nil
		}

		if p.isOutputClosed() && p.isErrClosed() {
			p.closeInput()
			err := p.cmd.Wait()
			if err != nil {
				if p.errSaver != nil {
					err = fmt.Errorf("%w\n%v", err, string(p.errSaver.Bytes()))
				}
				p.setError(err)
			} else {
				p.setError(io.EOF)
			}
		}

		p.readWait()
	}
}

func (p *process) readWait() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	for {
		if p.err != nil {
			return p.err
		}

		ready :=
			p.outBufs.Len() > 0 ||
				p.errBufs.Len() > 0 ||
				(!p.inClosed && p.inBufs.Len() <= minInBufferReady) ||
				(p.outClosed && p.errClosed)

		if ready {
			return nil
		}

		p.cvRead.Wait()
	}
}

func (p *process) isReadyForInput() bool {
	p.mu.Lock()
	defer p.mu.Unlock()
	return !p.inClosed && p.inBufs.Len() <= minInBufferReady
}

func (p *process) enqueueInput(buf *processBuffer) {
	p.mu.Lock()
	p.inBufs.PushBack(buf)
	p.mu.Unlock()
	p.cvInLoop.Signal()
}

func (p *process) hasOutput() bool {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.outBufs.Len() > 0
}

func (p *process) isOutputClosed() bool {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.outBufs.Len() == 0 && p.outClosed
}

func (p *process) dequeueOutput(b []byte) int {
	p.mu.Lock()

	total := 0
	for len(b) > 0 && p.outBufs.Len() > 0 {
		buf := p.outBufs.Front().Value.(*processBuffer)
		n := copy(b, buf.data)
		if n < len(buf.data) {
			buf.data = buf.data[n:]
		} else {
			p.outBufs.Remove(p.outBufs.Front())
			releaseProcessBuffer(buf)
		}
		total += n
		b = b[n:]
	}

	shouldSignal := p.outBufs.Len() <= minOutBufferReady

	p.mu.Unlock()

	if shouldSignal {
		p.cvOutLoop.Signal()
	}

	return total
}

func (p *process) closeInput() {
	p.mu.Lock()
	p.inClosed = true
	p.mu.Unlock()
	p.cvInLoop.Signal()
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

	case nitro.Iterator:
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

	if opt.Stderr != nil {
		stderr, ok := opt.Stderr.(io.Writer)
		if !ok {
			return nil, fmt.Errorf(
				"option stderr %q is not a writer", nitro.TypeName(opt.Stderr))
		}
		p.SetStderr(stderr)
	}

	err = p.Start()
	if err != nil {
		return nil, err
	}

	return []nitro.Value{p}, nil
}
