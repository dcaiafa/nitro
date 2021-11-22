package lib

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/dcaiafa/nitro"
	"github.com/dcaiafa/nitro/lib/core"
)

type File struct {
	*os.File
}

var _ /* implements */ nitro.Indexable = (*File)(nil)
var _ /* implements */ nitro.Callable = (*File)(nil)
var _ /* implements */ core.NativeReader = (*File)(nil)
var _ /* implements */ core.NativeWriter = (*File)(nil)

func (f *File) String() string { return fmt.Sprintf("File:%v", f.Name()) }
func (f *File) Type() string   { return "File" }

func (f *File) GetNativeReader() io.Reader {
	return f.File
}

func (f *File) GetNativeWriter() io.Writer {
	return f.File
}

func (f *File) EvalOp(op nitro.Op, operand nitro.Value) (nitro.Value, error) {
	return nil, fmt.Errorf("file does not support this operation")
}

func (f *File) Call(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 1 {
		return nil, core.ErrWriterCallUsage
	}

	reader, err := nitro.MakeReader(m, args[0])
	if err != nil {
		return nil, core.ErrWriterCallUsage
	}

	n, err := io.Copy(f.File, reader)
	if err != nil {
		return nil, err
	}

	return []nitro.Value{nitro.NewInt(n)}, nil
}

func (f *File) Index(key nitro.Value) (nitro.Value, error) {
	keyStr, ok := key.(nitro.String)
	if !ok {
		return nil, fmt.Errorf(
			"file cannot be indexed by %q",
			nitro.TypeName(key))
	}

	switch keyStr.String() {
	case "name":
		return nitro.NewString(f.Name()), nil

	default:
		return nil, fmt.Errorf(
			"file does not have method %q",
			keyStr.String())
	}
}

func (f *File) IndexRef(key nitro.Value) (nitro.ValueRef, error) {
	return nitro.ValueRef{}, fmt.Errorf("file is not assignable")
}

type openOptions struct {
	Read   bool   `nitro:"read"`
	Write  bool   `nitro:"write"`
	Append bool   `nitro:"append"`
	Create bool   `nitro:"create"`
	Excl   bool   `nitro:"excl"`
	Trunc  bool   `nitro:"trunc"`
	Perm   *int64 `nitro:"perm"`
}

var openOptionsConv core.Value2Structer

var errOpenUsage = errors.New(
	`invalid usage. Expected open(string, map?)`)

func open(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	var err error

	if len(args) != 1 && len(args) != 2 {
		return nil, errOpenUsage
	}

	filename, ok := args[0].(nitro.String)
	if !ok {
		return nil, errOpenUsage
	}

	var opts *openOptions
	if len(args) == 2 {
		opts = new(openOptions)
		optsMap, ok := args[1].(*nitro.Object)
		if !ok {
			return nil, errOpenUsage
		}
		err = openOptionsConv.Convert(optsMap, opts)
		if err != nil {
			return nil, err
		}
	}

	var f *os.File
	if opts == nil {
		f, err = os.Open(filename.String())
		if err != nil {
			return nil, err
		}
	} else {
		flags := 0
		var perm os.FileMode = 0666
		if opts.Read && !opts.Write {
			flags = flags | os.O_RDONLY
		} else if !opts.Read && opts.Write {
			flags = flags | os.O_WRONLY
		} else if opts.Read && opts.Write {
			flags = flags | os.O_RDWR
		} else {
			return nil, fmt.Errorf(
				`invalid options: "read" and/or "write" must be true`)
		}

		if opts.Append {
			flags = flags | os.O_APPEND
		}
		if opts.Create {
			flags = flags | os.O_CREATE
		}
		if opts.Excl {
			if !opts.Create {
				return nil, fmt.Errorf(
					`invalid options: "excl" also required "create" to be true`)
			}
			flags = flags | os.O_EXCL
		}
		if opts.Trunc {
			flags = flags | os.O_TRUNC
		}
		if opts.Perm != nil {
			perm = os.FileMode(*opts.Perm)
		}
		f, err = os.OpenFile(filename.String(), flags, perm)
		if err != nil {
			return nil, err
		}
	}

	return []nitro.Value{&File{f}}, nil
}

