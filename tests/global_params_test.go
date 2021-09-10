package tests

import (
	"testing"

	"github.com/dcaiafa/nitro/internal/vm"
)

func TestGlobalParams(t *testing.T) {
	RunSubPO(t, "set_all", `
		!param foo = 1
		!flag bar {type:"int"}

		print(foo, bar)
	`, map[string]vm.Value{
		"foo": vm.NewInt(3),
		"bar": vm.NewInt(2),
	}, `3 2`)

	RunSubPO(t, "use_default", `
		!param foo = 1
		!flag bar

		print(foo, bar)
	`, map[string]vm.Value{
		"bar": vm.NewInt(2),
	}, `1 2`)

	// Even though `len` is a global function, the parameter `len` is defined in
	// the module's scope.
	RunSubPO(t, "shadow_global", `
		!param len
		print(len)
	`, map[string]vm.Value{
		"len": vm.NewInt(2),
	}, `2`)
}
