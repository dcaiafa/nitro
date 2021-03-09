package tests

import "testing"

func TestArrayLiteral(t *testing.T) {
	RunSubO(t, "literal", `
		var a = [1, "hello", true]
		print(a)
	`, `[1 hello true]`)

	RunSubO(t, "empty_literal", `
		var a = []
		print(a)
	`, `[]`)

	RunSubO(t, "literal_with_exprs", `
		fn f(x)
			return [x, x+1, x+2]
		end
		print(f(10))
	`, `[10 11 12]`)

	RunSubO(t, "if", `
		fn f(even) 
			return [
			if even
				2
			end
			if not even
				3
			end
			if even
				4
			end
			if not even
				5
			end
			]
		end
		print(f(true))
		print(f(false))
	`, `
[2 4]
[3 5]
`)
}