type fileInfo struct {
	fs.FileInfo
}

func (i *fileInfo) String() string { return "<fileinfo>" }
func (i *fileInfo) Type() string   { return "fileinfo" }

func (i *fileInfo) EvalOp(op nitro.Op, operand nitro.Value) (nitro.Value, error) {
	return nil, errors.New("operation not supported")
}

func (i *fileInfo) Index(key nitro.Value) (nitro.Value, error) {
	keyStr, ok := key.(nitro.String)
	if !ok {
		return nil, fmt.Errorf(
			"fileinfo cannot be indexed by %q",
			nitro.TypeName(key))
	}

	switch keyStr.String() {
	case "name":
		return nitro.NewString(i.Name()), nil
	case "size":
		return nitro.NewInt(i.Size()), nil
	case "modtime":
		return NewTime(i.ModTime()), nil
	case "isdir":
		return nitro.NewBool(i.IsDir()), nil
	default:
		return nil, fmt.Errorf(
			"fileinfo does not have method %q",
			keyStr.String())
	}
}

func (i *fileInfo) IndexRef(key nitro.Value) (nitro.ValueRef, error) {
	return nitro.ValueRef{}, fmt.Errorf("fileinfo cannot be modified")
}

var errStatUsage = nitro.NewInvalidUsageError("stat(string)")

func stat(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 1 {
		return nil, errStatUsage
	}

	filepath, ok := args[0].(nitro.String)
	if !ok {
		return nil, errStatUsage
	}

	fi, err := os.Stat(filepath.String())
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []nitro.Value{nil}, nil
		}
		return nil, err
	}

	return []nitro.Value{&fileInfo{FileInfo: fi}}, nil
}

var errSeekUsage = errors.New(
	`invalid usage. Expected seek(seeker, int, string?)`)

func seek(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 2 && len(args) != 3 {
		return nil, errSeekUsage
	}

	seeker, ok := args[0].(io.Seeker)
	if !ok {
		return nil, errSeekUsage
	}

	offset, ok := args[1].(nitro.Int)
	if !ok {
		return nil, errSeekUsage
	}

	whence := os.SEEK_SET
	if len(args) == 3 {
		whenceArg, ok := args[2].(nitro.String)
		if !ok {
			return nil, errSeekUsage
		}

		switch whenceArg.String() {
		case "set":
			whence = os.SEEK_SET
		case "cur":
			whence = os.SEEK_CUR
		case "end":
			whence = os.SEEK_END
		default:
			return nil, fmt.Errorf(
				`%q is not a valid whence value. Valid values include: "set", "cur", "end"`,
				whenceArg.String())
		}
	}

	newOffset, err := seeker.Seek(offset.Int64(), whence)
	if err != nil {
		return nil, err
	}

	return []nitro.Value{nitro.NewInt(newOffset)}, nil
}

func closep(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) < 1 {
		return nil, errNotEnoughArgs
	}

	f, ok := args[0].(io.Closer)
	if !ok {
		return nil, fmt.Errorf(
			"expected argument 1 to be File, but it is %v",
			nitro.TypeName(args[0]))
	}

	err := f.Close()
	if err != nil && !errors.Is(err, os.ErrClosed) {
		return nil, err
	}

	return nil, nil
}

func create(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	filename, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}

	f, err := os.Create(filename)
	if err != nil {
		return nil, err
	}

	return []nitro.Value{&File{f}}, nil
}

var errReadUsage = errors.New(
	`invalid usage. Expected read(reader, int?)`)

