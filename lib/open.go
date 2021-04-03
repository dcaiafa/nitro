package lib

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/dcaiafa/nitro"
)

type File struct {
	*os.File
}

func (f *File) String() string { return fmt.Sprintf("<File:%v>", f.File.Name()) }
func (f *File) Type() string   { return "File" }

func fnFOpen(ctx context.Context, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
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

func fnFCreate(ctx context.Context, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
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

func fnFRemove(ctx context.Context, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
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

func fnFRemoveAll(ctx context.Context, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
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

func fnFCreateTemp(ctx context.Context, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
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
