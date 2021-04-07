package tests

import "testing"

func TestUnaryExpr(t *testing.T) {
	RunSubO(t, "not/true", `print(not true)`, `false`)
	RunSubO(t, "not/false", `print(not false)`, `true`)
	RunSubO(t, "not/nil", `print(not nil)`, `true`)
	RunSubO(t, "minus/int", `print(-1)`, `-1`)
	RunSubO(t, "minus/float", `print(-1.2)`, `-1.2`)
	RunSubO(t, "not/int", `print(not 0, not 123)`, `false false`)
	RunSubO(t, "not/float", `print(not (1.2 - 1.2), not 3.1415)`, `false false`)
	RunSubO(t, "not/string", `print(not "", not "hi")`, `false false`)
	RunSubO(t, "not/object", `print(not {}, not {a:1})`, `false false`)
	RunSubO(t, "not/array", `print(not [], not [1])`, `false false`)
}