func read(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	var err error

	if len(args) != 1 && len(args) != 2 {
		return nil, errReadUsage
	}

	reader, ok := args[0].(io.Reader)
	if !ok {
		return nil, errReadUsage
	}

	count := -1

	if len(args) == 2 {
		countArg, ok := args[1].(nitro.Int)
		if !ok {
			return nil, errReadUsage
		}
		count = int(countArg.Int64())
	}

	var data []byte
	if count == -1 {
		defer core.CloseReader(reader)
		data, err = ioutil.ReadAll(reader)
		if err != nil {
			return nil, err
		}
	} else {
		buf := make([]byte, count)
		n, err := io.ReadAtLeast(reader, buf, count)
		if err != nil && err != io.ErrUnexpectedEOF {
			return nil, err
		}
		data = buf[:n]
	}

	return []nitro.Value{nitro.NewString(string(data))}, nil
}

var errReadFileUsage = errors.New(
	`invalid usage. Expected read_file(string)`)

func readFile(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 1 {
		return nil, errReadFileUsage
	}

	filename, ok := args[0].(nitro.String)
	if !ok {
		return nil, errReadFileUsage
	}

	data, err := ioutil.ReadFile(filename.String())
	if err != nil {
		return nil, err
	}

	return []nitro.Value{nitro.NewString(string(data))}, nil
}

var errWriteFileUsage = errors.New(
	`invalid usage. Expected write_file(reader, string)`)

func writeFile(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 2 {
		return nil, errWriteFileUsage
	}

	src, err := nitro.MakeReader(m, args[0])
	if err != nil {
		return nil, errWriteFileUsage
	}

	filename, ok := args[1].(nitro.String)
	if !ok {
		return nil, errWriteFileUsage
	}

	dst, err := os.Create(filename.String())
	if err != nil {
		return nil, err
	}

	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return nil, err
	}

	err = dst.Close()
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func removeFile(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) < 1 {
		return nil, errNotEnoughArgs
	}

	var err error
	switch arg := args[0].(type) {
	case nitro.String:
		err = os.Remove(arg.String())

	case *File:
		arg.Close()
		err = os.Remove(arg.Name())

	default:
		return nil, fmt.Errorf(
			"expected argument 1 to be String or File, but it is %v",
			nitro.TypeName(args[0]))
	}

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []nitro.Value{nitro.False}, nil
		}
		return nil, err
	}

	return []nitro.Value{nitro.True}, nil
}

var errMoveFileUsage = nitro.NewInvalidUsageError("move_file(string, string)")

func moveFile(vm *nitro.VM, args []nitro.Value, nret int) ([]nitro.Value, error) {
	if len(args) != 2 {
		return nil, errMoveFileUsage
	}

	oldPath, ok := args[0].(nitro.String)
	if !ok {
		return nil, errMoveFileUsage
	}
	newPath, ok := args[1].(nitro.String)
	if !ok {
		return nil, errMoveFileUsage
	}

	err := os.Rename(oldPath.String(), newPath.String())
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func removeAll(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	path, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}

	err = os.RemoveAll(path)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

type createTempOptions struct {
	Dir     string `nitro:"dir"`
	Pattern string `nitro:"pattern"`
}

var createTempOptConv core.Value2Structer

var errCreateTempUsage = nitro.NewInvalidUsageError("create_temp(map?)")

func createTemp(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 0 && len(args) != 1 {
		return nil, errCreateTempUsage
	}

	var opt createTempOptions
	if len(args) == 1 {
		err := createTempOptConv.Convert(args[0], &opt)
		if err != nil {
			return nil, err
		}
	}

	f, err := ioutil.TempFile(opt.Dir, opt.Pattern)
	if err != nil {
		return nil, err
	}

	return []nitro.Value{&File{f}}, nil
}

func fileExists(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	path, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}
	_, err = os.Lstat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []nitro.Value{nitro.NewBool(false)}, nil
		}
		return nil, err
	}

	return []nitro.Value{nitro.NewBool(true)}, nil
}

func is_dir(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	path, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}
	fi, err := os.Lstat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []nitro.Value{nitro.NewBool(false)}, nil
		}
		return nil, err
	}

	return []nitro.Value{nitro.NewBool(fi.IsDir())}, nil
}

var errLsUsage = nitro.NewInvalidUsageError("ls(string)")

