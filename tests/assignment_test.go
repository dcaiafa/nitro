package tests

import "testing"

func TestAssignment(t *testing.T) {
	RunSubO(t, "single", `
		var a
		a = 1
		print(a)`,
		`1`)

	RunSubO(t, "multi", `
		func f(x) 
			return x - 2
		end
		var a, b, c
		a, b, c = 1, 1+1, f(5)
		print(a, b, c)`,
		`1 2 3`)

	RunSubO(t, "func_single", `
  	func foo(x) 
			return x + 1
		end
		var a
		a = foo(1)
		print(a)
	`, `2`)

	RunSubO(t, "func_multi", `
  	func foo(x) 
			return x, x+1, x+2
		end
		var a, b, c
		a, b, c = foo(1)
		print(a, b, c)
	`, `1 2 3`)

	RunSubO(t, "func_multi_discard", `
  	func foo(x) 
			return x, x+1, x+2
		end
		var a, b, c
		a = foo(1)
		print(a)
	`, `1`)
}
