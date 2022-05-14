package lib

import (
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/dcaiafa/nitro"
)

type RunOption struct {
	beforeCompile func(c *nitro.Compiler)
	beforeRun     func(vm *nitro.VM)
}

func WithParams(params map[string]nitro.Value) RunOption {
	return RunOption{
		beforeRun: func(vm *nitro.VM) {
			for n, v := range params {
				err := vm.SetParam(n, v)
				if err != nil {
					panic(err)
				}
			}
		},
	}
}

func WithFn(name string, f func(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error)) RunOption {
	return RunOption{
		beforeCompile: func(c *nitro.Compiler) {
			c.AddNativeFn(name, f)
		},
	}
}

type MemoryFileLoader map[string]string

func (fs MemoryFileLoader) LoadFile(name string) ([]byte, error) {
	data, ok := fs[name]
	if !ok {
		return nil, os.ErrNotExist
	}
	return []byte(data), nil
}

func run(prog string, opts ...RunOption) (output string, err error) {
	fs := make(MemoryFileLoader)
	fs["main.n"] = prog

	outBuilder := &strings.Builder{}

	compiler := nitro.NewCompiler(fs)
	compiler.SetDiag(true)

	for _, opt := range opts {
		if opt.beforeCompile != nil {
			opt.beforeCompile(compiler)
		}
	}

	RegisterAll(compiler)

	compiled, err := compiler.Compile("main.n", nitro.NewConsoleErrLogger())
	if err != nil {
		return "", err
	}

	vm := nitro.NewVM(compiled)

	for _, opt := range opts {
		if opt.beforeRun != nil {
			opt.beforeRun(vm)
		}
	}

	SetStdout(vm, outBuilder)

	err = vm.Run(nil)
	if err != nil {
		return "", err
	}

	output = strings.Trim(outBuilder.String(), "\r\n\t ")
	return output, nil
}

func RunO(t *testing.T, prog string, expectedOutput string, opts ...RunOption) {
	t.Helper()

	expectedOutput = strings.Trim(expectedOutput, "\r\n\t ")

	output, err := run(prog, opts...)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if output != expectedOutput {
		t.Fatalf("Expected output:\n%v\nActual:\n%v", expectedOutput, output)
	}
}

func RunSubO(t *testing.T, name string, prog string, expectedOutput string, opts ...RunOption) {
	t.Run(name, func(t *testing.T) {
		t.Helper()
		RunO(t, prog, expectedOutput, opts...)
	})
}

func RunErr(t *testing.T, prog string, expectedErr error, opts ...RunOption) {
	t.Helper()

	_, err := run(prog, opts...)
	if err == nil {
		t.Fatalf("Error expected but operation succeeded")
	}

	if expectedErr != nil && !errors.Is(err, expectedErr) {
		t.Fatalf("Expected error %v, but received %v", expectedErr, err)
	}
}

func RunSubErr(t *testing.T, name string, prog string, expectedErr error, opts ...RunOption) {
	t.Run(name, func(t *testing.T) {
		t.Helper()
		RunErr(t, prog, expectedErr, opts...)
	})
}
