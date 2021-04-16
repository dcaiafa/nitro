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

func (f *File) String() string { return fmt.Sprintf("<File:%v>", f.File.Name()) }
func (f *File) Type() string   { return "File" }

func open(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	filename, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	return []nitro.Value{&File{f}}, nil
}

func fcreate(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
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

func fremove(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
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

func write(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	if len(args) < 2 {
		return nil, errNotEnoughArgs
	}
	var writer io.Writer
	switch arg := args[0].(type) {
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

	reader, err := ToReader(m, args[1])
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(writer, reader)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func fremoveall(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
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

func fcreatetemp(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
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

func fexists(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
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

func fisdir(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
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

func flist(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	path, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	iter := &flistIter{
		root:    path,
		entries: entries,
	}

	return []nitro.Value{nitro.NewIterator(iter.Next, nil, 2)}, nil
}

type flistIter struct {
	root    string
	entries []fs.DirEntry
}

func (i *flistIter) Next(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
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

func fcopy(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
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

func fpathbase(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	path, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}

	base := filepath.Base(path)

	return []nitro.Value{nitro.NewString(base)}, nil
}

func fpathclean(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	path, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}

	cleanPath := filepath.Clean(path)

	return []nitro.Value{nitro.NewString(cleanPath)}, nil
}

func fpathdir(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	path, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}

	dir := filepath.Dir(path)

	return []nitro.Value{nitro.NewString(dir)}, nil
}

func fpathext(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	path, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}

	ext := filepath.Ext(path)

	return []nitro.Value{nitro.NewString(ext)}, nil
}

func fpathfromslash(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	path, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}

	fromSlash := filepath.FromSlash(path)

	return []nitro.Value{nitro.NewString(fromSlash)}, nil
}

func fpathjoin(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	if len(args) < 2 {
		return nil, errNotEnoughArgs
	}

	paths := make([]string, len(args))

	var err error
	for i := 0; i < len(args); i++ {
		paths[i], err = getStringArg(args, i)
		if err != nil {
			return nil, err
		}
	}

	path := filepath.Join(paths...)
	return []nitro.Value{nitro.NewString(path)}, nil
}
