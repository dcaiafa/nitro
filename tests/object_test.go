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
		print(tostring(a))
`, `{foo: "bar", other: 123, sub: {x: true, y: false}}`)

	RunSubO(t, "expr_value", `
		fn f(x) 
			return {
				foo: "bar"
				other: x
				sub: {
					x: x > 10
					[ "y" ]: false 
				}
			}
		end
		print(tostring(f(12)))
`, `{foo: "bar", other: 12, sub: {x: true, y: false}}`)

	RunSubO(t, "expr_key", `
		fn f(x) 
			return {
				foo: "bar"
				[ x ]: 123
				sub: {
					x: true
					[ "y" ]: false 
				}
			}
		end
		print(tostring(f("other")))
`, `{foo: "bar", other: 123, sub: {x: true, y: false}}`)

	RunSubO(t, "if", `
		fn f(a, b) 
			return {
				foo: "bar"
			if a then
				other: 123
				sub: {
				if b then
					x: true
				end
					[ "y" ]: false 
				}
			end
			}
    end
		print(tostring(f(false, false)))
		print(tostring(f(false, true)))
		print(tostring(f(true, false)))
		print(tostring(f(true, true)))
		`, `
{foo: "bar"}
{foo: "bar"}
{foo: "bar", other: 123, sub: {y: false}}
{foo: "bar", other: 123, sub: {x: true, y: false}}
`)

}
