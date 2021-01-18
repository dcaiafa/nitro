package tests

import "testing"

func TestAnd(t *testing.T) {
	RunSubO(t, "false_true", `print(0 and true)`, `0`)
	RunSubO(t, "true_false", `print(true and 0)`, `0`)
	RunSubO(t, "false_false", `print(false and false)`, `false`)
	RunSubO(t, "true_true", `print(true and 123)`, `123`)

	RunSubO(t, "short_circuit", `
		fn f(b)
			print("f", b)
			return b
		end
		print(f(false) and f(true))
`, `
f false
false`)

	RunSubO(t, "dont_short_circuit", `
		fn f(b)
			print("f", b)
			return b
		end
		print(f(true) and f(false))
`, `
f true
f false
false`)
}
