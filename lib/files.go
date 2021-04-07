package lib

import (
	"errors"
	"fmt"
	"io"
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

func fnFOpen(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
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

func fnFCreate(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
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

func fnFRemove(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
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

func fnFRemoveAll(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
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

func fnFCreateTemp(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
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

func fnFExists(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
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

func fnFList(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	path, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}

	var pattern string
	if len(args) >= 2 {
		pattern, err = getStringArg(args, 1)
		if err != nil {
			return nil, err
		}
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	result := nitro.NewArrayFromSlice(make([]nitro.Value, 0, len(entries)))

	for _, entry := range entries {
		if pattern != "" {
			ok, err := filepath.Match(pattern, entry.Name())
			if err != nil {
				return nil, err
			}
			if !ok {
				continue
			}
		}
		result.Push(nitro.NewString(filepath.Join(path, entry.Name())))
	}

	return []nitro.Value{result}, nil
}

func fnFCopy(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
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

func fnFPathBase(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	path, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}

	base := filepath.Base(path)

	return []nitro.Value{nitro.NewString(base)}, nil
}

func fnFPathClean(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	path, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}

	cleanPath := filepath.Clean(path)

	return []nitro.Value{nitro.NewString(cleanPath)}, nil
}

func fnFPathDir(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	path, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}

	dir := filepath.Dir(path)

	return []nitro.Value{nitro.NewString(dir)}, nil
}

func fnFPathExt(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	path, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}

	ext := filepath.Ext(path)

	return []nitro.Value{nitro.NewString(ext)}, nil
}

func fnFPathFromSlash(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	path, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}

	fromSlash := filepath.FromSlash(path)

	return []nitro.Value{nitro.NewString(fromSlash)}, nil
}

func fnFPathJoin(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
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
