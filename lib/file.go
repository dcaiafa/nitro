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

func (f *File) String() string { return fmt.Sprintf("File:%v", f.File.Name()) }
func (f *File) Type() string   { return "File" }

func (f *File) EvalBinOp(op nitro.BinOp, operand nitro.Value) (nitro.Value, error) {
	return nil, fmt.Errorf("file does not support this operation")
}

func (f *File) EvalUnaryMinus() (nitro.Value, error) {
	return nil, fmt.Errorf("file does not support this operation")
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

func read(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) < 1 {
		return nil, errNotEnoughArgs
	}

	var input io.Reader
	switch arg := args[0].(type) {
	case io.Reader:
		input = arg

	case nitro.String:
		f, err := os.Open(arg.String())
		if err != nil {
			return nil, err
		}
		input = f

	default:
		return nil, fmt.Errorf(
			"invalid argument %q. Expected Reader or String.",
			nitro.TypeName(arg))
	}

	defer CloseReader(input)

	data, err := ioutil.ReadAll(input)
	if err != nil {
		return nil, err
	}

	return []nitro.Value{nitro.NewString(string(data))}, nil
}

func write(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) < 2 {
		return nil, errNotEnoughArgs
	}

	reader, err := ToReader(m, args[0])
	if err != nil {
		return nil, err
	}

	var writer io.Writer
	switch arg := args[1].(type) {
	case nitro.String:
		f, err := os.Create(arg.String())
		if err != nil {
			return nil, err
		}
		defer f.Close()
		writer = f

	case io.Writer:
		writer = arg

	default:
		return nil, fmt.Errorf("invalid arg 2")
	}

	_, err = io.Copy(writer, reader)
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
