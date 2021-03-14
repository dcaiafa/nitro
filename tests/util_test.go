package tests

import (
	"context"
	"fmt"
	"os"
	"strings"
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

func valuesToInterface(values []nitro.Value) []interface{} {
	ivalues := make([]interface{}, len(values))
	for i, v := range values {
		switch v := v.(type) {
		case nitro.Int:
			ivalues[i] = v.Int64()
		case nitro.Float:
			ivalues[i] = v.Float64()
		case nitro.String:
			ivalues[i] = v.String()
		default:
			ivalues[i] = v
		}
	}
	return ivalues
}

func run(prog string, params map[string]nitro.Value) (output string, err error) {
	fs := make(MemoryFileSystem)
	fs["main"] = prog

	outBuilder := &strings.Builder{}

	compiler := nitro.NewCompiler(fs)
	compiler.SetDiag(true)

	compiler.AddExternalFn(
		"print",
		func(ctx context.Context, caps []nitro.ValueRef, args []nitro.Value) ([]nitro.Value, error) {
			iargs := valuesToInterface(args)
			fmt.Fprintln(outBuilder, iargs...)
			return nil, nil
		})

	compiler.AddExternalFn(
		"printf",
		func(ctx context.Context, caps []nitro.ValueRef, args []nitro.Value) ([]nitro.Value, error) {
			msg := args[0].(nitro.String)
			iargs := valuesToInterface(args[1:])
			fmt.Fprintf(outBuilder, msg.String()+"\n", iargs...)
			return nil, nil
		})

	compiled, err := compiler.Compile("main")
	if err != nil {
		return "", err
	}

	err = nitro.Run(context.Background(), compiled, params)
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
