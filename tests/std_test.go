package tests

import "testing"

func TestLen(t *testing.T) {
	RunSubO(t, "nil", `var a; print(len(a))`, `0`)
	RunSubO(t, "object", `var a = {a: 1, b: 3.1415, c: false}; print(len(a), len({}))`, `3 0`)
	RunSubO(t, "array", `var a = [1, 2]; print(len(a), len([]))`, `2 0`)
	RunSubO(t, "string", `var a = "foobar"; print(len(a), len(""))`, `6 0`)
}
