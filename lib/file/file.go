package file

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/dcaiafa/nitro"
	"github.com/dcaiafa/nitro/internal/vm"
	"github.com/dcaiafa/nitro/lib/core"
	"github.com/dcaiafa/nitro/lib/time"
)

//go:generate stubgen file.stubgen

type File struct {
	*os.File
}

var _ /* implements */ nitro.Indexable = (*File)(nil)
var _ /* implements */ nitro.Callable = (*File)(nil)
var _ /* implements */ core.NativeReader = (*File)(nil)
var _ /* implements */ core.NativeWriter = (*File)(nil)

func (f *File) String() string    { return fmt.Sprintf("File:%v", f.Name()) }
func (f *File) Type() string      { return "File" }
func (f *File) Traits() vm.Traits { return vm.TraitNone }

func (f *File) GetNativeReader() io.Reader {
	return f.File
}

func (f *File) GetNativeWriter() io.Writer {
	return f.File
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

func create0(vm *vm.VM, name string) (*File, error) {
	f, err := os.Create(name)
	if err != nil {
		return nil, err
	}
	file := &File{f}
	vm.RegisterCloser(file)
	return file, nil
}

func open0(vm *vm.VM, name string) (*File, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	file := &File{f}
	vm.RegisterCloser(file)
	return file, nil
}

func open1(vm *vm.VM, name string, opts *OpenOptions) (*File, error) {
	// TODO: permissions from Options.
	var perm os.FileMode = 0666
	flags := 0
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

	f, err := os.OpenFile(name, flags, perm)
	if err != nil {
		return nil, err
	}
	file := &File{f}
	vm.RegisterCloser(file)
	return file, nil
}

func stat0(vm *vm.VM, f *File) (*vm.Object, error) {
	fi, err := f.Stat()
	if err != nil {
		return nil, err
	}
	return fileInfoToMap(fi), nil
}

func stat1(vm *vm.VM, name string) (*vm.Object, error) {
	fi, err := os.Stat(name)
	if err != nil {
		return nil, err
	}
	return fileInfoToMap(fi), nil
}

func fileInfoToMap(fi os.FileInfo) *vm.Object {
	res := vm.NewObject()
	res.Put(vm.NewString("name"), vm.NewString(fi.Name()))
	res.Put(vm.NewString("size"), vm.NewInt(fi.Size()))
	res.Put(vm.NewString("mod_time"), time.NewTime(fi.ModTime()))
	res.Put(vm.NewString("is_dir"), vm.NewBool(fi.IsDir()))
	return res
}

func seek0(vm *vm.VM, f *File, offset int64, whenceStr string) (int64, error) {
	var whence int

	switch whenceStr {
	case "set":
		whence = os.SEEK_SET
	case "cur":
		whence = os.SEEK_CUR
	case "end":
		whence = os.SEEK_END
	default:
		return 0, fmt.Errorf(
			`%q is not a valid whence value. Valid values include: "set", "cur", "end"`,
			whenceStr)
	}

	newOffset, err := f.Seek(offset, whence)
	if err != nil {
		return 0, err
	}

	return newOffset, nil
}

func read0(vm *vm.VM, f *File, 

func fileRead(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) > 1 {
		return nil, errTooManyArgs
	} else if len(args) < 1 {
		return nil, errNotEnoughArgs
	}

	var err error
	var data []byte
	switch arg := args[0].(type) {
	case nitro.String:
		data, err = ioutil.ReadFile(arg.String())
		if err != nil {
			return nil, err
		}
	case io.Reader:
		defer core.CloseReader(arg)
		data, err = ioutil.ReadAll(arg)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errExpectedArg(0, args[0], "str", "reader")
	}

	return []nitro.Value{nitro.NewString(string(data))}, nil
}

func fileWriteTo(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) > 2 {
		return nil, errTooManyArgs
	}

	src, err := getReaderArg(m, args, 0)
	if err != nil {
		return nil, err
	}

	filename, err := getStringArg(args, 1)
	if err != nil {
		return nil, err
	}

	dst, err := os.Create(filename)
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

type createTempOptions struct {
	Dir     string `nitro:"dir"`
	Pattern string `nitro:"pattern"`
}

var createTempOptConv core.Value2Structer

func fileCreateTemp(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) > 1 {
		return nil, errTooManyArgs
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
