package tests

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/dcaiafa/nitro"
	"github.com/dcaiafa/nitro/internal/errlogger"
	"github.com/dcaiafa/nitro/internal/runtime"
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
	fs["main.n"] = prog

	outBuilder := &strings.Builder{}

	compiler := nitro.NewCompiler(fs)
	compiler.SetDiag(true)

	compiler.AddExternalFn(
		"print",
		func(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, expRetN int) ([]nitro.Value, error) {
			iargs := valuesToInterface(args)
			fmt.Fprintln(outBuilder, iargs...)
			return nil, nil
		})

	compiler.AddExternalFn(
		"printf",
		func(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, expRetN int) ([]nitro.Value, error) {
			msg := args[0].(nitro.String)
			iargs := valuesToInterface(args[1:])
			fmt.Fprintf(outBuilder, msg.String()+"\n", iargs...)
			return nil, nil
		})

	compiled, err := compiler.Compile("main.n", &errlogger.ConsoleErrLogger{})
	if err != nil {
		return "", err
	}

	machine := nitro.NewMachine(context.Background(), compiled)
	for n, v := range params {
		err = machine.SetParam(n, v)
		if err != nil {
			return "", err
		}
	}

	err = machine.Run()
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

func RunPO(t *testing.T, prog string, params map[string]runtime.Value, expectedOutput string) {
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

func RunSubPO(t *testing.T, name string, prog string, params map[string]runtime.Value, expectedOutput string) {
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
