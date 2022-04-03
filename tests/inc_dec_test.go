package tests

import "testing"

func TestIncDec(t *testing.T) {
	RunSubO(t, "inc", `var a = 2; a++; print(a)`, `3`)
	RunSubO(t, "inc_float", `var a = 2.5; a++; print(a)`, `3.5`)
	RunSubO(t, "sub", `var a = 2; a--; print(a)`, `1`)
	RunSubO(t, "sub_float", `var a = 2.5; a--; print(a)`, `1.5`)
}
