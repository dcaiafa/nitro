package lib

import (
	"context"
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/dcaiafa/nitro"
)

type MemoryFileLoader map[string]string

func (fs MemoryFileLoader) LoadFile(name string) ([]byte, error) {
	data, ok := fs[name]
	if !ok {
		return nil, os.ErrNotExist
	}
	return []byte(data), nil
}

func run(prog string, params map[string]nitro.Value) (output string, err error) {
	fs := make(MemoryFileLoader)
	fs["main.n"] = prog

	outBuilder := &strings.Builder{}

	compiler := nitro.NewCompiler(fs)
	compiler.SetDiag(true)

	compiler.AddNativeFn(
		"call",
		func(m *nitro.VM, caps []nitro.ValueRef, args []nitro.Value, nRet int) ([]nitro.Value, error) {
			callable := args[0].(nitro.Callable)
			return m.Call(callable, args, nRet)
		})

	RegisterAll(compiler)

	compiled, err := compiler.Compile("main.n", nitro.NewConsoleErrLogger())
	if err != nil {
		return "", err
	}

	vm := nitro.NewVM(context.Background(), compiled)
	for n, v := range params {
		err = vm.SetParam(n, v)
		if err != nil {
			return "", err
		}
	}

	SetStdout(vm, outBuilder)

	err = vm.Run()
	if err != nil {
		return "", err
	}

	output = strings.Trim(outBuilder.String(), "\r\n\t ")
	return output, nil
}

func RunO(t *testing.T, prog string, expectedOutput string) {
	t.Helper()

	expectedOutput = strings.Trim(expectedOutput, "\r\n\t ")

	output, err := run(prog, nil)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if output != expectedOutput {
		t.Fatalf("Expected output:\n%v\nActual:\n%v", expectedOutput, output)
	}
}

func RunSubO(t *testing.T, name string, prog string, expectedOutput string) {
	t.Run(name, func(t *testing.T) {
		t.Helper()
		RunO(t, prog, expectedOutput)
	})
}

func RunPO(t *testing.T, prog string, params map[string]nitro.Value, expectedOutput string) {
	t.Helper()

	expectedOutput = strings.Trim(expectedOutput, "\r\n\t ")

	output, err := run(prog, params)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if output != expectedOutput {
		t.Fatalf("Expected output:\n%v\nActual:\n%v", expectedOutput, output)
	}
}

func RunSubPO(t *testing.T, name string, prog string, params map[string]nitro.Value, expectedOutput string) {
	t.Run(name, func(t *testing.T) {
		t.Helper()
		RunPO(t, prog, params, expectedOutput)
	})
}

func RunErr(t *testing.T, prog string, expectedErr error) {
	t.Helper()

	_, err := run(prog, nil)
	if err == nil {
		t.Fatalf("Error expected but operation succeeded")
	}

	if !errors.Is(err, expectedErr) {
		t.Fatalf("Expected error %v, but received %v", expectedErr, err)
	}
}

func RunSubErr(t *testing.T, name string, prog string, expectedErr error) {
	t.Run(name, func(t *testing.T) {
		t.Helper()
		RunErr(t, prog, expectedErr)
	})
}
