package lib

import "testing"

func TestGet(t *testing.T) {
	RunSubO(t, "list", `[1, 2, 3] | get(1) | print`, `2`)
	RunSubO(t, "list_off_bounds", `[1, 2, 3] | get(4) | print`, `<nil>`)
	RunSubErr(t, "list_invalid_index_type", `[1, 2, 3] | get(3.14) | print`, nil)

	RunSubO(t, "map", `{ a: 1, b: "foo" } | get("b") | print`, `foo`)
	RunSubO(t, "map_not_found", `{ a: 1, b: "foo" } | get("c") | print`, `<nil>`)

	RunSubErr(t, "invalid_target", `3.1415 | get(1) | print`, nil)
}
