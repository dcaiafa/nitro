package tests

import "testing"

func TestObjectLiteral(t *testing.T) {
	RunSubO(t, "only_literals", `
		var a = {
			foo: "bar"
			other: 123
			sub: {
      	x: true
				[ "y" ]: false 
			}
		}
		print(a)
`, `{foo: "bar", other: 123, sub: {x: true, y: false}}`)

	RunSubO(t, "expr_value", `
		func f(x) {
			return {
				foo: "bar"
				other: x
				sub: {
					x: x > 10
					[ "y" ]: false 
				}
			}
		}
		print(f(12))
`, `{foo: "bar", other: 12, sub: {x: true, y: false}}`)

	RunSubO(t, "expr_key", `
		func f(x) {
			return {
				foo: "bar"
				[ x ]: 123
				sub: {
					x: true
					[ "y" ]: false 
				}
			}
		}
		print(f("other"))
`, `{foo: "bar", other: 123, sub: {x: true, y: false}}`)

	RunSubO(t, "keyword_keys", `
		var a = {
			and: 1
			else: 1
			false: 1
			func: 1
			for: 1
			if: 1
			in: 1
			meta: 1
			not: 1
			or: 1
			return: 1
			true: 1
			var: 1
			while: 1
		}
		print(a)
	`, `{and: 1, else: 1, false: 1, func: 1, for: 1, if: 1, in: 1, meta: 1, not: 1, or: 1, return: 1, true: 1, var: 1, while: 1}`)
}

func TestObjectMemberAccess(t *testing.T) {
	RunSubO(t, "rvalue", `
		var a = {
			foo: "bar"
			other: 123
			sub: {
      	x: true
				[ "y" ]: false 
			}
		}
		print(a.foo, a.invalid.reference, a.sub.y)
`, `bar <nil> false`)

	RunSubO(t, "lvalue", `
		var a = {
			foo: "bar"
			other: 123
			sub: {
      	x: true
				[ "y" ]: false 
			}
		}
		a.foo = "barr"
		a.sub.z = 123
		a.extra = { yay: "yo" }
		print(a)
`, `{foo: "barr", other: 123, sub: {x: true, y: false, z: 123}, extra: {yay: "yo"}}`)
}

func TestObjectOther(t *testing.T) {
	RunSubO(t, "has", `
		var a = { b: { c: 123 }, d: 3 }
		print(has(a, "b"), has(a, "c"), has(a.b, "c"), has(a.x.y.z, "w"))
	`, `true false true false`)

	RunSubO(t, "delete", `
		var a = {
			x: 1, y: 2, z: 3
		}
		delete(a, "x")
		delete(a, "x")
		print(a, len(a))
		a.x = 4
		print(a, len(a))
		delete(a, "x")
		delete(a, "y")
		delete(a, "z")
		print(a, len(a))
		a.w = 5
		delete(a.b, "nope")
		print(a, len(a))
	`, `
{y: 2, z: 3} 2
{y: 2, z: 3, x: 4} 3
{} 0
{w: 5} 1
`)
}
