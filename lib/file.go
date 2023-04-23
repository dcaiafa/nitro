package lib

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/dcaiafa/nitro"
	"github.com/dcaiafa/nitro/internal/vm"
	"github.com/dcaiafa/nitro/lib/core"
	"github.com/dcaiafa/nitro/lib/time"
)

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

func fileOpen(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	var err error

	if len(args) > 2 {
		return nil, errTooManyArgs
	}

	filename, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}

	var opts *openOptions
	if len(args) == 2 {
		opts = new(openOptions)
		optsMap, err := getObjectArg(args, 1)
		if err != nil {
			return nil, err
		}
		err = openOptionsConv.Convert(optsMap, opts)
		if err != nil {
			return nil, err
		}
	}

	var f *os.File
	if opts == nil {
		f, err = os.Open(filename)
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
		f, err = os.OpenFile(filename, flags, perm)
		if err != nil {
			return nil, err
		}
	}

	return []nitro.Value{&File{f}}, nil
}

func fileStat(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) > 1 {
		return nil, errTooManyArgs
	}

	filename, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []nitro.Value{nil}, nil
		}
		return nil, err
	}

	res := nitro.NewObject()
	res.Put(nitro.NewString("name"), nitro.NewString(fi.Name()))
	res.Put(nitro.NewString("size"), nitro.NewInt(fi.Size()))
	res.Put(nitro.NewString("mod_time"), time.NewTime(fi.ModTime()))
	res.Put(nitro.NewString("is_dir"), nitro.NewBool(fi.IsDir()))

	return []nitro.Value{res}, nil
}

func fileSeek(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) > 3 {
		return nil, errTooManyArgs
	}

	file, err := getFileArg(args, 0)
	if err != nil {
		return nil, err
	}

	offset, err := getIntArg(args, 1)
	if err != nil {
		return nil, err
	}

	whence := os.SEEK_SET
	if len(args) == 3 {
		whenceArg, err := getStringArg(args, 2)
		if err != nil {
			return nil, err
		}

		switch whenceArg {
		case "set":
			whence = os.SEEK_SET
		case "cur":
			whence = os.SEEK_CUR
		case "end":
			whence = os.SEEK_END
		default:
			return nil, fmt.Errorf(
				`%q is not a valid whence value. Valid values include: "set", "cur", "end"`,
				whenceArg)
		}
	}

	newOffset, err := file.Seek(offset, whence)
	if err != nil {
		return nil, err
	}

	return []nitro.Value{nitro.NewInt(newOffset)}, nil
}

func fileCreate(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
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
