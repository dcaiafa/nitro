package tests

import (
	"testing"

	"github.com/dcaiafa/nitro/internal/vm"
)

func TestGlobalParams(t *testing.T) {
	RunSubPO(t, "set_all", `
		meta param foo = 1
		meta param bar [type="int"]

		print(foo, bar)
	`, map[string]vm.Value{
		"foo": vm.NewInt(3),
		"bar": vm.NewInt(2),
	}, `3 2`)

	RunSubPO(t, "use_default", `
		meta param foo = 1
		meta param bar

		print(foo, bar)
	`, map[string]vm.Value{
		"bar": vm.NewInt(2),
	}, `1 2`)
}
