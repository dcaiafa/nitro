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
	fn h()
		x()
	end
	fn g()
		h()
	end
	fn f()
		g()
	end
	f()
`

	_, err := run(prog, nil)
	require.Error(t, err)

	var rerr *nitro.RuntimeError

	require.True(t, errors.As(err, &rerr))

	expectedStack := []nitro.Frame{
		{Filename: "main.ni", Line: 4},
		{Filename: "main.ni", Line: 7},
		{Filename: "main.ni", Line: 10},
		{Filename: "main.ni", Line: 12},
	}

	require.Equal(t, expectedStack, rerr.Stack)
}
