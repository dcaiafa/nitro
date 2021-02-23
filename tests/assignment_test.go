package tests

import "testing"

func TestAssignment(t *testing.T) {
	RunSubO(t, "single_literal", `
		var a
		a = 1
		print(a)`,
		`1`)

	RunSubO(t, "func", `
  	fn foo(x) 
			return x + 1
		end
		var a
		a = foo(1)
		print(a)
	`, `2`)
}
