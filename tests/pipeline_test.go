package tests

import "testing"

func TestPipeline(t *testing.T) {
	RunSubO(t, "simple", `
		func add(x, y) {
			return x + y
		}
		1 + 2 | add(10) | (&x->x-1)() | print("yeah")
	`, `12 yeah`)
}
