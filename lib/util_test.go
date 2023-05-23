package lib

import (
	"errors"
	"strings"
	"testing"

	"github.com/dcaiafa/nitro"
	"github.com/dcaiafa/nitro/internal/compiler"
	"github.com/dcaiafa/nitro/internal/export"
	"github.com/dcaiafa/nitro/internal/fs"
	"github.com/dcaiafa/nitro/internal/vm"
	libio "github.com/dcaiafa/nitro/lib/io"
)

func harnessCall(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	callable := args[0].(nitro.Callable)
	return m.Call(callable, args, nRet)
}

func compile(prog string) (*vm.Program, error) {
	compiler := compiler.New()

	RegisterAll(compiler)

	var harnessPackage = export.Exports{
		{N: "call", T: export.Func, F: harnessCall},
	}
	compiler.RegisterBuiltins("harness", harnessPackage)

	fs := fs.NewMem()
	fs.Put("main.n", []byte(prog))
	compiler.SetFS(fs)

	program, err := compiler.Compile("main.n")
	if err != nil {
		return nil, err
	}

	return program, nil
}

func run(prog string, params map[string]nitro.Value) (output string, err error) {
	compiled, err := compile(prog)
	if err != nil {
		return "", err
	}

	outBuilder := &strings.Builder{}

	vm := nitro.NewVM(compiled)
	libio.SetStdout(vm, outBuilder)

	for n, v := range params {
		err = vm.SetParam(n, v)
		if err != nil {
			return "", err
		}
	}

	err = vm.Run(nil)
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

func RunErr(t *testing.T, prog string, expectedErr error) {
	t.Helper()

	_, err := run(prog, nil)
	if err == nil {
		t.Fatalf("Error expected but operation succeeded")
	}

	if expectedErr != nil && !errors.Is(err, expectedErr) {
		t.Fatalf("Expected error %v, but received %v", expectedErr, err)
	}
}

func RunSubErr(t *testing.T, name string, prog string, expectedErr error) {
	t.Run(name, func(t *testing.T) {
		t.Helper()
		RunErr(t, prog, expectedErr)
	})
}
