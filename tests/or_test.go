package tests

import "testing"

func TestOr(t *testing.T) {
	RunSubO(t, "false_true", `print(0 or 123)`, `123`)
	RunSubO(t, "true_false", `print(123 or 0)`, `123`)
	RunSubO(t, "false_false", `print(false or 0)`, `0`)
	RunSubO(t, "true_true", `print(123 or 456)`, `123`)

	RunSubO(t, "short_circuit", `
		func f(b)
			print("f", b)
			return b
		end
		print(f(true) or f(true))
`, `
f true
true`)

	RunSubO(t, "dont_short_circuit", `
		func f(b)
			print("f", b)
			return b
		end
		print(f(false) or f(true))
`, `
f false
f true
true`)
}
