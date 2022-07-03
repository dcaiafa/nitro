package tests

import (
	"errors"
	"testing"

	"github.com/dcaiafa/nitro"
	"github.com/stretchr/testify/require"
)

func TestErrorStack(t *testing.T) {
	const prog = `
	var x
	func h() {
		x()
	}
	func g() {
		h()
	}
	func f() {
		g()
	}
	f()
`

	_, err := run(prog, nil)
	require.Error(t, err)

	var rerr *nitro.RuntimeError

	require.True(t, errors.As(err, &rerr))

	expectedStack := []nitro.Frame{
		{Filename: "main.n", Line: 4, Func: "h"},
		{Filename: "main.n", Line: 7, Func: "g"},
		{Filename: "main.n", Line: 10, Func: "f"},
		{Filename: "main.n", Line: 12, Func: "$main"},
	}

	require.Equal(t, expectedStack, rerr.Stack)
}
