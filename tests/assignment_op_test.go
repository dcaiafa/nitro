package tests

import "testing"

func TestAssignmentOp(t *testing.T) {
	RunSubO(t, "plus", `var a = 1; a += 2; print(a)`, `3`)
	RunSubO(t, "minus", `var a = 1; a -= 2; print(a)`, `-1`)
	RunSubO(t, "mul", `var a = 2; a *= 3; print(a)`, `6`)
	RunSubO(t, "div", `var a = 5; a /= 2.0; print(a)`, `2.5`)
	RunSubO(t, "mod", `var a = 5; a %= 2; print(a)`, `1`)
}