func ls(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 1 {
		return nil, errLsUsage
	}

	path, ok := args[0].(nitro.String)
	if !ok {
		return nil, errLsUsage
	}

	base, pattern := doublestar.SplitPattern(filepath.ToSlash(path.String()))
	if pattern == "" || pattern == "." {
		// This is a simple path. I.e. it does not include a pattern. Using
		// os.ReadDir is simpler/faster/leaner (no goroutines).
		entries, err := os.ReadDir(path.String())
		if err != nil {
			return nil, err
		}

		iter := &lsSimpleIter{
			root:    path.String(),
			entries: entries,
		}

		return []nitro.Value{nitro.NewIterator(iter.Next, nil, 2)}, nil
	}

	iter := newLSDoubleStarIter(base, pattern)
	return []nitro.Value{nitro.NewIterator(iter.Next, iter.Close, 2)}, nil
}

type lsDoubleStarIterEntry struct {
	path     string
	dirEntry fs.DirEntry
}

type lsDoubleStarIter struct {
	base    string
	pattern string
	outChan chan *lsDoubleStarIterEntry
	wg      sync.WaitGroup
	cancel  context.CancelFunc
}

func newLSDoubleStarIter(base, pattern string) *lsDoubleStarIter {
	i := &lsDoubleStarIter{
		base:    base,
		pattern: pattern,
		outChan: make(chan *lsDoubleStarIterEntry, 100),
	}

	var ctx context.Context
	ctx, i.cancel = context.WithCancel(context.Background())
	i.wg.Add(1)
	go i.run(ctx)

	return i
}

func (i *lsDoubleStarIter) run(ctx context.Context) {
	defer i.wg.Done()
	defer close(i.outChan)

	fsys := os.DirFS(i.base)
	doublestar.GlobWalk(fsys, i.pattern, func(path string, d fs.DirEntry) error {
		if d.Name() == "." || d.Name() == ".." {
			return ctx.Err()
		}

		select {
		case i.outChan <- &lsDoubleStarIterEntry{path, d}:
			return nil
		case <-ctx.Done():
			return ctx.Err()
		}
	})
}

func (i *lsDoubleStarIter) Next(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	entry, ok := <-i.outChan
	if !ok {
		i.cancel()
		return nil, nil
	}

	return []nitro.Value{
		nitro.NewString(filepath.FromSlash(filepath.Join(i.base, entry.path))),
		nitro.NewBool(entry.dirEntry.IsDir())}, nil
}

func (i *lsDoubleStarIter) Close(vm *nitro.VM) error {
	i.cancel()
	return nil
}

type lsSimpleIter struct {
	root    string
	entries []fs.DirEntry
}

func (i *lsSimpleIter) Next(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(i.entries) == 0 {
		return nil, nil
	}

	res := []nitro.Value{
		nitro.NewString(filepath.Join(i.root, i.entries[0].Name())),
		nitro.NewBool(i.entries[0].IsDir())}

	i.entries = i.entries[1:]
	return res, nil
}

func copyFile(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	to, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}
	from, err := getStringArg(args, 1)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(to)
	if err == nil && fi.IsDir() {
		to = filepath.Join(to, filepath.Base(from))
	}

	fromF, err := os.Open(from)
	if err != nil {
		return nil, err
	}
	defer fromF.Close()

	toF, err := os.Create(to)
	if err != nil {
		return nil, err
	}
	defer toF.Close()

	_, err = io.Copy(toF, fromF)
	if err != nil {
		return nil, err
	}

	err = toF.Close()
	if err != nil {
		return nil, err
	}

	return []nitro.Value{nitro.NewString(to)}, nil
}

var errMkdirAllUsage = nitro.NewInvalidUsageError("mkdir_all(string)")

func mkdirAll(vm *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 1 {
		return nil, errMkdirAllUsage
	}

	path, ok := args[0].(nitro.String)
	if !ok {
		return nil, errMkdirAllUsage
	}

	err := os.MkdirAll(path.String(), 0755)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
