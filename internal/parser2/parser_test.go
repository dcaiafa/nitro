package parser2

import (
	"testing"

	"github.com/dcaiafa/nitro/internal/errlogger"
	"github.com/stretchr/testify/require"
)

func TestHelloWorld(t *testing.T) {
	prog := `
	var a = {
	  hello: "world"
		a: 123
	if x > 123 then
	  if: true, end: false,
		elif: 123
		else: xxx
	end
	  d: 123
	}
		`

	module, err := Parse("test.nitro", prog, true, &errlogger.ConsoleErrLogger{})
	require.NoError(t, err)
	require.NotNil(t, module)
}
