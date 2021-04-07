package tests

import "testing"

func TestAnd(t *testing.T) {
	RunSubO(t, "false_true", `print(false and true)`, `false`)
	RunSubO(t, "true_false", `print(true and false)`, `false`)
	RunSubO(t, "false_false", `print(false and false)`, `false`)
	RunSubO(t, "true_true", `print(true and true)`, `true`)
	RunSubO(t, "nil_true", `print(nil and true)`, `<nil>`)
	RunSubO(t, "true_nil", `print(true and nil)`, `<nil>`)
	RunSubO(t, "nil_false", `print(nil and false)`, `<nil>`)
	RunSubO(t, "false_nil", `print(false and nil)`, `false`)
	RunSubO(t, "notnil_true", `print(0 and true)`, `true`)
	RunSubO(t, "true_notnil", `print(true and 0)`, `0`)
	RunSubO(t, "notnil_false", `print(0 and false)`, `false`)
	RunSubO(t, "false_notnil", `print(false and 0)`, `false`)

	RunSubO(t, "short_circuit", `
		func f(b) {
			print("f", b)
			return b
		}
		print(f(false) and f(true))
`, `
f false
false`)

	RunSubO(t, "dont_short_circuit", `
		func f(b) {
			print("f", b)
			return b
		}
		print(f(true) and f(false))
`, `
f true
f false
false`)
}
