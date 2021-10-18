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

	RunSubO(t, "if", `
		func f(even) {
			return [
			if even {
				2
			}
			if not even {
				3
			}
			if even {
				4
			}
			if not even {
				5
			}
			]
		}
		print(f(true))
		print(f(false))
	`, `
[2 4]
[3 5]
`)

	RunSubO(t, "if_else", `
		func f(x) {
			return [
				if x == "odd" {
					1, 3, 5
				} else if x == "even" {
					2, 4, 8
				} else {
					0
				}
			]
		}
		print(f("odd"), f("even"), f("other"))
		`, `[1 3 5] [2 4 8] [0]`)

	RunSubO(t, "for", `
		var a = [
			0
		for x in range(1, 10, 2) {
    	x
			-(x + 1)
		}
			99
		]
		print(a)
		`, `[0 1 -2 3 -4 5 -6 7 -8 9 -10 99]`)

	RunSubO(t, "for_multi", `
		var r = {
			names: [
			for f in [ "alice", "bob" ] {
				for l in [ "smith", "hunter" ] {
			{ first: f, last: l }
				}
			}
			]
		}
		print(r)
		`, `{names: [{first: "alice", last: "smith"} {first: "alice", last: "hunter"} {first: "bob", last: "smith"} {first: "bob", last: "hunter"}]}`)

	RunSubO(t, "add", `
		var a = ["foo"]
		a.add("bar")
		print(a)
	`, `[foo bar]`)

	RunSubO(t, "additer_list", `
		var a = ["foo"]
		a.add_iter(["bar", "baz"])
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
