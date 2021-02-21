package tests

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/dcaiafa/nitro"
	"github.com/dcaiafa/nitro/std"
)

type MemoryFileSystem map[string]string

func (fs MemoryFileSystem) ReadFile(name string) ([]byte, error) {
	data, ok := fs[name]
	if !ok {
		return nil, os.ErrNotExist
	}
	return []byte(data), nil
}

func valuesToInterface(values []nitro.Value) []interface{} {
	ivalues := make([]interface{}, len(values))
	for i, v := range values {
		ivalues[i] = v
	}
	return ivalues
}

func run(prog string) (output string, err error) {
	fs := make(MemoryFileSystem)
	fs["main"] = prog

	outBuilder := &strings.Builder{}

	compiler := nitro.NewCompiler(fs)

	compiler.RegisterExternalFn(
		"tostring", std.ToString)

	compiler.RegisterExternalFn(
		"print",
		func(ctx context.Context, args []nitro.Value) ([]nitro.Value, error) {
			iargs := valuesToInterface(args)
			fmt.Fprintln(outBuilder, iargs...)
			return nil, nil
		})

	compiler.RegisterExternalFn(
		"printf",
		func(ctx context.Context, args []nitro.Value) ([]nitro.Value, error) {
			msg := args[0].(nitro.String)
			iargs := valuesToInterface(args[1:])
			fmt.Fprintf(outBuilder, string(msg)+"\n", iargs...)
			return nil, nil
		})

	compiled, err := compiler.Compile("main")
	if err != nil {
		return "", err
	}

	err = nitro.Run(context.Background(), compiled)
	if err != nil {
		return "", err
	}

	output = strings.Trim(outBuilder.String(), "\r\n\t ")
	return output, nil
}

func RunO(t *testing.T, prog string, expectedOutput string) {
	t.Helper()

	expectedOutput = strings.Trim(expectedOutput, "\r\n\t ")

	output, err := run(prog)
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
