package parser2

import (
	"testing"

	"github.com/dcaiafa/nitro/internal/errlogger"
	"github.com/stretchr/testify/require"
)

func TestHelloWorld(t *testing.T) {
	prog := `
	print(a)
		`

	module, err := Parse("test.nitro", prog, true, &errlogger.ConsoleErrLogger{})
	require.NoError(t, err)
	require.NotNil(t, module)
}
