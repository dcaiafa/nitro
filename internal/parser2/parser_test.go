package parser2

import (
	"testing"

	"github.com/dcaiafa/nitro/internal/errlogger"
	"github.com/stretchr/testify/require"
)

func TestHelloWorld(t *testing.T) {
	prog := `
lines() | split("\t")

json_lines() | 
	join(json_lines("nodes.json"), ".metadata.node", ".node_name") |
	select(&e->e.metadata.labels.team=="t2")

select(
  join(
		json_lines(),
		json_lines("nodes.json"), ".metadata.node", ".node_name"),
	&e->e.metadata.labels.team=="t2")
		`

	module, err := Parse("test.nitro", prog, true, &errlogger.ConsoleErrLogger{})
	require.NoError(t, err)
	require.NotNil(t, module)
}
