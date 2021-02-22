package tests

import "testing"

func TestUnaryExpr(t *testing.T) {
	RunSubO(t, "minus/int", `print(-1)`, `-1`)
	RunSubO(t, "minus/float", `print(-1.2)`, `-1.2`)
	RunSubO(t, "not/int", `print(not 0, not 123)`, `true false`)
	RunSubO(t, "not/float", `print(not (1.2 - 1.2), not 3.1415)`, `true false`)
	RunSubO(t, "not/string", `print(not "", not "hi")`, `true false`)
	RunSubO(t, "not/object", `print(not {}, not {a:1})`, `true false`)
	RunSubO(t, "not/array", `print(not [], not [1])`, `true false`)
}
