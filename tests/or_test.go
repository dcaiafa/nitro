package tests

import "testing"

func TestOr(t *testing.T) {
	RunSubO(t, "false_true", `print(false or true)`, `true`)
	RunSubO(t, "true_false", `print(true or false)`, `true`)
	RunSubO(t, "false_false", `print(false or false)`, `false`)
	RunSubO(t, "true_true", `print(true or true)`, `true`)
	RunSubO(t, "nil_true", `print(nil or true)`, `true`)
	RunSubO(t, "true_nil", `print(true or nil)`, `true`)
	RunSubO(t, "nil_false", `print(nil or false)`, `false`)
	RunSubO(t, "false_nil", `print(false or nil)`, `<nil>`)
	RunSubO(t, "notnil_true", `print(0 or true)`, `0`)
	RunSubO(t, "true_notnil", `print(true or 0)`, `true`)
	RunSubO(t, "notnil_false", `print(0 or false)`, `0`)
	RunSubO(t, "false_notnil", `print(false or 0)`, `0`)

	RunSubO(t, "short_circuit", `
		func f(b) {
			print("f", b)
			return b
		}
		print(f(true) or f(true))
`, `
f true
true`)

	RunSubO(t, "dont_short_circuit", `
		func f(b) {
			print("f", b)
			return b
		}
		print(f(false) or f(true))
`, `
f false
f true
true`)
}
