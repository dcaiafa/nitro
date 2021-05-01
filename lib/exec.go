package lib

import (
	"container/list"
	"errors"
	"fmt"
	"io"
	osexec "os/exec"
	"strings"
	"sync"

	"github.com/dcaiafa/nitro"
)

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
	wg       sync.WaitGroup

	mu           sync.Mutex
	cvInLoop     *sync.Cond
	cvOutLoop    *sync.Cond
	cvRead       *sync.Cond
	inBufs       list.List
	outBufs      list.List
	inputClosed  bool
	outputClosed bool
	err          error
}

var _ io.Reader = (*process)(nil)

func newProcess(m *nitro.VM, cmd *osexec.Cmd, stdin io.Reader) *process {
	p := &process{
		cmd:   cmd,
		stdin: stdin,
	}
	p.cvInLoop = sync.NewCond(&p.mu)
	p.cvOutLoop = sync.NewCond(&p.mu)
	p.cvRead = sync.NewCond(&p.mu)
	p.inBufs.Init()
	p.outBufs.Init()
	return p
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
		p.inputClosed = true
	}

	p.errSaver = &prefixSuffixSaver{N: 512}
	p.cmd.Stderr = p.errSaver

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
	for {
		err := p.inLoopWait()
		if err != nil {
			break
		}
		//golog.Printf("inputLoop: awake")
		buf := p.dequeueInput()
		if buf == nil {
			break
		}
		//golog.Printf("inputLoop: inPipe.Write(%v)", len(buf.data))
		_, err = p.inPipe.Write(buf.data)
		if err != nil {
			//golog.Printf("inputLoop: inPipe.Write failed: %v", err)
			p.setError(err)
			break
		}

		releaseProcessBuffer(buf)
	}
}

func (p *process) inLoopWait() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	for {
		if p.err != nil {
			return p.err
		}

		ready := (p.inBufs.Len() > 0 || p.inputClosed)

		if ready {
			return nil
		}

		p.cvInLoop.Wait()
	}
}

func (p *process) outputLoop() {
	defer p.wg.Done()
	for {
		err := p.outLoopWait()
		if err != nil {
			break
		}
		//golog.Printf("outputLoop: awake")
		buf := newProcessBuffer()
		n, err := p.outPipe.Read(buf.data)
		if err == io.EOF {
			//golog.Printf("outputLoop: EOF")
			p.closeOutput()
			break
		} else if err != nil {
			//golog.Printf("outputLoop: err: %v", err)
			p.setError(err)
			break
		}
		//golog.Printf("outputLoop: received %v", n)
		buf.data = buf.data[:n]
		p.enqueueOutput(buf)
	}
}

func (p *process) outLoopWait() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	for {
		if p.err != nil {
			return p.err
		}

		ready := p.outBufs.Len() <= minOutBufferReady

		if ready {
			return nil
		}

		p.cvOutLoop.Wait()
	}
}

func (p *process) closeOutput() {
	//golog.Printf("closeOutput")
	p.mu.Lock()
	p.outputClosed = true
	p.mu.Unlock()
	p.cvRead.Signal()
}

func (p *process) enqueueOutput(buf *processBuffer) {
	//golog.Printf("enqueueOutput")
	p.mu.Lock()
	p.outBufs.PushBack(buf)
	p.mu.Unlock()
	p.cvRead.Signal()
}

func (p *process) dequeueInput() *processBuffer {
	//golog.Printf("dequeueInput")
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.inBufs.Len() == 0 {
		if !p.inputClosed {
			panic("dequeueInput called but no data")
		}
		return nil
	}

	buf := p.inBufs.Front().Value.(*processBuffer)
	p.inBufs.Remove(p.inBufs.Front())
	return buf
}

func (p *process) error() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	return p.err
}

func (p *process) setError(err error) {
	//golog.Printf("setError: %v", err)
	p.mu.Lock()
	if p.err == nil {
		p.err = err
	}
	p.mu.Unlock()
	p.cvRead.Signal()
	p.cvInLoop.Signal()
	p.cvOutLoop.Signal()
}

func (p *process) Read(b []byte) (int, error) {
	for {
		if p.error() != nil {
			p.cmd.Wait()
			p.wg.Wait()
			return 0, p.error()
		}

		if p.isReadyForInput() {
			buf := newProcessBuffer()
			n, err := p.stdin.Read(buf.data)
			if err == io.EOF {
				p.stdin = nil
				// TODO: close stdin
				p.closeInput()
				continue
			} else if err != nil {
				return 0, err
			}
			buf.data = buf.data[:n]
			//golog.Printf("Read: enqueing %v", n)
			p.enqueueInput(buf)
		} else if p.hasOutput() {
			n := p.dequeueOutput(b)
			//golog.Printf("Read: %v", n)
			return n, nil
		} else if p.isOutputClosed() {
			//golog.Printf("Read: output is closed")
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
		} else {
			p.readWait()
		}
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
			(p.outBufs.Len() > 0 || p.outputClosed) ||
				(!p.inputClosed && p.inBufs.Len() <= minInBufferReady)
		if ready {
			return nil
		}

		p.cvRead.Wait()
	}
}

func (p *process) isReadyForInput() bool {
	p.mu.Lock()
	defer p.mu.Unlock()
	return !p.inputClosed && p.inBufs.Len() <= 1
}

func (p *process) isInputClosed() bool {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.inputClosed
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
	return p.outBufs.Len() == 0 && p.outputClosed
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
	p.inputClosed = true
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

	err = p.Start()
	if err != nil {
		return nil, err
	}

	return []nitro.Value{p}, nil
}
