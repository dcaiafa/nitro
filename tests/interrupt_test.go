package tests

import (
	"errors"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/dcaiafa/nitro"
	"github.com/dcaiafa/nitro/lib"
	"github.com/stretchr/testify/require"
)

func TestInterrupt(t *testing.T) {
	code := `
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
	prog, err := compile(code)
	require.NoError(t, err)

	outBuilder := &strings.Builder{}

	vm := nitro.NewVM(prog)
	lib.SetStdout(vm, outBuilder)

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
