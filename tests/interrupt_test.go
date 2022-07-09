package tests

import (
	"errors"
	"fmt"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/dcaiafa/nitro"
	"github.com/dcaiafa/nitro/internal/errlogger"
	"github.com/stretchr/testify/require"
)

func TestInterrupt(t *testing.T) {
	prog := `
		func stuff() {
			try {
				while true {
				}
			} catch e {
				print(e.error)
			}
		}
		stuff()
		stuff()
		while true {
		}
`
	outBuilder := &strings.Builder{}

	funcReg := make(simpleFuncRegistry)
	funcReg["print"] =
		func(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
			iargs := valuesToInterface(args)
			fmt.Fprintln(outBuilder, iargs...)
			return nil, nil
		}

	compiler := nitro.NewCompiler()
	compiler.AddFuncRegistry(funcReg)
	compiled, err := compiler.CompileSimple(
		"main.n",
		[]byte(prog),
		&errlogger.ConsoleErrLogger{})
	require.NoError(t, err)

	vm := nitro.NewVM(compiled)

	go func() {
		time.Sleep(100 * time.Millisecond)
		vm.SignalError(errors.New("INT1"))
		time.Sleep(100 * time.Millisecond)
		vm.SignalError(errors.New("INT2"))
		time.Sleep(100 * time.Millisecond)
		vm.SignalError(io.EOF)
	}()

	err = vm.Run(nil)
	if err == nil || !errors.Is(err, io.EOF) {
		t.Fatalf("expected io.EOF, returned %v", err)
	}

	output := strings.Trim(outBuilder.String(), "\r\n\t ")

	expected := `
INT1
INT2
`
	expected = strings.Trim(output, "\r\n\t ")
	require.Equal(t, expected, output)
}
