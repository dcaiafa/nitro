package tests

import "testing"

func TestArrayLiteral(t *testing.T) {
	RunSubO(t, "literal", `
		var a = [1, "hello", nil, true]
		print(a)
	`, `[1 hello <nil> true]`)

	RunSubO(t, "empty_literal", `
		var a = []
		print(a)
	`, `[]`)

	RunSubO(t, "literal_with_exprs", `
		func f(x) {
			return [x, x+1, x+2]
		}
		print(f(10))
	`, `[10 11 12]`)

	RunSubO(t, "add", `
		var a = ["foo"]
		a | list.append("bar")
		print(a)
	`, `[foo bar]`)

	RunSubO(t, "additer_list", `
		var a = ["foo"]
		a | list.append_iter(["bar", "baz"])
		print(a)
	`, `[foo bar baz]`)

	RunSubO(t, "iter", `
		for e, i in ["a", "b", "c"] {
			print(e, i)
		}
	`, `
a 0
b 1
c 2
`)

}
