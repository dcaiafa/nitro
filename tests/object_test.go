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
		func f(x) 
			return {
				foo: "bar"
				other: x
				sub: {
					x: x > 10
					[ "y" ]: false 
				}
			}
		end
		print(f(12))
`, `{foo: "bar", other: 12, sub: {x: true, y: false}}`)

	RunSubO(t, "expr_key", `
		func f(x) 
			return {
				foo: "bar"
				[ x ]: 123
				sub: {
					x: true
					[ "y" ]: false 
				}
			}
		end
		print(f("other"))
`, `{foo: "bar", other: 123, sub: {x: true, y: false}}`)

	RunSubO(t, "if", `
		func f(a, b) 
			return {
				foo: "bar"
			if a
				other: 123
				sub: {
				if b
					x: true
				end
					[ "y" ]: false 
				}
			end
			}
    end
		print(f(false, false))
		print(f(false, true))
		print(f(true, false))
		print(f(true, true))
		`, `
{foo: "bar"}
{foo: "bar"}
{foo: "bar", other: 123, sub: {y: false}}
{foo: "bar", other: 123, sub: {x: true, y: false}}`)

	RunSubO(t, "keyword_keys", `
		var a = {
			and: 1
			else: 1
			end: 1
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
	`, `{and: 1, else: 1, end: 1, false: 1, func: 1, for: 1, if: 1, in: 1, meta: 1, not: 1, or: 1, return: 1, true: 1, var: 1, while: 1}`)
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
