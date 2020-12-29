package tests

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/dcaiafa/nitro"
)

type MemoryFileSystem map[string]string

func (fs MemoryFileSystem) ReadFile(name string) ([]byte, error) {
	data, ok := fs[name]
	if !ok {
		return nil, os.ErrNotExist
	}
	return []byte(data), nil
}

func Run(t *testing.T, prog string) {
	fs := make(MemoryFileSystem)
	fs["main"] = prog

	compiler := nitro.NewCompiler(fs)
	compiler.RegisterExternalFn(
		"print",
		func(ctx context.Context, args []nitro.Value) ([]nitro.Value, error) {
			fmt.Println(args)
			return nil, nil
		})

	compiled, err := compiler.Compile("main")
	if err != nil {
		t.Error(err.Error())
	}

	err = nitro.Run(context.Background(), compiled)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestTestTest(t *testing.T) {
	Run(t, "print(1, 2, 3)")
}
