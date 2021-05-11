package lib

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/dcaiafa/nitro"
)

type File struct {
	*os.File
}

var _ /* implements */ nitro.Indexable = (*File)(nil)
var _ /* implements */ nitro.Callable = (*File)(nil)

func (f *File) String() string { return fmt.Sprintf("File:%v", f.File.Name()) }
func (f *File) Type() string   { return "File" }

func (f *File) EvalOp(op nitro.Op, operand nitro.Value) (nitro.Value, error) {
	return nil, fmt.Errorf("file does not support this operation")
}

func (f *File) Call(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 1 {
		return nil, errWriterCallUsage
	}

	reader, err := ToReader(m, args[0])
	if err != nil {
		return nil, errWriterCallUsage
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
		return nitro.NewString(f.File.Name()), nil

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

var openOptionsConv Value2Structer

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
		err := openOptionsConv.Convert(optsMap, opts)
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
		defer CloseReader(reader)
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
	`invalid usage. Expected readfile(string)`)

func readfile(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
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

var errWriteToUsage = errors.New(
	`invalid usage. Expected writeto(reader, writer, int?)`)

func writeto(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 2 && len(args) != 3 {
		return nil, errWriteToUsage
	}

	src, err := ToReader(m, args[0])
	if err != nil {
		return nil, errWriteToUsage
	}

	dst, ok := args[1].(io.Writer)
	if !ok {
		return nil, errWriteToUsage
	}

	if len(args) == 3 {
		count, ok := args[2].(nitro.Int)
		if !ok {
			return nil, errWriteToUsage
		}

		n, err := io.CopyN(dst, src, count.Int64())
		if err != nil && err != io.EOF {
			return nil, err
		}

		return []nitro.Value{nitro.NewInt(n)}, nil
	} else {
		n, err := io.Copy(dst, src)
		if err != nil {
			return nil, err
		}
		return []nitro.Value{nitro.NewInt(n)}, nil
	}
}

var errWriteFileUsage = errors.New(
	`invalid usage. Expected writefile(reader, string)`)

func writefile(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 2 {
		return nil, errWriteFileUsage
	}

	src, err := ToReader(m, args[0])
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

func rm(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) < 1 {
		return nil, errNotEnoughArgs
	}

	switch arg := args[0].(type) {
	case nitro.String:
		err := os.Remove(arg.String())
		if err != nil {
			return nil, err
		}

	case *File:
		arg.Close()
		err := os.Remove(arg.Name())
		if err != nil {
			return nil, err
		}

	default:
		return nil, fmt.Errorf(
			"expected argument 1 to be String or File, but it is %v",
			nitro.TypeName(args[0]))
	}

	return nil, nil
}

func rmall(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
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

func createtemp(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	pattern, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}

	f, err := ioutil.TempFile("", pattern)
	if err != nil {
		return nil, err
	}

	return []nitro.Value{&File{f}}, nil
}

func fileexists(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
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

func isdir(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
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

func ls(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	path, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	iter := &lsIter{
		root:    path,
		entries: entries,
	}

	return []nitro.Value{nitro.NewIterator(iter.Next, 2)}, nil
}

type lsIter struct {
	root    string
	entries []fs.DirEntry
}

func (i *lsIter) Next(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(i.entries) == 0 {
		return []nitro.Value{nitro.NewBool(false), nil, nil}, nil
	}

	res := []nitro.Value{
		nitro.NewBool(true),
		nitro.NewString(filepath.Join(i.root, i.entries[0].Name())),
		nitro.NewBool(i.entries[0].IsDir())}

	i.entries = i.entries[1:]
	return res, nil
}

func cp(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	from, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}

	to, err := getStringArg(args, 1)
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
